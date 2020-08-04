package utils

import "strings"

type FASTAEntry struct {
	Name     string
	Sequence string
}

func ParseFasta(s string) []FASTAEntry {
	entries := []FASTAEntry{}
	segments := strings.Split(s, ">")
	segments = segments[1:]

	for _, segment := range segments {
		lines := strings.Split(segment, "\n")
		name := lines[0]
		sequence := strings.Join(lines[1:], "")
		entries = append(entries, FASTAEntry{
			Name:     name,
			Sequence: sequence,
		})
	}

	return entries
}
