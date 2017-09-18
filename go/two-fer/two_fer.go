package twofer

import "fmt"

// ShareWith prints "One for X, one for me." When X is a name or "you".
func ShareWith(name string) string {
	if len(name) == 0 {
		return "One for you, one for me."
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
