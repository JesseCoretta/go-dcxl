package dcxl

/*
type.go encompasses all types, constants and global variables
defined by this package.
*/

const (
	TwoDimensional   = `1.3.6.1.4.1.56521.101.3.2` // s. 3.2
	ThreeDimensional = `1.3.6.1.4.1.56521.101.3.3` // s. 3.3
)

/*
GetOrSetFunc is a first class (closure) function signature that users
may adopt in order to write custom "setter or getter" functions, thereby
allowing complete control over the creation or interrogation of a value
assigned to Registration or Registrant type instances.

All Set<*> and <*>GetFunc methods extended by Registration or Registrant
type instances allow the GetOrSetFunc type to be passed at runtime.

If no GetOrSetFunc is passed to a Set<*> method, the value is written
as-is (type assertion permitting) with no special processing.

In the context of Set<*> executions, the first input argument will be
the value to be written to the appropriate struct field. The second
input argument will be the (non-nil) POINTER receiver instance that
contains the target field(s) (i.e.: the object to which something is
being written).

A Set<*> function that interacts with a struct field of type []string
can allow append operations (with an individual input value), as well
as so-called "clobber" operations (with a slice ([]) input value) in
which any values already present are overwritten (clobbered).  If an
append operation is needed for multiple values, and clobbering is NOT
desired, the submission must be done in an iterative (looped) manner.
You have been warned.

In the context of Set<*> return values, the first return value will be
the (processed) value to be written. The second return value, an error,
will contain error information in the event of any encountered issues.

In the context of <*>GetFunc executions, the first input argument will
be the struct field value relevant to the executing method. This will
produce the value being "gotten" within functions/methods that conform
to the signature of this type. The second input argument will be the
non-nil Registration or Registrant instance being interrogated, which
may or may not be a POINTER instance.

In the context of <*>GetFunc return values, the first return value will
be the (processed) value being read. It will manifest as an instance of
'any', and will require manual type assertion. The second return value,
an error, will contain error information in the event of any encountered
issues.
*/
type GetOrSetFunc func(any, any) (any, error)

/*
DUAConfig contains important information necessary for effective DUA
configuration. This information could be set manually by the user, or
marshaled via the target DSA's Root DSE entry. For more information, see
draft-coretta-x660-ldap s. 3.5.1 (manual) and s. 3.5.2 (auto).

See also go-ldap/ldap/v3's Entry.Unmarshal method for auto-configuration
of (pointer!) instances of this type.

The DirectoryModel struct field MUST be populated at all times. Only two
valid string values exist for this field. As such, see the TwoDimensional
and ThreeDimensional constants provided by this package.

If neither the Registrations nor the Registrants struct fields are
populated with valid DNs, the instance is bogus.

If only the Registrations struct field is populated with a valid DN,
this indicates that authority information is neither stored nor managed
on the remote DSA.

If only the Registrants struct field is populated with a valid DN, this
indicates that OID registration information is neither stored nor managed
on the remote DSA (which is atypical and does not technically reflect the
intended use and procedures of draft-coretta-x660-ldap).

If the Registrations and Registrants struct fields are both populated with an
identical (and valid) DN, this indicates that both OID registration entries
and authority registrant entries are ONE IN THE SAME (so-called "combined
entries").

If the Registrations and Registrants struct fields are both populated with
valid (but different) DNs, this indicates that both OID registration entries
and authority registrant entries are managed and stored on the remote DSA,
but as so-called "dedicated entries".

Population of the ServiceEmails and/or ServiceURIs struct fields is
entirely optional and at the discretion of the directory owner(s).

A generic (and optional) map[string][]string instance, via the Settings
struct field, is available to enhance or augment client behavior.

OID: 1.3.6.1.4.1.56521.101.2.2.3, defined in s. 2.2.3.
*/
type DUAConfig struct {
	DirectoryModel string   `ldap:"rADirectoryModel"`   // s. 2.1.86, s. 3.2, s. 3.3
	Registrations  []string `ldap:"rARegistrationBase"` // s. 2.1.84
	Registrants    []string `ldap:"rARegistrantBase"`   // s. 2.1.85
	ServiceEmails  []string `ldap:"rAServiceMail"`      // s. 2.1.87
	ServiceURIs    []string `ldap:"rAServiceURIs"`      // s. 2.1.88

	// Not defined in draft-coretta-x660-ldap, but is
	// sensible to include for client optimization.
	Settings map[string][]string
}

/*
RootArc contains information either to be set upon, or derived from, an
LDAP entry that describes one (1) of three (3) possible top-level OID roots:

  - ITU-T (0)
  - ISO (1)
  - Joint-ISO-ITU-T (2)

See also go-ldap/ldap's Entry.Unmarshal method for auto-population of
(pointer!) instances of this type.

When defining entries in the directory that reflect one of the above
root arcs, one of the following DN conventions are recommended without
regard for any specific directory model (dimension):

  - n=[0|1|2],<rARegistrationBase>, or ...
  - numberForm=[0|1|2],<rARegistrationBase>, or ...
  - identifier=[itu-t|iso|joint-iso-itu-t],<rARegistrationBase>, or ...
  - unicodeValue=[ITU-T|ISO|Joint-ISO-ITU-T],<rARegistrationBase>

Though, it is worth mentioning the first example (n=[0|1|2]) is an ideal
choice for 3D models, as this allows seamless conversion with fewer
conditional checks in the programmatic sense. YMMV.

OID: 1.3.6.1.4.1.56521.101.2.2.1, defined in s. 2.2.1.
*/
type RootArc struct {
	R_DN       string   `ldap:"dn"`
	R_N        string   `ldap:"n"`
	R_Desc     string   `ldap:"description"`
	R_ASN1Not  string   `ldap:"asn1Notation"`
	R_Id       string   `ldap:"identifier"`
	R_Created  string   `ldap:"registrationCreated"`
	R_NaNF     string   `ldap:"nameAndNumberForm"`
	R_LeftArc  string   `ldap:"leftArc"`
	R_RightArc string   `ldap:"rightArc"`
	R_Modified []string `ldap:"registrationModified"`
	R_SubArc   []string `ldap:"subArc"`
	R_StdNF    []string `ldap:"stdNameForm"`
	R_IRI      []string `ldap:"iRI"`
	R_UVal     []string `ldap:"unicodeValue"`
	R_AddlId   []string `ldap:"additionalIdentifier"`
	R_Info     []string `ldap:"registrationInformation"`
	R_URI      []string `ldap:"registrationURI"`
	R_FAuthyDN []string `ldap:"firstAuthority"`   // used only for dedicated registrant entries
	R_CAuthyDN []string `ldap:"currentAuthority"` // used only for dedicated registrant entries

	// Fields for COMBINED registration + registrant
	// entries ONLY (e.g.: NOT dedicated). Note there are
	// drawbacks with using COMBINED entries and is a practice
	// that is generally discouraged outside of special
	// corner-cases.
	R_FAuthy *FirstAuthority
	R_CAuthy *CurrentAuthority

	R_DUAConfig *DUAConfig
}

