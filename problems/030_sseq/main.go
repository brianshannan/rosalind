package main

import (
	"flag"
	"fmt"
	"strconv"
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
	sequence := entries[0].Sequence
	subSequence := entries[1].Sequence

	currSubSeqIdx := 0
	var foundIndices []string
	for i, char := range sequence {
		if byte(char) == subSequence[currSubSeqIdx] {
			foundIndices = append(foundIndices, strconv.Itoa(i+1))
			currSubSeqIdx++
			if len(subSequence) == currSubSeqIdx {
				break
			}
		}
	}

	fmt.Println(strings.Join(foundIndices, " "))
}
