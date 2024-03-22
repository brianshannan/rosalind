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
	mainSequence := entries[0].Sequence

	for _, entry := range entries[1:] {
		mainSequence = strings.Replace(mainSequence, entry.Sequence, "", -1)
	}

	rna := utils.DNAToRNA(mainSequence)
	protein := utils.RNAToProtein(rna)
	fmt.Println(protein)
}
