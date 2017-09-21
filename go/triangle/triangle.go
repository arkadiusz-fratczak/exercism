package triangle

import "math"

const testVersion = 3

func KindFromSides(a, b, c float64) Kind {
	if a+b < c || a+c < b || b+c < a || !isValidNumber(a) || !isValidNumber(b) || !isValidNumber(c) {
		return NaT
	} else if a == b && b == c {
		return Equ
	} else if a != b && a != c && b != c {
		return Sca
	}
	return Iso
}
func isValidNumber(n float64) bool {
	return n > 0 && n < math.MaxFloat64
}

type Kind byte

var NaT Kind = 0 // not a triangle
var Equ Kind = 1 // equilateral
var Iso Kind = 2 // isosceles
var Sca Kind = 3 // scalene
