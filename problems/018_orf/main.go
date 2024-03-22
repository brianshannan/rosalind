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
		proteins := transcribeRNAToProtein(rna)
		for _, protein := range proteins {
			results[protein] = struct{}{}
		}
	}
}

func transcribeRNAToProtein(rna string) []string {
	var proteins []string
	var currentProteins []string
	for i := 0; i < len(rna); i += 3 {
		if i+3 > len(rna) {
			break
		}

		codon := rna[i : i+3]
		if utils.IsStartCodon(codon) {
			currentProteins = append(currentProteins, "")
		}
		if len(currentProteins) > 0 && utils.IsStopCodon(codon) {
			proteins = append(proteins, currentProteins...)
			currentProteins = nil
		}

		for i, currentProtein := range currentProteins {
			currentProteins[i] = currentProtein + utils.CodonToAA(codon)
		}
	}
	return proteins
}
