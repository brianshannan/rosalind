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
	if len(entries) != 1 {
		panic("unexpected number of entries")
	}
	entry := entries[0]

	for i := 0; i < len(entry.Sequence); i++ {
		checkForPalidromesAtPosition(entry.Sequence, i, 4, 12)
	}
}

func checkForPalidromesAtPosition(dna string, start int, minLength, maxLength int) {
	for i := minLength; i <= maxLength; i++ {
		if start+i > len(dna) {
			return
		}

		if dna[start:start+i] == utils.DNAReverseCompliment(dna[start:start+i]) {
			fmt.Printf("%v %v\n", start+1, i)
		}
	}
}
