package main

import (
	"fmt"
	"os"
)

var specFile string

func findOpenAPIFile() (string, error) {
	// Wenn per Flag angegeben, direkt prüfen
	if specFile != "" {
		if _, err := os.Stat(specFile); err == nil {
			return specFile, nil
		}
		return "", fmt.Errorf("specified file not found: %s", specFile)
	}

	// Sonst Kandidatenliste durchsuchen
	candidates := []string{
		"OpenAPI.yaml",
		"openapi.yaml",
		"openapi.yml",
		"OpenAPI.yml",
	}

	for _, file := range candidates {
		if _, err := os.Stat(file); err == nil {
			return file, nil
		}
	}

	return "", fmt.Errorf("no openapi file found, use -f <file> to specify one")
}