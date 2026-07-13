package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

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

func schemaDetail(name string) (string, error) {
	spec, err := loadSpec()
	if err != nil {
		return "", err
	}

	if spec.Components == nil {
		return "", fmt.Errorf("spec has no components")
	}

	schemaRef, ok := spec.Components.Schemas[name]
	if !ok {
		available := make([]string, 0, len(spec.Components.Schemas))
		for k := range spec.Components.Schemas {
			available = append(available, k)
		}
		return "", fmt.Errorf("schema %q not found. Available schemas: %s", name, strings.Join(available, ", "))
	}

	data, err := json.MarshalIndent(schemaRef, "", "  ")
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Schema: %s\n\n%s", name, string(data)), nil
}

func endpointDetail(method, path string) (string, error) {
	spec, err := loadSpec()
	if err != nil {
		return "", err
	}

	if spec.Paths == nil {
		return "", fmt.Errorf("spec has no paths")
	}

	pathItem := spec.Paths.Find(path)
	if pathItem == nil {
		available := make([]string, 0)
		for p := range spec.Paths.Map() {
			available = append(available, p)
		}
		return "", fmt.Errorf("path %q not found. Available paths: %s", path, strings.Join(available, ", "))
	}

	var op *openapi3.Operation
	switch strings.ToUpper(method) {
	case "GET":
		op = pathItem.Get
	case "POST":
		op = pathItem.Post
	case "PUT":
		op = pathItem.Put
	case "DELETE":
		op = pathItem.Delete
	case "PATCH":
		op = pathItem.Patch
	}

	if op == nil {
		return "", fmt.Errorf("method %s not found for path %q", strings.ToUpper(method), path)
	}

	data, err := json.MarshalIndent(op, "", "  ")
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Endpoint: %s %s\n\n%s", strings.ToUpper(method), path, string(data)), nil
}

func parseEndpointURI(uri string) (method, path string) {
	rest := strings.TrimPrefix(uri, "openapi://endpoint/")
	parts := strings.SplitN(rest, "/", 2)
	if len(parts) == 2 {
		return parts[0], "/" + parts[1]
	}
	return "", ""
}
