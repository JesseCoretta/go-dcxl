package dcxl

/*
DN returns the string-based LDAP Distinguished Name
value, or a zero string if unset.
*/
func (r RootArc) DN() string {
	return r.R_DN
}

/*
DNGetFunc executes the GetFunc instance and returns its own
return values. The 'any' value will require type assertion in
order to be accessed reliably. An error is returned if issues
arise.
*/
func (r RootArc) DNGetFunc(getfunc GetFunc) (any, error) {
	if getfunc == nil {
		return nil, errorf("Nil %T provided; aborting", getfunc)
	}

	return getfunc(r.R_DN, r)
}

/*
SetDN assigns the string value as to the receiver's R_DN field.
*/
func (r *RootArc) SetDN(X any, setfunc ...SetFunc) error {
	if len(setfunc) == 0 {
		if assert, ok := X.(string); ok {
			r.R_DN = assert
			return nil
		}
		return errorf("Unsupported DN type %T provided without SetFunc instance", X)
	}

	v, err := setfunc[0](X, r)
	if err != nil {
		return err
	}

	return r.SetDN(v)
}

/*
DN returns the string-based LDAP Distinguished Name
value, or a zero string if unset.
*/
func (r SubArc) DN() string {
	return r.R_DN
}

/*
DNGetFunc executes the GetFunc instance and returns its own
return values. The 'any' value will require type assertion in
order to be accessed reliably. An error is returned if issues
arise.
*/
func (r SubArc) DNGetFunc(getfunc GetFunc) (any, error) {
	if getfunc == nil {
		return nil, errorf("Nil %T provided; aborting", getfunc)
	}

	return getfunc(r.R_DN, r)
}

/*
SetDN assigns the string value as to the receiver's R_DN field.
*/
func (r *SubArc) SetDN(X any, setfunc ...SetFunc) error {
	if len(setfunc) == 0 {
		if assert, ok := X.(string); ok {
			r.R_DN = assert
			return nil
		}
		return errorf("Unsupported DN type %T provided without SetFunc instance", X)
	}

	v, err := setfunc[0](X, r)
	if err != nil {
		return err
	}

	return r.SetDN(v)
}

/*
Description returns the string description assigned to the
registration, else a zero string if unset.
*/
func (r RootArc) Description() string {
	return r.R_Desc
}

/*
DescriptionGetFunc executes the GetFunc instance and returns its own
return values. The 'any' value will require type assertion in order
to be accessed reliably.
*/
func (r RootArc) DescriptionGetFunc(getfunc GetFunc) (any, error) {
	if getfunc == nil {
		return nil, errorf("Nil %T provided; aborting", getfunc)
	}

	return getfunc(r.R_Desc, r)
}

/*
SetDescription assigns the string value as to the receiver's R_Desc field.
*/
func (r *RootArc) SetDescription(X any, setfunc ...SetFunc) error {
	if len(setfunc) == 0 {
		if assert, ok := X.(string); ok {
			r.R_Desc = assert
			return nil
		}
		return errorf("Unsupported Description type %T provided without SetFunc instance", X)
	}

	v, err := setfunc[0](X, r)
	if err != nil {
		return err
	}

	return r.SetDescription(v)
}

/*
Description returns the string description assigned to the
registration, else a zero string if unset.
*/
func (r SubArc) Description() string {
	return r.R_Desc
}

/*
DescriptionGetFunc executes the GetFunc instance and returns its own
return values. The 'any' value will require type assertion in order
to be accessed reliably.
*/
func (r SubArc) DescriptionGetFunc(getfunc GetFunc) (any, error) {
	if getfunc == nil {
		return nil, errorf("Nil %T provided; aborting", getfunc)
	}

	return getfunc(r.R_Desc, r)
}

/*
SetDescription assigns the string value as to the receiver's R_Desc field.
*/
func (r *SubArc) SetDescription(X any, setfunc ...SetFunc) error {
	if len(setfunc) == 0 {
		if assert, ok := X.(string); ok {
			r.R_Desc = assert
			return nil
		}
		return errorf("Unsupported Description type %T provided without SetFunc instance", X)
	}

	v, err := setfunc[0](X, r)
	if err != nil {
		return err
	}

	return r.SetDescription(v)
}

/*
N returns the number form, or primary identifier, of
the receiver. If unset, a zero string is returned.
*/
func (r RootArc) N() string {
	return r.R_N
}

/*
NGetFunc executes the GetFunc instance and returns its own return
values. The 'any'
*/
func (r RootArc) NGetFunc(getfunc GetFunc) (any, error) {
	if getfunc == nil {
		return nil, errorf("Nil %T provided; aborting", getfunc)
	}
	return getfunc(r.R_N, r)
}

/*
SetN assigns the string value to the receiver's R_N field.
*/
func (r *RootArc) SetN(X any, setfunc ...SetFunc) error {
	if len(setfunc) == 0 {
		if assert, ok := X.(string); ok && isNumber(assert) {
			r.R_N = assert
			return nil
		}
		return errorf("Unsupported N type %T provided without SetFunc instance", X)
	}

	v, err := setfunc[0](X, r)
	if err != nil {
		return err
	}

	return r.SetN(v)
}

/*
N() returns the number form, or primary identifier, of
the receiver. If unset, a zero string is returned.
*/
func (r SubArc) N() string {
	return r.R_N
}

/*
NGetFunc executes the GetFunc instance and returns its own return
values. The 'any'
*/
func (r SubArc) NGetFunc(getfunc GetFunc) (any, error) {
	if getfunc == nil {
		return nil, errorf("Nil %T provided; aborting", getfunc)
	}
	return getfunc(r.R_N, r)
}

/*
SetN assigns the string value to the receiver's R_N field.
*/
func (r *SubArc) SetN(X any, setfunc ...SetFunc) error {
	if len(setfunc) == 0 {
		if assert, ok := X.(string); ok && isNumber(assert) {
			r.R_N = assert
			return nil
		}
		return errorf("Unsupported N type %T provided without SetFunc instance", X)
	}

	v, err := setfunc[0](X, r)
	if err != nil {
		return err
	}

	return r.SetN(v)
}

/*
DotNotation returns a zero string, as the dotNotation
attribute type is not applicable to entries of objectClass
'x660RootArc.
*/
func (r RootArc) DotNotation() string { return `` }

func (r RootArc) DotNotationGetFunc(getfunc GetFunc) (any, error) {
	return nil, errorf("DotNotation does not apply to %T instances", r)
}

/*
SetDotNotation does not perform any useful task, as root
registrations do not have dotNotation values within LDAP
due to syntax criteria. This method exists to satisfy Go
interface requirements.
*/
func (r *RootArc) SetDotNotation(o any, setfunc ...SetFunc) error {
	return errorf("DotNotation not applicable to %T", r)
}

/*
DotNotation returns the string dotNotation value
assigned to the receiver, or a zero string if unset.
*/
func (r SubArc) DotNotation() string {
	return r.R_DotNot
}

func (r SubArc) DotNotationGetFunc(getfunc GetFunc) (any, error) {
	if getfunc == nil {
		return nil, errorf("Nil %T provided; aborting", getfunc)
	}
	return getfunc(r.R_DotNot, r)
}

/*
SetDotNotation assigns the string value to the receiver's R_DotNot field.
*/
func (r *SubArc) SetDotNotation(X any, setfunc ...SetFunc) error {
	if len(setfunc) == 0 {
		if assert, ok := X.(string); ok {
			r.R_DotNot = assert
			return nil
		}
		return errorf("Unsupported N type %T provided without SetFunc instance", X)
	}

	v, err := setfunc[0](X, r)
	if err != nil {
		return err
	}

	return r.SetDotNotation(v)
}

/*
ASN1Notation returns the string asn1Notation value
assigned to the receiver, or a zero string if unset.
*/
func (r RootArc) ASN1Notation() string {
	return r.R_ASN1Not
}

/*
SetASN1Notation assigns the string value to the receiver's
R_ASN1Not field.
*/
func (r *RootArc) SetASN1Notation(X any, setfunc ...SetFunc) error {
	if len(setfunc) == 0 {
		if assert, ok := X.(string); ok {
			r.R_ASN1Not = assert
			return nil
		}
		return errorf("Unsupported ASN.1 Notation type %T provided without SetFunc instance", X)
	}

	v, err := setfunc[0](X, r)
	if err != nil {
		return err
	}

	return r.SetASN1Notation(v)
}

func (r RootArc) ASN1NotationGetFunc(getfunc GetFunc) (any, error) {
	if getfunc == nil {
		return nil, errorf("Nil %T provided; aborting", getfunc)
	}

	return getfunc(r.R_ASN1Not, r)
}

/*
ASN1Notation returns the string asn1Notation value
assigned to the receiver, or a zero string if unset.
*/
func (r SubArc) ASN1Notation() string {
	return r.R_ASN1Not
}

/*
SetASN1Notation assigns the string value to the receiver's
R_ASN1Not field.
*/
func (r *SubArc) SetASN1Notation(X any, setfunc ...SetFunc) error {
	if len(setfunc) == 0 {
		if assert, ok := X.(string); ok {
			r.R_ASN1Not = assert
			return nil
		}
		return errorf("Unsupported ASN.1 Notation type %T provided without SetFunc instance", X)
	}

	v, err := setfunc[0](X, r)
	if err != nil {
		return err
	}

	return r.SetASN1Notation(v)
}

func (r SubArc) ASN1NotationGetFunc(getfunc GetFunc) (any, error) {
	if getfunc == nil {
		return nil, errorf("Nil %T provided; aborting", getfunc)
	}

	return getfunc(r.R_ASN1Not, r)
}

/*
NameAndNumberForm returns the string nameAndNumberForm
value assigned to the receiver, or a zero string if unset.
*/
func (r RootArc) NameAndNumberForm() string {
	return r.R_NaNF
}

/*
SetNameAndNumberForm assigns the string value to the receiver's
R_NaNF field.
*/
func (r *RootArc) SetNameAndNumberForm(X any, setfunc ...SetFunc) error {
	if len(setfunc) == 0 {
		if assert, ok := X.(string); ok {
			r.R_NaNF = assert
			return nil
		}
		return errorf("Unsupported NameAndNumberForm type %T provided without SetFunc instance", X)
	}

	v, err := setfunc[0](X, r)
	if err != nil {
		return err
	}

	return r.SetNameAndNumberForm(v)
}

func (r RootArc) NameAndNumberFormGetFunc(getfunc GetFunc) (any, error) {
	if getfunc == nil {
		return nil, errorf("Nil %T provided; aborting", getfunc)
	}

	return getfunc(r.R_NaNF, r)
}

/*
NameAndNumberForm returns the string nameAndNumberForm
value assigned to the receiver, or a zero string if unset.
*/
func (r SubArc) NameAndNumberForm() string {
	return r.R_NaNF
}

/*
SetNameAndNumberForm assigns the string value to the receiver's
R_NaNF field.
*/
func (r *SubArc) SetNameAndNumberForm(X any, setfunc ...SetFunc) error {
	if len(setfunc) == 0 {
		if assert, ok := X.(string); ok {
			r.R_NaNF = assert
			return nil
		}
		return errorf("Unsupported NameAndNumberForm type %T provided without SetFunc instance", X)
	}

	v, err := setfunc[0](X, r)
	if err != nil {
		return err
	}

	return r.SetNameAndNumberForm(v)
}

func (r SubArc) NameAndNumberFormGetFunc(getfunc GetFunc) (any, error) {
	if getfunc == nil {
		return nil, errorf("Nil %T provided; aborting", getfunc)
	}

	return getfunc(r.R_NaNF, r)
}

/*
LeftArc returns the string leftArc DN value assigned to the
receiver, or a zero string if unset.
*/
func (r RootArc) LeftArc() string {
	return r.R_LeftArc
}

/*
SetLeftArc assigns the string value to the receiver's R_LeftArc
field.
*/
func (r *RootArc) SetLeftArc(X any, setfunc ...SetFunc) error {
	if len(setfunc) == 0 {
		if assert, ok := X.(string); ok {
			r.R_LeftArc = assert
			return nil
		}
		return errorf("Unsupported LeftArc type %T provided without SetFunc instance", X)
	}

	v, err := setfunc[0](X, r)
	if err != nil {
		return err
	}

	return r.SetLeftArc(v)
}

func (r RootArc) LeftArcGetFunc(getfunc GetFunc) (any, error) {
	if getfunc == nil {
		return nil, errorf("Nil %T provided; aborting", getfunc)
	}

	return getfunc(r.R_LeftArc, r)
}

/*
LeftArc turns the string leftArc DN value assigned to the
receiver, or a zero string if unset.
*/
func (r SubArc) LeftArc() string {
	return r.R_LeftArc
}

/*
SetLeftArc assigns the string value to the receiver's R_LeftArc
field.
*/
func (r *SubArc) SetLeftArc(X any, setfunc ...SetFunc) error {
	if len(setfunc) == 0 {
		if assert, ok := X.(string); ok {
			r.R_LeftArc = assert
			return nil
		}
		return errorf("Unsupported LeftArc type %T provided without SetFunc instance", X)
	}

	v, err := setfunc[0](X, r)
	if err != nil {
		return err
	}

	return r.SetLeftArc(v)
}

func (r SubArc) LeftArcGetFunc(getfunc GetFunc) (any, error) {
	if getfunc == nil {
		return nil, errorf("Nil %T provided; aborting", getfunc)
	}

	return getfunc(r.R_LeftArc, r)
}

/*
FirstArc returns a zero string, and exists only to satisfy go
interface requirements. The 'firstArc' attributeType is
not applicable to entries of objectClass 'x660RootArc'.
*/
func (r RootArc) FirstArc() string { return `` }

/*
SetFirstArc performs no useful task, and exists only to satisfy
Go interface requirements.
*/
func (r *RootArc) SetFirstArc(X any, setfunc ...SetFunc) error {
	return errorf("FirstArc not applicable to %T", r)
}

func (r RootArc) FirstArcGetFunc(getfunc GetFunc) (any, error) {
	return nil, errorf("FirstArc not applicable to %T", r)
}

/*
FirstArc returns the string firstArc DN value assigned to the
receiver, or a zero string if unset.
*/
func (r SubArc) FirstArc() string {
	return r.R_FirstArc
}

/*
SetFirstArc assigns the string value to the receiver's R_FirstArc
field.
*/
func (r *SubArc) SetFirstArc(X any, setfunc ...SetFunc) error {
	if len(setfunc) == 0 {
		if assert, ok := X.(string); ok {
			r.R_FirstArc = assert
			return nil
		}
		return errorf("Unsupported FirstArc type %T provided without SetFunc instance", X)
	}

	v, err := setfunc[0](X, r)
	if err != nil {
		return err
	}

	return r.SetFirstArc(v)
}

func (r SubArc) FirstArcGetFunc(getfunc GetFunc) (any, error) {
	if getfunc == nil {
		return nil, errorf("Nil %T provided; aborting", getfunc)
	}

	return getfunc(r.R_FirstArc, r)
}

/*
FinalArc returns a zero string, and exists only to satisfy go
interface requirements. The 'finalArc' attributeType is
not applicable to entries of objectClass 'x660RootArc'.
*/
func (r RootArc) FinalArc() string { return `` }

