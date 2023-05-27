/*
Package dcxl implements draft-coretta-x660-ldap-08, an Internet Draft of which I am the author.

# Draft Information

The Internet Draft (henceforth referred to as "the ID") can be viewed here: https://datatracker.ietf.org/doc/html/draft-coretta-x660-ldap. It is currently EXPIRED, pending review. There may or may not be an update of this ID in the near future.

A text copy of the ID is also included in the package directory.

# Advisory

No past or present version of the ID is meant to be used in any production environment and -- because it is not yet approved for an RFC track (and may never be) -- it should be considered PURELY EXPERIMENTAL and totally devoid of any official standing or acceptance.

That said, this specification COULD very well allow some nifty functionality that would be beneficial to any individual, entity or institution that deals with or manages ASN.1 object identifiers regularly, and requires a central (non-ORS) means of OID resolution / interrogation. This package exists to provide abstract access to such capabilities through the Go language.

# Early Release

This package is under heavy, active development. In addition to the above advisory, adopters should expect some unannounced volatility with regards to the state of the package.

# Versioning

This package may include subdirectories, one containing a versioned reference implementation that corresponds to the terms and precepts of the same document version. One will exist for each new release of the document if and when one is released. Previous versions will remain intact for backwards compatibility, however pull requests will not be accepted. The package contents in the top-most package directory represent the latest release, and will be identical to the highest version-numbered folder.

Pull requests are always welcome for the latest release.

# Synopsis

The ID describes a possible means for storing the OID spectrum (in whole or in part) within an X.500 directory. It is effectively an "LDAP alternative" to an ORS implementation, which would typically be DNS-based.

To that end, the ID introduces myriad attribute types and a small handful of objectClass definitions -- each facilitating the storage of information pertaining to certain abstract types defined within ITU-T X-Series X.660, et al.

# No LDAP Functionality

This package, although designed to work well in LDAP-related operations, DOES NOT import the go-ldap/ldap package. Users interacting with an X.500 Directory System Agent (DSA, a.k.a.: an "LDAP server") within the context of the ID are expected to craft their own Directory User Agent (DUA, a.k.a.: an "LDAP client").

# Features

This package offers the following:

• Total compliance with the terms of the ID -- all eighty eight (88) attributeTypes, four (4) objectClasses and two (2) dimensional abstracts defined in the ID are available via this package through any instances of the Registration and Registrant types

• Plentiful documentation with useful examples

• Complete interface support, negating the need for manual type assertion involving any Registration or Registrant types during composition and interrogation procedures

• Extensible design; user-authored "Get" and "Set" closure functions are supported (but not required) at virtually every point, allowing limitless control over the contents and presentation of relevant objects/values

• Seamless compatibility with *ldap.NewEntry using the map-populating Unmarshal method, which is extended through any Registration or Registrant type instance

• Portable configuration type (per s. 2.2.4 of the ID), allowing critical information and configuration settings to be kept within individual Registration/Registrant instances for efficient and reliable operation

# Enhanced Operation

For those dealing with ASN.1 Notation, Dot Notation, NumberForm and NameAndNumberForm values, OR for situations requiring uint128 NumberForm support, consider my objectid package (https://pkg.go.dev/github.com/JesseCoretta/go-objectid) for inclusion in any custom Set/Get functions in use. It can offer a potent advantage in scenarios relevant to this very package, and extends several very useful methods to greatly reduce the tedium (and error-prone nature) of this subject matter.

For those interested in the actual schema definitions relating to this ID, see https://github.com/JesseCoretta/draft-coretta-x660-ldap. The schemata is available for download in one (1) of three (3) distinct schema formats: OpenLDAP, ApacheDS and Netscape.
*/
package dcxl
