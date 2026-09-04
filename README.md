# phonenumberlib

A small Go wrapper around [nyaruka/phonenumbers](https://github.com/nyaruka/phonenumbers), a Go port of Google's libphonenumber.

## Versioning

The upstream module is versioned. Because its current major version is v2, Go
requires `/v2` in its module and import paths. This repository pins it in
`go.mod`:

```go
require github.com/nyaruka/phonenumbers/v2 v2.0.11
```

## Usage

```go
import "github.com/golib-hub/phonenumberlib"

// Parse a national number using its ISO 3166-1 alpha-2 region.
num, err := phonenumberlib.Parse("8886418722", "US", phonenumberlib.VersionV2_0_11)
if err != nil {
	// Handle invalid input or an unknown region.
}

countryCode := num.GetCountryCode()
nationalNumber := num.GetNationalNumber()
```

`Parse` returns the stable `entity.PhoneNumber` type, so callers do not depend
on an upstream version-specific phone-number type.
