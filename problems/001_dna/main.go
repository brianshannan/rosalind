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

	aCount, cCount, gCount, tCount := 0, 0, 0, 0
	for _, char := range data {
		if char == 'A' {
			aCount++
		} else if char == 'C' {
			cCount++
		} else if char == 'G' {
			gCount++
		} else if char == 'T' {
			tCount++
		}
	}

	fmt.Printf("%d %d %d %d\n", aCount, cCount, gCount, tCount)
}
