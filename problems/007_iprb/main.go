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

	// homozygous dominant (k)
	// heterozygous (m)
	// homozygous recessive (n)
	// let t = k + m + n
	// k + k -> 1 dominant so (k/t * (k-1)/(t-1))
	// k + m -> 1 dominant so (k/t * m/(t-1))
	// k + n -> 1 dominant so (k/t * n/(t-1))
	// m + k -> 1 dominant so (m/t * k/(t-1))
	// m + m -> 3/4 dominant so 3/4 * (m/t * (m-1)/(t-1))
	// m + n -> 1/2 dominant so 1/2 * (m/t * n/(t-1))
	// n + k -> 1 dominant so (n/t * k/(t-1))
	// n + m -> 1/2 dominant so 1/2 * (n/t * m/(t-1))
	// n + n -> 0 dominant so 0

	nums := strings.Split(string(data), " ")
	k, err := strconv.ParseFloat(nums[0], 64)
	if err != nil {
		panic(err)
	}
	m, err := strconv.ParseFloat(nums[1], 64)
	if err != nil {
		panic(err)
	}
	n, err := strconv.ParseFloat(nums[2], 64)
	if err != nil {
		panic(err)
	}
	t := k + m + n

	var prob float64
	prob += k / t * (k - 1) / (t - 1)
	prob += k / t * m / (t - 1)
	prob += k / t * n / (t - 1)
	prob += m / t * k / (t - 1)
	prob += 0.75 * (m / t * (m - 1) / (t - 1))
	prob += 0.5 * (m / t * n / (t - 1))
	prob += n / t * k / (t - 1)
	prob += 0.5 * (n / t * m / (t - 1))

	fmt.Printf("%f\n", prob)
}
