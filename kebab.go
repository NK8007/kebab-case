package main

import (
	"strings"
	"unicode"
)

// ToKebabCase converts s to kebab-case: lowercase words separated by
// single hyphens. Word boundaries are spaces, punctuation, underscores,
// hyphens, lower-to-upper case changes ("helloWorld"), the end of an
// acronym ("HTTPServer" -> "http-server") and letter/digit changes
// ("version2" -> "version-2").
func ToKebabCase(s string) string {
	runes := []rune(s)
	var b strings.Builder
	b.Grow(len(s) + 4)

	pendingHyphen := false
	for i, r := range runes {
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
			// Any separator ends the current word.
			pendingHyphen = b.Len() > 0
			continue
		}

		if i > 0 && b.Len() > 0 && !pendingHyphen {
			prev := runes[i-1]
			switch {
			case unicode.IsUpper(r) && (unicode.IsLower(prev) || unicode.IsDigit(prev)):
				// helloWorld, v2Beta
				pendingHyphen = true
			case unicode.IsUpper(r) && unicode.IsUpper(prev) &&
				i+1 < len(runes) && unicode.IsLower(runes[i+1]):
				// HTTPServer: hyphen before the "S"
				pendingHyphen = true
			case unicode.IsDigit(r) != unicode.IsDigit(prev):
				// version2, 2fast
				pendingHyphen = true
			}
		}

		if pendingHyphen {
			b.WriteByte('-')
			pendingHyphen = false
		}
		b.WriteRune(unicode.ToLower(r))
	}
	return b.String()
}
