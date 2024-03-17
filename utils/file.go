package utils

import (
	"errors"
	"flag"
	"io"
	"os"
)

func ReadFileFromArg() ([]byte, error) {
	filePath := flag.Arg(0)
	if filePath == "" {
		return nil, errors.New("no file provided")
	}

	f, err := os.Open(filePath)
	defer f.Close()
	if err != nil {
		return nil, errors.New("unable to open file path")
	}

	data, err := io.ReadAll(f)
	if err != nil {
		return nil, errors.New("unable to read file")
	}

	return data, nil
}

func ReadFileFromArgStripTrailingNewline() ([]byte, error) {
	data, err := ReadFileFromArg()
	if err != nil {
		return nil, err
	}

	if data[len(data)-1] == '\n' {
		data = data[:len(data)-1]
	}

	return data, nil
}
