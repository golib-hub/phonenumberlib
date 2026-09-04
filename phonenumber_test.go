package phonenumberlib

import (
	"testing"

	"github.com/nyaruka/phonenumbers/v2"
)

func TestParseUSNumber(t *testing.T) {
	number, err := Parse("8886418722", "US")
	if err != nil {
		t.Fatalf("Parse() error = %v", err)
	}

	if !phonenumbers.IsValidNumber(number) {
		t.Fatal("Parse() returned an invalid US number")
	}

	if got, want := phonenumbers.Format(number, phonenumbers.E164), "+18886418722"; got != want {
		t.Errorf("E.164 format = %q, want %q", got, want)
	}
}
