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

	num, err := strconv.Atoi(string(data))
	if err != nil {
		panic(err)
	}

	numPermutations := 1
	values := make([]string, num)
	for i := 1; i <= num; i++ {
		numPermutations *= i
		values[i-1] = strconv.Itoa(i)
	}
	fmt.Println(numPermutations)

	outputFn := func(values []string) {
		fmt.Println(strings.Join(values, " "))
	}
	generate(num, values, outputFn)
}

// heaps algorithm
func generate(k int, currentValues []string, outputFn func(values []string)) {
	if k == 1 {
		outputFn(currentValues)
		return
	}

	generate(k-1, currentValues, outputFn)
	for i := 0; i < k-1; i++ {
		if k%2 == 0 {
			currentValues[i], currentValues[k-1] = currentValues[k-1], currentValues[i]
		} else {
			currentValues[0], currentValues[k-1] = currentValues[k-1], currentValues[0]
		}
		generate(k-1, currentValues, outputFn)
	}
}
