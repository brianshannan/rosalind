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

	line := strings.Split(string(data), "\n")[0]
	vals := strings.Split(line, " ")

	n, err := strconv.Atoi(vals[0])
	if err != nil {
		panic(err)
	}
	k, err := strconv.Atoi(vals[1])
	if err != nil {
		panic(err)
	}

	fmt.Println(numPartialPermutationsModulo(n, k, 1000000))
}

func numPartialPermutationsModulo(n, k, m int) int {
	result := 1
	for i := 0; i < k; i++ {
		result *= n - i
		result %= m
	}
	return result
}
