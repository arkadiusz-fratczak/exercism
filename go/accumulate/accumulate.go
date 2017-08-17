package accumulate

const testVersion = 1

func Accumulate(stringSlice []string, f func(string) string) []string {
	result := make([]string, len(stringSlice))
	for i, s := range stringSlice {
		result[i] = f(s)
	}
	return result
}
