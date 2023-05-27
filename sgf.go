package dcxl

/*
sgf.go contains stock SetFunc/GetFunc-qualified functions for use
during composition or interrogation of Registration and/or Registrant
type instances, as well as to offer practical examples for use in the
creation of such CUSTOM function instances by the user.
*/

import "time"

/*
GeneralizedTimeToTime converts one or more string-based generalizedTime
values into a UTC-aligned time.Time instances, which are then added as
slices to an instance of []time.Time (as an interface type), which is
returned alongside an error. This function qualifies for the SetFunc
and GetFunc type signatures.

For more information about functions such as this one, as well as
information on writing your own speciality functions/methods, see
the documentation for the SetFunc and GetFunc closure types.
*/
func GeneralizedTimeToTime(X, R any) (T any, err error) {
	switch tv := X.(type) {
	case string:
		if t, ok := genTimeToTime(tv); ok {
			T = t
			return
		}
	case []string:
		var Tx []time.Time
		for i := 0; i < len(tv); i++ {
			if t, ok := genTimeToTime(tv[i]); ok && !t.IsZero() {
				Tx = append(Tx, t)
			}
		}

		if len(Tx) > 0 {
			T = Tx
			return
		}
	default:
		err = errorf("Unsupported time format (expecting string or []string, got %T)", tv)
		return
	}

	err = errorf("Failed to parse input values, or no values were found")
	return
}

/*
TimeToGeneralizedTime converts an instance of time.Time into a string
value instance of generalizedTime, which is returned alonside an error.
This function qualifies for the SetFunc and GetFunc type signatures.

For more information about functions such as this one, as well as
information on writing your own speciality functions/methods, see
the documentation for the SetFunc and GetFunc closure types.
*/
func TimeToGeneralizedTime(X, R any) (G any, err error) {
	switch tv := X.(type) {
	case time.Time:
		if g, ok := timeToGenTime(tv); ok {
			G = g
			return
		}
	case []time.Time:
		var Gx []string
		for i := 0; i < len(tv); i++ {
			if g, ok := timeToGenTime(tv[i]); ok {
				Gx = append(Gx, g)
			}
		}

		if len(Gx) > 0 {
			G = Gx
			return
		}
	default:
		err = errorf("Unsupported time format (expecting time.Time or []time.Time, got %T)", tv)
		return
	}

	err = errorf("Failed to parse input values, or no values were found")
	return
}

/*
DotNotToDN2D returns a string-based LDAP distinguished name value
(dn) based upon the contents of the input ASN.1 dotNotation value
(X) alongside an error. This function qualifies for the SetFunc
and GetFunc type signature.

This conforms to the two dimensional DN syntax, as described in
draft-coretta-x660-ldap-08, section 3.2. This function will output
a distinguished name value that uses the dotNotation for the RDN AT.
Individual numberForms present within the dotNotation are verified
as non-negative numbers, but are not modified.

For more information about functions such as this one, as well as
information on writing your own speciality functions/methods, see
the documentation for the SetFunc and GetFunc closure types.
*/
func DotNotToDN2D(X, R any) (dn any, err error) {
	r, ok := R.(Registration)
	if !ok {
		err = errorf("Type assertion failed for '%T'", R)
		return
	}

	var duaConf *DUAConfig = r.DUAConfig()
	if duaConf == nil {
		err = errorf("%T is nil", duaConf)
		return
	}

	// We want at least one Registration Base
	// and our model MUST be 3D. Return error
	// value otherwise.
	if len(duaConf.Registrations) == 0 || duaConf.DirectoryModel != TwoDimensional {
		err = errorf("Invalid %T configuration", r)
		return
	}

	// store the original OID here, as we'll
	// just slap it on our composite DN later.
	var O string

	// Make sure input X type is supported, else bail.
	switch tv := X.(type) {
	case string:
		O = tv
	default:
		// This stock function only allows string-based OIDs as
		// input. If you need something more specialized, such
		// as asn1.ObjectIdentifier, write your own SetFunc :)
		err = errorf("Unsupported OID type (%T)", tv)
		return
	}

	// Just scan the slices and verify as
	// a number; no alterations needed.
	var D []string = split(O, `.`) // temporary storage for verification
	for i := 0; i < len(D); i++ {
		if !isNumber(D[i]) {
			err = errorf("Bogus numberForm value '%v' (slice[%d])", D[i-1], i-1)
			return
		}
	}

	// Prepare return value.
	dn = `dotNotation=` + O + `,` + duaConf.Registrations[0]
	return
}

