package main

import (
	"flag"
	"fmt"
	"math/big"

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
	sequence := entries[0].Sequence

	cgCount := 0
	for _, char := range sequence {
		if char == 'C' || char == 'G' {
			cgCount++
		}
	}
	auCount := len(sequence) - cgCount

	cgCombos := factorial(cgCount / 2)
	auCombos := factorial(auCount / 2)
	combos := big.NewInt(1).Mul(cgCombos, auCombos)
	fmt.Println(combos.String())
}

func factorial(n int) *big.Int {
	result := big.NewInt(1)
	for i := 2; i <= n; i++ {
		result.Mul(result, big.NewInt(int64(i)))
	}
	return result
}
