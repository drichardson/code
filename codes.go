// Package codes generates random codes from a set of characters.
package codes

import (
	"crypto/rand"
)

const decimalDigits = "0123456789"

// NewRandomDecimalString produces a string made up of
// n somewhat randomly selected decmial digits.
func NewRandomDecimalString(n int) (string, error) {
	r := make([]byte, n)
	_, err := rand.Read(r)
	if err != nil {
		return "", err
	}

	result := make([]byte, n)
	for i := range result {
		result[i] = decimalDigits[int(r[i])%len(decimalDigits)]
	}
	return string(result), nil
}
