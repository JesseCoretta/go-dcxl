package dcxl

import (
	"fmt"
	"time"
)

func ExampleRootArc_ModifyTime() {
	var X *SubArc = new(SubArc)
	X.SetDN(`n=11,n=1,n=2,n=101,n=56521,n=1,n=4,n=1,n=6,n=3,n=1,ou=Registrations,o=rA`)
	X.SetN(`11`)
	X.SetASN1Notation(`{iso(1) org(3) dod(6) internet(1) private(4) enterprise(1) 56521 101 2 1 11}`)
	X.SetURI(`http://ldap.example.com`)
	X.SetURI(`http://ra.example.com`)
	X.SetModifyTime(`19711103040901Z`)
	X.SetModifyTime(`19860110120110Z`)
	X.SetModifyTime(`20080414090555Z`)

	fmt.Printf("%s\n", X.ModifyTime()[0])
	// Output: 19711103040901Z
}

func ExampleSubArc_ModifyTime() {
	var X *SubArc = new(SubArc)
	X.SetDN(`n=11,n=1,n=2,n=101,n=56521,n=1,n=4,n=1,n=6,n=3,n=1,ou=Registrations,o=rA`)
	X.SetN(`11`)
	X.SetASN1Notation(`{iso(1) org(3) dod(6) internet(1) private(4) enterprise(1) 56521 101 2 1 11}`)
	X.SetDotNotation(`1.3.6.1.4.1.56521.2.1.11`)
	X.SetModifyTime(`19711103040901Z`)
	X.SetModifyTime(`19860110120110Z`)
	X.SetModifyTime(`20080414090555Z`)

	fmt.Printf("%s\n", X.ModifyTime()[0])
	// Output: 19711103040901Z
}

func ExampleSubArc_CreateTime() {
	var X *SubArc = new(SubArc)
	X.SetDN(`n=11,n=1,n=2,n=101,n=56521,n=1,n=4,n=1,n=6,n=3,n=1,ou=Registrations,o=rA`)
	X.SetN(`11`)
	X.SetASN1Notation(`{iso(1) org(3) dod(6) internet(1) private(4) enterprise(1) 56521 101 2 1 11}`)
	X.SetDotNotation(`1.3.6.1.4.1.56521.2.1.11`)
	X.SetCreateTime(`19860110120110Z`)

	fmt.Printf("%s\n", X.CreateTime())
	// Output: 19860110120110Z
}

func ExampleRootArc_ObjectClass() {
	var X *RootArc = new(RootArc)
	X.SetDN(`n=1,ou=Registrations,o=rA`)
	X.SetN(`1`)
	X.SetUnicodeValue(`ISO`)
	X.SetASN1Notation(`{iso(1)}`)

	fmt.Printf("%s\n", X.ObjectClass())
	// Output: x660RootArc
}

func ExampleSubArc_ObjectClass() {
	var X *SubArc = new(SubArc)
	X.SetDN(`n=11,n=1,n=2,n=101,n=56521,n=1,n=4,n=1,n=6,n=3,n=1,ou=Registrations,o=rA`)
	X.SetN(`11`)
	X.SetASN1Notation(`{iso(1) org(3) dod(6) internet(1) private(4) enterprise(1) 56521 101 2 1 11}`)
	X.SetDotNotation(`1.3.6.1.4.1.56521.2.1.11`)
	X.SetCreateTime(`19860110120110Z`)

	fmt.Printf("%s\n", X.ObjectClass())
	// Output: x660SubArc
}

func ExampleRootArc_Kind() {
	var X *RootArc = new(RootArc)
	X.SetDN(`n=1,ou=Registrations,o=rA`)
	X.SetN(`1`)
	X.SetUnicodeValue(`ISO`)
	X.SetASN1Notation(`{iso(1)}`)

	fmt.Printf("%s\n", X.Kind())
	// Output: STRUCTURAL
}

