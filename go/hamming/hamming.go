package hamming

import "fmt"

const testVersion = 6

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return -1, fmt.Errorf("Lenght of %s and %s differs: %d != %d", a, b, len(a), len(b))
	}
	distance := 0
	aRunes := []rune(a)
	bRunes := []rune(b)
	for i, v := range aRunes {
		if v != bRunes[i] {
			distance++
		}
	}
	return distance, nil
}
