package utils

import (
	"errors"
	"flag"
	"io/ioutil"
	"os"
)

func ReadFileFromArg() (string, error) {
	filePath := flag.Arg(0)
	if filePath == "" {
		return "", errors.New("no file provided")
	}

	f, err := os.Open(filePath)
	if err != nil {
		return "", errors.New("unable to open file path")
	}

	data, err := ioutil.ReadAll(f)
	if err != nil {
		return "", errors.New("unable to read file")
	}

	return string(data), nil
}
