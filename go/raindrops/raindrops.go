package raindrops

import "strconv"

const testVersion = 3

func Convert(n int) string {
	result := ""
	if n % 3 == 0 {
		result += "Pling"
	}
	if n % 5 == 0 {
		result += "Plang"
	}
	if n % 7 == 0 {
		result += "Plong"
	}
	if len(result) == 0 {
		result += strconv.Itoa(n)
	}
	return result
}