/*
SetFinalArc performs no useful task, and exists only to satisfy
Go interface requirements.
*/
func (r *RootArc) SetFinalArc(X any, setfunc ...SetFunc) error {
	return errorf("FinalArc not applicable to %T", r)
}

func (r RootArc) FinalArcGetFunc(getfunc GetFunc) (any, error) {
	return nil, errorf("FinalArc not applicable to %T", r)
}

/*
FinalArc returns the string finalArc DN value assigned to the
receiver, or a zero string if unset.
*/
func (r SubArc) FinalArc() string {
	return r.R_FinalArc
}

/*
SetFinalArc assigns the string value to the receiver's R_FinalArc
field.
*/
func (r *SubArc) SetFinalArc(X any, setfunc ...SetFunc) error {
	if len(setfunc) == 0 {
		if assert, ok := X.(string); ok {
			r.R_FinalArc = assert
			return nil
		}
		return errorf("Unsupported FinalArc type %T provided without SetFunc instance", X)
	}

	v, err := setfunc[0](X, r)
	if err != nil {
		return err
	}
	return r.SetFinalArc(v)
}

func (r SubArc) FinalArcGetFunc(getfunc GetFunc) (any, error) {
	if getfunc == nil {
		return nil, errorf("Nil %T provided; aborting", getfunc)
	}

	return getfunc(r.R_FinalArc, r)
}

/*
RightArc returns the string rightArc DN value assigned to the
receiver, or a zero string if unset.
*/
func (r RootArc) RightArc() string {
	return r.R_RightArc
}

/*
SetRightArc assigns the string value to the receiver's R_RightArc
field.
*/
func (r *RootArc) SetRightArc(X any, setfunc ...SetFunc) error {
	if len(setfunc) == 0 {
		if assert, ok := X.(string); ok {
			r.R_RightArc = assert
			return nil
		}
		return errorf("Unsupported RightArc type %T provided without SetFunc instance", X)
	}

	v, err := setfunc[0](X, r)
	if err != nil {
		return err
	}

	return r.SetRightArc(v)
}

func (r RootArc) RightArcGetFunc(getfunc GetFunc) (any, error) {
	if getfunc == nil {
		return nil, errorf("Nil %T provided; aborting", getfunc)
	}

	return getfunc(r.R_RightArc, r)
}

/*
RightArc returns the string rightArc DN value assigned to the
receiver, or a zero string if unset.
*/
func (r SubArc) RightArc() string {
	return r.R_RightArc
}

/*
SetRightArc assigns the string value to the receiver's R_RightArc
field.
*/
func (r *SubArc) SetRightArc(X any, setfunc ...SetFunc) error {
	if len(setfunc) == 0 {
		if assert, ok := X.(string); ok {
			r.R_RightArc = assert
			return nil
		}
		return errorf("Unsupported RightArc type %T provided without SetFunc instance", X)
	}

	v, err := setfunc[0](X, r)
	if err != nil {
		return err
	}

	return r.SetRightArc(v)
}

func (r SubArc) RightArcGetFunc(getfunc GetFunc) (any, error) {
	if getfunc == nil {
		return nil, errorf("Nil %T provided; aborting", getfunc)
	}

	return getfunc(r.R_RightArc, r)
}

/*
TopArc returns a zero string, and exists only to satisfy go
interface requirements. RootArc instances ARE the "top"
nodes for all other respective sub arcs, and thus they
themselves have no superior nodes.
*/
func (r RootArc) TopArc() string { return `` }

/*
SetTopArc does not perform any useful task, as a root
arc registration is a "top arc" in of itself.
*/
func (r *RootArc) SetTopArc(X any, setfunc ...SetFunc) error {
	return errorf("TopArc not applicable to %T", r)
}

func (r RootArc) TopArcGetFunc(getfunc GetFunc) (any, error) {
	return nil, errorf("TopArc not applicable to %T", r)
}

/*
TopArc returns the string topArc DN value assigned to the
receiver, or a zero string if unset.
*/
func (r SubArc) TopArc() string {
	return r.R_TopArc
}

/*
SetTopArc assigns the string value to the receiver's R_TopArc
field.
*/
func (r *SubArc) SetTopArc(X any, setfunc ...SetFunc) error {
	if len(setfunc) == 0 {
		if assert, ok := X.(string); ok {
			r.R_TopArc = assert
			return nil
		}
		return errorf("Unsupported TopArc type %T provided without SetFunc instance", X)
	}

	v, err := setfunc[0](X, r)
	if err != nil {
		return err
	}
	return r.SetTopArc(v)
}

func (r SubArc) TopArcGetFunc(getfunc GetFunc) (any, error) {
	if getfunc == nil {
		return nil, errorf("Nil %T provided; aborting", getfunc)
	}

	return getfunc(r.R_TopArc, r)
}

/*
SupArc returns a zero string, and exists only to satisfy
go interface requirements. RootArc have no superior (parent)
nodes.
*/
func (r RootArc) SupArc() string { return `` }

/*
SetSupArc does not perform any useful task, as a root
arc registration cannot have a superior (parent).
*/
func (r *RootArc) SetSupArc(X any, setfunc ...SetFunc) error {
	return errorf("SupArc not applicable to %T", r)
}

func (r RootArc) SupArcGetFunc(getfunc GetFunc) (any, error) {
	return nil, errorf("SupArc not applicable to %T", r)
}

/*
SupArc returns the string supArc DN value assigned to the
receiver, or a zero string if unset.
*/
func (r SubArc) SupArc() string {
	return r.R_SupArc
}

/*
SetSupArc assigns the string value to the receiver's R_SupArc
field.
*/
func (r *SubArc) SetSupArc(X any, setfunc ...SetFunc) error {
	if len(setfunc) == 0 {
		if assert, ok := X.(string); ok {
			r.R_SupArc = assert
			return nil
		}
		return errorf("Unsupported SupArc type %T provided without SetFunc instance", X)
	}

	v, err := setfunc[0](X, r)
	if err != nil {
		return err
	}
	return r.SetSupArc(v)
}

func (r SubArc) SupArcGetFunc(getfunc GetFunc) (any, error) {
	if getfunc == nil {
		return nil, errorf("Nil %T provided; aborting", getfunc)
	}

	return getfunc(r.R_SupArc, r)
}

/*
Identifier returns the string identifier, or name
form, of the receiver. If unset, a zero string is
returned.
*/
func (r RootArc) Identifier() string {
	return r.R_Id
}

/*
SetIdentifier assigns the string value to the receiver's R_Id
field.
*/
func (r *RootArc) SetIdentifier(X any, setfunc ...SetFunc) error {
	if len(setfunc) == 0 {
		if assert, ok := X.(string); ok {
			r.R_Id = assert
			return nil
		}
		return errorf("Unsupported Identifier type %T provided without SetFunc instance", X)
	}

	v, err := setfunc[0](X, r)
	if err != nil {
		return err
	}
	return r.SetIdentifier(v)
}

func (r RootArc) IdentifierGetFunc(getfunc GetFunc) (any, error) {
	if getfunc == nil {
		return nil, errorf("Nil %T provided; aborting", getfunc)
	}

	return getfunc(r.R_Id, r)
}

/*
Identifier returns the string identifier, or name
form, of the receiver. If unset, a zero string is
returned.
*/
func (r SubArc) Identifier() string {
	return r.R_Id
}

/*
SetIdentifier assigns the string value to the receiver's R_Id
field.
*/
func (r *SubArc) SetIdentifier(X any, setfunc ...SetFunc) error {
	if len(setfunc) == 0 {
		if assert, ok := X.(string); ok {
			r.R_Id = assert
			return nil
		}
		return errorf("Unsupported Identifier type %T provided without SetFunc instance", X)
	}

	v, err := setfunc[0](X, r)
	if err != nil {
		return err
	}

	return r.SetIdentifier(v)
}

func (r SubArc) IdentifierGetFunc(getfunc GetFunc) (any, error) {
	if getfunc == nil {
		return nil, errorf("Nil %T provided; aborting", getfunc)
	}

	return getfunc(r.R_Id, r)
}

/*
AdditionalIdentifier returns the string slice values
assigned to the receiver, each defining an alternative
identifier by which the registration is known.
*/
func (r RootArc) AdditionalIdentifier() []string {
	return r.R_AddlId
}

/*
SetAdditionalIdentifier appends one or more string slice values
to the receiver's R_AddlId field.
*/
func (r *RootArc) SetAdditionalIdentifier(X any, setfunc ...SetFunc) error {
	if len(setfunc) == 0 {
		if assert, ok := X.(string); ok {
			r.R_AddlId = append(r.R_AddlId, assert) // append
			return nil
		} else if asserts, oks := X.([]string); oks {
			r.R_AddlId = asserts // clobber
			return nil
		}
		return errorf("Unsupported AdditionalIdentifier type %T provided without SetFunc instance", X)
	}

	v, err := setfunc[0](X, r)
	if err != nil {
		return err
	}

	return r.SetAdditionalIdentifier(v)
}

func (r RootArc) AdditionalIdentifierGetFunc(getfunc GetFunc) (any, error) {
	if getfunc == nil {
		return nil, errorf("Nil %T provided; aborting", getfunc)
	}

	return getfunc(r.R_AddlId, r)
}

/*
AdditionalIdentifier returns the string slice values
assigned to the receiver, each defining an alternative
identifier by which the registration is known.
*/
func (r SubArc) AdditionalIdentifier() []string {
	return r.R_AddlId
}

/*
SetAdditionalIdentifier appends one or more string slice values
to the receiver's R_AddlId field.
*/
func (r *SubArc) SetAdditionalIdentifier(X any, setfunc ...SetFunc) error {
	if len(setfunc) == 0 {
		if assert, ok := X.(string); ok {
			r.R_AddlId = append(r.R_AddlId, assert) // append
			return nil
		} else if asserts, oks := X.([]string); oks {
			r.R_AddlId = asserts // clobber
			return nil
		}
		return errorf("Unsupported AdditionalIdentifier type %T provided without SetFunc instance", X)
	}

	v, err := setfunc[0](X, r)
	if err != nil {
		return err
	}

	return r.SetAdditionalIdentifier(v)
}

func (r SubArc) AdditionalIdentifierGetFunc(getfunc GetFunc) (any, error) {
	if getfunc == nil {
		return nil, errorf("Nil %T provided; aborting", getfunc)
	}

	return getfunc(r.R_AddlId, r)
}

/*
CurrentAuthority returns the string slice values assigned to
the receiver, each referencing a specific currentAuthority
(x660Registrant) entry by DN.
*/
func (r RootArc) CurrentAuthority() []string {
	return r.R_CAuthyDN
}

/*
SetAdditionalIdentifier appends one or more string slice values
to the receiver's R_AddlId field.
*/
func (r *RootArc) SetCurrentAuthority(X any, setfunc ...SetFunc) error {
	if len(setfunc) == 0 {
		if assert, ok := X.(string); ok {
			r.R_CAuthyDN = append(r.R_CAuthyDN, assert) // append
			return nil
		} else if asserts, oks := X.([]string); oks {
			r.R_CAuthyDN = asserts // clobber
			return nil
		}
		return errorf("Unsupported CurrentAuthority type %T provided without SetFunc instance", X)
	}

	v, err := setfunc[0](X, r)
	if err != nil {
		return err
	}

	return r.SetCurrentAuthority(v)
}

func (r RootArc) CurrentAuthorityGetFunc(getfunc GetFunc) (any, error) {
	if getfunc == nil {
		return nil, errorf("Nil %T provided; aborting", getfunc)
	}

	return getfunc(r.R_CAuthyDN, r)
}

/*
SetCombinedCurrentAuthority assigns a non-nil instance of *CurrentAuthority
to the receiver's R_CAuthy field. This is only to be used in scenarios that
leverage the so-called "Combined Entries" technique, and is generally only
recommended for tiny and/or sparse implementations of draft-coretta-x660-ldap.
*/
func (r *RootArc) SetCombinedCurrentAuthority(c *CurrentAuthority) {
	if c != nil {
		r.R_CAuthy = c
	}
}

/*
CombinedCurrentAuthority returns an instance of *CurrentAuthority present
within the receiver's R_CAuthy field. If unset, nil is returned. Use of
Combined Entries is generally not recommended.
*/
func (r RootArc) CombinedCurrentAuthority() *CurrentAuthority { return r.R_CAuthy }

/*
CurrentAuthority returns the string slice values assigned to
the receiver, each referencing a specific currentAuthority
(x660Registrant) entry by DN.
*/
func (r SubArc) CurrentAuthority() []string {
	return r.R_CAuthyDN
}

/*
SetCurrentAuthority appends one or more string slice DN values
to the receiver's R_CAuthyDN field.
*/
func (r *SubArc) SetCurrentAuthority(X any, setfunc ...SetFunc) error {
	if len(setfunc) == 0 {
		if assert, ok := X.(string); ok {
			r.R_CAuthyDN = append(r.R_CAuthyDN, assert) // append
			return nil
		} else if asserts, oks := X.([]string); oks {
			r.R_CAuthyDN = asserts // clobber
			return nil
		}
		return errorf("Unsupported CurrentAuthority type %T provided without SetFunc instance", X)
	}

	v, err := setfunc[0](X, r)
	if err != nil {
		return err
	}

	return r.SetCurrentAuthority(v)
}

func (r SubArc) CurrentAuthorityGetFunc(getfunc GetFunc) (any, error) {
	if getfunc == nil {
		return nil, errorf("Nil %T provided; aborting", getfunc)
	}

	return getfunc(r.R_CAuthyDN, r)
}

/*
SetCombinedCurrentAuthority assigns a non-nil instance of *CurrentAuthority
to the receiver's R_CAuthy field. This is only to be used in scenarios that
leverage the so-called "Combined Entries" technique, and is generally only
recommended for tiny and/or sparse implementations of draft-coretta-x660-ldap.
*/
func (r *SubArc) SetCombinedCurrentAuthority(c *CurrentAuthority) {
	if c != nil {
		r.R_CAuthy = c
	}
}

/*
CombinedFirstAuthority returns an instance of *FirstAuthority present
within the receiver's R_FAuthy field. If unset, nil is returned. This
*/
func (r SubArc) CombinedCurrentAuthority() *CurrentAuthority { return r.R_CAuthy }

/*
FirstAuthority returns the string slice values assigned to
the receiver, each referencing a specific firstAuthority
(x660Registrant) entry by DN.
*/
func (r RootArc) FirstAuthority() []string {
	return r.R_FAuthyDN
}

/*
CombinedFirstAuthority returns an instance of *FirstAuthority present
within the receiver's R_FAuthy field. If unset, nil is returned. Use of
Combined Entries is generally not recommended.
*/
func (r RootArc) CombinedFirstAuthority() *FirstAuthority { return r.R_FAuthy }

/*
SetFirstAuthority appends one or more string slice DN values
to the receiver's R_FAuthyDN field.
*/
func (r *RootArc) SetFirstAuthority(X any, setfunc ...SetFunc) error {
	if len(setfunc) == 0 {
		if assert, ok := X.(string); ok {
			r.R_FAuthyDN = append(r.R_FAuthyDN, assert) // append
			return nil
		} else if asserts, oks := X.([]string); oks {
			r.R_FAuthyDN = asserts // clobber
			return nil
		}
		return errorf("Unsupported CurrentAuthority type %T provided without SetFunc instance", X)
	}

	v, err := setfunc[0](X, r)
	if err != nil {
		return err
	}

	return r.SetFirstAuthority(v)
}

