package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/brianshannan/rosalind/utils"
)

func main() {
	flag.Parse()
	data, err := utils.ReadFileFromArg()
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(data), "\n")
	line1 := lines[0]
	line2 := lines[1]

	if len(line1) != len(line2) {
		panic("mismatched lengths")
	}

	mismatchCount := 0
	for i := 0; i < len(line1); i++ {
		if line1[i] != line2[i] {
			mismatchCount++
		}
	}
	fmt.Printf("%d\n", mismatchCount)
}