/*
DNtoDotNot2D returns a dotNotation-based ASN.1 Object Identifier
(id) based upon the contents of the input string distinguished
name value (X) alongside an error. This function qualifies for
the SetFunc type signature.

This conforms to the two dimensional DN syntax, as described in
draft-coretta-x660-ldap-08, section 3.2. This function expects
the use of dotNotation in the RDN.

For more information about functions such as this one, as well as
information on writing your own speciality functions/methods, see
the documentation for the SetFunc closure type.
*/
func DNToDotNot2D(X, R any) (id any, err error) {

	var D string
	switch tv := X.(type) {
	case string:
		D = tv
	default:
		// This stock function only allows string-based DNs as
		// input. If you need something more specialized, such
		// as *ldap.DN, write your own SetFunc :)
		err = errorf("Unsupported DN type (%T)", tv)
		return
	}

	var N string
	if idx := idxRune(D, ','); idx != -1 {
		N = replaceAll(lc(D[:idx]), `dotnotation=`, ``)
	}

	// In case the input DN was whacky,
	// or flat-out zero, let's put a
	// stop to this madness now.
	if len(N) == 0 {
		err = errorf("Failed to parse input DN '%v'", D)
		return
	}

	// Just scan the slices and verify as
	// a number; no alterations needed.
	var S []string = split(N, `.`)
	for i := 0; i < len(S); i++ {
		if !isNumber(S[i]) {
			err = errorf("Bogus numberForm value '%v' (slice[%d])", S[i-1], i-1)
			return
		}
	}

	id = N
	return
}

/*
DotNotToDN3D returns a string-based LDAP distinguished name value
(dn) based upon the contents of the input ASN.1 dotNotation value
(X) alongside an error. This function qualifies for the SetFunc
type signature.

This conforms to the three dimensional DN syntax, as described in
draft-coretta-x660-ldap-08, section 3.3. This function will output
relative distinguished name values, each of whom describe specific
numberForm values, using the preferred AT descriptor 'n'.

For more information about functions such as this one, as well as
information on writing your own speciality functions/methods, see
the documentation for the SetFunc closure type.
*/
func DotNotToDN3D(X, R any) (dn any, err error) {
	r, ok := R.(Registration)
	if !ok {
		err = errorf("Type assertion failed for '%T'", R)
		return
	}

	var duaConf *DUAConfig = r.DUAConfig()
	if duaConf == nil {
		err = errorf("%T is nil", duaConf)
		return
	}

	// We want at least one Registration Base
	// and our model MUST be 3D. Return error
	// value otherwise.
	if len(duaConf.Registrations) == 0 || duaConf.DirectoryModel != ThreeDimensional {
		err = errorf("Invalid %T configuration", r)
		return
	}

	var D []string // store a portion of the original dotNotation

	// Make sure input X type is supported, else bail.
	switch tv := X.(type) {
	case string:
		D = split(tv, `.`)
	default:
		// This stock function only allows string-based DNs as
		// input. If you need something more specialized, such
		// as *ldap.DN, write your own SetFunc :)
		err = errorf("Unsupported OID type (%T)", tv)
		return
	}

	var nfs []string

	// reverse iterate over our D slices,
	// verifying each slice as a number
	// and appending the composite RDN value
	// (n=N) to the nfs string slice type.
	for i := len(D); i > 0; i-- {
		if !isNumber(D[i-1]) {
			err = errorf("Bogus numberForm value '%v' (slice[%d])", D[i-1], i-1)
			return
		}
		nfs = append(nfs, `n=`+D[i-1])
	}

	// prepare return value, and just grab
	// the first reg. base ...
	dn = join(nfs, `,`) + `,` + duaConf.Registrations[0]
	return
}

