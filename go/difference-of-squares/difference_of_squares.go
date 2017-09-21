package diffsquares

const testVersion = 1

func SquareOfSums(n int) int {
	sum := n * (n + 1) / 2
	return sum * sum
}

func SumOfSquares(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i * i
	}
	return sum
}

func SumOfSquares2(n int) int {
	sum := 0
	prevPow := 0
	nextPow := 0
	for i := 0; i < n; i++ {
		nextPow = prevPow + 1 + i*2
		sum += nextPow
		prevPow = nextPow
	}
	return sum
}

func Difference(n int) int {
	return SquareOfSums(n) - SumOfSquares(n)
}
