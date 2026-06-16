package main

import (
	"context"
	"flag"
	"log"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func main() {
	flag.StringVar(&specFile, "f", "", "path to OpenAPI spec file")
	flag.Parse()

	s := server.NewMCPServer(
		"mcp-dredger",
		"0.1.0",
	)

	s.AddResource(mcp.NewResource(
		"openapi://spec",
		"Full OpenAPI specification",
	), func(ctx context.Context, req mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
		content, err := fullSpec()
		if err != nil {
			return nil, err
		}
		return []mcp.ResourceContents{
			mcp.TextResourceContents{
				URI:      "openapi://spec",
				MIMEType: "application/yaml",
				Text:     content,
			},
		}, nil
	})

	s.AddResource(mcp.NewResource(
		"openapi://summary",
		"Summary of schemas and endpoints",
	), func(ctx context.Context, req mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
		content, err := summary()
		if err != nil {
			return nil, err
		}
		return []mcp.ResourceContents{
			mcp.TextResourceContents{
				URI:      "openapi://summary",
				MIMEType: "text/plain",
				Text:     content,
			},
		}, nil
	})

	log.Fatal(server.ServeStdio(s))
}