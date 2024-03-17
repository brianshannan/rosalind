package utils

import "bytes"

func DNAToRNA(dna string) string {
	dnaBytes := bytes.Clone([]byte(dna))
	for pos, char := range dnaBytes {
		if char == 'T' {
			dnaBytes[pos] = 'U'
		}
	}
	return string(dnaBytes)
}

func DNAReverseCompliment(dna string) string {
	dnaBytes := bytes.Clone([]byte(dna))
	reverse(dnaBytes)
	dnaCompliment(dnaBytes)
	return string(dnaBytes)
}

func reverse(data []byte) {
	for i := 0; i < len(data)/2; i++ {
		data[i], data[len(data)-i-1] = data[len(data)-i-1], data[i]
	}
}

func dnaCompliment(data []byte) {
	for pos, char := range data {
		if char == 'A' {
			data[pos] = 'T'
		} else if char == 'T' {
			data[pos] = 'A'
		} else if char == 'C' {
			data[pos] = 'G'
		} else if char == 'G' {
			data[pos] = 'C'
		}
	}
}
