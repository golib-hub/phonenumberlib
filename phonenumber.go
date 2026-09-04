// Package phonenumberlib provides phone-number parsing backed by libphonenumber.
package phonenumberlib

import "github.com/nyaruka/phonenumbers/v2"

// Parse parses input using region when input does not include a country calling
// code. Region must be an ISO 3166-1 alpha-2 country code, such as "US".
func Parse(input, region string) (*phonenumbers.PhoneNumber, error) {
	return phonenumbers.Parse(input, region)
}
