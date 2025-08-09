// Package morestrings provides functions for manipulating UTF-8 strings.
//
// It includes utilities such as reversing rune order, uppercasing,
// and more complex transformations.
package morestrings

func ReverseRunes(s string) string {
	r := []rune(s)

	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
