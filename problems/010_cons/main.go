package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"

	"github.com/brianshannan/rosalind/utils"
)

func main() {
	flag.Parse()
	data, err := utils.ReadFileFromArg()
	if err != nil {
		panic(err)
	}

	entries := utils.ParseFasta(string(data))

	profile := make([][]int, 4)
	sequenceLen := len(entries[0].Sequence)
	for i := 0; i < 4; i++ {
		profile[i] = make([]int, sequenceLen)
	}

	for _, entry := range entries {
		for i, char := range entry.Sequence {
			charNum := dnaCharToInt(char)
			profile[charNum][i]++
		}
	}

	consensus := ""
	for i := 0; i < sequenceLen; i++ {
		maxCharNum := 0
		maxCount := 0

		for charNum := 0; charNum < 4; charNum++ {
			if maxCount < profile[charNum][i] {
				maxCharNum = charNum
				maxCount = profile[charNum][i]
			}
		}

		consensus += intToDnaChar(maxCharNum)
	}

	fmt.Printf("%s\n", consensus)

	for i := 0; i < 4; i++ {
		sslice := intSliceToStringSlice(profile[i])
		fmt.Printf("%s: %s\n", intToDnaChar(i), strings.Join(sslice, " "))
	}
}

func dnaCharToInt(r rune) int {
	if r == 'A' {
		return 0
	} else if r == 'C' {
		return 1
	} else if r == 'G' {
		return 2
	} else if r == 'T' {
		return 3
	}
	return -1
}

func intToDnaChar(i int) string {
	if i == 0 {
		return "A"
	} else if i == 1 {
		return "C"
	} else if i == 2 {
		return "G"
	} else if i == 3 {
		return "T"
	}
	return ""
}

func intSliceToStringSlice(islice []int) []string {
	sslice := make([]string, len(islice))
	for i := 0; i < len(islice); i++ {
		sslice[i] = strconv.Itoa(islice[i])
	}
	return sslice
}