/*
SubArc contains information either to be set upon, or derived from, an
LDAP entry that describes a non-root arc registration.

See also go-ldap/ldap's Entry.Unmarshal method for auto-population of
(pointer!) instances of this type.

OID: 1.3.6.1.4.1.56521.101.2.2.2, defined in s. 2.2.2.
*/
type SubArc struct {
	R_DN         string   `ldap:"dn"`
	R_N          string   `ldap:"n"`
	R_Desc       string   `ldap:"description"`
	R_DotNot     string   `ldap:"dotNotation"`
	R_ASN1Not    string   `ldap:"asn1Notation"`
	R_Id         string   `ldap:"identifier"`
	R_Created    string   `ldap:"registrationCreated"`
	R_Range      string   `ldap:"registrationRange"`
	R_Status     string   `ldap:"registrationStatus"`
	R_LeafNode   string   `ldap:"isLeafNode"`
	R_Frozen     string   `ldap:"isFrozen"`
	R_NaNF       string   `ldap:"nameAndNumberForm"`
	R_SupArc     string   `ldap:"supArc"`
	R_TopArc     string   `ldap:"topArc"`
	R_LeftArc    string   `ldap:"leftArc"`
	R_FirstArc   string   `ldap:"firstArc"`
	R_RightArc   string   `ldap:"rightArc"`
	R_FinalArc   string   `ldap:"finalArc"`
	R_SubArc     []string `ldap:"subArc"`
	R_LongArc    []string `ldap:"longArc"` // only permitted for subArcs of Joint-ISO-ITU-T (2).
	R_Modified   []string `ldap:"registrationModified"`
	R_DiscloseTo []string `ldap:"discloseTo"`
	R_StdNF      []string `ldap:"stdNameForm"`
	R_IRI        []string `ldap:"iRI"`
	R_UVal       []string `ldap:"unicodeValue"` // almost always single-val
	R_AddlId     []string `ldap:"additionalIdentifier"`
	R_Info       []string `ldap:"registrationInformation"`
	R_URI        []string `ldap:"registrationURI"`
	R_FAuthyDN   []string `ldap:"firstAuthority"`   // used only for dedicated registrant entries
	R_CAuthyDN   []string `ldap:"currentAuthority"` // used only for dedicated registrant entries
	R_SAuthyDN   []string `ldap:"sponsor"`          // used only for dedicated registrant entries

	// Fields for COMBINED registration + registrant
	// entries ONLY (e.g.: NOT dedicated). Note there are
	// drawbacks with using COMBINED entries and is a practice
	// that is generally discouraged outside of special
	// corner-cases.
	R_FAuthy *FirstAuthority
	R_CAuthy *CurrentAuthority
	R_SAuthy *Sponsor

	R_DUAConfig *DUAConfig
}

