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
	sequences := utils.SliceMap(entries, func(entry utils.FASTAEntry) string {
		return entry.Sequence
	})

	superstring := sequences[0]
	possibilities := sequences[1:]
	for i := 0; i < len(sequences)-1; i++ {
		superstring, possibilities = findBestOverlap(superstring, possibilities)
	}
	fmt.Println(superstring)
}

// findBestOverlap finds the best overlap match between the target string
// and the strings in possibilities. Returns the superstring of target and
// the best match, and the new possibilities pool, with the best match removed.
func findBestOverlap(target string, possibilities []string) (string, []string) {
	bestMatchLength := 0
	bestMatchIdx := -1
	bestSuperstring := ""

	checkOverlap := func(idx int, first, second string) {
		overlap := findOverlap(first, second)
		if overlap > bestMatchLength {
			bestMatchLength = overlap
			bestMatchIdx = idx
			bestSuperstring = first + second[overlap:]
		}
	}

	for i, possibility := range possibilities {
		checkOverlap(i, target, possibility)
		checkOverlap(i, possibility, target)
	}

	possibilities[bestMatchIdx] = possibilities[len(possibilities)-1]
	possibilities = possibilities[:len(possibilities)-1]
	return bestSuperstring, possibilities
}

// findOverlap finds the maximum overlap between the suffix of
// string 1 and the prefix of string 1.
func findOverlap(first, second string) int {
	minLength := min(len(first), len(second))
	firstOffset := len(first) - minLength
	for i := 0; i < minLength; i++ {
		if first[firstOffset+i:] == second[:minLength-i] {
			return minLength - i
		}
	}
	return 0
}