func (r RootArc) FirstAuthorityGetFunc(getfunc GetFunc) (any, error) {
	if getfunc == nil {
		return nil, errorf("Nil %T provided; aborting", getfunc)
	}

	return getfunc(r.R_FAuthyDN, r)
}

/*
SetCombinedFirstAuthority assigns a non-nil instance of *FirstAuthority
to the receiver's R_CAuthy field. This is only to be used in scenarios that
leverage the so-called "Combined Entries" technique, and is generally only
recommended for tiny and/or sparse implementations of draft-coretta-x660-ldap.
*/
func (r *RootArc) SetCombinedFirstAuthority(c *FirstAuthority) {
	if c != nil {
		r.R_FAuthy = c
	}
}

/*
FirstAuthority returns the string slice values assigned to
the receiver, each referencing a specific firstAuthority
(x660Registrant) entry by DN.
*/
func (r SubArc) FirstAuthority() []string {
	return r.R_FAuthyDN
}

/*
CombinedFirstAuthority returns an instance of *FirstAuthority present
within the receiver's R_FAuthy field. If unset, nil is returned.  Use
of Combined Entries is generally not recommended.
*/
func (r SubArc) CombinedFirstAuthority() *FirstAuthority { return r.R_FAuthy }

/*
SetFirstAuthority appends one or more string slice DN values
to the receiver's R_FAuthyDN field.
*/
func (r *SubArc) SetFirstAuthority(X any, setfunc ...SetFunc) error {
	if len(setfunc) == 0 {
		if assert, ok := X.(string); ok {
			r.R_FAuthyDN = append(r.R_FAuthyDN, assert) // append
			return nil
		} else if asserts, oks := X.([]string); oks {
			r.R_FAuthyDN = asserts // clobber
			return nil
		}
		return errorf("Unsupported CurrentAuthority type %T provided without SetFunc instance", X)
	}

	v, err := setfunc[0](X, r)
	if err != nil {
		return err
	}

	return r.SetFirstAuthority(v)
}

func (r SubArc) FirstAuthorityGetFunc(getfunc GetFunc) (any, error) {
	if getfunc == nil {
		return nil, errorf("Nil %T provided; aborting", getfunc)
	}

	return getfunc(r.R_FAuthyDN, r)
}

/*
SetCombinedFirstAuthority assigns a non-nil instance of *FirstAuthority
to the receiver's R_CAuthy field. This is only to be used in scenarios that
leverage the so-called "Combined Entries" technique, and is generally only
recommended for tiny and/or sparse implementations of draft-coretta-x660-ldap.
*/
func (r *SubArc) SetCombinedFirstAuthority(c *FirstAuthority) {
	if c != nil {
		r.R_FAuthy = c
	}
}

/*
Sponsors returns the an empty string slice, as root registration
do not possess any sponsorship mechanics themselves. This method
exists only to satisfy go interface requirements.
*/
func (r RootArc) Sponsor() []string {
	return []string{}
}

/*
CombinedSponsor always returns nil, as instances of *RootArc do
not possess any sponsorship mechanics themselves. This method
only exists to satisfy Go interface requirements.
*/
func (r RootArc) CombinedSponsor() *Sponsor { return nil }

/*
SetSponsor does not perform any useful task, as Sponsors
do not apply to root arc registrations. This method exists
only to satisfy Go interface requirements.
*/
func (r *RootArc) SetSponsor(X any, setfunc ...SetFunc) error {
	return errorf("Sponsor not applicable to %T", r)
}

func (r RootArc) SponsorGetFunc(getfunc GetFunc) (any, error) {
	return nil, errorf("Sponsor not applicable to %T", r)
}

/*
SetCombinedSponsor does not perform any useful task, as instances of
*RootArc instances do not possess the capability of storing Sponsor
related information. This method only exists to satisfy Go interface
requirements.
*/
func (r *RootArc) SetCombinedSponsor(c *Sponsor) {}

/*
Sponsor returns the string slice values assigned to
the receiver, each referencing a specific sponsor
(x660Registrant) entry by DN.
*/
func (r SubArc) Sponsor() []string {
	return r.R_SAuthyDN
}

/*
CombinedSponsor returns an instance of *Sponsor present within the
receiver's R_SAuthy field. If unset, nil is returned. Use of
Combined Entries is generally not recommended.
*/
func (r SubArc) CombinedSponsor() *Sponsor { return r.R_SAuthy }

/*
SetSponsor appends one or more string slice DN values
to the receiver's R_SAuthyDN field.
*/
func (r *SubArc) SetSponsor(X any, setfunc ...SetFunc) error {
	if len(setfunc) == 0 {
		if assert, ok := X.(string); ok {
			r.R_SAuthyDN = append(r.R_SAuthyDN, assert) // append
			return nil
		} else if asserts, oks := X.([]string); oks {
			r.R_SAuthyDN = asserts // clobber
			return nil
		}
		return errorf("Unsupported Sponsor type %T provided without SetFunc instance", X)
	}

	v, err := setfunc[0](X, r)
	if err != nil {
		return err
	}

	return r.SetSponsor(v)
}

func (r SubArc) SponsorGetFunc(getfunc GetFunc) (any, error) {
	if getfunc == nil {
		return nil, errorf("Nil %T provided; aborting", getfunc)
	}

	return getfunc(r.R_SAuthyDN, r)
}

/*
SetCombinedSponsor assigns a non-nil instance of *Sponsor to the
receiver's R_SAuthy field. This is only to be used in scenarios
that leverage the so-called "Combined Entries" technique, and is
generally only recommended for tiny and/or sparse implementations
of draft-coretta-x660-ldap. Note this will only have any effect
when the underlying receiver instance is *SubArc.
*/
func (r *SubArc) SetCombinedSponsor(c *Sponsor) {
	if c != nil {
		r.R_SAuthy = c
	}
}

/*
DiscloseTo returns the an empty string slice, as root
registrations do not support the discloseTo attribute
type.
*/
func (r RootArc) DiscloseTo() []string {
	return []string{}
}

/*
SetDiscloseTo does not perform any useful task and
exists only to satisfy Go interface requirements.
*/
func (r *RootArc) SetDiscloseTo(X any, setfunc ...SetFunc) error {
	return errorf("DiscloseTo not applicable to %T", r)
}

func (r RootArc) DiscloseToGetFunc(getfunc GetFunc) (any, error) {
	return nil, errorf("DiscloseTo not applicable to %T", r)
}

/*
DiscloseTo returns the string slice values assigned to
the receiver, each defining an identity, group or some
other DN-based reference related to the 'discloseTo'
attribute type mechanics.
*/
func (r SubArc) DiscloseTo() []string {
	return r.R_DiscloseTo
}

/*
SetDiscloseTo appends one or more string slice DN values
to the receiver's R_DiscloseTo field.
*/
func (r *SubArc) SetDiscloseTo(X any, setfunc ...SetFunc) error {
	if len(setfunc) == 0 {
		if assert, ok := X.(string); ok {
			r.R_DiscloseTo = append(r.R_DiscloseTo, assert) // append
			return nil
		} else if asserts, oks := X.([]string); oks {
			r.R_DiscloseTo = asserts // clobber
			return nil
		}
		return errorf("Unsupported DiscloseTo type %T provided without SetFunc instance", X)
	}

	v, err := setfunc[0](X, r)
	if err != nil {
		return err
	}

	return r.SetDiscloseTo(v)
}

func (r SubArc) DiscloseToGetFunc(getfunc GetFunc) (any, error) {
	if getfunc == nil {
		return nil, errorf("Nil %T provided; aborting", getfunc)
	}

	return getfunc(r.R_DiscloseTo, r)
}

/*
IRI returns the string slice values assigned to the
receiver, each defining an internationalized resource
identifier.
*/
func (r RootArc) IRI() []string {
	return r.R_IRI
}

/*
SetIRI appends one or more string slice DN values
to the receiver's R_IRI field.
*/
func (r *RootArc) SetIRI(X any, setfunc ...SetFunc) error {
	if len(setfunc) == 0 {
		if assert, ok := X.(string); ok {
			r.R_IRI = append(r.R_IRI, assert) // append
			return nil
		} else if asserts, oks := X.([]string); oks {
			r.R_IRI = asserts // clobber
			return nil
		}
		return errorf("Unsupported IRI type %T provided without SetFunc instance", X)
	}

	v, err := setfunc[0](X, r)
	if err != nil {
		return err
	}

	return r.SetIRI(v)
}

func (r RootArc) IRIGetFunc(getfunc GetFunc) (any, error) {
	if getfunc == nil {
		return nil, errorf("Nil %T provided; aborting", getfunc)
	}

	return getfunc(r.R_IRI, r)
}

/*
IRI returns the string slice values assigned to the
receiver, each defining an internationalized resource
identifier.
*/
func (r SubArc) IRI() []string {
	return r.R_IRI
}

/*
SetIRI appends one or more string slice DN values
to the receiver's R_IRI field.
*/
func (r *SubArc) SetIRI(X any, setfunc ...SetFunc) error {
	if len(setfunc) == 0 {
		if assert, ok := X.(string); ok {
			r.R_IRI = append(r.R_IRI, assert) // append
			return nil
		} else if asserts, oks := X.([]string); oks {
			r.R_IRI = asserts // clobber
			return nil
		}
		return errorf("Unsupported IRI type %T provided without SetFunc instance", X)
	}

	v, err := setfunc[0](X, r)
	if err != nil {
		return err
	}

	return r.SetIRI(v)
}

func (r SubArc) IRIGetFunc(getfunc GetFunc) (any, error) {
	if getfunc == nil {
		return nil, errorf("Nil %T provided; aborting", getfunc)
	}

	return getfunc(r.R_IRI, r)
}

/*
UnicodeValue returns the string slice values assigned to
the receiver, each defining a unicodeValue assigned to
the registration.
*/
func (r RootArc) UnicodeValue() []string {
	return r.R_UVal
}

/*
SetUnicodeValue appends one or more string slice DN values
to the receiver's R_UVal field.
*/
func (r *RootArc) SetUnicodeValue(X any, setfunc ...SetFunc) error {
	if len(setfunc) == 0 {
		if assert, ok := X.(string); ok {
			r.R_UVal = append(r.R_UVal, assert) // append
			return nil
		} else if asserts, oks := X.([]string); oks {
			r.R_UVal = asserts // clobber
			return nil
		}
		return errorf("Unsupported UnicodeValue type %T provided without SetFunc instance", X)
	}

	v, err := setfunc[0](X, r)
	if err != nil {
		return err
	}

	return r.SetUnicodeValue(v)
}

func (r RootArc) UnicodeValueGetFunc(getfunc GetFunc) (any, error) {
	if getfunc == nil {
		return nil, errorf("Nil %T provided; aborting", getfunc)
	}

	return getfunc(r.R_UVal, r)
}

/*
UnicodeValue returns the string slice values assigned to
the receiver, each defining a unicodeValue assigned to
the registration.
*/
func (r SubArc) UnicodeValue() []string {
	return r.R_UVal
}

/*
SetUnicodeValue appends one or more string slice DN values
to the receiver's R_UVal field.
*/
func (r *SubArc) SetUnicodeValue(X any, setfunc ...SetFunc) error {
	if len(setfunc) == 0 {
		if assert, ok := X.(string); ok {
			r.R_UVal = append(r.R_UVal, assert) // append
			return nil
		} else if asserts, oks := X.([]string); oks {
			r.R_UVal = asserts // clobber
			return nil
		}
		return errorf("Unsupported UnicodeValue type %T provided without SetFunc instance", X)
	}

	v, err := setfunc[0](X, r)
	if err != nil {
		return err
	}

	return r.SetUnicodeValue(v)
}

func (r SubArc) UnicodeValueGetFunc(getfunc GetFunc) (any, error) {
	if getfunc == nil {
		return nil, errorf("Nil %T provided; aborting", getfunc)
	}

	return getfunc(r.R_UVal, r)
}

/*
Info returns the string slice values assigned to the
receiver, each defining an arbitrary block of textual
information assigned to the registration.
*/
func (r RootArc) Info() []string {
	return r.R_Info
}

/*
SetInfo assigns one or more string slice values to the
receiver's R_Info field.
*/
func (r *RootArc) SetInfo(X any, setfunc ...SetFunc) error {
	if len(setfunc) == 0 {
		if assert, ok := X.(string); ok {
			r.R_Info = append(r.R_Info, assert) // append
			return nil
		} else if asserts, oks := X.([]string); oks {
			r.R_Info = asserts // clobber
			return nil
		}
		return errorf("Unsupported Information type %T provided without SetFunc instance", X)
	}

	v, err := setfunc[0](X, r)
	if err != nil {
		return err
	}

	return r.SetInfo(v)
}

func (r RootArc) InfoGetFunc(getfunc GetFunc) (any, error) {
	if getfunc == nil {
		return nil, errorf("Nil %T provided; aborting", getfunc)
	}

	return getfunc(r.R_Info, r)
}

/*
Info returns the string slice values assigned to the
receiver, each defining an arbitrary block of textual
information assigned to the registration.
*/
func (r SubArc) Info() []string {
	return r.R_Info
}

/*
SetInfo assigns one or more string slice values to the
receiver's R_Info field.
*/
func (r *SubArc) SetInfo(X any, setfunc ...SetFunc) error {
	if len(setfunc) == 0 {
		if assert, ok := X.(string); ok {
			r.R_Info = append(r.R_Info, assert) // append
			return nil
		} else if asserts, oks := X.([]string); oks {
			r.R_Info = asserts // clobber
			return nil
		}
		return errorf("Unsupported Information type %T provided without SetFunc instance", X)
	}

	v, err := setfunc[0](X, r)
	if err != nil {
		return err
	}

	return r.SetInfo(v)
}

func (r SubArc) InfoGetFunc(getfunc GetFunc) (any, error) {
	if getfunc == nil {
		return nil, errorf("Nil %T provided; aborting", getfunc)
	}

	return getfunc(r.R_Info, r)
}

/*
StdNameForm returns the string slice values assigned to the
receiver, each defining a standardized name form assigned to
the registration.
*/
func (r RootArc) StdNameForm() []string {
	return r.R_StdNF
}

/*
SetStdNameForm assigns one or more string slice values to the
receiver's R_StdNF field.
*/
func (r *RootArc) SetStdNameForm(X any, setfunc ...SetFunc) error {
	if len(setfunc) == 0 {
		if assert, ok := X.(string); ok {
			r.R_StdNF = append(r.R_StdNF, assert) // append
			return nil
		} else if asserts, oks := X.([]string); oks {
			r.R_StdNF = asserts // clobber
			return nil
		}
		return errorf("Unsupported StdNameForm type %T provided without SetFunc instance", X)
	}

	v, err := setfunc[0](X, r)
	if err != nil {
		return err
	}

	return r.SetStdNameForm(v)
}

func (r RootArc) StdNameFormGetFunc(getfunc GetFunc) (any, error) {
	if getfunc == nil {
		return nil, errorf("Nil %T provided; aborting", getfunc)
	}

	return getfunc(r.R_StdNF, r)
}

/*
StdNameForm returns the string slice values assigned to the
receiver, each defining a standardized name form assigned to
the registration.
*/
func (r SubArc) StdNameForm() []string {
	return r.R_StdNF
}

