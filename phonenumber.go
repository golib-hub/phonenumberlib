// Package phonenumberlib provides version-independent phone-number parsing.
package phonenumberlib

import (
	"fmt"

	"github.com/golib-hub/phonenumberlib/entity"
	"github.com/nyaruka/phonenumbers/v2"
)

// Version identifies an embedded upstream phonenumbers release.
type Version string

const (
	// VersionV2_0_11 uses the embedded behavior and metadata from v2.0.11.
	VersionV2_0_11 Version = "v2.0.11"
)

// UnsupportedVersionError reports a version that is not embedded in this build.
type UnsupportedVersionError struct {
	Version Version
}

func (err *UnsupportedVersionError) Error() string {
	return fmt.Sprintf("phonenumberlib: unsupported version %q", err.Version)
}

// Parse parses input using region when input does not include a country calling
// code. Region must be an ISO 3166-1 alpha-2 country code, such as "US". The
// version selects the embedded upstream parsing behavior.
//
// Every version adapter returns the same entity.PhoneNumber type.
func Parse(input, region string, version Version) (*entity.PhoneNumber, error) {
	if version != VersionV2_0_11 {
		return nil, &UnsupportedVersionError{Version: version}
	}

	number, err := phonenumbers.Parse(input, region)
	if err != nil {
		return nil, err
	}

	return toPhoneNumber(number), nil
}

func toPhoneNumber(number *phonenumbers.PhoneNumber) *entity.PhoneNumber {
	if number == nil {
		return nil
	}

	result := &entity.PhoneNumber{
		CountryCode:                  number.CountryCode,
		NationalNumber:               number.NationalNumber,
		Extension:                    number.Extension,
		ItalianLeadingZero:           number.ItalianLeadingZero,
		NumberOfLeadingZeros:         number.NumberOfLeadingZeros,
		RawInput:                     number.RawInput,
		PreferredDomesticCarrierCode: number.PreferredDomesticCarrierCode,
	}
	if number.CountryCodeSource != nil {
		source := entity.CountryCodeSource(*number.CountryCodeSource)
		result.CountryCodeSource = &source
	}

	return result
}
