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

	lines := strings.Split(string(data), "\n")
	s := lines[0]
	sub := lines[1]
	matches := []string{}

	for i := 0; i < len(s); i++ {
		if checkMatch(s[i:], sub) {
			matches = append(matches, fmt.Sprintf("%d", i+1))
		}
	}

	fmt.Printf("%s\n", strings.Join(matches, " "))
}

func checkMatch(s, sub string) bool {
	if len(s) < len(sub) {
		return false
	}

	for i := 0; i < len(sub); i++ {
		if s[i] != sub[i] {
			return false
		}
	}

	return true
}
