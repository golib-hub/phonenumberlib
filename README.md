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
import (
	"github.com/golib-hub/phonenumberlib"
	"github.com/nyaruka/phonenumbers/v2"
)

// Parse a national number using its ISO 3166-1 alpha-2 region.
num, err := phonenumberlib.Parse("8886418722", "US")
if err != nil {
	// Handle invalid input or an unknown region.
}

valid := phonenumbers.IsValidNumber(num)
formatted := phonenumbers.Format(num, phonenumbers.NATIONAL)
```

The wrapper exposes the upstream `PhoneNumber` type, so formatting and
validation use the upstream package directly.
