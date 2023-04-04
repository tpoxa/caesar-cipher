package caesar

import (
	"strings"
	"unicode"
)

type Cipher struct {
	alphabet []rune
}

func NewCipher(alphabet []rune) *Cipher {
	return &Cipher{alphabet: alphabet}
}

func (c Cipher) Encrypt(text string, shifts int) string {
	return c.shift(text, shifts)
}

func (c Cipher) Decrypt(text string, shifts int) string {
	return c.shift(text, -shifts)
}

// Takes a string and shifts each character by the provided amount of letters.
// Supports unicode-based encodings.
func (c Cipher) shift(input string, shift int) string {
	// prevent range overflow
	shift %= len(c.alphabet)
	var result strings.Builder
	for _, r := range input {
		if idx := c.getAlphabetPos(unicode.ToLower(r)); idx >= 0 {
			// rune exists in alphabet
			isUpper := unicode.IsUpper(r)
			// ensure we wont't go negative
			idx = idx + shift + len(c.alphabet)
			// replace it with the shifted rune
			r = c.alphabet[idx%len(c.alphabet)]
			if isUpper {
				// if the original rune was uppercased, then change the case to upper
				r = unicode.ToUpper(r)
			}
		}
		result.WriteRune(r)
	}
	return result.String()
}

func (c Cipher) getAlphabetPos(r int32) int {
	for k, v := range c.alphabet {
		if v == r {
			return k
		}
	}
	return -1
}
