package utils

func Factorial(n int) int {
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result
}

func IntPow(base, exp int) int {
	result := 1
	for exp > 0 {
		// if we're odd, multiply result by our base then apply even logic
		// if we're even, square our base and half our exponent
		if exp%2 == 1 {
			result *= base
		}
		base *= base
		exp /= 2
	}
	return result
}
