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

	proteinString := string(data)
	possibilities := 1
	for _, protein := range proteinString {
		numCodons := len(proteinToCodon[protein])
		possibilities *= numCodons
		possibilities %= 1000000
	}
	possibilities *= 3 // for the stop codon
	possibilities %= 1000000

	fmt.Println(possibilities)
}

var proteinToCodon = map[rune][]string{
	'A': {"GCG", "GCU", "GCA", "GCC"},
	'C': {"UGC", "UGU"},
	'D': {"GAU", "GAC"},
	'E': {"GAA", "GAG"},
	'F': {"UUU", "UUC"},
	'G': {"GGG", "GGC", "GGU", "GGA"},
	'H': {"CAC", "CAU"},
	'I': {"AUU", "AUA", "AUC"},
	'K': {"AAG", "AAA"},
	'L': {"CUG", "CUA", "CUU", "UUG", "CUC", "UUA"},
	'M': {"AUG"},
	'N': {"AAU", "AAC"},
	'P': {"CCG", "CCU", "CCA", "CCC"},
	'Q': {"CAG", "CAA"},
	'R': {"AGA", "CGC", "AGG", "CGU", "CGG", "CGA"},
	'S': {"UCA", "UCG", "UCC", "AGU", "AGC", "UCU"},
	'T': {"ACA", "ACG", "ACC", "ACU"},
	'V': {"GUC", "GUG", "GUA", "GUU"},
	'W': {"UGG"},
	'Y': {"UAC", "UAU"},
	//"Stop": {"UAA", "UGA", "UAG"},
}
