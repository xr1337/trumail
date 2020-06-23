package verifier

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmailWithPlus(t *testing.T) {
	type test struct {
		email    string
		username string
		domain   string
		address  string
		hash     string
	}

	testCases := []test{
		{"email_username+a@domain.com", "email_username+a", "domain.com", "email_username+a@domain.com", "842da636835e1db9be6f0609610fc439"},
		{"email_username@domain.com", "email_username", "domain.com", "email_username@domain.com", "629b2a45027be2158761fecb17eb79d6"},
		{"email.username@domain.com", "email.username", "domain.com", "email.username@domain.com", "80f59a9e93602984c43085a1e9412d9a"},
		{"email_username@DoMAIn.CoM", "email_username", "domain.com", "email_username@domain.com", "629b2a45027be2158761fecb17eb79d6"},
		{"EMAIL_USERNAME@DOMAIN.COM", "EMAIL_USERNAME", "domain.com", "EMAIL_USERNAME@domain.com", "94d8a553082c902d086c47bd40ccf3c1"},
	}
	for _, c := range testCases {
		address, err := ParseAddress(c.email)

		assert.Nil(t, err)
		assert.Equal(t, address.Username, c.username)
		assert.Equal(t, address.Domain, c.domain)
		assert.Equal(t, address.Address, c.address)
		assert.Equal(t, address.MD5Hash, c.hash)
	}
}

func TestParseAddressInvalidEmail(t *testing.T) {
	email := "email_username@"
	address, err := ParseAddress(email)

	assert.Error(t, err)
	assert.Nil(t, address)
}