/*
Registration encompasses arc definition instances of the
RootArc and SubArc struct types.
*/
type Registration interface {
	// ObjectClass returns either `x660RootArc` or `x660SubArc`
	// depending on the underlying instance type.
	ObjectClass() string

	// Kind returns `STRUCTURAL` as a convenient means of identifying
	// the nature of the objectClass in terms of entry structure.
	Kind() string

	// DN returns the string distinguished name as derived from the
	// underlying instance, or a zero string if unset.
	DN() string

	// SetDN assigns the string distinguished name value to the
	// receiver's R_DN field.
	SetDN(any, ...GetOrSetFunc) error

	// DNGetFunc returns an interface value alongside an error. A
	// non-nil instance of GetOrSetFunc processes the appropriate
	// underlying struct field into the desired type.
	DNGetFunc(GetOrSetFunc) (any, error)

	// N returns the numberForm of the registration. This is always
	// required without exception.
	N() string

	// SetN assigns the provided string numberForm value to the
	// receiver's R_N field.
	SetN(any, ...GetOrSetFunc) error

	// NGetFunc returns an interface value alongside an error. A non-nil
	// instance of GetOrSetFunc processes the appropriate underlying struct
	// field into the desired type.
	NGetFunc(GetOrSetFunc) (any, error)

	// DotNotation returns the dot-delimited ASN.1 Object Identifier
	// as a string value from the underlying instance, or a zero string
	// if unset. Note that while ThreeDimensional implementations allow
	// the absence of such a value, TwoDimensional implementations will
	// be required to ensure this be non-zero at all times. Please also
	// note that DotNotation will ALWAYS return a zero-string for any
	// registration that is RootArc (x660RootArc), as such registrations
	// do not syntactically qualify for actual "dotted notation", as there
	// is only one digit.
	DotNotation() string

	// SetDotNotation assigns the string-based dotNotation object
	// identifier value to the receiver's R_DotNot field.
	SetDotNotation(any, ...GetOrSetFunc) error

	// DotNotationGetFunc returns an interface value alongside an error. A
	// non-nil instance of GetOrSetFunc processes the appropriate underlying
	// struct field into the desired type.
	DotNotationGetFunc(GetOrSetFunc) (any, error)

	// ASN1Notation returns the string form of an ASN.1 Object Identifier
	// as a sequence of NameAndNumberForms as derived from the underlying
	// instance.
	ASN1Notation() string

	// SetASN1Notation assigns the string-based ASN.1 Notation object
	// identifier value to the receiver's R_ASN1Not field.
	SetASN1Notation(any, ...GetOrSetFunc) error

	// ASN1NotationGetFunc returns an interface value alongside an error. A
	// non-nil instance of GetOrSetFunc processes the appropriate underlying
	// struct field into the desired type.
	ASN1NotationGetFunc(GetOrSetFunc) (any, error)

	// Identifier returns the primary nameForm of the registration, if set,
	// else a zero string is returned.
	Identifier() string

	// SetIdentifier assigns the string identifier (nameForm) value
	// to the receiver's R_Id field.
	SetIdentifier(any, ...GetOrSetFunc) error

	// IdentifierGetFunc returns an interface value alongside an error. A
	// non-nil instance of GetOrSetFunc processes the appropriate underlying
	// struct field into the desired type.
	IdentifierGetFunc(GetOrSetFunc) (any, error)

	// NameAndNumberForm returns either "identifier(Number)" or "number",
	// if set, within the registration else a zero string is returned.
	NameAndNumberForm() string

	// SetNameAndNumberForm assigns the string nameAndNumberForm value
	// to the receiver's R_NaNF field. Though a bare numberForm value
	// is technically legal for this field, N is preferred for that
	// purpose. This method should be used to assign a complete value
	// of nameAndNumberForm (such as "enterprise(1)") to the receiver
	// instead.
	SetNameAndNumberForm(any, ...GetOrSetFunc) error

	// NameAndNumberFormGetFunc returns an interface value alongside an error.
	// A non-nil instance of GetOrSetFunc processes the appropriate underlying
	// struct field into the desired type.
	NameAndNumberFormGetFunc(GetOrSetFunc) (any, error)

	// LeftArc returns the leftArc value assigned to the underlying instance,
	// if set, else a zero string is returned. This should be interpreted
	// as the nearest-left arc registration DN in relation to the current.
	LeftArc() string

	// SetLeftArc assigns a string distinguished name value to the receiver's
	// R_LeftArc field.
	SetLeftArc(any, ...GetOrSetFunc) error

	// LeftArcGetFunc returns an interface value alongside an error. A
	// non-nil instance of GetOrSetFunc processes the appropriate underlying
	// struct field into the desired type.
	LeftArcGetFunc(GetOrSetFunc) (any, error)

	// RightArc returns the rightArc value assigned to the underlying instance,
	// if set, else a zero string is returned. This should be interpreted
	// as the nearest-right arc registration DN in relation to the current.
	RightArc() string

	// SetRightArc assigns a string distinguished name value to the receiver's
	// R_RightArc field.
	SetRightArc(any, ...GetOrSetFunc) error

	// RightArcGetFunc returns an interface value alongside an error. A
	// non-nil instance of GetOrSetFunc processes the appropriate underlying
	// struct field into the desired type.
	RightArcGetFunc(GetOrSetFunc) (any, error)

	// FirstArc returns the farthest-right arc registration DN in relation to
	// the current, if set, else a zero string is returned. In numerical terms,
	// this would be the lowest numberForm (N) among sibling registrations.
	FirstArc() string

	// SetFirstArc assigns a string distinguished name value to the receiver's
	// R_FirstArc field. This will only have any effect if the underlying
	// receiver instance is *dcxl.SubArc.
	SetFirstArc(any, ...GetOrSetFunc) error

	// FirstArcGetFunc returns an interface value alongside an error. A
	// non-nil instance of GetOrSetFunc processes the appropriate underlying
	// struct field into the desired type.
	FirstArcGetFunc(GetOrSetFunc) (any, error)

	// FinalArc returns the farthest-left arc registration DN in relation to
	// the current, if set, else a zero string is returned. In numerical terms,
	// this would be the highest numberForm (N) among sibling registrations.
	FinalArc() string

	// SetFinalArc assigns a string distinguished name value to the receiver's
	// R_FinalArc field. This will only have any effect if the underlying
	// receiver instance is *dcxl.SubArc.
	SetFinalArc(any, ...GetOrSetFunc) error

	// FinalArcGetFunc returns an interface value alongside an error. A
	// non-nil instance of GetOrSetFunc processes the appropriate underlying
	// struct field into the desired type.
	FinalArcGetFunc(GetOrSetFunc) (any, error)

	// TopArc returns the absolute top-most root registration DN in relation to
	// the current if set, else a zero string is returned.
	TopArc() string

	// SetTopArc assigns a string distinguished name value to the receiver's
	// R_TopArc field. This will only have any effect if the underlying
	// receiver instance is *dcxl.SubArc.
	SetTopArc(any, ...GetOrSetFunc) error

	// TopArcGetFunc returns an interface value alongside an error. A
	// non-nil instance of GetOrSetFunc processes the appropriate underlying
	// struct field into the desired type.
	TopArcGetFunc(GetOrSetFunc) (any, error)
	// SupArc returns the immediate parent registration DN in relation to
	// the current, if set, else a zero string is returned. Note this will
	// only apply to SubArc (x660SubArc) registrations.
	SupArc() string

	// SetSupArc assigns a string distinguished name value to the receiver's
	// R_SupArc field. This will only have any effect if the underlying
	// receiver instance is *dcxl.SubArc.
	SetSupArc(any, ...GetOrSetFunc) error

	// SupArcGetFunc returns an interface value alongside an error. A
	// non-nil instance of GetOrSetFunc processes the appropriate underlying
	// struct field into the desired type.
	SupArcGetFunc(GetOrSetFunc) (any, error)

	// Range returns the assigned string form of a ranged allocation
	// declaration if set, else a zero string is returned.
	//
	// A value of -1 means "from N to infinity" (or is an
	// "indefinite range").
	//
	// A value of zero (0) has no meaning and should be interpreted
	// as a lack of any range allocation.
	//
	// A value that is greater than zero (0) and less than N should
	// be considered wholly invalid. N always represents the low end
	// of a range.
	//
	// A value that is greater than both zero (0) and N should be
	// interpreted as a proper range of definite length (e.g.:
	// 55 (n, range start) through 1500 (registrationRange, range end)).
	//
	// Note this only applies to SubArc (x660SubArc) registrations,
	// as root registrations never have range allocations.
	Range() string

	// SetRange assigns the string value, which can be any non-zero
	// positive number OR negative one (-1). In the case of a positive
	// number, the value MUST be greater than the value of the receiver's
	// R_N field.
	SetRange(any, ...GetOrSetFunc) error

	// RangeGetFunc returns an interface value alongside an error. A
	// non-nil instance of GetOrSetFunc processes the appropriate
	// underlying struct field into the desired type.
	RangeGetFunc(GetOrSetFunc) (any, error)

	// LeafNode returns a boolean value that defines whether the
	// registration is considered a leaf-node (in that it can never
	// have children, and that applications SHOULD NOT attempt to
	// enumerate any children below the registration). Note this
	// only applies to SubArc (x660SubArc) registrations, as
	// root registrations can never possibly be leaf-nodes.
	LeafNode() bool

	// SetLeafNode assigns the LDAP string-equivalent of TRUE
	// and FALSE, crafted from the input boolean value, to the
	// receiver's R_LeafNode field.
	SetLeafNode(any, ...GetOrSetFunc) error

	// LeafNodeGetFunc returns an interface value alongside an error. A
	// non-nil instance of GetOrSetFunc processes the appropriate
	// underlying struct field into the desired type.
	LeafNodeGetFunc(GetOrSetFunc) (any, error)

	// Frozen returns a boolean value that defines whether the
	// registration is considered frozen (in that any existing
	// child registrations may be enumerated, but that no more
	// child registrations will be created). Note this will only
	// apply to SubArc (x660SubArc) registrations, as none of
	// the roots can ever be frozen.
	Frozen() bool

	// SetFrozen assigns the LDAP string-equivalent of TRUE
	// and FALSE, crafted from the input boolean value, to the
	// receiver's R_Frozen field.
	SetFrozen(any, ...GetOrSetFunc) error

	// FrozenGetFunc returns an interface value alongside an error. A
	// non-nil instance of GetOrSetFunc processes the appropriate
	// underlying struct field into the desired type.
	FrozenGetFunc(GetOrSetFunc) (any, error)

	// Status returns zero or more string "status" values, as
	// derived from the registration instance.
	Status() string

	// SetStatus assigns the string value to the receiver's
	// R_Status field.
	SetStatus(any, ...GetOrSetFunc) error

	// StatusGetFunc returns an interface value alongside an error. A
	// non-nil instance of GetOrSetFunc processes the appropriate
	// underlying struct field into the desired type.
	StatusGetFunc(GetOrSetFunc) (any, error)

	// Description returns the description assigned to the
	// registration, else a zero string if unset.
	Description() string

	// SetDescription assigns the string value to the
	// receiver's R_Desc field.
	SetDescription(any, ...GetOrSetFunc) error

	// DescriptionGetFunc returns an interface value alongside an error.
	// A non-nil instance of GetOrSetFunc processes the appropriate
	// underlying struct field into the desired type.
	DescriptionGetFunc(GetOrSetFunc) (any, error)

	// SubArc returns zero or more string registration
	// arc DNs, each representing a child registration of the
	// current. Note that in certain portions of the OID spectrum,
	// this can return a great many registrations. Depending on
	// the implementation around draft-coretta-x660-ldap, this
	// could be populated through use of the subArc attribute type,
	// or as a result of a (possibly intensive!) onelevel LDAP
	// search to enumerate and record relevant child entries. For
	// performance reasons, use of the subArc attribute type
	// is strongly recommended in large implementations.
	SubArc() []string

	// SetSubArc appends one or more string distinguished
	// name values to the receiver's R_SubArc field.
	SetSubArc(any, ...GetOrSetFunc) error

	// SubArcGetFunc returns an interface value alongside an error. A
	// non-nil instance of GetOrSetFunc processes the appropriate
	// underlying struct field into the desired type.
	SubArcGetFunc(GetOrSetFunc) (any, error)

	// LongArc returns slices of longArc values if assigned to the
	// registration instance, else a zero string slice is returned.
	// Note that use of a LongArc is only appropriate for SubArc
	// (x660SubArc) registrations that fall under the Joint-ISO-ITU-T
	// root, and should be considered invalid anywhere else.
	LongArc() []string

	// SetLongArc appends one or more string long arc values to
	// the receiver's R_LongArc field.
	SetLongArc(any, ...GetOrSetFunc) error

	// LongArcGetFunc returns an interface value alongside an error. A
	// non-nil instance of GetOrSetFunc processes the appropriate
	// underlying struct field into the desired type.
	LongArcGetFunc(GetOrSetFunc) (any, error)

	// Create time returns the registrationCreated Generalized Time value
	CreateTime() string

	// SetCreateTime assigns the string generalizedTime value
	// value to the receiver's R_Created field.
	SetCreateTime(any, ...GetOrSetFunc) error

	// ModifyTime returns zero or more string generalizedTime instances
	// each representing a previous point in time at which the underlying
	// registration instance was modified.
	ModifyTime() []string

	// SetModifyTime appends one or more string generalizedTime value
	// to the receiver.
	SetModifyTime(any, ...GetOrSetFunc) error

	// DiscloseTo returns zero or more string DN values, each
	// representative of a discloseTo value. Note this will only
	// apply to SubArc (x660SubArc) registrations.
	DiscloseTo() []string

	// SetDiscloseTo appends one or more string distinguished
	// name values to the receiver R_DiscloseTo field.
	SetDiscloseTo(any, ...GetOrSetFunc) error

	// DiscloseToGetFunc returns an interface value alongside an
	// error. A non-nil instance of GetOrSetFunc processes the
	// appropriate underlying struct field into the desired type.
	DiscloseToGetFunc(GetOrSetFunc) (any, error)

	// StdNameForm returns zero or more slices of standardized
	// name form values assigned to the registration instance.
	StdNameForm() []string

	// SetStdNameForm appends one or more string standardized
	// name form values to the receiver's R_StdNF field.
	SetStdNameForm(any, ...GetOrSetFunc) error

	// StdNameFormGetFunc returns an interface value alongside an
	// an error. A non-nil instance of GetOrSetFunc processes the
	// appropriate underlying struct field into the desired type.
	StdNameFormGetFunc(GetOrSetFunc) (any, error)

	// IRI returns zero or more internationalized resource
	// identifier string values assigned to the registration
	// instance.
	IRI() []string

	// SetIRI appends one or more string internationalized
	// resource identifier values to the receiver's R_IRI
	// field.
	SetIRI(any, ...GetOrSetFunc) error

	// IRIGetFunc returns an interface value alongside an error. A
	// non-nil instance of GetOrSetFunc processes the appropriate
	// underlying struct field into the desired type.
	IRIGetFunc(GetOrSetFunc) (any, error)

	// UnicodeValue returns zero or more unicodeValue slice
	// instances as string values. Note that in almost every
	// case, this should be single-valued, however there have
	// been a few exceptions to this rule over the years.
	UnicodeValue() []string

	// SetUnicodeValue appends one or more string values
	// to the receiver's R_UVal field.
	SetUnicodeValue(any, ...GetOrSetFunc) error

	// UnicodeValueGetFunc returns an interface value alongside
	// an error. A non-nil instance of GetOrSetFunc processes the
	// appropriate underlying struct field into the desired type.
	UnicodeValueGetFunc(GetOrSetFunc) (any, error)

	// AdditionalIdentifier returns zero or more so-called
	// alternative identifiers assigned to the registration
	// instance.

	AdditionalIdentifier() []string

	// SetAdditionalIdentifier appends one or more string
	// values to the receiver's R_AddlId field.
	SetAdditionalIdentifier(any, ...GetOrSetFunc) error

	// AdditionalIdentifierGetFunc returns an interface value
	// alongside an error. A non-nil instance of GetOrSetFunc
	// processes the appropriate underlying struct field into
	// the desired type.
	AdditionalIdentifierGetFunc(GetOrSetFunc) (any, error)

	// Info returns zero or more string values, each representing
	// a single informational text block describing the nature and
	// status of the registration instance.
	Info() []string

	// SetInfo appends one or more string values to the receiver's
	// R_Info field.
	SetInfo(any, ...GetOrSetFunc) error

	// InfoGetFunc returns an interface value alongside an error. A non-nil
	// instance of GetOrSetFunc processes the appropriate underlying struct
	// field into the desired type.
	InfoGetFunc(GetOrSetFunc) (any, error)

	// URI returns zero or more uniform resource identifier
	// values assigned to the registration instance. Generally
	// these would be web pages that the user may be directed
	// to for some reason.
	URI() []string

	// SetURI appends one or more string values to the receiver's
	// R_URI field.
	SetURI(any, ...GetOrSetFunc) error

	// URIGetFunc returns an interface value alongside an error. A non-nil
	// instance of GetOrSetFunc processes the appropriate underlying struct
	// field into the desired type.
	URIGetFunc(GetOrSetFunc) (any, error)

	// FirstAuthority returns zero or more LDAP string distinguished
	// name values assigned to the registration instance, each of
	// which represent a so-called First (Or Previous) registration
	// authority responsible for the current registration. This only
	// applies to so-called "dedicated" or "non-combined" registrant
	// entries.
	FirstAuthority() []string

	// CombinedFirstAuthority returns the underlying instance of
	// *FirstAuthority assigned to the receiver when Combined Entries
	// are in use. This is generally not recommended aside from special
	// corner cases.
	CombinedFirstAuthority() *FirstAuthority

	// SetFirstAuthority appends one or more string distinguished
	// name values to the receiver's R_Firsts field.
	SetFirstAuthority(any, ...GetOrSetFunc) error

	// FirstAuthorityGetFunc returns an interface value alongside an
	// error. A non-nil instance of GetOrSetFunc processes the appropriate
	// underlying struct field into the desired type.
	FirstAuthorityGetFunc(GetOrSetFunc) (any, error)

	// SetCombinedFirstAuthority assigns a non-nil instance of
	// *FirstAuthority to the receiver. This only applies to
	// so-called "Combined Entries", which are generally not
	// recommended.
	SetCombinedFirstAuthority(*FirstAuthority)

	// CurrentAuthority returns zero or more LDAP string distinguished
	// name values assigned to the registration instance, each of which
	// represent a current registration authorityresponsible for the
	// current registration. This only applies to so-called "dedicated"
	// or "non-combined" registrant entries.
	CurrentAuthority() []string

	// CombinedCurrentAuthority returns the underlying instance of
	// *CurrentAuthority assigned to the receiver when Combined Entries
	// are in use. This is generally not recommended aside from special
	// corner cases.
	CombinedCurrentAuthority() *CurrentAuthority

	// SetCurrentAuthority appends one or more string distinguished
	// name values to the receiver's R_Currents field.
	SetCurrentAuthority(any, ...GetOrSetFunc) error

	// CurrentAuthorityGetFunc returns an interface value alongside
	// an error. A non-nil instance of GetOrSetFunc processes the
	// appropriate underlying struct field into the desired type.
	CurrentAuthorityGetFunc(GetOrSetFunc) (any, error)

	// SetCombinedCurrentAuthority assigns a non-nil instance of
	// *CurrentAuthority to the receiver. This only applies to
	// so-called "Combined Entries", which are generally not
	// recommended.
	SetCombinedCurrentAuthority(*CurrentAuthority)

	// Sponsor returns zero or more LDAP string distinguished
	// name values assigned to the registration instance, each
	// of which represent a sponsoring entity for the current
	// registration. This only applies to "dedicated" or "non-
	// combined" registrant entries, and only for dcxl.SubArc
	// (or x660SubArc) registrations.
	Sponsor() []string

	// CombinedSponsor returns the underlying instance of *Sponsor
	// assigned to the receiver when Combined Entries are in use.
	// This is generally not recommended aside from special corner
	// cases.
	CombinedSponsor() *Sponsor

	// SetSponsor appends one or more string distinguished
	// name values to the receiver's R_Sponsors field. This will
	// only have an effect if the underlying receiver instance is
	// *dcxl.SubArc.
	SetSponsor(any, ...GetOrSetFunc) error

	// SponsorGetFunc returns an interface value alongside an error.
	// A non-nil instance of GetOrSetFunc processes the appropriate
	// underlying struct field into the desired type.
	SponsorGetFunc(GetOrSetFunc) (any, error)

	// SetCombinedSponsor assigns a non-nil instance of *Sponsor
	// to the receiver. This only applies to so-called "Combined
	// Entries" with an underlying instance of *SubArc. Use of
	// Combined Entries is generally not recommended. This has no
	// effect for underlying instances of *RootArc.
	SetCombinedSponsor(*Sponsor)

	// Unmarshal, once it executes Valid, will unmarshal the
	// underlying registration instance into an instance of
	// map[string][]string, which can be fed to ldap.NewEntry.
	Unmarshal() map[string][]string

	// DUAConfig returns an instance of *DUAConfig assigned
	// to the receiver, else nil if unset.
	DUAConfig() *DUAConfig

	// SetDUAConfig assigns an instance of *DUAConfig to
	// the receiver.
	SetDUAConfig(*DUAConfig)
}

