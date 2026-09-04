package phonenumberlib

import (
	"errors"
	"testing"
)

func TestParseUSNumber(t *testing.T) {
	number, err := Parse("8886418722", "US", VersionV2_0_11)
	if err != nil {
		t.Fatalf("Parse() error = %v", err)
	}

	if got, want := number.GetCountryCode(), int32(1); got != want {
		t.Errorf("country code = %d, want %d", got, want)
	}

	if got, want := number.GetNationalNumber(), uint64(8886418722); got != want {
		t.Errorf("national number = %d, want %d", got, want)
	}
}

func TestParseRejectsUnsupportedVersion(t *testing.T) {
	_, err := Parse("8886418722", "US", Version("v2.0.10"))

	var versionError *UnsupportedVersionError
	if !errors.As(err, &versionError) {
		t.Fatalf("Parse() error = %v, want UnsupportedVersionError", err)
	}
	if got, want := versionError.Version, Version("v2.0.10"); got != want {
		t.Errorf("unsupported version = %q, want %q", got, want)
	}
}
