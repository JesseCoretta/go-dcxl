package dcxl

import (
	"fmt"
	"io"
	"reflect"
	"strconv"
	"strings"
	"time"
	"unicode"
)

const datefmt = `20060102150405` // for gen. timestamp -> time.Time

var (
	sprintf  func(string, ...any) string                  = fmt.Sprintf
	fprintf  func(io.Writer, string, ...any) (int, error) = fmt.Fprintf
	fprintln func(io.Writer, string, ...any) (int, error) = fmt.Fprintf
	printf   func(string, ...any) (int, error)            = fmt.Printf

	atoi func(string) (int, error) = strconv.Atoi
	itoa func(int) string          = strconv.Itoa

	fields     func(string) []string               = strings.Fields
	hasPrefix  func(string, string) bool           = strings.HasPrefix
	hasSuffix  func(string, string) bool           = strings.HasSuffix
	idxRune    func(string, rune) int              = strings.IndexRune
	join       func([]string, string) string       = strings.Join
	lc         func(string) string                 = strings.ToLower
	uc         func(string) string                 = strings.ToUpper
	split      func(string, string) []string       = strings.Split
	eq         func(string, string) bool           = strings.EqualFold
	contains   func(string, string) bool           = strings.Contains
	splitAfter func(string, string) []string       = strings.SplitAfter
	splitN     func(string, string, int) []string  = strings.SplitN
	trimS      func(string) string                 = strings.TrimSpace
	trimL      func(string, string) string         = strings.TrimLeft
	trimR      func(string, string) string         = strings.TrimRight
	replaceAll func(string, string, string) string = strings.ReplaceAll

	isLetter func(rune) bool = unicode.IsLetter
	isDigit  func(rune) bool = unicode.IsDigit

	since func(time.Time) time.Duration = time.Since
	until func(time.Time) time.Duration = time.Until
	now   func() time.Time              = time.Now

	typeOf func(any) reflect.Type  = reflect.TypeOf
	valOf  func(any) reflect.Value = reflect.ValueOf

	altnames map[string]string = map[string]string{
		`numberForm`: `n`,          // s. 2.1.1.
		`nameForm`:   `identifier`, // s. 2.1.6.
	}
)

/*
timeToGenTime converts a time.Time value into a generalizedTime string
value, which is returned alongside a success-indicative boolean value.
*/
func timeToGenTime(gt time.Time) (v string, ok bool) {
	// Don't waste time on
	// null time values
	if gt.IsZero() {
		return
	}

	// If time.Time value (gt) is not
	// already set to UTC, do so now.
	if gt.Location() != time.UTC {
		gt = gt.UTC()
	}

	// Use fmt.Sprintf alias to compose
	// the string-based generalizedTime
	// zulu value (v)
	v = sprintf("%04d%02d%02d%02d%02d%02dZ",
		gt.Year(),   // 4
		gt.Month(),  // 2
		gt.Day(),    // 2
		gt.Hour(),   // 2
		gt.Minute(), // 2
		gt.Second()) // 2 = ( 14 + 1 = 15 )

	// If not exactly 15 bytes in
	// length, we're boned.
	ok = len(v) == 15

	return
}

/*
genTimeToTime converts a generalizedTime value into an instance of
time.Time, which is returned alongside a success-indicative boolean value.

NOTE: this function will understand both UTC time AND UTC-offset time for
input, but will enforce UTC time output.
*/
func genTimeToTime(gt string) (v time.Time, ok bool) {
	// The absolute minimum valid length
	// of generalized time should be
	// fifteen (15) bytes.
	if len(gt) < 15 {
		return
	}

	var (
		zulu byte   = byte('Z') // UTC
		uoff string = `-0700`   // UTC offset
		err  error
	)

	if gt[len(gt)-1] == zulu {
		v, err = time.Parse(datefmt+string(zulu), gt)
	} else {
		v, err = time.Parse(datefmt+string(uoff), gt)
		v = v.UTC()
	}
	ok = !v.IsZero() && err == nil

	return
}

func isNumber(n string) bool {
	if len(n) == 0 {
		return false
	}

	for i := 0; i < len(n); i++ {
		if !isDigit(rune(n[i])) {
			return false
		}
	}
	return true
}

/*
isPtr returns a boolean value indicative of whether kind
reflection revealed the presence of a pointer type.
*/
func isPtr(x any) bool {
	return typeOf(x).Kind() == reflect.Ptr
}

/*
isStruct returns a boolean value indicative of whether
kind reflection revealed the presence of a struct type.
*/
func isStruct(x any) bool {
	return typeOf(x).Kind() == reflect.Struct
}

func getReflectInstances(a any) (t reflect.Type, v reflect.Value, ok bool) {
	ot := typeOf(a) // reflect.Type
	ov := valOf(a)  // reflect.Value

	// If this is a pointer instance,
	// we'll need to reference the
	// *element* of the pointer.
	if isPtr(a) {
		ot = ot.Elem()
		ov = ov.Elem()
	}

	// Whether pointer or not, this
	// function only handles structs.
	if !isStruct(ov) {
		return
	}

	t = ot
	v = ov
	ok = true

	return
}