/*
SetStdNameForm assigns one or more string slice values to the
receiver's R_StdNF field.
*/
func (r *SubArc) SetStdNameForm(X any, setfunc ...SetFunc) error {
	if len(setfunc) == 0 {
		if assert, ok := X.(string); ok {
			r.R_StdNF = append(r.R_StdNF, assert) // append
			return nil
		} else if asserts, oks := X.([]string); oks {
			r.R_StdNF = asserts // clobber
			return nil
		}
		return errorf("Unsupported StdNameForm type %T provided without SetFunc instance", X)
	}

	v, err := setfunc[0](X, r)
	if err != nil {
		return err
	}

	return r.SetStdNameForm(v)
}

func (r SubArc) StdNameFormGetFunc(getfunc GetFunc) (any, error) {
	if getfunc == nil {
		return nil, errorf("Nil %T provided; aborting", getfunc)
	}

	return getfunc(r.R_StdNF, r)
}

/*
URI returns the string slice values assigned to the receiver,
each defining a uniform resource identifier assigned to the
registration.
*/
func (r RootArc) URI() []string {
	return r.R_StdNF
}

/*
SetURI assigns one or more string slice values to the
receiver's R_URI field.
*/
func (r *RootArc) SetURI(X any, setfunc ...SetFunc) error {
	if len(setfunc) == 0 {
		if assert, ok := X.(string); ok {
			r.R_URI = append(r.R_URI, assert) // append
			return nil
		} else if asserts, oks := X.([]string); oks {
			r.R_URI = asserts // clobber
			return nil
		}
		return errorf("Unsupported URI type %T provided without SetFunc instance", X)
	}

	v, err := setfunc[0](X, r)
	if err != nil {
		return err
	}

	return r.SetURI(v)
}

func (r RootArc) URIGetFunc(getfunc GetFunc) (any, error) {
	if getfunc == nil {
		return nil, errorf("Nil %T provided; aborting", getfunc)
	}

	return getfunc(r.R_URI, r)
}

/*
URI returns the string slice values assigned to the receiver,
each defining a uniform resource identifier assigned to the
registration.
*/
func (r SubArc) URI() []string {
	return r.R_StdNF
}

/*
SetURI assigns one or more string slice values to the
receiver's R_URI field.
*/
func (r *SubArc) SetURI(X any, setfunc ...SetFunc) error {
	if len(setfunc) == 0 {
		if assert, ok := X.(string); ok {
			r.R_URI = append(r.R_URI, assert) // append
			return nil
		} else if asserts, oks := X.([]string); oks {
			r.R_URI = asserts // clobber
			return nil
		}
		return errorf("Unsupported URI type %T provided without SetFunc instance", X)
	}

	v, err := setfunc[0](X, r)
	if err != nil {
		return err
	}

	return r.SetURI(v)
}

func (r SubArc) URIGetFunc(getfunc GetFunc) (any, error) {
	if getfunc == nil {
		return nil, errorf("Nil %T provided; aborting", getfunc)
	}

	return getfunc(r.R_URI, r)
}

/*
LongArc returns an empty string slice, as root registrations
do not support use of the longArc.
*/
func (r RootArc) LongArc() []string {
	return []string{}
}

/*
SetLongArc assigns one or more string slice values to the
receiver's R_LongArc field.
*/
func (r *RootArc) SetLongArc(X any, setfunc ...SetFunc) error {
	return errorf("LongArc not applicable to %T", r)
}

func (r RootArc) LongArcGetFunc(getfunc GetFunc) (any, error) {
	return nil, errorf("LongArc not applicable to %T", r)
}

/*
LongArc returns the string slice values assigned to the
receiver, each defining a longArc associated with the
registration.

Not that only sub arcs below Joint-ISO-ITU-T (2) are
permitted to possess longArc values. Assignment of a
longArc to an ITU-T (0) or ISO (1) registration is
considered illegal.
*/
func (r SubArc) LongArc() []string {
	return r.R_LongArc
}

/*
SetLongArc assigns one or more string slice values to the
receiver's R_LongArc field.
*/
func (r *SubArc) SetLongArc(X any, setfunc ...SetFunc) error {
	if len(setfunc) == 0 {
		if assert, ok := X.(string); ok {
			r.R_LongArc = append(r.R_LongArc, assert) // append
			return nil
		} else if asserts, oks := X.([]string); oks {
			r.R_LongArc = asserts // clobber
			return nil
		}
		return errorf("Unsupported LongArc type %T provided without SetFunc instance", X)
	}

	v, err := setfunc[0](X, r)
	if err != nil {
		return err
	}

	return r.SetLongArc(v)
}

func (r SubArc) LongArcGetFunc(getfunc GetFunc) (any, error) {
	if getfunc == nil {
		return nil, errorf("Nil %T provided; aborting", getfunc)
	}

	return getfunc(r.R_LongArc, r)
}

/*
Range returns a zero string, and exists only to satisfy go
interface requirements. The registrationRange attribute type
is not applicable to entries of objectClass 'x660RootArc'.
*/
func (r RootArc) Range() string { return `` }

/*
SetRange performs no useful task, as root registrations
cannot define ranges. This method exists only to satisfy
Go interface requirements.
*/
func (r *RootArc) SetRange(X any, setfunc ...SetFunc) error {
	return errorf("Range not applicable to %T", r)
}

func (r RootArc) RangeGetFunc(getfunc GetFunc) (any, error) {
	return nil, errorf("Range not applicable to %T", r)
}

/*
Range returns the string registrationRange value
assigned to the receiver, or a zero string if unset.
*/
func (r SubArc) Range() string {
	return r.R_Range
}

/*
SetRange assigns a string value, which can be any unsigned
number OR a negative -1.
*/
func (r *SubArc) SetRange(X any, setfunc ...SetFunc) error {
	if len(setfunc) == 0 {
		if assert, ok := X.(string); ok {
			r.R_Range = assert
			return nil
		}
		return errorf("Unsupported Range type %T provided without SetFunc instance", X)
	}

	v, err := setfunc[0](X, r)
	if err != nil {
		return err
	}

	return r.SetRange(v)
}

func (r SubArc) RangeGetFunc(getfunc GetFunc) (any, error) {
	if getfunc == nil {
		return nil, errorf("Nil %T provided; aborting", getfunc)
	}

	return getfunc(r.R_Range, r)
}

/*
LeafNode returns false, and exists only to satisfy go interface
requirements. Entries of objectClass 'x660RootArc' will never
be leaf-nodes.
*/
func (r RootArc) LeafNode() bool {
	return false
}

/*
SetLeafNode performs no useful task. This method exists only to
satisfy Go interface requirements, and does not apply to root
registrations.
*/
func (r *RootArc) SetLeafNode(X any, setfunc ...SetFunc) error {
	return errorf("Leaf node not applicable to %T", r)
}

func (r RootArc) LeafNodeGetFunc(getfunc GetFunc) (any, error) {
	return nil, errorf("Leaf node not applicable to %T", r)
}

/*
LeafNode returns a boolean value indicative of whether the
receiver has been marked as a leaf-node, or false if unset.
*/
func (r SubArc) LeafNode() bool {
	return r.R_LeafNode == `TRUE`
}

/*
SetLeafNode assigns the provided boolean value to the receiver's
R_LeafNode field.
*/
func (r *SubArc) SetLeafNode(X any, setfunc ...SetFunc) error {
	if len(setfunc) == 0 {
		if assert, ok := X.(string); ok {
			r.R_LeafNode = assert
			return nil
		} else if assertb, okb := X.(bool); okb {
			switch assertb {
			case true:
				r.R_LeafNode = `TRUE`
			default:
				r.R_LeafNode = ``
			}
			return nil
		}
		return errorf("Unsupported LeafNode type %T provided without SetFunc instance", X)
	}

	v, err := setfunc[0](X, r)
	if err != nil {
		return err
	}

	return r.SetLeafNode(v)
}

func (r SubArc) LeafNodeGetFunc(getfunc GetFunc) (any, error) {
	if getfunc == nil {
		return nil, errorf("Nil %T provided; aborting", getfunc)
	}

	return getfunc(r.R_LeafNode, r)
}

/*
Frozen returns false, and exists only to satisfy go interface
requirements. Entries of objectClass 'x660RootArc' will never
be frozen.
*/
func (r RootArc) Frozen() bool {
	return false
}

/*
SetFrozen performs no useful task. This method exists only to
satisfy Go interface requirements, and does not apply to root
registrations.
*/
func (r *RootArc) SetFrozen(X any, setfunc ...SetFunc) error {
	return errorf("Frozen not applicable to %T", r)
}

func (r RootArc) FrozenNodeGetFunc(getfunc GetFunc) (any, error) {
	return nil, errorf("Leaf node not applicable to %T", r)
}

/*
Frozen returns a boolean value indicative of whether the
receiver has been marked as frozen, or false if unset.
*/
func (r SubArc) Frozen() bool {
	return r.R_Frozen == `TRUE`
}

/*
SetFrozen assigns the provided boolean value to the receiver's
R_Frozen field.
*/
func (r *SubArc) SetFrozen(X any, setfunc ...SetFunc) error {
	if len(setfunc) == 0 {
		if assert, ok := X.(string); ok {
			r.R_Frozen = assert
			return nil
		} else if assertb, okb := X.(bool); okb {
			switch assertb {
			case true:
				r.R_Frozen = `TRUE`
			default:
				r.R_Frozen = ``
			}
			return nil
		}
		return errorf("Unsupported Frozen type %T provided without SetFunc instance", X)
	}

	v, err := setfunc[0](X, r)
	if err != nil {
		return err
	}

	return r.SetFrozen(v)
}

func (r SubArc) FrozenGetFunc(getfunc GetFunc) (any, error) {
	if getfunc == nil {
		return nil, errorf("Nil %T provided; aborting", getfunc)
	}

	return getfunc(r.R_Frozen, r)
}

/*
SubArc returns zero or more string DN values, each
describing a 'subArc' value assigned to the receiver.
*/
func (r RootArc) SubArc() []string {
	return r.R_SubArc
}

/*
SetSubArc appends one or more string DN values to
the receiver's R_SubArc field.
*/
func (r *RootArc) SetSubArc(X any, setfunc ...SetFunc) error {
	if len(setfunc) == 0 {
		if assert, ok := X.(string); ok {
			r.R_SubArc = append(r.R_SubArc, assert) // append
			return nil
		} else if asserts, oks := X.([]string); oks {
			r.R_SubArc = asserts // clobber
			return nil
		}
		return errorf("Unsupported SubArc type %T provided without SetFunc instance", X)
	}

	v, err := setfunc[0](X, r)
	if err != nil {
		return err
	}

	return r.SetSubArc(v)
}

func (r RootArc) SubArcGetFunc(getfunc GetFunc) (any, error) {
	if getfunc == nil {
		return nil, errorf("Nil %T provided; aborting", getfunc)
	}

	return getfunc(r.R_SubArc, r)
}

/*
SubArc returns zero or more string DN values, each
describing a 'subArc' value assigned to the receiver.
*/
func (r SubArc) SubArc() []string {
	return r.R_SubArc
}

/*
SetSubArc appends one or more string DN values to
the receiver's R_SubArc field.
*/
func (r *SubArc) SetSubArc(X any, setfunc ...SetFunc) error {
	if len(setfunc) == 0 {
		if assert, ok := X.(string); ok {
			r.R_SubArc = append(r.R_SubArc, assert) // append
			return nil
		} else if asserts, oks := X.([]string); oks {
			r.R_SubArc = asserts // clobber
			return nil
		}
		return errorf("Unsupported SubArc type %T provided without SetFunc instance", X)
	}

	v, err := setfunc[0](X, r)
	if err != nil {
		return err
	}

	return r.SetSubArc(v)
}

func (r SubArc) SubArcGetFunc(getfunc GetFunc) (any, error) {
	if getfunc == nil {
		return nil, errorf("Nil %T provided; aborting", getfunc)
	}

	return getfunc(r.R_SubArc, r)
}

/*
Status returns an empty string, and exists only to satisfy
go interface requirements. The registrationStatus attribute
type is not applicable to entries of objectClass 'x660RootArc'.
*/
func (r RootArc) Status() string { return `` }

/*
SetStatus performs no useful task, and exists only to satisfy
Go interface requirements. Status values do not apply to root
registrations.
*/
func (r *RootArc) SetStatus(X any, setfunc ...SetFunc) error {
	return errorf("Status not applicable to %T", r)
}

func (r RootArc) StatusGetFunc(getfunc GetFunc) (any, error) {
	return nil, errorf("Leaf node not applicable to %T", r)
}

/*
Status returns a string value that defines the current status
of the receiver.
*/
func (r SubArc) Status() string {
	return r.R_Status
}

/*
SetStatus assigns a string value to the receiver's R_Status
field.
*/
func (r *SubArc) SetStatus(X any, setfunc ...SetFunc) error {
	if len(setfunc) == 0 {
		if assert, ok := X.(string); ok {
			r.R_Status = assert
			return nil
		}
		return errorf("Unsupported Status type %T provided without SetFunc instance", X)
	}

	v, err := setfunc[0](X, r)
	if err != nil {
		return err
	}

	return r.SetStatus(v)
}

func (r SubArc) StatusGetFunc(getfunc GetFunc) (any, error) {
	if getfunc == nil {
		return nil, errorf("Nil %T provided; aborting", getfunc)
	}

	return getfunc(r.R_Status, r)
}

/*
Unmarshal transports values from the receiver into an instance
of map[string][]string, which can subsequently be fed to go-ldap's
NewEntry function.
*/
func (r RootArc) Unmarshal() map[string][]string {
	return toMap(r)
}

/*
Unmarshal transports values from the receiver into an instance
of map[string][]string, which can subsequently be fed to go-ldap's
NewEntry function.
*/
func (r SubArc) Unmarshal() map[string][]string {
	return toMap(r)
}

/*
Unmarshal is a convenience method that returns slices of map[string][]string
instances, each representative of an individual Registration interface type
that was deemed valid for unmarshaling to map[string][]string.
*/
func (r Registrations) Unmarshal() (maps []map[string][]string) {
	maps = make([]map[string][]string, 0)
	for i := 0; i < len(r); i++ {
		if um := r[i].Unmarshal(); len(um) > 0 {
			maps = append(maps, um)
		}
	}
	return
}

/*
Unmarshal is a convenience method that returns slices of map[string][]string
instances, each representative of an individual Registrant interface type
that was deemed valid for unmarshaling to map[string][]string.
*/
func (r Registrants) Unmarshal() (maps []map[string][]string) {
	maps = make([]map[string][]string, 0)
	for i := 0; i < len(r); i++ {
		if um := r[i].Unmarshal(); len(um) > 0 {
			maps = append(maps, um)
		}
	}

	return
}

/*
ObjectClass is a convenience method that returns the string name of
the associated objectClass for the receiver.

This method can be used as an interface-friendly alternative to type
assertion of any instance encompassed by the Arc empty interface type.
*/
func (r RootArc) ObjectClass() string {
	return `x660RootArc`
}

/*
ObjectClass is a convenience method that returns the string name of
the associated objectClass for the receiver.

This method can be used as an interface-friendly alternative to type
assertion of any instance encompassed by the Arc empty interface type.
*/
func (r SubArc) ObjectClass() string {
	return `x660SubArc`
}

