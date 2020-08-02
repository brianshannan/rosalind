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

	reverse(data)
	compliment(data)
	fmt.Println(string(data))
}

func reverse(data []byte) {
	for i := 0; i < len(data)/2; i++ {
		data[i], data[len(data)-i-1] = data[len(data)-i-1], data[i]
	}
}

func compliment(data []byte) {
	for pos, char := range data {
		if char == 'A' {
			data[pos] = 'T'
		} else if char == 'T' {
			data[pos] = 'A'
		} else if char == 'C' {
			data[pos] = 'G'
		} else if char == 'G' {
			data[pos] = 'C'
		}
	}
}
