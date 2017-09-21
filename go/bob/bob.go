package bob

import (
	"strings"
	"unicode"
)

const testVersion = 3

func Hey(sentence string) string {
	if containsOnlySpaces(sentence) {
		return "Fine. Be that way!"
	} else if containsLetters(sentence) && containsOnlyUpperLetters(sentence) {
		return "Whoa, chill out!"
	} else if strings.HasSuffix(strings.TrimSpace(sentence), "?") {
		return "Sure."
	}
	return "Whatever."
}

func containsOnlySpaces(sentence string) bool {
	for _, r := range sentence {
		if !unicode.IsSpace(r) {
			return false
		}
	}
	return true
}

func containsLetters(sentence string) bool {
	for _, r := range sentence {
		if unicode.IsLetter(r) {
			return true
		}
	}
	return false
}

func containsOnlyUpperLetters(sentence string) bool {
	for _, r := range sentence {
		if unicode.IsLetter(r) && unicode.IsLower(r) {
			return false
		}
	}
	return true
}
