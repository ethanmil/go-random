package main

import (
	"os"
	"strings"
)

type File []string

func NewFile(path string) (File, error) {
	fileBytes, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return strings.Split(string(fileBytes), "\n"), nil
}