/*
Registrations contains slices of Registration interface types, each of
which will be one of RootArc or SubArc struct types.
*/
type Registrations []Registration

/*
Registrant encompasses authority definition instances of the
FirstAuthority, CurrentAuthority and Sponsor struct types.
*/
type Registrant interface {
	// ObjectClass always returns `x660Registrant`
	// as a convenient means of identifying the
	// objectClass value in an entry.
	ObjectClass() string

	// Kind always returns `AUXILIARY` as a means of
	// identifying the kind of objectClass.
	Kind() string

	// Type returns one of `First`, `Current` or `Sponsor`
	// as a convenient means of identifying the nature
	// of a particular Registrant (x660Registrant) instance.
	Type() string

	// DN returns the string distinguished name as derived from the
	// underlying instance, or a zero string if unset.
	DN() string

	// SetDN assigns the string distinguished name value to the
	// receiver's R_DN field.
	SetDN(any, ...GetOrSetFunc) error

	// DNGetFunc returns an interface value alongside an error. A
	// non-nil instance of GetOrSetFunc processes the appropriate
	// underlying struct field into the desired type.
	DNGetFunc(GetOrSetFunc) (any, error)

	// RegistrantID returns the registrantID assigned to the registrant
	// instance, or a zero string if unset.
	RegistrantID() string

	// SetRegistrantID assigns the registrant ID string value to
	// the receiver's R_Id field
	SetRegistrantID(any, ...GetOrSetFunc) error

	// RegistrantIDGetFunc returns an interface value alongside an
	// error. A non-nil instance of GetOrSetFunc processes the
	// appropriate underlying struct field into the desired type.
	RegistrantIDGetFunc(GetOrSetFunc) (any, error)

	// CN returns the common name assigned to the registrant
	// instance, or a zero string if unset.
	CN() string

	// SetCN assigns the string value to the receiver's R_CN field.
	SetCN(any, ...GetOrSetFunc) error

	// CNGetFunc returns an interface value alongside an error. A
	// non-nil instance of GetOrSetFunc processes the appropriate
	// underlying struct field into the desired type.
	CNGetFunc(GetOrSetFunc) (any, error)

	// L returns the l (locality) assigned to the registrant
	// instance, or a zero string if unset.
	L() string

	// SetL assigns the string value to the receiver's R_L field.
	SetL(any, ...GetOrSetFunc) error

	// LGetFunc returns an interface value alongside an error. A
	// non-nil instance of GetOrSetFunc processes the appropriate
	// underlying struct field into the desired type.
	LGetFunc(GetOrSetFunc) (any, error)

	// O returns the organization name assigned to the registrant
	// instance, or a zero string if unset.
	O() string

	// SetO assigns the string value to the receiver's R_O field.
	SetO(any, ...GetOrSetFunc) error

	// OGetFunc returns an interface value alongside an error. A
	// non-nil instance of GetOrSetFunc processes the appropriate
	// underlying struct field into the desired type.
	OGetFunc(GetOrSetFunc) (any, error)

	// C returns the [2 character] country code assigned to the
	// registrant instance, or a zero string if unset.
	C() string

	// SetC assigns the string value to the receiver's R_C field.
	SetC(any, ...GetOrSetFunc) error

	// CGetFunc returns an interface value alongside an error. A
	// non-nil instance of GetOrSetFunc processes the appropriate
	// underlying struct field into the desired type.
	CGetFunc(GetOrSetFunc) (any, error)

	// CO returns the full country name assigned to the registrant
	// instance, or a zero string if unset.
	CO() string

	// SetCO assigns the string value to the receiver's R_CO field.
	SetCO(any, ...GetOrSetFunc) error

	// COGetFunc returns an interface value alongside an error. A
	// non-nil instance of GetOrSetFunc processes the appropriate
	// underlying struct field into the desired type.
	COGetFunc(GetOrSetFunc) (any, error)

	// ST returns the full state or province name assigned to the
	// registrant instance, or a zero string if unset.
	ST() string

	// SetST assigns the string value to the receiver's R_ST field.
	SetST(any, ...GetOrSetFunc) error

	// STGetFunc returns an interface value alongside an error. A
	// non-nil instance of GetOrSetFunc processes the appropriate
	// underlying struct field into the desired type.
	STGetFunc(GetOrSetFunc) (any, error)

	// Tel returns the telephone number assigned to the registrant
	// instance, or a zero string if unset.
	Tel() string

	// SetTel assigns the string value to the receiver's R_Tel field.
	SetTel(any, ...GetOrSetFunc) error

	// TelGetFunc returns an interface value alongside an error. A
	// non-nil instance of GetOrSetFunc processes the appropriate
	// underlying struct field into the desired type.
	TelGetFunc(GetOrSetFunc) (any, error)

	// Fax returns the facsimile telephone number assigned to the
	// registrant instance, or a zero string if unset.
	Fax() string

	// SetFax assigns the string value to the receiver's R_Fax field.
	SetFax(any, ...GetOrSetFunc) error

	// FaxGetFunc returns an interface value alongside an error. A
	// non-nil instance of GetOrSetFunc processes the appropriate
	// underlying struct field into the desired type.
	FaxGetFunc(GetOrSetFunc) (any, error)

	// Title returns the acknowledged title of the individual or
	// institution assigned to the registrant instance, or a
	// zero string if unset.
	Title() string

	// SetTitle assigns the string value to the receiver's R_Title
	// field.
	SetTitle(any, ...GetOrSetFunc) error

	// TitleGetFunc returns an interface value alongside an error.
	// A non-nil instance of GetOrSetFunc processes the appropriate
	// underlying struct field into the desired type.
	TitleGetFunc(GetOrSetFunc) (any, error)

	// Email returns the email address of the individual or
	// institution assigned to the registrant instance, or a
	// zero string if unset.
	Email() string

	// SetEmail assigns the string value to the receiver's R_Email
	// field.
	SetEmail(any, ...GetOrSetFunc) error

	// EmailGetFunc returns an interface value alongside an error.
	// A non-nil instance of GetOrSetFunc processes the appropriate
	// underlying struct field into the desired type.
	EmailGetFunc(GetOrSetFunc) (any, error)

	// POBox returns the box office box number assigned to the
	// registrant instance, or a zero string if unset.
	POBox() string

	// SetPOBox assigns the string value to the receiver's R_POBox
	// field.
	SetPOBox(any, ...GetOrSetFunc) error

	// POBoxGetFunc returns an interface value alongside an error. A
	// non-nil instance of GetOrSetFunc processes the appropriate
	// underlying struct field into the desired type.
	POBoxGetFunc(GetOrSetFunc) (any, error)

	// PostalAddress returns the postal address number assigned
	// to the registrant instance, or a zero string if unset.
	PostalAddress() string

	// SetPostalAddress assigns the string value to the receiver's
	// R_PAddr field.
	SetPostalAddress(any, ...GetOrSetFunc) error

	// PostalAddressGetFunc returns an interface value alongside an
	// error. A non-nil instance of GetOrSetFunc processes the
	// appropriate underlying struct field into the desired type.
	PostalAddressGetFunc(GetOrSetFunc) (any, error)

	// PostalCode returns the postal code assigned to the
	// registrant instance, or a zero string if unset.
	PostalCode() string

	// SetPostalCode assigns the string value to the receiver's
	// R_PCode field.
	SetPostalCode(any, ...GetOrSetFunc) error

	// PostalCodeGetFunc returns an interface value alongside an error.
	// A non-nil instance of GetOrSetFunc processes the appropriate
	// underlying struct field into the desired type.
	PostalCodeGetFunc(GetOrSetFunc) (any, error)

	// Mobile returns the mobile telephone number assigned to the
	// registrant instance, or a zero string if unset.
	Mobile() string

	// SetMobile assigns the string value to the receiver's R_Mobile field.
	SetMobile(any, ...GetOrSetFunc) error

	// MobileGetFunc returns an interface value alongside an error. A
	// non-nil instance of GetOrSetFunc processes the appropriate
	// underlying struct field into the desired type.
	MobileGetFunc(GetOrSetFunc) (any, error)

	// Street returns the street address assigned to the registrant
	// instance, or a zero string if unset.
	Street() string

	// SetStreet assigns the string value to the receiver's R_Street field.
	SetStreet(any, ...GetOrSetFunc) error

	// StreetGetFunc returns an interface value alongside an error. A
	// non-nil instance of GetOrSetFunc processes the appropriate
	// underlying struct field into the desired type.
	StreetGetFunc(GetOrSetFunc) (any, error)

	// URI returns zero or more uniform resource identifier values
	// assigned to the registrant instance.
	URI() []string

	// SetURI assigns the string value to the receiver's R_URI field.
	SetURI(any, ...GetOrSetFunc) error

	// URIGetFunc returns an interface value alongside an error. A
	// non-nil instance of GetOrSetFunc processes the appropriate
	// underlying struct field into the desired type.
	URIGetFunc(GetOrSetFunc) (any, error)

	// StartTime returns a string generalizedTime value
	// indicative of the date and time at which the registrant
	// is known to have begun its authoritative responsibilities
	// upon one or more registrations.
	StartTime() string

	// SetStartTime assigns a string-based generalizedTime OR
	// to the receiver's R_StartTime field.
	SetStartTime(any, ...GetOrSetFunc) error

	// EndTime returns a string generalizedTime value indicative
	// of the date and time at which the registrant is known to
	// have ceased its authoritative responsibilities upon one or
	// more registrations.
	//
	// Note this only applies to FirstAuthority and Sponsor
	// registrants, as CurrentAuthority instances ostensibly
	// have no "end date", as they should still be current.
	EndTime() string

	// SetEndTime assigns a string-based generalizedTime to
	// the receiver's R_EndTime field.
	SetEndTime(any, ...GetOrSetFunc) error

	// Unmarshal, once it executes Valid, will unmarshal the
	// underlying registrant instance into an instance of
	// map[string][]string, which can be fed to ldap.NewEntry.
	Unmarshal() map[string][]string

	// DUAConfig returns an instance of *DUAConfig assigned
	// to the receiver, else nil if unset.
	DUAConfig() *DUAConfig

	// SetDUAConfig assigns an instance of *DUAConfig to
	// the receiver.
	SetDUAConfig(*DUAConfig)
}

