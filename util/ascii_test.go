package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestASCII(t *testing.T) {
	t.Run("rune to ascii code", func(t *testing.T) {
		assert.Equal(t, 65, CharToASCII('A'))
		assert.Equal(t, 90, CharToASCII('Z'))
		assert.Equal(t, 97, CharToASCII('a'))
		assert.Equal(t, 122, CharToASCII('z'))
	})

	t.Run("byte to ascii code", func(t *testing.T) {
		assert.Equal(t, 65, CharToASCII(byte('A')))
		assert.Equal(t, 90, CharToASCII(byte('Z')))
		assert.Equal(t, 97, CharToASCII(byte('a')))
		assert.Equal(t, 122, CharToASCII(byte('z')))
	})

	t.Run("ascii code to rune", func(t *testing.T) {
		assert.Equal(t, 'A', ASCIIToChar[rune](65))
		assert.Equal(t, 'Z', ASCIIToChar[rune](90))
		assert.Equal(t, 'a', ASCIIToChar[rune](97))
		assert.Equal(t, 'z', ASCIIToChar[rune](122))
	})

	t.Run("ascii code to byte", func(t *testing.T) {
		assert.Equal(t, byte('A'), ASCIIToChar[byte](65))
		assert.Equal(t, byte('Z'), ASCIIToChar[byte](90))
		assert.Equal(t, byte('a'), ASCIIToChar[byte](97))
		assert.Equal(t, byte('z'), ASCIIToChar[byte](122))
	})
}
