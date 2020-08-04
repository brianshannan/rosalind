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

	nums := strings.Split(string(data), " ")

	// 1. AA-AA -> 1 dominant
	// 2. AA-Aa -> 1 dominant
	// 3. AA-aa -> 1 dominant
	// 4. Aa-Aa -> 3/4 dominant
	// 5. Aa-aa -> 1/2 dominant
	// 6. aa-aa -> 0 dominant
	type1Pairs, err := strconv.Atoi(nums[0])
	if err != nil {
		panic(err)
	}
	type2Pairs, err := strconv.Atoi(nums[1])
	if err != nil {
		panic(err)
	}
	type3Pairs, err := strconv.Atoi(nums[2])
	if err != nil {
		panic(err)
	}
	type4Pairs, err := strconv.Atoi(nums[3])
	if err != nil {
		panic(err)
	}
	type5Pairs, err := strconv.Atoi(nums[4])
	if err != nil {
		panic(err)
	}

	expectedDominant := 0.0
	expectedDominant += float64(type1Pairs)
	expectedDominant += float64(type2Pairs)
	expectedDominant += float64(type3Pairs)
	expectedDominant += 0.75 * float64(type4Pairs)
	expectedDominant += 0.5 * float64(type5Pairs)
	expectedDominant *= 2

	fmt.Printf("%f\n", expectedDominant)

}