/*
Kind returns the static string value `STRUCTURAL` merely as a convenient means
to declare what kind of objectClass the Registrant receiver describes.
*/
func (r RootArc) Kind() string {
	return `STRUCTURAL`
}

/*
Kind returns the static string value `STRUCTURAL` merely as a convenient means
to declare what kind of objectClass the Registrant receiver describes.
*/
func (r SubArc) Kind() string {
	return `STRUCTURAL`
}

/*
DN returns the distinguished name value assigned to the receiver.
*/
func (r FirstAuthority) DN() string {
	return r.R_DN
}

/*
SetDN assigns the provided string value to the receiver's R_DN field.
*/
func (r *FirstAuthority) SetDN(X any, setfunc ...SetFunc) error {
	if len(setfunc) == 0 {
		if assert, ok := X.(string); ok {
			r.R_DN = assert
			return nil
		}
		return errorf("Unsupported DN type %T provided without SetFunc instance", X)
	}

	v, err := setfunc[0](X, r)
	if err != nil {
		return err
	}

	return r.SetDN(v)
}

func (r FirstAuthority) DNGetFunc(getfunc GetFunc) (any, error) {
	if getfunc == nil {
		return nil, errorf("Nil %T provided; aborting", getfunc)
	}

	return getfunc(r.R_DN, r)
}

/*
DN returns the distinguished name value assigned to the receiver.
*/
func (r CurrentAuthority) DN() string {
	return r.R_DN
}

/*
SetDN assigns the provided string value to the receiver's R_DN field.
*/
func (r *CurrentAuthority) SetDN(X any, setfunc ...SetFunc) error {
	if len(setfunc) == 0 {
		if assert, ok := X.(string); ok {
			r.R_DN = assert
			return nil
		}
		return errorf("Unsupported DN type %T provided without SetFunc instance", X)
	}

	v, err := setfunc[0](X, r)
	if err != nil {
		return err
	}

	return r.SetDN(v)
}

func (r CurrentAuthority) DNGetFunc(getfunc GetFunc) (any, error) {
	if getfunc == nil {
		return nil, errorf("Nil %T provided; aborting", getfunc)
	}

	return getfunc(r.R_DN, r)
}

/*
DN returns the distinguished name value assigned to the receiver.
*/
func (r Sponsor) DN() string {
	return r.R_DN
}

/*
SetDN assigns the provided string value to the receiver's R_DN field.
*/
func (r *Sponsor) SetDN(X any, setfunc ...SetFunc) error {
	if len(setfunc) == 0 {
		if assert, ok := X.(string); ok {
			r.R_DN = assert
			return nil
		}
		return errorf("Unsupported DN type %T provided without SetFunc instance", X)
	}

	v, err := setfunc[0](X, r)
	if err != nil {
		return err
	}

	return r.SetDN(v)
}

func (r Sponsor) DNGetFunc(getfunc GetFunc) (any, error) {
	if getfunc == nil {
		return nil, errorf("Nil %T provided; aborting", getfunc)
	}

	return getfunc(r.R_DN, r)
}

/*
RegistrantID returns the registrantID value assigned to the receiver.
*/
func (r FirstAuthority) RegistrantID() string {
	return r.R_Id
}

func (r FirstAuthority) RegistrantIDGetFunc(getfunc GetFunc) (any, error) {
	if getfunc == nil {
		return nil, errorf("Nil %T provided; aborting", getfunc)
	}

	return getfunc(r.R_Id, r)
}

/*
RegistrantID returns the registrantID value assigned to the receiver.
*/
func (r CurrentAuthority) RegistrantID() string {
	return r.R_Id
}

func (r CurrentAuthority) RegistrantIDGetFunc(getfunc GetFunc) (any, error) {
	if getfunc == nil {
		return nil, errorf("Nil %T provided; aborting", getfunc)
	}

	return getfunc(r.R_Id, r)
}

/*
RegistrantID returns the registrantID value assigned to the receiver.
*/
func (r Sponsor) RegistrantID() string {
	return r.R_Id
}

func (r Sponsor) RegistrantIDGetFunc(getfunc GetFunc) (any, error) {
	if getfunc == nil {
		return nil, errorf("Nil %T provided; aborting", getfunc)
	}

	return getfunc(r.R_Id, r)
}

/*
CreateTime returns a string generalizedTime value assigned to the receiver's
R_Created field.
*/
func (r RootArc) CreateTime() string {
	return r.R_Created
}

func (r RootArc) CreateTimeGetFunc(getfunc GetFunc) (any, error) {
	if getfunc == nil {
		return nil, errorf("Nil %T provided; aborting", getfunc)
	}

	return getfunc(r.R_Created, r)
}

/*
SetCreateTime assigns a string-based generalizedTime value to the receiver's R_Created field.
*/
func (r *RootArc) SetCreateTime(X any, setfunc ...SetFunc) error {
	if len(setfunc) == 0 {
		if assert, ok := X.(string); ok {
			r.R_Created = assert
			return nil
		}
		return errorf("Unsupported CreateTime type %T provided without SetFunc instance", X)
	}

	v, err := setfunc[0](X, r)
	if err != nil {
		return err
	}

	return r.SetCreateTime(v)
}

/*
CreateTime returns a string generalizedTime value assigned to the receiver's
R_Created field.
*/
func (r SubArc) CreateTime() string {
	return r.R_Created
}

func (r SubArc) CreateTimeGetFunc(getfunc GetFunc) (any, error) {
	if getfunc == nil {
		return nil, errorf("Nil %T provided; aborting", getfunc)
	}

	return getfunc(r.R_Created, r)
}

/*
SetCreateTime assigns a string-based generalizedTime value instance to the receiver's R_Created field.
*/
func (r *SubArc) SetCreateTime(X any, setfunc ...SetFunc) error {
	if len(setfunc) == 0 {
		if assert, ok := X.(string); ok {
			r.R_Created = assert
			return nil
		}
		return errorf("Unsupported CreateTime type %T provided without SetFunc instance", X)
	}

	v, err := setfunc[0](X, r)
	if err != nil {
		return err
	}

	return r.SetCreateTime(v)
}

/*
ModifyTime returns slices of string generalizedTime values, each
reflecting a time and date at which the receiver was reported to
be modified.
*/
func (r RootArc) ModifyTime() []string {
	return r.R_Modified
}

func (r RootArc) ModifyTimeGetFunc(getfunc GetFunc) (any, error) {
	if getfunc == nil {
		return nil, errorf("Nil %T provided; aborting", getfunc)
	}

	return getfunc(r.R_Modified, r)
}

/*
SetModifyTime appends one or more instances of string-based generalizedTime
value to the receiver's R_Modified field.
*/
func (r *RootArc) SetModifyTime(X any, setfunc ...SetFunc) error {
	if len(setfunc) == 0 {
		if assert, ok := X.(string); ok {
			r.R_Modified = append(r.R_Modified, assert) // append
			return nil
		} else if asserts, oks := X.([]string); oks {
			r.R_Modified = asserts // clobber
			return nil
		}
		return errorf("Unsupported ModifyTime type %T provided without SetFunc instance", X)
	}

	v, err := setfunc[0](X, r)
	if err != nil {
		return err
	}

	return r.SetModifyTime(v)
}

/*
ModifyTime returns slices of string generalizedTime values, each
reflecting a time and date at which the receiver was reported to
be modified.
*/
func (r SubArc) ModifyTime() (T []string) {
	return r.R_Modified
}

func (r SubArc) ModifyTimeGetFunc(getfunc GetFunc) (any, error) {
	if getfunc == nil {
		return nil, errorf("Nil %T provided; aborting", getfunc)
	}

	return getfunc(r.R_Modified, r)
}

/*
SetModifyTime appends one or more instances of string-based generalizedTime
value to the receiver's R_Modified field.
*/
func (r *SubArc) SetModifyTime(X any, setfunc ...SetFunc) error {
	if len(setfunc) == 0 {
		if assert, ok := X.(string); ok {
			r.R_Modified = append(r.R_Modified, assert) // append
			return nil
		} else if asserts, oks := X.([]string); oks {
			r.R_Modified = asserts // clobber
			return nil
		}
		return errorf("Unsupported ModifyTime type %T provided without SetFunc instance", X)
	}

	v, err := setfunc[0](X, r)
	if err != nil {
		return err
	}

	return r.SetModifyTime(v)
}

/*
CN returns the common name value assigned to the receiver.
*/
func (r FirstAuthority) CN() string {
	return r.R_CN
}

/*
SetCN assigns the provided string value to the receiver's
R_CN field.
*/
func (r *FirstAuthority) SetCN(X any, setfunc ...SetFunc) error {
	if len(setfunc) == 0 {
		if assert, ok := X.(string); ok {
			r.R_CN = assert
			return nil
		}
		return errorf("Unsupported CN type %T provided without SetFunc instance", X)
	}

	v, err := setfunc[0](X, r)
	if err != nil {
		return err
	}

	return r.SetCN(v)
}

func (r FirstAuthority) CNGetFunc(getfunc GetFunc) (any, error) {
	if getfunc == nil {
		return nil, errorf("Nil %T provided; aborting", getfunc)
	}

	return getfunc(r.R_CN, r)
}

/*
CN returns the common name value assigned to the receiver.
*/
func (r CurrentAuthority) CN() string {
	return r.R_CN
}

/*
SetCN assigns the provided string value to the receiver's
R_CN field.
*/
func (r *CurrentAuthority) SetCN(X any, setfunc ...SetFunc) error {
	if len(setfunc) == 0 {
		if assert, ok := X.(string); ok {
			r.R_CN = assert
			return nil
		}
		return errorf("Unsupported CN type %T provided without SetFunc instance", X)
	}

	v, err := setfunc[0](X, r)
	if err != nil {
		return err
	}

	return r.SetCN(v)
}

func (r CurrentAuthority) CNGetFunc(getfunc GetFunc) (any, error) {
	if getfunc == nil {
		return nil, errorf("Nil %T provided; aborting", getfunc)
	}

	return getfunc(r.R_CN, r)
}

/*
CN returns the common name value assigned to the receiver.
*/
func (r Sponsor) CN() string {
	return r.R_CN
}

/*
SetCN assigns the provided string value to the receiver's
R_CN field.
*/
func (r *Sponsor) SetCN(X any, setfunc ...SetFunc) error {
	if len(setfunc) == 0 {
		if assert, ok := X.(string); ok {
			r.R_CN = assert
			return nil
		}
		return errorf("Unsupported CN type %T provided without SetFunc instance", X)
	}

	v, err := setfunc[0](X, r)
	if err != nil {
		return err
	}

	return r.SetCN(v)
}

func (r Sponsor) CNGetFunc(getfunc GetFunc) (any, error) {
	if getfunc == nil {
		return nil, errorf("Nil %T provided; aborting", getfunc)
	}

	return getfunc(r.R_CN, r)
}

/*
L returns the locality name value assigned to the receiver.
*/
func (r FirstAuthority) L() string {
	return r.R_L
}

/*
SetL assigns the provided string value to the receiver's
R_L field.
*/
func (r *FirstAuthority) SetL(X any, setfunc ...SetFunc) error {
	if len(setfunc) == 0 {
		if assert, ok := X.(string); ok {
			r.R_L = assert
			return nil
		}
		return errorf("Unsupported L type %T provided without SetFunc instance", X)
	}

	v, err := setfunc[0](X, r)
	if err != nil {
		return err
	}

	return r.SetL(v)
}

func (r FirstAuthority) LGetFunc(getfunc GetFunc) (any, error) {
	if getfunc == nil {
		return nil, errorf("Nil %T provided; aborting", getfunc)
	}

	return getfunc(r.R_L, r)
}

/*
L returns the locality name value assigned to the receiver.
*/
func (r CurrentAuthority) L() string {
	return r.R_L
}

/*
SetL assigns the provided string value to the receiver's
R_L field.
*/
func (r *CurrentAuthority) SetL(X any, setfunc ...SetFunc) error {
	if len(setfunc) == 0 {
		if assert, ok := X.(string); ok {
			r.R_L = assert
			return nil
		}
		return errorf("Unsupported L type %T provided without SetFunc instance", X)
	}

	v, err := setfunc[0](X, r)
	if err != nil {
		return err
	}

	return r.SetL(v)
}

func (r CurrentAuthority) LGetFunc(getfunc GetFunc) (any, error) {
	if getfunc == nil {
		return nil, errorf("Nil %T provided; aborting", getfunc)
	}

	return getfunc(r.R_L, r)
}

/*
L returns the locality name value assigned to the receiver.
*/
func (r Sponsor) L() string {
	return r.R_L
}

/*
SetL assigns the provided string value to the receiver's
R_L field.
*/
func (r *Sponsor) SetL(X any, setfunc ...SetFunc) error {
	if len(setfunc) == 0 {
		if assert, ok := X.(string); ok {
			r.R_L = assert
			return nil
		}
		return errorf("Unsupported L type %T provided without SetFunc instance", X)
	}

	v, err := setfunc[0](X, r)
	if err != nil {
		return err
	}

	return r.SetL(v)
}

func (r Sponsor) LGetFunc(getfunc GetFunc) (any, error) {
	if getfunc == nil {
		return nil, errorf("Nil %T provided; aborting", getfunc)
	}

	return getfunc(r.R_L, r)
}

/*
O returns the organization name value assigned to the receiver.
*/
func (r FirstAuthority) O() string {
	return r.R_O
}

/*
SetO assigns the provided string value to the receiver's
R_O field.
*/
func (r *FirstAuthority) SetO(X any, setfunc ...SetFunc) error {
	if len(setfunc) == 0 {
		if assert, ok := X.(string); ok {
			r.R_O = assert
			return nil
		}
		return errorf("Unsupported O type %T provided without SetFunc instance", X)
	}

	v, err := setfunc[0](X, r)
	if err != nil {
		return err
	}

	return r.SetO(v)
}

func (r FirstAuthority) OGetFunc(getfunc GetFunc) (any, error) {
	if getfunc == nil {
		return nil, errorf("Nil %T provided; aborting", getfunc)
	}

	return getfunc(r.R_O, r)
}

/*
O returns the organization name value assigned to the receiver.
*/
func (r CurrentAuthority) O() string {
	return r.R_O
}

/*
SetO assigns the provided string value to the receiver's
R_O field.
*/
func (r *CurrentAuthority) SetO(X any, setfunc ...SetFunc) error {
	if len(setfunc) == 0 {
		if assert, ok := X.(string); ok {
			r.R_O = assert
			return nil
		}
		return errorf("Unsupported O type %T provided without SetFunc instance", X)
	}

	v, err := setfunc[0](X, r)
	if err != nil {
		return err
	}

	return r.SetO(v)
}

func (r CurrentAuthority) OGetFunc(getfunc GetFunc) (any, error) {
	if getfunc == nil {
		return nil, errorf("Nil %T provided; aborting", getfunc)
	}

	return getfunc(r.R_O, r)
}

/*
O returns the organization name value assigned to the receiver.
*/
func (r Sponsor) O() string {
	return r.R_O
}

/*
SetO assigns the provided string value to the receiver's
R_O field.
*/
func (r *Sponsor) SetO(X any, setfunc ...SetFunc) error {
	if len(setfunc) == 0 {
		if assert, ok := X.(string); ok {
			r.R_O = assert
			return nil
		}
		return errorf("Unsupported O type %T provided without SetFunc instance", X)
	}

	v, err := setfunc[0](X, r)
	if err != nil {
		return err
	}

	return r.SetO(v)
}

