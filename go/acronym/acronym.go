package acronym

import (
	"unicode"
)

const testVersion = 3

func Abbreviate(longName string) string {
	longNameRunes := []rune(longName)
	result := []rune{longNameRunes[0]}
	for i, v := range []rune(longName) {
		if isSpaceOrDash(v) && i+1 < len(longNameRunes) {
			result = append(result, unicode.ToUpper(longNameRunes[i+1]))
		}
	}
	return string(result)
}

func isSpaceOrDash(v rune) bool {
	return unicode.IsSpace(v) || v == '-'
}
