package utils

import "strconv"

func IntSliceToStringSlice(ints []int) []string {
	strings := make([]string, len(ints))
	for i, ival := range ints {
		strings[i] = strconv.Itoa(ival)
	}
	return strings
}