func (r Sponsor) OGetFunc(getfunc GetFunc) (any, error) {
	if getfunc == nil {
		return nil, errorf("Nil %T provided; aborting", getfunc)
	}

	return getfunc(r.R_O, r)
}

/*
C returns the 2-letter country code value assigned to the receiver.
*/
func (r FirstAuthority) C() string {
	return r.R_C
}

/*
SetC assigns the provided string value to the receiver's
R_C field.
*/
func (r *FirstAuthority) SetC(X any, setfunc ...SetFunc) error {
	if len(setfunc) == 0 {
		if assert, ok := X.(string); ok {
			r.R_C = assert
			return nil
		}
		return errorf("Unsupported C type %T provided without SetFunc instance", X)
	}

	v, err := setfunc[0](X, r)
	if err != nil {
		return err
	}

	return r.SetC(v)
}

func (r FirstAuthority) CGetFunc(getfunc GetFunc) (any, error) {
	if getfunc == nil {
		return nil, errorf("Nil %T provided; aborting", getfunc)
	}

	return getfunc(r.R_C, r)
}

/*
C returns the 2-letter country code value assigned to the receiver.
*/
func (r CurrentAuthority) C() string {
	return r.R_C
}

/*
SetC assigns the provided string value to the receiver's
R_C field.
*/
func (r *CurrentAuthority) SetC(X any, setfunc ...SetFunc) error {
	if len(setfunc) == 0 {
		if assert, ok := X.(string); ok {
			r.R_C = assert
			return nil
		}
		return errorf("Unsupported C type %T provided without SetFunc instance", X)
	}

	v, err := setfunc[0](X, r)
	if err != nil {
		return err
	}

	return r.SetC(v)
}

func (r CurrentAuthority) CGetFunc(getfunc GetFunc) (any, error) {
	if getfunc == nil {
		return nil, errorf("Nil %T provided; aborting", getfunc)
	}

	return getfunc(r.R_C, r)
}

/*
C returns the 2-letter country code value assigned to the receiver.
*/
func (r Sponsor) C() string {
	return r.R_C
}

/*
SetC assigns the provided string value to the receiver's
R_C field.
*/
func (r *Sponsor) SetC(X any, setfunc ...SetFunc) error {
	if len(setfunc) == 0 {
		if assert, ok := X.(string); ok {
			r.R_C = assert
			return nil
		}
		return errorf("Unsupported C type %T provided without SetFunc instance", X)
	}

	v, err := setfunc[0](X, r)
	if err != nil {
		return err
	}

	return r.SetC(v)
}

func (r Sponsor) CGetFunc(getfunc GetFunc) (any, error) {
	if getfunc == nil {
		return nil, errorf("Nil %T provided; aborting", getfunc)
	}

	return getfunc(r.R_C, r)
}

/*
CO returns the so-called "friendly country name" value assigned to the receiver.
*/
func (r FirstAuthority) CO() string {
	return r.R_CO
}

/*
SetCO assigns the provided string value to the receiver's
R_CO field.
*/
func (r *FirstAuthority) SetCO(X any, setfunc ...SetFunc) error {
	if len(setfunc) == 0 {
		if assert, ok := X.(string); ok {
			r.R_CO = assert
			return nil
		}
		return errorf("Unsupported CO type %T provided without SetFunc instance", X)
	}

	v, err := setfunc[0](X, r)
	if err != nil {
		return err
	}

	return r.SetCO(v)
}

func (r FirstAuthority) COGetFunc(getfunc GetFunc) (any, error) {
	if getfunc == nil {
		return nil, errorf("Nil %T provided; aborting", getfunc)
	}

	return getfunc(r.R_CO, r)
}

/*
CO returns the so-called "friendly country name" value assigned to the receiver.
*/
func (r CurrentAuthority) CO() string {
	return r.R_CO
}

/*
SetCO assigns the provided string value to the receiver's
R_CO field.
*/
func (r *CurrentAuthority) SetCO(X any, setfunc ...SetFunc) error {
	if len(setfunc) == 0 {
		if assert, ok := X.(string); ok {
			r.R_CO = assert
			return nil
		}
		return errorf("Unsupported CO type %T provided without SetFunc instance", X)
	}

	v, err := setfunc[0](X, r)
	if err != nil {
		return err
	}

	return r.SetCO(v)
}

func (r CurrentAuthority) COGetFunc(getfunc GetFunc) (any, error) {
	if getfunc == nil {
		return nil, errorf("Nil %T provided; aborting", getfunc)
	}

	return getfunc(r.R_CO, r)
}

/*
CO returns the so-called "friendly country name" value assigned to the receiver.
*/
func (r Sponsor) CO() string {
	return r.R_CO
}

/*
SetCO assigns the provided string value to the receiver's
R_CO field.
*/
func (r *Sponsor) SetCO(X any, setfunc ...SetFunc) error {
	if len(setfunc) == 0 {
		if assert, ok := X.(string); ok {
			r.R_CO = assert
			return nil
		}
		return errorf("Unsupported CO type %T provided without SetFunc instance", X)
	}

	v, err := setfunc[0](X, r)
	if err != nil {
		return err
	}

	return r.SetCO(v)
}

func (r Sponsor) COGetFunc(getfunc GetFunc) (any, error) {
	if getfunc == nil {
		return nil, errorf("Nil %T provided; aborting", getfunc)
	}

	return getfunc(r.R_CO, r)
}

/*
ST returns the state or province name value assigned to the receiver.
*/
func (r FirstAuthority) ST() string {
	return r.R_ST
}

/*
SetST assigns the provided string value to the receiver's
R_ST field.
*/
func (r *FirstAuthority) SetST(X any, setfunc ...SetFunc) error {
	if len(setfunc) == 0 {
		if assert, ok := X.(string); ok {
			r.R_ST = assert
			return nil
		}
		return errorf("Unsupported ST type %T provided without SetFunc instance", X)
	}

	v, err := setfunc[0](X, r)
	if err != nil {
		return err
	}

	return r.SetST(v)
}

func (r FirstAuthority) STGetFunc(getfunc GetFunc) (any, error) {
	if getfunc == nil {
		return nil, errorf("Nil %T provided; aborting", getfunc)
	}

	return getfunc(r.R_ST, r)
}

/*
ST returns the state or province name value assigned to the receiver.
*/
func (r CurrentAuthority) ST() string {
	return r.R_ST
}

/*
SetST assigns the provided string value to the receiver's
R_ST field.
*/
func (r *CurrentAuthority) SetST(X any, setfunc ...SetFunc) error {
	if len(setfunc) == 0 {
		if assert, ok := X.(string); ok {
			r.R_ST = assert
			return nil
		}
		return errorf("Unsupported ST type %T provided without SetFunc instance", X)
	}

	v, err := setfunc[0](X, r)
	if err != nil {
		return err
	}

	return r.SetST(v)
}

func (r CurrentAuthority) STGetFunc(getfunc GetFunc) (any, error) {
	if getfunc == nil {
		return nil, errorf("Nil %T provided; aborting", getfunc)
	}

	return getfunc(r.R_ST, r)
}

/*
ST returns the state or province name value assigned to the receiver.
*/
func (r Sponsor) ST() string {
	return r.R_ST
}

/*
SetST assigns the provided string value to the receiver's
R_ST field.
*/
func (r *Sponsor) SetST(X any, setfunc ...SetFunc) error {
	if len(setfunc) == 0 {
		if assert, ok := X.(string); ok {
			r.R_ST = assert
			return nil
		}
		return errorf("Unsupported ST type %T provided without SetFunc instance", X)
	}

	v, err := setfunc[0](X, r)
	if err != nil {
		return err
	}

	return r.SetST(v)
}

func (r Sponsor) STGetFunc(getfunc GetFunc) (any, error) {
	if getfunc == nil {
		return nil, errorf("Nil %T provided; aborting", getfunc)
	}

	return getfunc(r.R_ST, r)
}

/*
Tel returns the telephone number value assigned to the receiver.
*/
func (r FirstAuthority) Tel() string {
	return r.R_Tel
}

/*
SetTel assigns the provided string value to the receiver's
R_Tel field.
*/
func (r *FirstAuthority) SetTel(X any, setfunc ...SetFunc) error {
	if len(setfunc) == 0 {
		if assert, ok := X.(string); ok {
			r.R_Tel = assert
			return nil
		}
		return errorf("Unsupported Tel type %T provided without SetFunc instance", X)
	}

	v, err := setfunc[0](X, r)
	if err != nil {
		return err
	}

	return r.SetTel(v)
}

func (r FirstAuthority) TelGetFunc(getfunc GetFunc) (any, error) {
	if getfunc == nil {
		return nil, errorf("Nil %T provided; aborting", getfunc)
	}

	return getfunc(r.R_Tel, r)
}

/*
Tel returns the telephone number value assigned to the receiver.
*/
func (r CurrentAuthority) Tel() string {
	return r.R_Tel
}

/*
SetTel assigns the provided string value to the receiver's
R_Tel field.
*/
func (r *CurrentAuthority) SetTel(X any, setfunc ...SetFunc) error {
	if len(setfunc) == 0 {
		if assert, ok := X.(string); ok {
			r.R_Tel = assert
			return nil
		}
		return errorf("Unsupported Tel type %T provided without SetFunc instance", X)
	}

	v, err := setfunc[0](X, r)
	if err != nil {
		return err
	}

	return r.SetTel(v)
}

func (r CurrentAuthority) TelGetFunc(getfunc GetFunc) (any, error) {
	if getfunc == nil {
		return nil, errorf("Nil %T provided; aborting", getfunc)
	}

	return getfunc(r.R_Tel, r)
}

/*
Tel returns the telephone number value assigned to the receiver.
*/
func (r Sponsor) Tel() string {
	return r.R_Tel
}

/*
SetTel assigns the provided string value to the receiver's
R_Tel field.
*/
func (r *Sponsor) SetTel(X any, setfunc ...SetFunc) error {
	if len(setfunc) == 0 {
		if assert, ok := X.(string); ok {
			r.R_Tel = assert
			return nil
		}
		return errorf("Unsupported Tel type %T provided without SetFunc instance", X)
	}

	v, err := setfunc[0](X, r)
	if err != nil {
		return err
	}

	return r.SetTel(v)
}

func (r Sponsor) TelGetFunc(getfunc GetFunc) (any, error) {
	if getfunc == nil {
		return nil, errorf("Nil %T provided; aborting", getfunc)
	}

	return getfunc(r.R_Tel, r)
}

/*
Fax returns the facsimile telephone number value assigned to the receiver.
*/
func (r FirstAuthority) Fax() string {
	return r.R_Fax
}

/*
SetFax assigns the provided string value to the receiver's
R_Fax field.
*/
func (r *FirstAuthority) SetFax(X any, setfunc ...SetFunc) error {
	if len(setfunc) == 0 {
		if assert, ok := X.(string); ok {
			r.R_Fax = assert
			return nil
		}
		return errorf("Unsupported Fax type %T provided without SetFunc instance", X)
	}

	v, err := setfunc[0](X, r)
	if err != nil {
		return err
	}

	return r.SetFax(v)
}

func (r FirstAuthority) FaxGetFunc(getfunc GetFunc) (any, error) {
	if getfunc == nil {
		return nil, errorf("Nil %T provided; aborting", getfunc)
	}

	return getfunc(r.R_Fax, r)
}

/*
Fax returns the facsimile telephone number value assigned to the receiver.
*/
func (r CurrentAuthority) Fax() string {
	return r.R_Fax
}

/*
SetFax assigns the provided string value to the receiver's
R_Fax field.
*/
func (r *CurrentAuthority) SetFax(X any, setfunc ...SetFunc) error {
	if len(setfunc) == 0 {
		if assert, ok := X.(string); ok {
			r.R_Fax = assert
			return nil
		}
		return errorf("Unsupported Fax type %T provided without SetFunc instance", X)
	}

	v, err := setfunc[0](X, r)
	if err != nil {
		return err
	}

	return r.SetFax(v)
}

func (r CurrentAuthority) FaxGetFunc(getfunc GetFunc) (any, error) {
	if getfunc == nil {
		return nil, errorf("Nil %T provided; aborting", getfunc)
	}

	return getfunc(r.R_Fax, r)
}

/*
Fax returns the facsimile telephone number value assigned to the receiver.
*/
func (r Sponsor) Fax() string {
	return r.R_Fax
}

/*
SetFax assigns the provided string value to the receiver's
R_Fax field.
*/
func (r *Sponsor) SetFax(X any, setfunc ...SetFunc) error {
	if len(setfunc) == 0 {
		if assert, ok := X.(string); ok {
			r.R_Fax = assert
			return nil
		}
		return errorf("Unsupported Fax type %T provided without SetFunc instance", X)
	}

	v, err := setfunc[0](X, r)
	if err != nil {
		return err
	}

	return r.SetFax(v)
}

func (r Sponsor) FaxGetFunc(getfunc GetFunc) (any, error) {
	if getfunc == nil {
		return nil, errorf("Nil %T provided; aborting", getfunc)
	}

	return getfunc(r.R_Fax, r)
}

/*
Title returns the title value assigned to the receiver.
*/
func (r FirstAuthority) Title() string {
	return r.R_Title
}

/*
SetTitle assigns the provided string value to the receiver's
R_Title field.
*/
func (r *FirstAuthority) SetTitle(X any, setfunc ...SetFunc) error {
	if len(setfunc) == 0 {
		if assert, ok := X.(string); ok {
			r.R_Title = assert
			return nil
		}
		return errorf("Unsupported Title type %T provided without SetFunc instance", X)
	}

	v, err := setfunc[0](X, r)
	if err != nil {
		return err
	}

	return r.SetTitle(v)
}

func (r FirstAuthority) TitleGetFunc(getfunc GetFunc) (any, error) {
	if getfunc == nil {
		return nil, errorf("Nil %T provided; aborting", getfunc)
	}

	return getfunc(r.R_Title, r)
}

/*
Title returns the title value assigned to the receiver.
*/
func (r CurrentAuthority) Title() string {
	return r.R_Title
}

/*
SetTitle assigns the provided string value to the receiver's
R_Title field.
*/
func (r *CurrentAuthority) SetTitle(X any, setfunc ...SetFunc) error {
	if len(setfunc) == 0 {
		if assert, ok := X.(string); ok {
			r.R_Title = assert
			return nil
		}
		return errorf("Unsupported Title type %T provided without SetFunc instance", X)
	}

	v, err := setfunc[0](X, r)
	if err != nil {
		return err
	}

	return r.SetTitle(v)
}

func (r CurrentAuthority) TitleGetFunc(getfunc GetFunc) (any, error) {
	if getfunc == nil {
		return nil, errorf("Nil %T provided; aborting", getfunc)
	}

	return getfunc(r.R_Title, r)
}

/*
Title returns the title value assigned to the receiver.
*/
func (r Sponsor) Title() string {
	return r.R_Title
}

/*
SetTitle assigns the provided string value to the receiver's
R_Title field.
*/
func (r *Sponsor) SetTitle(X any, setfunc ...SetFunc) error {
	if len(setfunc) == 0 {
		if assert, ok := X.(string); ok {
			r.R_Title = assert
			return nil
		}
		return errorf("Unsupported Title type %T provided without SetFunc instance", X)
	}

	v, err := setfunc[0](X, r)
	if err != nil {
		return err
	}

	return r.SetTitle(v)
}