/*
DNtoDotNot3D returns a dotNotation-based ASN.1 Object Identifier
(id) based upon the contents of the input string distinguished
name value (X) alongside an error. This function qualifies for
the SetFunc type signature.

This conforms to the three dimensional DN syntax, as described in
draft-coretta-x660-ldap-08, section 3.3. This function offers
positive support for the RDN AT descriptor 'n' as well as its
more "distinguished" descriptor alias 'numberForm'.

For more information about functions such as this one, as well as
information on writing your own speciality functions/methods, see
the documentation for the SetFunc closure type.
*/
func DNToDotNot3D(X, R any) (id any, err error) {

	// Make sure type R is indeed what we expect.
	// This function really only applies to actual
	// x660RootArc and x660SubArc registrations,
	// each of which qualify for the Registration
	// interface type.
	r, ok := R.(Registration)
	if !ok {
		err = errorf("Type assertion failed for '%T'", R)
		return
	}

	// Grab our DUAConfig instance if defined,
	// else throw an error. This is required
	// because we need to know the reg. base
	// string value, as well asthe directory
	// model in use, so as to use it for value
	// termination during processing.
	var duaConf *DUAConfig = r.DUAConfig()
	if duaConf == nil {
		err = errorf("%T is nil", duaConf)
		return
	}

	// We want at least one Registration Base
	// and our model MUST be 3D. Return error
	// value otherwise.
	if len(duaConf.Registrations) == 0 || duaConf.DirectoryModel != ThreeDimensional {
		err = errorf("Invalid %T configuration", r)
		return
	}

	var D string // store a portion of the original dn
	var L int    // store slice count for alloc

	// Make sure input X type is supported, else bail.
	switch tv := X.(type) {
	case string:
		// Chomp root suffix
		D = replaceAll(tv, duaConf.Registrations[0], ``)

		// Count the number of delimiters remaining
		// in the RDN sequence and use this to note
		// the desired size of our slice type. This
		// is necessary because we'll be "appending"
		// in reverse.
		for i := 0; i < len(D); i++ {
			if D[i] == byte(',') {
				L++
			}
		}

		// One or fewer elements? Bail out.
		if L < 2 {
			err = errorf("Invalid OID '%v' (cannot be shorter than two (2) elements, got %d)", D, L)
			return
		}
	default:
		// This stock function only allows string-based DNs as
		// input. If you need something more specialized, such
		// as *ldap.DN, write your own SetFunc :)
		err = errorf("Unsupported DN type (%T)", tv)
		return
	}

	var nfs []string = make([]string, L, L) // numberForm slice alloc
	var lastidx int                         // breadcrumbs for loops of indefinite length
	var ct int                              // progressive iteration counter for reverse indices

	// Loop through each RDN in the DN, convert to
	// NumberForm. If even one (1) RDN is bogus, then
	// crash the party and return nil. Note that in
	// this loop, we're parsing a "backwards 3D DN"
	// into a forward dotNotation value, hence the
	// fixed slice array and LIFO slice indices.
	for {
		// A return index of negative one (-1) means
		// the party is over, whether we succeeded or
		// not ...
		if idx := idxRune(D[lastidx:], ','); idx != -1 {

			// Grab the current RDN value using
			// start and end indices. Also use
			// this opportunity to fold the RDN
			// case to lower for simplicity.
			//
			// Remember, the desired return val
			// will be a dotNotation OID, so we
			// don't need alpha chars.
			val := lc(D[lastidx : lastidx+idx])

			// Strip out the attributeType and
			// equals (=) symbol, we just want
			// the decimal value.
			nf := replaceAll(val, `n=`, ``)

			// We need to make sure that the above
			// string replacement actually produced
			// a change. If not, try to take steps
			// to mitigate.
			if len(nf) == len(val) {
				// No change in val length? Its possible the
				// user might be using "numberform=N,..." vs.
				// "n=N,..."  so we'll be sure to prepare for
				// that possibility.
				nf = replaceAll(val, `numberform=`, ``)

				// Still no change? We might be done
				// processing, OR the original input
				// DN is garbage. Either way, we bail.
				if len(nf) == len(val) {
					break
				}
			}

			// If what we got was NOT strictly,
			// numerical, then we're boned.
			if !isNumber(nf) {
				err = errorf("Bogus numberForm value '%v' (idx[%d:%d])",
					val, lastidx, lastidx+idx)
				return
			}

			// We seem to have obtained a number,
			// so we'll add it to the next available
			// slice index, but we'll use contrary
			// ordering (Left=Right) during value
			// assignment, as this is the custom
			// for so-called three dimensional DN
			// values per section 3.3 of the ID.
			nfs[len(nfs)-ct-1] = nf

			// make a note of this iteration's index
			// value by saving it as lastidx. Add one
			// (1) to account for comma delimiters.
			lastidx += idx + 1

			// Increment our iteration counter.
			ct++

			// Proceed to next iteration
			continue
		}

		// We're done processing.
		break
	}

	// prepare return value
	id = join(nfs, `.`)
	return
}
