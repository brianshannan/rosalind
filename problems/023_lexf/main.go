package main

import (
	"flag"
	"fmt"
	"slices"
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
	alphabet := strings.Split(lines[0], " ")
	slices.Sort(alphabet)

	desiredLength, err := strconv.Atoi(lines[1])
	if err != nil {
		panic(err)
	}

	generate(alphabet, nil, desiredLength)
}

func generate(alphabet []string, currentValue []string, desiredLength int) {
	if len(currentValue) == desiredLength {
		fmt.Println(strings.Join(currentValue, ""))
		return
	}

	currentValue = append(currentValue, "")
	for _, letter := range alphabet {
		currentValue[len(currentValue)-1] = letter
		generate(alphabet, currentValue, desiredLength)
	}
}
