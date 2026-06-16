package main

import (
	"os"

	"github.com/getkin/kin-openapi/openapi3"
)

func loadSpec() (*openapi3.T, error) {
	file, err := findOpenAPIFile()
	if err != nil {
		return nil, err
	}

	loader := openapi3.NewLoader()

	return loader.LoadFromFile(file)
}

func fullSpec() (string, error) {
	file, err := findOpenAPIFile()
	if err != nil {
		return "", err
	}

	b, err := os.ReadFile(file)
	if err != nil {
		return "", err
	}

	return string(b), nil
}