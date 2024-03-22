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

	length, err := strconv.Atoi(string(data))
	if err != nil {
		panic(err)
	}

	values := make([]int, length)
	for i := 0; i < length; i++ {
		values[i] = i + 1
	}

	signedVariants := utils.IntPow(2, length)
	fmt.Println(signedVariants * utils.Factorial(length))

	for i := 0; i < signedVariants; i++ {
		outputPermutations(values)
		makeNextSignedVariation(values)
	}

}

func outputPermutations(values []int) {
	strValues := utils.SliceMap(values, strconv.Itoa)
	utils.GeneratePermutation(strValues, func(values []string) {
		fmt.Println(strings.Join(values, " "))
	})
}

func makeNextSignedVariation(values []int) {
	for i := 0; i < len(values); i++ {
		values[i] *= -1
		if values[i] < 0 {
			return
		}
	}
}
