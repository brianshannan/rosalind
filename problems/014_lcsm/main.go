package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/brianshannan/rosalind/utils"
)

func main() {
	flag.Parse()
	data, err := utils.ReadFileFromArgStripTrailingNewline()
	if err != nil {
		panic(err)
	}

	entries := utils.ParseFasta(string(data))

	// This is not an efficient way to do this (O(n^3)), but it
	// should be fine for small datasets.
	s := entries[0].Sequence
	maximalSubstring := ""
	for i := 0; i < len(s); i++ {
		for j := i + 1; j <= len(s); j++ {
			substring := s[i:j]
			fullMatch := true

			for _, entry := range entries[1:] {
				if !strings.Contains(entry.Sequence, substring) {
					fullMatch = false
					break
				}
			}

			if fullMatch && len(substring) > len(maximalSubstring) {
				maximalSubstring = substring
			}
		}
	}

	fmt.Println(maximalSubstring)
}
