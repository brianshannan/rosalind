package main

import (
	"flag"
	"fmt"

	"github.com/brianshannan/rosalind/utils"
)

func main() {
	flag.Parse()
	data, err := utils.ReadFileFromArg()
	if err != nil {
		panic(err)
	}

	for pos, char := range data {
		if char == 'T' {
			data[pos] = 'U'
		}
	}

	fmt.Println(string(data))
}
