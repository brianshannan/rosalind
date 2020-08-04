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

	entries := utils.ParseFasta(string(data))

	prefixMap := map[string][]string{}
	for _, entry := range entries {
		prefix := entry.Sequence[:3]
		if _, ok := prefixMap[prefix]; ok {
			prefixMap[prefix] = append(prefixMap[prefix], entry.Name)
		} else {
			prefixMap[prefix] = []string{entry.Name}
		}
	}

	for _, entry := range entries {
		suffix := entry.Sequence[len(entry.Sequence)-3:]
		if matches, ok := prefixMap[suffix]; ok {
			for _, match := range matches {
				if entry.Name == match {
					continue
				}
				fmt.Printf("%s %s\n", entry.Name, match)
			}
		}
	}
}
