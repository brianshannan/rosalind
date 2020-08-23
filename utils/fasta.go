package utils

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

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

func DownloadFasta(id string) ([]FASTAEntry, error) {
	resp, err := http.Get(fmt.Sprintf("https://www.uniprot.org/uniprot/%s.fasta", id))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 300 {
		return nil, fmt.Errorf("received status code %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return ParseFasta(string(body)), nil
}
