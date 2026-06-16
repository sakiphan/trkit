// Package text provides Turkish-locale-aware case conversion and slugification.
//
// The plain strings.ToUpper / strings.ToLower get the dotted/dotless i wrong
// for Turkish: strings.ToUpper("istanbul") yields "ISTANBUL" instead of
// "İSTANBUL". The standard library can do it correctly via
// strings.ToUpperSpecial(unicode.TurkishCase, s), but that incantation is easy
// to forget. Upper and Lower wrap it so the correct behaviour is the default,
// while Title, Capitalize and Slug add helpers that the standard library does
// not provide.
package text

import (
	"strings"
	"unicode"
)

// Upper returns s with all letters in upper case using Turkish rules:
// 'i' becomes 'İ' and 'ı' becomes 'I'.
func Upper(s string) string {
	return strings.ToUpperSpecial(unicode.TurkishCase, s)
}

// Lower returns s with all letters in lower case using Turkish rules:
// 'I' becomes 'ı' and 'İ' becomes 'i'.
func Lower(s string) string {
	return strings.ToLowerSpecial(unicode.TurkishCase, s)
}

// Title upper-cases the first letter of each whitespace-separated word and
// lower-cases the rest, using Turkish rules.
func Title(s string) string {
	words := strings.Fields(s)
	for i, w := range words {
		words[i] = Capitalize(w)
	}
	return strings.Join(words, " ")
}

// Capitalize upper-cases the first letter of s and lower-cases the rest, using
// Turkish rules.
func Capitalize(s string) string {
	if s == "" {
		return s
	}
	r := []rune(s)
	first := Upper(string(r[0]))
	rest := Lower(string(r[1:]))
	return first + rest
}

// slugMap holds the ASCII transliteration of Turkish letters used by Slug.
var slugMap = map[rune]rune{
	'ç': 'c', 'ğ': 'g', 'ı': 'i', 'ö': 'o', 'ş': 's', 'ü': 'u',
	'Ç': 'c', 'Ğ': 'g', 'İ': 'i', 'I': 'i', 'Ö': 'o', 'Ş': 's', 'Ü': 'u',
}

// Slug converts s into a URL-friendly slug: lower case, Turkish letters
// transliterated to ASCII, and any run of other characters collapsed into a
// single hyphen. Leading and trailing hyphens are trimmed.
//
//	Slug("Çağrı'nın İşi") == "cagrinin-isi"
func Slug(s string) string {
	var b strings.Builder
	b.Grow(len(s))
	lastHyphen := false
	for _, r := range s {
		// Apostrophes and quotes are dropped so contractions stay joined:
		// "Çağrı'nın" -> "cagrinin", not "cagri-nin".
		if r == '\'' || r == '’' || r == '"' {
			continue
		}
		if m, ok := slugMap[r]; ok {
			r = m
		} else {
			r = unicode.ToLower(r)
		}
		switch {
		case r >= 'a' && r <= 'z', r >= '0' && r <= '9':
			b.WriteRune(r)
			lastHyphen = false
		default:
			if !lastHyphen {
				b.WriteByte('-')
				lastHyphen = true
			}
		}
	}
	return strings.Trim(b.String(), "-")
}
