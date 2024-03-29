// Package morestrings implements additional functions to manipulate UTF-8
// encoded strings, beyond what is provided in the standard "strings" package.
package morestrings

import "fmt"

// ReverseRunes returns its argument string reversed rune-wise left to right.
func ReverseRunes(s string) string {
	r := []rune(s)

	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}

	return string(r)
}

// Print illegal UTF-8 encoding
// In range, erroneous encodings consume 1 byte and produce the replacement rune U+FFFD
// `rune` is Go's type and terminology for a single unicode code point (1 byte)
func IllegalChars() {
	someString := "日本\x80語"

	for pos, char := range someString {
		fmt.Printf("Character %#U starts at byte position %d\n", char, pos)
	}
}