/*
Registrants contains slices of Registrant interface types, each of
which will be one of FirstAuthority, CurrentAuthority or Sponsor
struct types.
*/
type Registrants []Registrant

/*
FirstAuthority describes an initial or previous registration
authority.
*/
type FirstAuthority struct {
	R_DN        string   `ldap:"dn"`
	R_Id        string   `ldap:"registrantID"`
	R_L         string   `ldap:"firstAuthorityLocality"`
	R_O         string   `ldap:"firstAuthorityOrg"`
	R_C         string   `ldap:"firstAuthorityCountryCode"`
	R_CO        string   `ldap:"firstAuthorityCountryName"`
	R_ST        string   `ldap:"firstAuthorityState"`
	R_CN        string   `ldap:"firstAuthorityCommonName"`
	R_Tel       string   `ldap:"firstAuthorityTelephone"`
	R_Fax       string   `ldap:"firstAuthorityFax"`
	R_Title     string   `ldap:"firstAuthorityTitle"`
	R_Email     string   `ldap:"firstAuthorityEmail"`
	R_POBox     string   `ldap:"firstAuthorityPOBox"`
	R_PCode     string   `ldap:"firstAuthorityPostalCode"`
	R_PAddr     string   `ldap:"firstAuthorityPostalAddress"`
	R_Street    string   `ldap:"firstAuthorityStreet"`
	R_Mobile    string   `ldap:"firstAuthorityMobile"`
	R_StartTime string   `ldap:"firstAuthorityStartTimestamp"`
	R_EndTime   string   `ldap:"firstAuthorityEndTimestamp"`
	R_URI       []string `ldap:"firstAuthorityURI"`
	R_DUAConfig *DUAConfig
}

