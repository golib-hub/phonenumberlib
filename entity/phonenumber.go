// Package entity contains stable, version-independent public data types.
package entity

// CountryCodeSource identifies how a parsed number's country code was derived.
// Its values mirror libphonenumber's stable phone-number protocol fields.
type CountryCodeSource int32

const (
	CountryCodeSourceUnspecified               CountryCodeSource = 0
	CountryCodeSourceFromNumberWithPlusSign    CountryCodeSource = 1
	CountryCodeSourceFromNumberWithIDD         CountryCodeSource = 5
	CountryCodeSourceFromNumberWithoutPlusSign CountryCodeSource = 10
	CountryCodeSourceFromDefaultCountry        CountryCodeSource = 20
)

// PhoneNumber is the stable, version-independent representation returned by
// this library. It intentionally contains no upstream implementation value.
//
// Pointer fields preserve the presence semantics of libphonenumber's
// PhoneNumber protocol message without making this type a protobuf message.
type PhoneNumber struct {
	CountryCode                  *int32             `json:"country_code,omitempty"`
	NationalNumber               *uint64            `json:"national_number,omitempty"`
	Extension                    *string            `json:"extension,omitempty"`
	ItalianLeadingZero           *bool              `json:"italian_leading_zero,omitempty"`
	NumberOfLeadingZeros         *int32             `json:"number_of_leading_zeros,omitempty"`
	RawInput                     *string            `json:"raw_input,omitempty"`
	CountryCodeSource            *CountryCodeSource `json:"country_code_source,omitempty"`
	PreferredDomesticCarrierCode *string            `json:"preferred_domestic_carrier_code,omitempty"`
}

// GetCountryCode returns the country calling code, or zero when absent.
func (number *PhoneNumber) GetCountryCode() int32 {
	if number != nil && number.CountryCode != nil {
		return *number.CountryCode
	}
	return 0
}

// GetNationalNumber returns the national significant number, or zero when absent.
func (number *PhoneNumber) GetNationalNumber() uint64 {
	if number != nil && number.NationalNumber != nil {
		return *number.NationalNumber
	}
	return 0
}

// GetExtension returns the extension, or an empty string when absent.
func (number *PhoneNumber) GetExtension() string {
	if number != nil && number.Extension != nil {
		return *number.Extension
	}
	return ""
}

// GetItalianLeadingZero reports whether the number retains an Italian leading zero.
func (number *PhoneNumber) GetItalianLeadingZero() bool {
	return number != nil && number.ItalianLeadingZero != nil && *number.ItalianLeadingZero
}

// GetNumberOfLeadingZeros returns the number of retained leading zeros.
func (number *PhoneNumber) GetNumberOfLeadingZeros() int32 {
	if number != nil && number.NumberOfLeadingZeros != nil {
		return *number.NumberOfLeadingZeros
	}
	return 1
}

// GetRawInput returns the original input when it was retained.
func (number *PhoneNumber) GetRawInput() string {
	if number != nil && number.RawInput != nil {
		return *number.RawInput
	}
	return ""
}

// GetCountryCodeSource returns how the country code was derived.
func (number *PhoneNumber) GetCountryCodeSource() CountryCodeSource {
	if number != nil && number.CountryCodeSource != nil {
		return *number.CountryCodeSource
	}
	return CountryCodeSourceUnspecified
}

// GetPreferredDomesticCarrierCode returns the preferred domestic carrier code.
func (number *PhoneNumber) GetPreferredDomesticCarrierCode() string {
	if number != nil && number.PreferredDomesticCarrierCode != nil {
		return *number.PreferredDomesticCarrierCode
	}
	return ""
}