func (r Sponsor) TitleGetFunc(getfunc GetFunc) (any, error) {
	if getfunc == nil {
		return nil, errorf("Nil %T provided; aborting", getfunc)
	}

	return getfunc(r.R_Title, r)
}

/*
Email returns the email address value assigned to the receiver.
*/
func (r FirstAuthority) Email() string {
	return r.R_Email
}

/*
SetEmail assigns the provided string value to the receiver's
R_Email field.
*/
func (r *FirstAuthority) SetEmail(X any, setfunc ...SetFunc) error {
	if len(setfunc) == 0 {
		if assert, ok := X.(string); ok {
			r.R_Email = assert
			return nil
		}
		return errorf("Unsupported Email type %T provided without SetFunc instance", X)
	}

	v, err := setfunc[0](X, r)
	if err != nil {
		return err
	}

	return r.SetEmail(v)
}

func (r FirstAuthority) EmailGetFunc(getfunc GetFunc) (any, error) {
	if getfunc == nil {
		return nil, errorf("Nil %T provided; aborting", getfunc)
	}

	return getfunc(r.R_Email, r)
}

/*
Email returns the email address value assigned to the receiver.
*/
func (r CurrentAuthority) Email() string {
	return r.R_Email
}

/*
SetEmail assigns the provided string value to the receiver's
R_Email field.
*/
func (r *CurrentAuthority) SetEmail(X any, setfunc ...SetFunc) error {
	if len(setfunc) == 0 {
		if assert, ok := X.(string); ok {
			r.R_Email = assert
			return nil
		}
		return errorf("Unsupported Email type %T provided without SetFunc instance", X)
	}

	v, err := setfunc[0](X, r)
	if err != nil {
		return err
	}

	return r.SetEmail(v)
}

func (r CurrentAuthority) EmailGetFunc(getfunc GetFunc) (any, error) {
	if getfunc == nil {
		return nil, errorf("Nil %T provided; aborting", getfunc)
	}

	return getfunc(r.R_Email, r)
}

/*
Email returns the email address value assigned to the receiver.
*/
func (r Sponsor) Email() string {
	return r.R_Email
}

/*
SetEmail assigns the provided string value to the receiver's
R_Email field.
*/
func (r *Sponsor) SetEmail(X any, setfunc ...SetFunc) error {
	if len(setfunc) == 0 {
		if assert, ok := X.(string); ok {
			r.R_Email = assert
			return nil
		}
		return errorf("Unsupported Email type %T provided without SetFunc instance", X)
	}

	v, err := setfunc[0](X, r)
	if err != nil {
		return err
	}

	return r.SetEmail(v)
}

func (r Sponsor) EmailGetFunc(getfunc GetFunc) (any, error) {
	if getfunc == nil {
		return nil, errorf("Nil %T provided; aborting", getfunc)
	}

	return getfunc(r.R_Email, r)
}

/*
POBox returns the postal office box value assigned to the receiver.
*/
func (r FirstAuthority) POBox() string {
	return r.R_POBox
}

/*
SetPOBox assigns the provided string value to the receiver's
R_POBox field.
*/
func (r *FirstAuthority) SetPOBox(X any, setfunc ...SetFunc) error {
	if len(setfunc) == 0 {
		if assert, ok := X.(string); ok {
			r.R_POBox = assert
			return nil
		}
		return errorf("Unsupported POBox type %T provided without SetFunc instance", X)
	}

	v, err := setfunc[0](X, r)
	if err != nil {
		return err
	}

	return r.SetPOBox(v)
}

func (r FirstAuthority) POBoxGetFunc(getfunc GetFunc) (any, error) {
	if getfunc == nil {
		return nil, errorf("Nil %T provided; aborting", getfunc)
	}

	return getfunc(r.R_POBox, r)
}

/*
POBox returns the postal office box value assigned to the receiver.
*/
func (r CurrentAuthority) POBox() string {
	return r.R_POBox
}

/*
SetPOBox assigns the provided string value to the receiver's
R_POBox field.
*/
func (r *CurrentAuthority) SetPOBox(X any, setfunc ...SetFunc) error {
	if len(setfunc) == 0 {
		if assert, ok := X.(string); ok {
			r.R_POBox = assert
			return nil
		}
		return errorf("Unsupported POBox type %T provided without SetFunc instance", X)
	}

	v, err := setfunc[0](X, r)
	if err != nil {
		return err
	}

	return r.SetPOBox(v)
}

func (r CurrentAuthority) POBoxGetFunc(getfunc GetFunc) (any, error) {
	if getfunc == nil {
		return nil, errorf("Nil %T provided; aborting", getfunc)
	}

	return getfunc(r.R_POBox, r)
}

/*
POBox returns the postal office box value assigned to the receiver.
*/
func (r Sponsor) POBox() string {
	return r.R_POBox
}

/*
SetPOBox assigns the provided string value to the receiver's
R_POBox field.
*/
func (r *Sponsor) SetPOBox(X any, setfunc ...SetFunc) error {
	if len(setfunc) == 0 {
		if assert, ok := X.(string); ok {
			r.R_POBox = assert
			return nil
		}
		return errorf("Unsupported POBox type %T provided without SetFunc instance", X)
	}

	v, err := setfunc[0](X, r)
	if err != nil {
		return err
	}

	return r.SetPOBox(v)
}

func (r Sponsor) POBoxGetFunc(getfunc GetFunc) (any, error) {
	if getfunc == nil {
		return nil, errorf("Nil %T provided; aborting", getfunc)
	}

	return getfunc(r.R_POBox, r)
}

/*
PostalAddress returns the postal address value assigned to the receiver.
*/
func (r FirstAuthority) PostalAddress() string {
	return r.R_PAddr
}

/*
SetPostalAddress assigns the provided string value to the receiver's
R_PAddr field.
*/
func (r *FirstAuthority) SetPostalAddress(X any, setfunc ...SetFunc) error {
	if len(setfunc) == 0 {
		if assert, ok := X.(string); ok {
			r.R_PAddr = assert
			return nil
		}
		return errorf("Unsupported PostalAddress type %T provided without SetFunc instance", X)
	}

	v, err := setfunc[0](X, r)
	if err != nil {
		return err
	}

	return r.SetPostalAddress(v)
}

func (r FirstAuthority) PostalAddressGetFunc(getfunc GetFunc) (any, error) {
	if getfunc == nil {
		return nil, errorf("Nil %T provided; aborting", getfunc)
	}

	return getfunc(r.R_PAddr, r)
}

/*
PostalAddress returns the postal address value assigned to the receiver.
*/
func (r CurrentAuthority) PostalAddress() string {
	return r.R_PAddr
}

/*
SetPostalAddress assigns the provided string value to the receiver's
R_PAddr field.
*/
func (r *CurrentAuthority) SetPostalAddress(X any, setfunc ...SetFunc) error {
	if len(setfunc) == 0 {
		if assert, ok := X.(string); ok {
			r.R_PAddr = assert
			return nil
		}
		return errorf("Unsupported PostalAddress type %T provided without SetFunc instance", X)
	}

	v, err := setfunc[0](X, r)
	if err != nil {
		return err
	}

	return r.SetPostalAddress(v)
}

func (r CurrentAuthority) PostalAddressGetFunc(getfunc GetFunc) (any, error) {
	if getfunc == nil {
		return nil, errorf("Nil %T provided; aborting", getfunc)
	}

	return getfunc(r.R_PAddr, r)
}

/*
PostalAddress returns the postal address value assigned to the receiver.
*/
func (r Sponsor) PostalAddress() string {
	return r.R_PAddr
}

/*
SetPostalAddress assigns the provided string value to the receiver's
R_PAddr field.
*/
func (r *Sponsor) SetPostalAddress(X any, setfunc ...SetFunc) error {
	if len(setfunc) == 0 {
		if assert, ok := X.(string); ok {
			r.R_PAddr = assert
			return nil
		}
		return errorf("Unsupported PostalAddress type %T provided without SetFunc instance", X)
	}

	v, err := setfunc[0](X, r)
	if err != nil {
		return err
	}

	return r.SetPostalAddress(v)
}

func (r Sponsor) PostalAddressGetFunc(getfunc GetFunc) (any, error) {
	if getfunc == nil {
		return nil, errorf("Nil %T provided; aborting", getfunc)
	}

	return getfunc(r.R_PAddr, r)
}

/*
PostalCode returns the postal code value assigned to the receiver.
*/
func (r FirstAuthority) PostalCode() string {
	return r.R_PCode
}

/*
SetPostalCode assigns the provided string value to the receiver's
R_PCode field.
*/
func (r *FirstAuthority) SetPostalCode(X any, setfunc ...SetFunc) error {
	if len(setfunc) == 0 {
		if assert, ok := X.(string); ok {
			r.R_PCode = assert
			return nil
		}
		return errorf("Unsupported PostalCode type %T provided without SetFunc instance", X)
	}

	v, err := setfunc[0](X, r)
	if err != nil {
		return err
	}

	return r.SetPostalCode(v)
}

func (r FirstAuthority) PostalCodeGetFunc(getfunc GetFunc) (any, error) {
	if getfunc == nil {
		return nil, errorf("Nil %T provided; aborting", getfunc)
	}

	return getfunc(r.R_PCode, r)
}

/*
PostalCode returns the postal code value assigned to the receiver.
*/
func (r CurrentAuthority) PostalCode() string {
	return r.R_PCode
}

/*
SetPostalCode assigns the provided string value to the receiver's
R_PCode field.
*/
func (r *CurrentAuthority) SetPostalCode(X any, setfunc ...SetFunc) error {
	if len(setfunc) == 0 {
		if assert, ok := X.(string); ok {
			r.R_PCode = assert
			return nil
		}
		return errorf("Unsupported PostalCode type %T provided without SetFunc instance", X)
	}

	v, err := setfunc[0](X, r)
	if err != nil {
		return err
	}

	return r.SetPostalCode(v)
}

func (r CurrentAuthority) PostalCodeGetFunc(getfunc GetFunc) (any, error) {
	if getfunc == nil {
		return nil, errorf("Nil %T provided; aborting", getfunc)
	}

	return getfunc(r.R_PCode, r)
}

/*
PostalCode returns the postal code value assigned to the receiver.
*/
func (r Sponsor) PostalCode() string {
	return r.R_PCode
}

/*
SetPostalCode assigns the provided string value to the receiver's
R_PCode field.
*/
func (r *Sponsor) SetPostalCode(X any, setfunc ...SetFunc) error {
	if len(setfunc) == 0 {
		if assert, ok := X.(string); ok {
			r.R_PCode = assert
			return nil
		}
		return errorf("Unsupported PostalCode type %T provided without SetFunc instance", X)
	}

	v, err := setfunc[0](X, r)
	if err != nil {
		return err
	}

	return r.SetPostalCode(v)
}

func (r Sponsor) PostalCodeGetFunc(getfunc GetFunc) (any, error) {
	if getfunc == nil {
		return nil, errorf("Nil %T provided; aborting", getfunc)
	}

	return getfunc(r.R_PCode, r)
}

/*
Mobile returns the mobile telephone number value assigned to the receiver.
*/
func (r FirstAuthority) Mobile() string {
	return r.R_Mobile
}

/*
SetMobile assigns the provided string value to the receiver's
R_Mobile field.
*/
func (r *FirstAuthority) SetMobile(X any, setfunc ...SetFunc) error {
	if len(setfunc) == 0 {
		if assert, ok := X.(string); ok {
			r.R_Mobile = assert
			return nil
		}
		return errorf("Unsupported Mobile type %T provided without SetFunc instance", X)
	}

	v, err := setfunc[0](X, r)
	if err != nil {
		return err
	}

	return r.SetMobile(v)
}

func (r FirstAuthority) MobileGetFunc(getfunc GetFunc) (any, error) {
	if getfunc == nil {
		return nil, errorf("Nil %T provided; aborting", getfunc)
	}

	return getfunc(r.R_Mobile, r)
}

/*
Mobile returns the mobile telephone number value assigned to the receiver.
*/
func (r CurrentAuthority) Mobile() string {
	return r.R_Mobile
}

/*
SetMobile assigns the provided string value to the receiver's
R_Mobile field.
*/
func (r *CurrentAuthority) SetMobile(X any, setfunc ...SetFunc) error {
	if len(setfunc) == 0 {
		if assert, ok := X.(string); ok {
			r.R_Mobile = assert
			return nil
		}
		return errorf("Unsupported Mobile type %T provided without SetFunc instance", X)
	}

	v, err := setfunc[0](X, r)
	if err != nil {
		return err
	}

	return r.SetMobile(v)
}

func (r CurrentAuthority) MobileGetFunc(getfunc GetFunc) (any, error) {
	if getfunc == nil {
		return nil, errorf("Nil %T provided; aborting", getfunc)
	}

	return getfunc(r.R_Mobile, r)
}

/*
Mobile returns the mobile telephone number value assigned to the receiver.
*/
func (r Sponsor) Mobile() string {
	return r.R_Mobile
}

/*
SetMobile assigns the provided string value to the receiver's
R_Mobile field.
*/
func (r *Sponsor) SetMobile(X any, setfunc ...SetFunc) error {
	if len(setfunc) == 0 {
		if assert, ok := X.(string); ok {
			r.R_Mobile = assert
			return nil
		}
		return errorf("Unsupported Mobile type %T provided without SetFunc instance", X)
	}

	v, err := setfunc[0](X, r)
	if err != nil {
		return err
	}

	return r.SetMobile(v)
}

func (r Sponsor) MobileGetFunc(getfunc GetFunc) (any, error) {
	if getfunc == nil {
		return nil, errorf("Nil %T provided; aborting", getfunc)
	}

	return getfunc(r.R_Mobile, r)
}

/*
Street returns the street value assigned to the receiver.
*/
func (r FirstAuthority) Street() string {
	return r.R_Street
}

/*
SetStreet assigns the provided string value to the receiver's
R_Street field.
*/
func (r *FirstAuthority) SetStreet(X any, setfunc ...SetFunc) error {
	if len(setfunc) == 0 {
		if assert, ok := X.(string); ok {
			r.R_Street = assert
			return nil
		}
		return errorf("Unsupported Street type %T provided without SetFunc instance", X)
	}

	v, err := setfunc[0](X, r)
	if err != nil {
		return err
	}

	return r.SetStreet(v)
}

func (r FirstAuthority) StreetGetFunc(getfunc GetFunc) (any, error) {
	if getfunc == nil {
		return nil, errorf("Nil %T provided; aborting", getfunc)
	}

	return getfunc(r.R_Street, r)
}

/*
Street returns the street value assigned to the receiver.
*/
func (r CurrentAuthority) Street() string {
	return r.R_Street
}

/*
SetStreet assigns the provided string value to the receiver's
R_Street field.
*/
func (r *CurrentAuthority) SetStreet(X any, setfunc ...SetFunc) error {
	if len(setfunc) == 0 {
		if assert, ok := X.(string); ok {
			r.R_Street = assert
			return nil
		}
		return errorf("Unsupported Street type %T provided without SetFunc instance", X)
	}

	v, err := setfunc[0](X, r)
	if err != nil {
		return err
	}

	return r.SetStreet(v)
}

