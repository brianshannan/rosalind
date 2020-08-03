package main

import (
	"flag"
	"fmt"

	"github.com/brianshannan/rosalind/utils"
)

func main() {
	flag.Parse()
	data, err := utils.ReadFileFromArg()
	if err != nil {
		panic(err)
	}

	s := string(data)
	proteinString := ""
	for i := 0; i < len(s); i += 3 {
		codon := rnaToCodon(s[i : i+3])
		if codon == "Stop" {
			break
		}
		proteinString += codon
	}

	fmt.Println(proteinString)
}

var rnaCodonMap map[string]string = map[string]string{
	"UUU": "F",
	"CUU": "L",
	"AUU": "I",
	"GUU": "V",
	"UUC": "F",
	"CUC": "L",
	"AUC": "I",
	"GUC": "V",
	"UUA": "L",
	"CUA": "L",
	"AUA": "I",
	"GUA": "V",
	"UUG": "L",
	"CUG": "L",
	"AUG": "M",
	"GUG": "V",
	"UCU": "S",
	"CCU": "P",
	"ACU": "T",
	"GCU": "A",
	"UCC": "S",
	"CCC": "P",
	"ACC": "T",
	"GCC": "A",
	"UCA": "S",
	"CCA": "P",
	"ACA": "T",
	"GCA": "A",
	"UCG": "S",
	"CCG": "P",
	"ACG": "T",
	"GCG": "A",
	"UAU": "Y",
	"CAU": "H",
	"AAU": "N",
	"GAU": "D",
	"UAC": "Y",
	"CAC": "H",
	"AAC": "N",
	"GAC": "D",
	"UAA": "Stop",
	"CAA": "Q",
	"AAA": "K",
	"GAA": "E",
	"UAG": "Stop",
	"CAG": "Q",
	"AAG": "K",
	"GAG": "E",
	"UGU": "C",
	"CGU": "R",
	"AGU": "S",
	"GGU": "G",
	"UGC": "C",
	"CGC": "R",
	"AGC": "S",
	"GGC": "G",
	"UGA": "Stop",
	"CGA": "R",
	"AGA": "R",
	"GGA": "G",
	"UGG": "W",
	"CGG": "R",
	"AGG": "R",
	"GGG": "G",
}

func rnaToCodon(rna string) string {
	return rnaCodonMap[rna]
}
