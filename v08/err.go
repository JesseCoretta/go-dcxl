package dcxl

import "errors"

/*
err.go contains predefined error instances that
describe certain known aberrant conditions.
*/

var (
	RegistrationValidityErr,
	UnsupportedInputTypeErr,
	IllegalASN1NotationErr,
	RegistrantValidityErr,
	DUAConfigValidityErr,
	IllegalNumberFormErr,
	InvalidDimensionErr,
	NilRegistrationErr,
	IllegalLongArcErr,
	MismatchedLeafErr,
	NilRegistrantErr,
	IllegalRootErr,
	InvalidOIDErr,
	InvalidDNErr error
)

func init() {
	RegistrationValidityErr = errorf("Registration instance did not pass validity checks")
	UnsupportedInputTypeErr = errorf("Unsupported input type")
	IllegalASN1NotationErr = errorf("ASN.1 Notation value is malformed or zero-length")
	RegistrantValidityErr = errorf("Registrant instance did not pass validity checks")
	DUAConfigValidityErr = errorf("DUAConfig instance did not pass validity checks")
	IllegalNumberFormErr = errorf("N (Number Form) is malformed or zero length")
	InvalidDimensionErr = errorf("Unknown dimension; must be TwoDimensional or ThreeDimensional")
	NilRegistrationErr = errorf("Registration instance is nil")
	MismatchedLeafErr = errorf("Mismatched NumberForm with leaf node of ASN.1 and/or DotNotation")
	IllegalLongArcErr = errorf("LongArc cannot be applied to this registration type or root")
	NilRegistrantErr = errorf("Registration instance is nil")
	IllegalRootErr = errorf("Illegal root (must be 0, 1 or 2)")
	InvalidOIDErr = errorf("OID value is malformed or zero length")
	InvalidDNErr = errorf("DN value is malformed or zero length")
}

func errorf(msg any, x ...any) error {
	switch tv := msg.(type) {
	case string:
		if len(tv) > 0 {
			return errors.New(sprintf(tv, x...))
		}
	case error:
		if tv != nil {
			return errors.New(sprintf(tv.Error(), x...))
		}
	}

	return nil
}
