package utils

func TranscribeRNAToProtein(rna string) []string {
	var proteins []string
	var currentProteins []string
	for i := 0; i < len(rna); i += 3 {
		if i+3 > len(rna) {
			break
		}

		codon := rna[i : i+3]
		if IsStartCodon(codon) {
			currentProteins = append(currentProteins, "")
		}
		if len(currentProteins) > 0 && IsStopCodon(codon) {
			proteins = append(proteins, currentProteins...)
			currentProteins = nil
		}

		for i, currentProtein := range currentProteins {
			currentProteins[i] = currentProtein + rnaCodonMap[codon]
		}
	}
	return proteins
}

func IsStartCodon(codon string) bool {
	return codon == "AUG"
}

func IsStopCodon(codon string) bool {
	return rnaCodonMap[codon] == "Stop"
}

var rnaCodonMap = map[string]string{
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
