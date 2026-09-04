package entity

import (
	"encoding/json"
	"testing"
)

func TestPhoneNumberGetters(t *testing.T) {
	countryCode := int32(39)
	nationalNumber := uint64(212345678)
	extension := "42"
	italianLeadingZero := true
	leadingZeros := int32(2)
	rawInput := "+39 02 1234 5678"
	source := CountryCodeSourceFromNumberWithPlusSign
	carrierCode := "3"

	number := &PhoneNumber{
		CountryCode:                  &countryCode,
		NationalNumber:               &nationalNumber,
		Extension:                    &extension,
		ItalianLeadingZero:           &italianLeadingZero,
		NumberOfLeadingZeros:         &leadingZeros,
		RawInput:                     &rawInput,
		CountryCodeSource:            &source,
		PreferredDomesticCarrierCode: &carrierCode,
	}

	if got, want := number.GetCountryCode(), countryCode; got != want {
		t.Errorf("GetCountryCode() = %d, want %d", got, want)
	}
	if got, want := number.GetNationalNumber(), nationalNumber; got != want {
		t.Errorf("GetNationalNumber() = %d, want %d", got, want)
	}
	if got, want := number.GetExtension(), extension; got != want {
		t.Errorf("GetExtension() = %q, want %q", got, want)
	}
	if !number.GetItalianLeadingZero() {
		t.Error("GetItalianLeadingZero() = false, want true")
	}
	if got, want := number.GetNumberOfLeadingZeros(), leadingZeros; got != want {
		t.Errorf("GetNumberOfLeadingZeros() = %d, want %d", got, want)
	}
	if got, want := number.GetRawInput(), rawInput; got != want {
		t.Errorf("GetRawInput() = %q, want %q", got, want)
	}
	if got, want := number.GetCountryCodeSource(), source; got != want {
		t.Errorf("GetCountryCodeSource() = %d, want %d", got, want)
	}
	if got, want := number.GetPreferredDomesticCarrierCode(), carrierCode; got != want {
		t.Errorf("GetPreferredDomesticCarrierCode() = %q, want %q", got, want)
	}
}

func TestPhoneNumberGettersUseDefaults(t *testing.T) {
	var number *PhoneNumber

	if got := number.GetCountryCode(); got != 0 {
		t.Errorf("GetCountryCode() = %d, want 0", got)
	}
	if got := number.GetNationalNumber(); got != 0 {
		t.Errorf("GetNationalNumber() = %d, want 0", got)
	}
	if got := number.GetExtension(); got != "" {
		t.Errorf("GetExtension() = %q, want empty", got)
	}
	if number.GetItalianLeadingZero() {
		t.Error("GetItalianLeadingZero() = true, want false")
	}
	if got := number.GetNumberOfLeadingZeros(); got != 1 {
		t.Errorf("GetNumberOfLeadingZeros() = %d, want 1", got)
	}
	if got := number.GetRawInput(); got != "" {
		t.Errorf("GetRawInput() = %q, want empty", got)
	}
	if got := number.GetCountryCodeSource(); got != CountryCodeSourceUnspecified {
		t.Errorf("GetCountryCodeSource() = %d, want unspecified", got)
	}
	if got := number.GetPreferredDomesticCarrierCode(); got != "" {
		t.Errorf("GetPreferredDomesticCarrierCode() = %q, want empty", got)
	}
}

func TestPhoneNumberJSON(t *testing.T) {
	countryCode := int32(1)
	number := PhoneNumber{CountryCode: &countryCode}

	encoded, err := json.Marshal(number)
	if err != nil {
		t.Fatalf("Marshal() error = %v", err)
	}

	if got, want := string(encoded), `{"country_code":1}`; got != want {
		t.Errorf("Marshal() = %s, want %s", got, want)
	}
}