func (r CurrentAuthority) StreetGetFunc(getfunc GetFunc) (any, error) {
	if getfunc == nil {
		return nil, errorf("Nil %T provided; aborting", getfunc)
	}

	return getfunc(r.R_Street, r)
}

/*
Street returns the street value assigned to the receiver.
*/
func (r Sponsor) Street() string {
	return r.R_Street
}

/*
SetStreet assigns the provided string value to the receiver's
R_Street field.
*/
func (r *Sponsor) SetStreet(X any, setfunc ...SetFunc) error {
	if len(setfunc) == 0 {
		if assert, ok := X.(string); ok {
			r.R_Street = assert
			return nil
		}
		return errorf("Unsupported Street type %T provided without SetFunc instance", X)
	}

	v, err := setfunc[0](X, r)
	if err != nil {
		return err
	}

	return r.SetStreet(v)
}

func (r Sponsor) StreetGetFunc(getfunc GetFunc) (any, error) {
	if getfunc == nil {
		return nil, errorf("Nil %T provided; aborting", getfunc)
	}

	return getfunc(r.R_Street, r)
}

/*
URI returns slices of string URIs assigned to the receiver.
*/
func (r FirstAuthority) URI() []string {
	return r.R_URI
}

/*
SetURI appends one or more string slice value to the receiver's
R_URI field.
*/
func (r *FirstAuthority) SetURI(X any, setfunc ...SetFunc) error {
	if len(setfunc) == 0 {
		if assert, ok := X.(string); ok {
			r.R_URI = append(r.R_URI, assert) // append
			return nil
		} else if asserts, oks := X.([]string); oks {
			r.R_URI = asserts // clobber
			return nil
		}
		return errorf("Unsupported URI type %T provided without SetFunc instance", X)
	}

	v, err := setfunc[0](X, r)
	if err != nil {
		return err
	}

	return r.SetURI(v)
}

func (r FirstAuthority) URIGetFunc(getfunc GetFunc) (any, error) {
	if getfunc == nil {
		return nil, errorf("Nil %T provided; aborting", getfunc)
	}

	return getfunc(r.R_URI, r)
}

/*
URI returns slices of string URIs assigned to the receiver.
*/
func (r CurrentAuthority) URI() []string {
	return r.R_URI
}

/*
SetURI appends one or more string slice value to the receiver's
R_URI field.
*/
func (r *CurrentAuthority) SetURI(X any, setfunc ...SetFunc) error {
	if len(setfunc) == 0 {
		if assert, ok := X.(string); ok {
			r.R_URI = append(r.R_URI, assert) // append
			return nil
		} else if asserts, oks := X.([]string); oks {
			r.R_URI = asserts // clobber
			return nil
		}
		return errorf("Unsupported URI type %T provided without SetFunc instance", X)
	}

	v, err := setfunc[0](X, r)
	if err != nil {
		return err
	}

	return r.SetURI(v)
}

func (r CurrentAuthority) URIGetFunc(getfunc GetFunc) (any, error) {
	if getfunc == nil {
		return nil, errorf("Nil %T provided; aborting", getfunc)
	}

	return getfunc(r.R_URI, r)
}

/*
URI returns slices of string URIs assigned to the receiver.
*/
func (r Sponsor) URI() []string {
	return r.R_URI
}

/*
SetURI appends one or more string slice value to the receiver's
R_URI field.
*/
func (r *Sponsor) SetURI(X any, setfunc ...SetFunc) error {
	if len(setfunc) == 0 {
		if assert, ok := X.(string); ok {
			r.R_URI = append(r.R_URI, assert) // append
			return nil
		} else if asserts, oks := X.([]string); oks {
			r.R_URI = asserts // clobber
			return nil
		}
		return errorf("Unsupported URI type %T provided without SetFunc instance", X)
	}

	v, err := setfunc[0](X, r)
	if err != nil {
		return err
	}

	return r.SetURI(v)
}

func (r Sponsor) URIGetFunc(getfunc GetFunc) (any, error) {
	if getfunc == nil {
		return nil, errorf("Nil %T provided; aborting", getfunc)
	}

	return getfunc(r.R_URI, r)
}

/*
StartTime returns the string-based generalizedTime value that reflects the time
at which the receiver was (or will be) officiated.
*/
func (r FirstAuthority) StartTime() string {
	return r.R_StartTime
}

/*
SetStartTime appends one or more string slice value to the receiver's
R_StartTime field.
*/
func (r *FirstAuthority) SetStartTime(X any, setfunc ...SetFunc) error {
	if len(setfunc) == 0 {
		if assert, ok := X.(string); ok {
			r.R_StartTime = assert
			return nil
		}
		return errorf("Unsupported StartTime type %T provided without SetFunc instance", X)
	}

	v, err := setfunc[0](X, r)
	if err != nil {
		return err
	}

	return r.SetStartTime(v)
}

/*
StartTime returns the string-based generalizedTime value that reflects the time
at which the receiver was (or will be) officiated.
*/
func (r CurrentAuthority) StartTime() string {
	return r.R_StartTime
}

/*
SetStartTime appends one or more string slice value to the receiver's
R_StartTime field.
*/
func (r *CurrentAuthority) SetStartTime(X any, setfunc ...SetFunc) error {
	if len(setfunc) == 0 {
		if assert, ok := X.(string); ok {
			r.R_StartTime = assert
			return nil
		}
		return errorf("Unsupported StartTime type %T provided without SetFunc instance", X)
	}

	v, err := setfunc[0](X, r)
	if err != nil {
		return err
	}

	return r.SetStartTime(v)
}

/*
StartTime returns the string-based generalizedTime value that reflects
the time at which the receiver was (or will be) officiated.
*/
func (r Sponsor) StartTime() string {
	return r.R_StartTime
}

/*
SetStartTime appends one or more string slice value to the receiver's
R_StartTime field.
*/
func (r *Sponsor) SetStartTime(X any, setfunc ...SetFunc) error {
	if len(setfunc) == 0 {
		if assert, ok := X.(string); ok {
			r.R_StartTime = assert
			return nil
		}
		return errorf("Unsupported StartTime type %T provided without SetFunc instance", X)
	}

	v, err := setfunc[0](X, r)
	if err != nil {
		return err
	}

	return r.SetStartTime(v)
}

/*
EndTime returns the string-based generalizedTime value that reflects the time
at which the receiver was (or will be) terminated.
*/
func (r FirstAuthority) EndTime() string {
	return r.R_EndTime
}

/*
SetRegistrantID assigns the provided string value to the
receiver's R_Id field.
*/
func (r *FirstAuthority) SetRegistrantID(X any, setfunc ...SetFunc) error {
	if len(setfunc) == 0 {
		if assert, ok := X.(string); ok {
			r.R_Id = assert
			return nil
		}
		return errorf("Unsupported RegistrantID type %T provided without SetFunc instance", X)
	}

	v, err := setfunc[0](X, r)
	if err != nil {
		return err
	}

	return r.SetRegistrantID(v)
}

/*
SetEndTime appends one or more string slice value to the receiver's
R_EndTime field.
*/
func (r *FirstAuthority) SetEndTime(X any, setfunc ...SetFunc) error {
	if len(setfunc) == 0 {
		if assert, ok := X.(string); ok {
			r.R_EndTime = assert
			return nil
		}
		return errorf("Unsupported EndTime type %T provided without SetFunc instance", X)
	}

	v, err := setfunc[0](X, r)
	if err != nil {
		return err
	}

	return r.SetEndTime(v)
}

/*
EndTime returns a bogus generalizedTime value. A current authority registrant, because it is
active, has no end time. This method exists for this type (CurrentAuthority) ONLY TO
satisfy go interface requirements and is thoroughly useless.
*/
func (r CurrentAuthority) EndTime() string {
	return ``
}

/*
SetRegistrantID assigns the provided string value to the
receiver's R_Id field.
*/
func (r *CurrentAuthority) SetRegistrantID(X any, setfunc ...SetFunc) error {
	if len(setfunc) == 0 {
		if assert, ok := X.(string); ok {
			r.R_Id = assert
			return nil
		}
		return errorf("Unsupported RegistrantID type %T provided without SetFunc instance", X)
	}

	v, err := setfunc[0](X, r)
	if err != nil {
		return err
	}

	return r.SetRegistrantID(v)
}

/*
SetEndTime performs no useful task, as instances of *CurrentAuthority do not have
an R_EndTime field.

See the Relegate function for cases where an instance of *CurrentAuthority needs
to be "converted" into an instance of *FirstAuthority, which would then make the
R_EndTime field available.
*/
func (r *CurrentAuthority) SetEndTime(X any, setfunc ...SetFunc) error {
	return errorf("EndTime is not applicable to %T", r)
}

/*
EndTime returns the string-based generalizedTime value that reflects the time
at which the receiver was (or will be) terminated.
*/
func (r Sponsor) EndTime() string {
	return r.R_EndTime
}

/*
SetRegistrantID assigns the provided string value to the
receiver's R_Id field.
*/
func (r *Sponsor) SetRegistrantID(X any, setfunc ...SetFunc) error {
	if len(setfunc) == 0 {
		if assert, ok := X.(string); ok {
			r.R_Id = assert
			return nil
		}
		return errorf("Unsupported RegistrantID type %T provided without SetFunc instance", X)
	}

	v, err := setfunc[0](X, r)
	if err != nil {
		return err
	}

	return r.SetRegistrantID(v)
}

/*
SetEndTime appends one or more string slice value to the receiver's
R_EndTime field.
*/
func (r *Sponsor) SetEndTime(X any, setfunc ...SetFunc) error {
	if len(setfunc) == 0 {
		if assert, ok := X.(string); ok {
			r.R_EndTime = assert
			return nil
		}
		return errorf("Unsupported EndTime type %T provided without SetFunc instance", X)
	}

	v, err := setfunc[0](X, r)
	if err != nil {
		return err
	}

	return r.SetEndTime(v)
}

/*
ObjectClass returns the static string value `x660Registrant` merely as an alternative
to tested type assertion.
*/
func (r FirstAuthority) ObjectClass() string {
	return `x660Registrant`
}

/*
ObjectClass returns the static string value `x660Registrant` merely as an alternative
to tested type assertion.
*/
func (r CurrentAuthority) ObjectClass() string {
	return `x660Registrant`
}

/*
ObjectClass returns the static string value `x660Registrant` merely as an alternative
to tested type assertion.
*/
func (r Sponsor) ObjectClass() string {
	return `x660Registrant`
}

/*
Type returns the static string value `firstAuthority` as a convenient means of testing
which kind of Registrant type is described by the receiver (FirstAuthority).
*/
func (r FirstAuthority) Type() string {
	return `firstAuthority`
}

/*
Type returns the static string value `currentAuthority` as a convenient means of testing
which kind of Registrant type is described by the receiver (CurrentAuthority).
*/
func (r CurrentAuthority) Type() string {
	return `currentAuthority`
}

/*
Type returns the static string value `sponsor` as a convenient means of testing which
kind of Registrant type is described by the receiver (Sponsor).
*/
func (r Sponsor) Type() string {
	return `sponsor`
}

/*
Kind returns the static string value `AUXILIARY` merely as a convenient means
to declare what kind of objectClass the Registrant receiver describes.
*/
func (r FirstAuthority) Kind() string {
	return `AUXILIARY`
}

/*
Kind returns the static string value `AUXILIARY` merely as a convenient means
to declare what kind of objectClass the Registrant receiver describes.
*/
func (r CurrentAuthority) Kind() string {
	return `AUXILIARY`
}

/*
Kind returns the static string value `AUXILIARY` merely as a convenient means
to declare what kind of objectClass the Registrant receiver describes.
*/
func (r Sponsor) Kind() string {
	return `AUXILIARY`
}

/*
Unmarshal transports values from the receiver into an instance
of map[string][]string, which can subsequently be fed to go-ldap's
NewEntry function.
*/
func (r FirstAuthority) Unmarshal() map[string][]string {
	return toMap(r)
}

/*
Unmarshal transports values from the receiver into an instance
of map[string][]string, which can subsequently be fed to go-ldap's
NewEntry function.
*/
func (r CurrentAuthority) Unmarshal() map[string][]string {
	return toMap(r)
}

/*
Unmarshal transports values from the receiver into an instance
of map[string][]string, which can subsequently be fed to go-ldap's
NewEntry function.
*/
func (r Sponsor) Unmarshal() map[string][]string {
	return toMap(r)
}

/*
Valid returns a boolean value indicative of whether the receiver configuration
instance is considered contextually valid and usable.
*/
func (r *DUAConfig) Valid() (valid bool) {
	if r == nil {
		return
	}

	switch r.DirectoryModel {
	case TwoDimensional, ThreeDimensional:
		// ok
	default:
		return
	}

	valid = len(r.Registrations)+len(r.Registrants) > 0
	return
}

/*
NewDUAConfig returns an initialized instance of *DUAConfig containing
an initialized map[string][]string Settings value.
*/
func NewDUAConfig() *DUAConfig {
	return &DUAConfig{Settings: make(map[string][]string, 0)}
}

/*
SetDUAConfig assigns the input *DUAConfig (d) to the receiver's
R_DUAConfig struct field.
*/
func (r *RootArc) SetDUAConfig(d *DUAConfig) {
	r.R_DUAConfig = d
}

/*
DUAConfig returns the *DUAConfig instance assigned to the
receiver, if set, else nil is returned.
*/
func (r RootArc) DUAConfig() *DUAConfig {
	return r.R_DUAConfig
}

/*
SetDUAConfig assigns the input *DUAConfig (d) to the receiver's
R_DUAConfig struct field.
*/
func (r *SubArc) SetDUAConfig(d *DUAConfig) {
	r.R_DUAConfig = d
}

/*
DUAConfig returns the *DUAConfig instance assigned to the
receiver, if set, else nil is returned.
*/
func (r SubArc) DUAConfig() *DUAConfig {
	return r.R_DUAConfig
}

/*
SetDUAConfig assigns the input *DUAConfig (d) to the receiver's
R_DUAConfig struct field.
*/
func (r *FirstAuthority) SetDUAConfig(d *DUAConfig) {
	r.R_DUAConfig = d
}

/*
DUAConfig returns the *DUAConfig instance assigned to the
receiver, if set, else nil is returned.
*/
func (r FirstAuthority) DUAConfig() *DUAConfig {
	return r.R_DUAConfig
}

/*
SetDUAConfig assigns the input *DUAConfig (d) to the receiver's
R_DUAConfig struct field.
*/
func (r *CurrentAuthority) SetDUAConfig(d *DUAConfig) {
	r.R_DUAConfig = d
}

/*
DUAConfig returns the *DUAConfig instance assigned to the
receiver, if set, else nil is returned.
*/
func (r CurrentAuthority) DUAConfig() *DUAConfig {
	return r.R_DUAConfig
}

/*
SetDUAConfig assigns the input *DUAConfig (d) to the receiver's
R_DUAConfig struct field.
*/
func (r *Sponsor) SetDUAConfig(d *DUAConfig) {
	r.R_DUAConfig = d
}

/*
DUAConfig returns the *DUAConfig instance assigned to the
receiver, if set, else nil is returned.
*/
func (r Sponsor) DUAConfig() *DUAConfig {
	return r.R_DUAConfig
}
