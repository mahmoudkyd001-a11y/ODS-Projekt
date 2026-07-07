package main

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/mark3labs/mcp-go/mcp"
)

func validateHandler(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	specPath := req.GetString("spec_path", "")
	if specPath == "" {
		var err error
		specPath, err = findOpenAPIFile()
		if err != nil {
			return &mcp.CallToolResult{
				Content: []mcp.Content{mcp.NewTextContent("No spec file specified and no default found: " + err.Error())},
				IsError: true,
			}, nil
		}
	}

	loader := openapi3.NewLoader()
	doc, err := loader.LoadFromFile(specPath)
	if err != nil {
		return &mcp.CallToolResult{
			Content: []mcp.Content{mcp.NewTextContent("Failed to load spec: " + err.Error())},
			IsError: true,
		}, nil
	}

	err = doc.Validate(ctx)
	if err != nil {
		return &mcp.CallToolResult{
			Content: []mcp.Content{mcp.NewTextContent("Validation errors:\n" + err.Error())},
		}, nil
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{mcp.NewTextContent(fmt.Sprintf("Spec %q is valid. No errors found.", specPath))},
	}, nil
}

func listExamplesHandler(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	dir := examplesPath

	var results []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		ext := strings.ToLower(filepath.Ext(path))
		if ext == ".yaml" || ext == ".yml" || ext == ".json" {
			specType := detectSpecTypeFromFile(path)
			results = append(results, fmt.Sprintf("- %s (%s)", path, specType))
		}
		return nil
	})
	if err != nil {
		return &mcp.CallToolResult{
			Content: []mcp.Content{mcp.NewTextContent("Failed to list examples: " + err.Error())},
			IsError: true,
		}, nil
	}

	if len(results) == 0 {
		return &mcp.CallToolResult{
			Content: []mcp.Content{mcp.NewTextContent("No example specs found in " + dir)},
		}, nil
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{mcp.NewTextContent("Available example specs:\n" + strings.Join(results, "\n"))},
	}, nil
}

func detectSpecTypeFromFile(path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		return "unknown"
	}
	n := len(data)
	if n > 1024 {
		n = 1024
	}
	text := strings.ToLower(string(data[:n]))
	if strings.Contains(text, "asyncapi") {
		return "AsyncAPI"
	}
	if strings.Contains(text, "openapi") || strings.Contains(text, "swagger") {
		return "OpenAPI"
	}
	return "unknown"
}

func generateHandler(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	specPath, err := req.RequireString("spec_path")
	if err != nil {
		return &mcp.CallToolResult{
			Content: []mcp.Content{mcp.NewTextContent("spec_path is required")},
			IsError: true,
		}, nil
	}

	outputPath := req.GetString("output_path", "src")
	projectName := req.GetString("project_name", "default")
	database := req.GetBool("database", false)
	frontend := req.GetBool("frontend", false)

	dredgerBin, err := findDredgerBinary()
	if err != nil {
		return &mcp.CallToolResult{
			Content: []mcp.Content{mcp.NewTextContent("Could not find dredger binary: " + err.Error())},
			IsError: true,
		}, nil
	}

	args := []string{"generate", specPath, "-o", outputPath, "-n", projectName}
	if database {
		args = append(args, "-D")
	}
	if frontend {
		args = append(args, "-f")
	}

	cmd := exec.CommandContext(ctx, dredgerBin, args...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return &mcp.CallToolResult{
			Content: []mcp.Content{mcp.NewTextContent(fmt.Sprintf("Generation failed:\n%s\n%s", err.Error(), string(output)))},
			IsError: true,
		}, nil
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{mcp.NewTextContent(fmt.Sprintf("Project generated successfully at %s\n%s", outputPath, string(output)))},
	}, nil
}

func findDredgerBinary() (string, error) {
	self, err := os.Executable()
	if err != nil {
		return "", err
	}
	projectRoot := filepath.Dir(filepath.Dir(self))

	candidates := []string{
		filepath.Join(projectRoot, "build", "dredger"),
		filepath.Join(projectRoot, "build", "dredger.exe"),
		filepath.Join(projectRoot, "dredger"),
		filepath.Join(projectRoot, "dredger.exe"),
	}
	for _, c := range candidates {
		if _, err := os.Stat(c); err == nil {
			return c, nil
		}
	}

	if p, err := exec.LookPath("dredger"); err == nil {
		return p, nil
	}

	return "", fmt.Errorf("dredger binary not found in %s/build/ or PATH", projectRoot)
}