/*
CurrentAuthority describes an active registration authority.
Note that, unlike FirstAuthority and Sponsor types, this type
DOES NOT possess an EndTime struct field.
*/
type CurrentAuthority struct {
	R_DN        string   `ldap:"dn"`
	R_Id        string   `ldap:"registrantID"`
	R_L         string   `ldap:"currentAuthorityLocality"`
	R_O         string   `ldap:"currentAuthorityOrg"`
	R_C         string   `ldap:"currentAuthorityCountryCode"`
	R_CO        string   `ldap:"currentAuthorityCountryName"`
	R_ST        string   `ldap:"currentAuthorityState"`
	R_CN        string   `ldap:"currentAuthorityCommonName"`
	R_Tel       string   `ldap:"currentAuthorityTelephone"`
	R_Fax       string   `ldap:"currentAuthorityFax"`
	R_Title     string   `ldap:"currentAuthorityTitle"`
	R_Email     string   `ldap:"currentAuthorityEmail"`
	R_POBox     string   `ldap:"currentAuthorityPOBox"`
	R_PCode     string   `ldap:"currentAuthorityPostalCode"`
	R_PAddr     string   `ldap:"currentAuthorityPostalAddress"`
	R_Street    string   `ldap:"currentAuthorityStreet"`
	R_Mobile    string   `ldap:"currentAuthorityMobile"`
	R_StartTime string   `ldap:"currentAuthorityStartTimestamp"`
	R_URI       []string `ldap:"currentAuthorityURI"`
	R_DUAConfig *DUAConfig
}

