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

	results := map[string]struct{}{}
	findProteins(entry.Sequence, results)
	findProteins(utils.DNAReverseCompliment(entry.Sequence), results)
	for result := range results {
		fmt.Println(result)
	}
}

func findProteins(dna string, results map[string]struct{}) {
	for i := 0; i < 3; i++ {
		readingFrame := dna[i:]
		rna := utils.DNAToRNA(readingFrame)
		proteins := utils.TranscribeRNAToProtein(rna)
		for _, protein := range proteins {
			results[protein] = struct{}{}
		}
	}
}
