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
	pairsPerLitter, err := strconv.Atoi(nums[1])
	if err != nil {
		panic(err)
	}

	totalPairs := 1
	babyPairs := 1

	for i := 1; i < months; i++ {
		maturePairs := totalPairs - babyPairs
		babyPairs = maturePairs * pairsPerLitter
		totalPairs += babyPairs
	}

	fmt.Println(totalPairs)
}