/*
Sponsor describes a past or present sponsoring authority.
*/
type Sponsor struct {
	R_DN        string   `ldap:"dn"`
	R_Id        string   `ldap:"registrantID"`
	R_L         string   `ldap:"sponsorLocality"`
	R_O         string   `ldap:"sponsorOrg"`
	R_C         string   `ldap:"sponsorCountryCode"`
	R_CO        string   `ldap:"sponsorCountryName"`
	R_ST        string   `ldap:"sponsorState"`
	R_CN        string   `ldap:"sponsorCommonName"`
	R_Tel       string   `ldap:"sponsorTelephone"`
	R_Fax       string   `ldap:"sponsorFax"`
	R_Title     string   `ldap:"sponsorTitle"`
	R_Email     string   `ldap:"sponsorEmail"`
	R_POBox     string   `ldap:"sponsorPOBox"`
	R_PCode     string   `ldap:"sponsorPostalCode"`
	R_PAddr     string   `ldap:"sponsorPostalAddress"`
	R_Street    string   `ldap:"sponsorStreet"`
	R_Mobile    string   `ldap:"sponsorMobile"`
	R_StartTime string   `ldap:"sponsorStartTimestamp"`
	R_EndTime   string   `ldap:"sponsorEndTimestamp"`
	R_URI       []string `ldap:"sponsorURI"`
	R_DUAConfig *DUAConfig
}
