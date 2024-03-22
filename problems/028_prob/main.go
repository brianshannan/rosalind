package main

import (
	"flag"
	"fmt"
	"math"
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

	lines := strings.Split(string(data), "\n")
	sequence := lines[0]

	gcProbabilities := utils.SliceMap(strings.Split(lines[1], " "), func(input string) float64 {
		probability, err := strconv.ParseFloat(input, 64)
		if err != nil {
			panic(err)
		}
		return probability
	})

	seqProbabilities := utils.SliceMap(gcProbabilities, func(gcProb float64) string {
		seqProb := calculateSequenceProbabilityLog10(sequence, gcProb)
		return strconv.FormatFloat(seqProb, 'f', 5, 64)
	})
	fmt.Println(strings.Join(seqProbabilities, " "))
}

func calculateSequenceProbabilityLog10(sequence string, gcProb float64) float64 {
	// log10(x * y) = log10(x) + log10(y)
	probability := 0.0
	for _, char := range sequence {
		if char == 'C' || char == 'G' {
			probability += math.Log10(gcProb / 2)
		} else {
			probability += math.Log10((1 - gcProb) / 2)
		}
	}

	return probability
}
