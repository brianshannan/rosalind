package utils

import "strconv"

func SliceMap[T, U any](input []T, fn func(T) U) []U {
	output := make([]U, len(input))
	for i, inputVal := range input {
		output[i] = fn(inputVal)
	}
	return output
}

func IntSliceToStringSlice(ints []int) []string {
	return SliceMap(ints, strconv.Itoa)
}
