package main

import (
	"flag"
	"fmt"
	"math"
	"math/big"
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
	k, err := strconv.Atoi(nums[0])
	if err != nil {
		panic(err)
	}
	n, err := strconv.Atoi(nums[1])
	if err != nil {
		panic(err)
	}

	// AA + Aa -> 0.5 AA 0.5 Aa
	// Aa + Aa -> 0.25 AA 0.5 Aa 0.25 aa
	// aa + Aa -> 0.5 Aa 0.5 aa
	// steady state is 0.25 AA, 0.5 Aa, 0.25 aa

	genoAaBb := 0.5 * 0.5
	kGenerationPop := int(math.Pow(2, float64(k)))
	atLeastNumAaBb := 0.0

	// Probabibility of at least n AaBb is the sum (n to 2^k) probability of n
	for i := n; i <= kGenerationPop; i++ {
		chooseki := getMChooseN(kGenerationPop, i)
		choosekiF := big.NewFloat(0.0).SetInt(chooseki)
		choosekiF.Mul(choosekiF, big.NewFloat(math.Pow(genoAaBb, float64(i))*math.Pow(1-genoAaBb, float64(kGenerationPop-i))))
		choosekiF64, _ := choosekiF.Float64()
		atLeastNumAaBb += choosekiF64
	}

	fmt.Printf("%.20f\n", atLeastNumAaBb)
}

func getMChooseN(m, n int) *big.Int {
	numerator := factorial(m)
	denomenator := big.NewInt(0).Mul(factorial(m-n), factorial(n))
	return big.NewInt(0).Div(numerator, denomenator)
}

func factorial(x int) *big.Int {
	fact := big.NewInt(1)
	for i := 2; i <= x; i++ {
		fact.Mul(fact, big.NewInt(int64(i)))
	}
	return fact
}