/*
toMap processes the provided interface type (a) into
an instance of map[string][]string, which is then
returned. Ultimately, this function exists in order
to provide an effective means of turning one of the
structs defined in this package into something that
can be fed to ldap.NewEntry.
*/
func toMap(a any) (m map[string][]string) {
	ot, ov, ok := getReflectInstances(a)
	if !ok {
		return
	}

	// prepare for any objectClasses
	// we'll want to record. This will
	// work whether or not combined
	// entries are in use.
	m = make(map[string][]string, 0)
	m[`objectClass`] = make([]string, 0)
	m[`objectClass`] = append(m[`objectClass`], `top`) // always include top of superchain

	// Look for a method called 'ObjectClass' with
	// a null input signature and a string output
	// signature. If found, run it, and append the
	// value returned to our OC list.
	method := ov.MethodByName(`ObjectClass`)
	if meth, ok := method.Interface().(func() string); ok {
		// If the return value from meth() (an objectClass
		// name) is not already present in the map value
		// slice, then add it.
		if o := meth(); !strInSlice(o, m[`objectClass`]) {
			m[`objectClass`] = append(m[`objectClass`], o)
		}
	}

	// Begin iterating each struct field
	for i := 0; i < ot.NumField(); i++ {
		ft := ot.Field(i)
		fv := ov.Field(i)

		// IsZero would mean, in this context,
		// that the AT was not specified, so
		// we'll skip it.
		if fv.IsZero() {
			continue
		}

		// Look for the `ldap` tag and
		// extract the AT name if one
		// is present. When dealing with
		// embedded fields, such as when
		// doing so-called "Combined"
		// entries, there may not be a tag
		// but we still need to process it
		// in some way (as such, we won't
		// bail outright).
		var at string
		tag, found := extractLTag(ft.Name, a)
		if found {
			// Don't include the DN, as ldap.NewEntry
			// expects the DN as a separate argument.
			if eq(tag[0], `dn`) {
				continue
			}

			var ok bool
			if at, ok = altnames[tag[0]]; !ok {
				at = tag[0]
			}
		}

		// This package introduces struct types that
		// really only contain one (1) of three (3)
		// possible types: string, []string and
		// embedded structs (at the time of this
		// writing, three Registrant struct types
		// qualify for this).
		switch fv.Kind() {
		case reflect.String:
			// We enclose the single string value
			// in a string slice to accommdate the
			// map signature required by ldap.NewEntry.
			if found {
				m[at] = []string{fv.String()}
			}
		case reflect.Slice:
			// Risky, I know ... but slice types
			// should always be []string in this
			// package for the foreseeable future.
			if found {
				m[at] = fv.Interface().([]string)
			}
		case reflect.Struct:
			m = structToMap(fv, m)
		}
	}

	return
}

func structToMap(src reflect.Value, m map[string][]string) map[string][]string {
	// Launch a new iteration of this function,
	// but using the nested (embedded) type we've
	// arrived at. This HAS to be one of the three
	// defined Registrant interface types, else m2
	// will be an empty map.
	if m2 := toMap(src.Interface()); len(m2) > 0 {
		for k, v := range m2 {
			if eq(k, `objectClass`) {
				// Append to previous iteration's
				// OC list, rather than clobbering
				// it in this iteration. Ignore
				// duplicates in the event of more
				// than one embedded value being
				// present (otherwise, N of the same
				// OC would be appended, which your
				// DSA almost certainly wouldn't like).
				for el := 0; el < len(v); el++ {
					if !strInSlice(v[el], m[k]) {
						m[k] = append(m[k], v[el])
					}
				}
			} else {
				// There shouldn't be anything there,
				// so I'm not worried about clobbering.
				m[k] = v
			}
		}
	}

	return m
}

/*
extractLTag extracts the `ldap:"..."` tag sequence and processes it
into string slices, which are returned alongside a success-indicative
boolean value. The `ldap` tag is defined in the go-ldap package, and
is used here to accommodate (a) smooth marshaling into a suitable map
instance for submission to ldap.NewEntry, and (b) smooth unmarshaling
from an ldap.Entry into one of the struct types defined in this package.
*/
func extractLTag(fname string, x any) (tag []string, ok bool) {
	t := typeOf(x)
	if isPtr(x) {
		t = t.Elem()
		if !isStruct(t) {
			return
		}
	}

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if f.Name != fname {
			continue
		}
		if xt, found := f.Tag.Lookup(`ldap`); found {
			tag = split(xt, `,`)
			break
		}
	}

	ok = len(tag) > 0

	return
}

/*
strInSlice returns a boolean value indicative of whether
the specified string (str) was present within the slice
(sl). Matching is conducted using the eq alias, which is
strings.EqualFold.
*/
func strInSlice(str string, sl []string) bool {
	for i := 0; i < len(sl); i++ {
		if eq(str, sl[i]) {
			return true
		}
	}

	return false
}
