package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/brianshannan/rosalind/utils"
)

func main() {
	flag.Parse()
	data, err := utils.ReadFileFromArg()
	if err != nil {
		panic(err)
	}

	segments := strings.Split(string(data), ">")
	segments = segments[1:]

	var (
		maxName       string
		maxPercentage float64
	)

	for _, segment := range segments {
		lines := strings.Split(segment, "\n")
		name := lines[0]
		dnaString := strings.Join(lines[1:], "")

		cgCount := 0
		for _, char := range dnaString {
			if char == 'C' || char == 'G' {
				cgCount++
			}
		}
		percentage := float64(cgCount) / float64(len(dnaString))
		if percentage > maxPercentage {
			maxPercentage = percentage
			maxName = name
		}
	}
	fmt.Printf("%s\n%f\n", maxName, 100*maxPercentage)
}
