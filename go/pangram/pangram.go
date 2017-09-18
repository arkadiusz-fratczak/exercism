package pangram

import (
	"strings"
	"unicode"
)

const testVersion = 2
const alphabetLength = 26

func IsPangram(word string) bool {
	alphabet := make(map[rune]struct{})
	for _, l := range strings.ToLower(word) {
		if unicode.IsLetter(l) {
			alphabet[l] = struct{}{}
		}
	}
	return len(alphabet) == alphabetLength
}
