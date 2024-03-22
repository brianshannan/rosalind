package main

import (
	"flag"
	"fmt"

	"github.com/brianshannan/rosalind/utils"
)

func main() {
	flag.Parse()
	data, err := utils.ReadFileFromArgStripTrailingNewline()
	if err != nil {
		panic(err)
	}

	entries := utils.ParseFasta(string(data))
	seq1, seq2 := entries[0].Sequence, entries[1].Sequence

	numTransitions := 0
	numTransversions := 0
	for i := 0; i < len(seq1); i++ {
		if isTransition(seq1[i], seq2[i]) {
			numTransitions++
		}
		if isTransversion(seq1[i], seq2[i]) {
			numTransversions++
		}
	}
	println(numTransitions, numTransversions)
	fmt.Println(float64(numTransitions) / float64(numTransversions))
}

func isTransition(a, b byte) bool {
	switch a {
	case 'A':
		return b == 'G'
	case 'G':
		return b == 'A'
	case 'C':
		return b == 'T'
	case 'T':
		return b == 'C'
	}
	return false
}

func isTransversion(a, b byte) bool {
	switch a {
	case 'A', 'G':
		return b == 'C' || b == 'T'
	case 'C', 'T':
		return b == 'A' || b == 'G'
	}
	return false
}
