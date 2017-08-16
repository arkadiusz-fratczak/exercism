package hamming

import "fmt"

const testVersion = 6

type DNAStrandsLenghtDiffer struct {
	Strand1, Strand2 string
}

func (e *DNAStrandsLenghtDiffer) Error() string {
	return fmt.Sprintf("Lenght of %s and %s differs: %d != %d", e.Strand1, e.Strand2, len(e.Strand1), len(e.Strand2))
}

func newDNAStrandsLenghtDiffer(s1, s2 string) *DNAStrandsLenghtDiffer {
	return &DNAStrandsLenghtDiffer{s1, s2}
}

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return -1, newDNAStrandsLenghtDiffer(a, b)
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