func ExampleSubArc_Kind() {
	var X *SubArc = new(SubArc)
	X.SetDN(`n=11,n=1,n=2,n=101,n=56521,n=1,n=4,n=1,n=6,n=3,n=1,ou=Registrations,o=rA`)
	X.SetN(`11`)
	X.SetASN1Notation(`{iso(1) org(3) dod(6) internet(1) private(4) enterprise(1) 56521 101 2 1 11}`)
	X.SetDotNotation(`1.3.6.1.4.1.56521.2.1.11`)
	X.SetCreateTime(`19860110120110Z`)

	fmt.Printf("%s\n", X.Kind())
	// Output: STRUCTURAL
}

func ExampleSubArc_SetDotNotation_withSetFunc3D() {
	var X *SubArc = new(SubArc)
	//var X Registration = new(SubArc)
	X.SetDN(`n=11,n=1,n=2,n=101,n=56521,n=1,n=4,n=1,n=6,n=3,n=1,ou=Registrations,o=rA`)
	X.SetDUAConfig(&DUAConfig{
		Registrations:  []string{`ou=Registrations,o=rA`},
		DirectoryModel: ThreeDimensional,
	})

	if err := X.SetDotNotation(X.DN(), DNToDotNot3D); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%s\n", X.DotNotation())
	// Output: 1.3.6.1.4.1.56521.101.2.1.11
}

func ExampleSubArc_SetDotNotation_withSetFunc2D() {
	var X *SubArc = new(SubArc)
	//var X Registration = new(SubArc)
	X.SetDN(`dotNotation=1.3.6.1.4.1.56521.101.2.1.11,ou=Registrations,o=rA`)
	X.SetDUAConfig(&DUAConfig{
		Registrations:  []string{`ou=Registrations,o=rA`},
		DirectoryModel: TwoDimensional,
	})

	if err := X.SetDotNotation(X.DN(), DNToDotNot2D); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%s\n", X.DotNotation())
	// Output: 1.3.6.1.4.1.56521.101.2.1.11
}

func ExampleSubArc_SetDN_withSetFunc3D() {
	var X *SubArc = new(SubArc)
	X.SetDotNotation(`1.3.6.1.4.1.56521.101.2.1.11`)
	X.SetDUAConfig(&DUAConfig{
		Registrations:  []string{`ou=Registrations,o=rA`},
		DirectoryModel: ThreeDimensional,
	})

	if err := X.SetDN(X.DotNotation(), DotNotToDN3D); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%s\n", X.DN())
	// Output: n=11,n=1,n=2,n=101,n=56521,n=1,n=4,n=1,n=6,n=3,n=1,ou=Registrations,o=rA
}

func ExampleSubArc_SetDN_withSetFunc2D() {
	var X *SubArc = new(SubArc)
	X.SetDotNotation(`1.3.6.1.4.1.56521.101.2.1.11`)
	X.SetDUAConfig(&DUAConfig{
		Registrations:  []string{`ou=Registrations,o=rA`},
		DirectoryModel: TwoDimensional,
	})

	if err := X.SetDN(X.DotNotation(), DotNotToDN2D); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%s\n", X.DN())
	// Output: dotNotation=1.3.6.1.4.1.56521.101.2.1.11,ou=Registrations,o=rA
}

func ExampleRootArc_ModifyTimeGetFunc_generalizedTimeToTime() {
	var X *RootArc = new(RootArc)

	// give our instance a string val
	X.SetModifyTime(`20080114214613Z`)

	// Hand the ModifyTimeGetFunc method the
	// desired GetFunc function instance to
	// be executed.
	t, err := X.ModifyTimeGetFunc(GeneralizedTimeToTime)
	if err != nil {
		fmt.Println(err)
		return
	}

	if T, ok := t.(time.Time); ok && !T.IsZero() {
		fmt.Printf("%s", T)
	} else if Ts, oks := t.([]time.Time); oks && len(Ts) > 0 {
		fmt.Printf("%s", Ts[0])
	}
	// Output: 2008-01-14 21:46:13 +0000 UTC
}

