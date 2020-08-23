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
	data, err := utils.ReadFileFromArgStripTrailingNewline()
	if err != nil {
		panic(err)
	}

	ids := strings.Split(string(data), "\n")

	for _, id := range ids {
		fastas, err := utils.DownloadFasta(id)
		if err != nil {
			panic(err)
		}
		if len(fastas) != 1 {
			panic(fmt.Errorf("found %d fastas, not 1", len(fastas)))
		}

		matchingIdxs := findMotifIdxs(fastas[0].Sequence, 4, func(s string) bool {
			return s[0] == 'N' && s[1] != 'P' && (s[2] == 'S' || s[2] == 'T') && s[3] != 'P'
		})
		if len(matchingIdxs) == 0 {
			continue
		}

		matchingStartIdxs := make([]string, len(matchingIdxs))
		for i, matchingIdx := range matchingIdxs {
			matchingStartIdxs[i] = strconv.Itoa(matchingIdx)
		}

		fmt.Println(id)
		fmt.Println(strings.Join(matchingStartIdxs, " "))
	}
}

func findMotifIdxs(s string, mlen int, f func(string) bool) []int {
	idxs := []int{}
	for i := 0; i < len(s)-mlen; i++ {
		if f(s[i : i+mlen]) {
			idxs = append(idxs, i+1)
		}
	}
	return idxs
}
