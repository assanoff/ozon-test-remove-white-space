package main

import (
	"strings"
	"unicode"
)

// THIS IS A SOLUTION
// Strings are immutable in GO that's why we need to construct
// a new string using a string.Builder
func RemoveWhiteSpace(str string) string {
	//avoid allocation
	if str == "" {
		return ""
	}
	var b strings.Builder
	b.Grow(len(str))
	wasSpace, isSpace := true, false
	var curr rune

	for _, r := range str {
		curr = r
		isSpace = unicode.IsSpace(curr)
		if isSpace && wasSpace {
			continue
		}
		wasSpace = isSpace
		b.WriteRune(curr)
	}
	// remove last space accurate added after not-whitespace-char
	if unicode.IsSpace(curr) && b.Len() > 1 {
		return b.String()[0 : b.Len()-1]
	}
	return b.String()
}

// ONLY FOR BENCHMARK
// using stdlib not optimal for performance
func StandardizeSpacesFields(s string) string {
	return strings.Join(strings.Fields(s), " ")
}