func ExampleSubArc_dedicated() {
	var X *SubArc = new(SubArc)
	X.SetDN(`n=11,n=1,n=2,n=101,n=56521,n=1,n=4,n=1,n=6,n=3,n=1,ou=Registrations,o=rA`)
	X.SetN(`11`)
	X.SetASN1Notation(`{iso(1) org(3) dod(6) internet(1) private(4) enterprise(1) 56521 101 2 1 11}`)
	X.SetDotNotation(`1.3.6.1.4.1.56521.2.1.11`)
	X.SetCreateTime(`20210110120110Z`)
	X.SetFirstAuthority(`registrantID=Jesse Coretta,ou=Registrants,o=rA`)
	X.SetCurrentAuthority(`registrantID=ABC Authority,ou=Registrants,o=rA`)

	fmt.Printf("%s (first), %s (current)\n", X.FirstAuthority()[0], X.CurrentAuthority()[0])
	// Output: registrantID=Jesse Coretta,ou=Registrants,o=rA (first), registrantID=ABC Authority,ou=Registrants,o=rA (current)
}

func ExampleSubArc_combined() {
	var A *FirstAuthority = new(FirstAuthority)
	A.SetCN(`Mister Authority`)
	A.SetO(`Authority, Co.`)

	var X *SubArc = new(SubArc)
	X.SetDN(`n=11,n=1,n=2,n=101,n=56521,n=1,n=4,n=1,n=6,n=3,n=1,ou=Registrations,o=rA`)
	X.SetN(`11`)
	X.SetASN1Notation(`{iso(1) org(3) dod(6) internet(1) private(4) enterprise(1) 56521 101 2 1 11}`)
	X.SetDotNotation(`1.3.6.1.4.1.56521.2.1.11`)
	X.SetCreateTime(`19860110120110Z`)
	X.SetCombinedFirstAuthority(A)

	fmt.Printf("%s at %s\n", X.CombinedFirstAuthority().CN(), X.CombinedFirstAuthority().O())
	// Output: Mister Authority at Authority, Co.
}

func ExampleSubArc_Unmarshal() {
	var X *SubArc = new(SubArc)
	X.SetDN(`n=11,n=1,n=2,n=101,n=56521,n=1,n=4,n=1,n=6,n=3,n=1,ou=Registrations,o=rA`)
	X.SetN(`11`)
	X.SetASN1Notation(`{iso(1) org(3) dod(6) internet(1) private(4) enterprise(1) 56521 101 2 1 11}`)

	m := X.Unmarshal()
	fmt.Printf("%s\n", m[`asn1Notation`][0])
	// Output: {iso(1) org(3) dod(6) internet(1) private(4) enterprise(1) 56521 101 2 1 11}
}

func Registrants_Unmarshal() {
	sp := new(Sponsor)
	sp.SetDN(`registrantID=430a8727-c8b0-4734-83d7-0a104ab00a69,ou=Registrants,o=rA`)
	sp.SetCN(`Mister Sponsor`)
	sp.SetO(`Benevolent Sponsors, Inc.`)
	sp.SetTel(`+1 555 555 0134`)
	sp.SetEmail(`sponsors@example.com`)

	fa := new(FirstAuthority)
	fa.SetDN(`registrantID=90b29aac-7771-1249-01a6-143f5d4d2500,ou=Registrants,o=rA`)
	fa.SetCN(`Jesse Coretta`)
	fa.SetTel(`+1 555 555 9812`)

	var regs Registrants = Registrants{sp, fa}

	maps := regs.Unmarshal()
	fmt.Printf("len:%d\n", len(maps))
	// Output: 2
}

func ExamplegenTimeToTime() {
	t := `20080114154613-0600`
	if T, ok := genTimeToTime(t); ok {
		fmt.Printf("%s\n", T)
	}
	// Output: 2008-01-14 21:46:13 +0000 UTC
}

func ExampletimeToGenTime() {
	// Use the GenTimeToTime example for brevity.
	t := `20080114154613-0600`
	T, _ := genTimeToTime(t) // parse into time.Time

	// Use the same time, but with utc (skip
	// bool checking for brevity).
	tu := `20080114214613Z`
	Tu, _ := genTimeToTime(tu) // parse into time.Time

	// Parse the respective times back into
	// string-based generalizedTime values.
	nt, ok := timeToGenTime(T)
	ntu, oku := timeToGenTime(Tu)

	// They should be valid and identical
	if ok && oku {
		fmt.Printf("Times are valid and equal: %t\n", nt == ntu)
	}
	// Output: Times are valid and equal: true
}
