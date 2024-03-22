package utils

// heaps algorithm https://en.wikipedia.org/wiki/Heap%27s_algorithm
func GeneratePermutation[T any](values []T, outputFn func(values []T)) {
	valuesCopy := make([]T, len(values))
	copy(valuesCopy, values)

	generatePermutation(len(valuesCopy), valuesCopy, outputFn)
}

func generatePermutation[T any](k int, currentValues []T, outputFn func(values []T)) {
	if k == 1 {
		outputFn(currentValues)
		return
	}

	generatePermutation(k-1, currentValues, outputFn)
	for i := 0; i < k-1; i++ {
		if k%2 == 0 {
			currentValues[i], currentValues[k-1] = currentValues[k-1], currentValues[i]
		} else {
			currentValues[0], currentValues[k-1] = currentValues[k-1], currentValues[0]
		}
		generatePermutation(k-1, currentValues, outputFn)
	}
}
