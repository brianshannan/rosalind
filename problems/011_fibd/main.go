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

	if data[len(data)-1] == '\n' {
		data = data[:len(data)-1]
	}

	nums := strings.Split(string(data), " ")
	months, err := strconv.Atoi(nums[0])
	if err != nil {
		panic(err)
	}
	lifetime, err := strconv.Atoi(nums[1])
	if err != nil {
		panic(err)
	}

	totalPairs := 1
	pairs := make([]int, lifetime)
	pairs[0] = 1

	for i := 1; i < months; i++ {
		maturePairs := totalPairs - pairs[0]

		totalPairs -= pairs[lifetime-1]
		shiftRight(pairs)
		pairs[0] = maturePairs
		totalPairs += maturePairs
	}

	fmt.Println(totalPairs)
}

func shiftRight(arr []int) {
	for i := len(arr) - 1; i > 0; i-- {
		arr[i] = arr[i-1]
	}
}
