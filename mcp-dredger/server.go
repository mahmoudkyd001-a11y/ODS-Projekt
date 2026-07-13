package main

import (
	"context"
	"flag"
	"log"
	"strings"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

var examplesPath string

func main() {
	flag.StringVar(&specFile, "f", "", "path to OpenAPI spec file")
	flag.StringVar(&asyncSpecFile, "a", "", "path to AsyncAPI spec file")
	flag.StringVar(&examplesPath, "examples", "examples", "path to examples directory")
	flag.Parse()

	s := server.NewMCPServer(
		"mcp-dredger",
		"0.2.0",
	)

	// --- Resources ---

	s.AddResource(mcp.NewResource(
		"openapi://spec",
		"Full OpenAPI specification",
		mcp.WithMIMEType("application/yaml"),
	), func(ctx context.Context, req mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
		content, err := fullSpec()
		if err != nil {
			return nil, err
		}
		return []mcp.ResourceContents{
			mcp.TextResourceContents{URI: "openapi://spec", MIMEType: "application/yaml", Text: content},
		}, nil
	})

	s.AddResource(mcp.NewResource(
		"openapi://summary",
		"Summary of schemas and endpoints",
		mcp.WithMIMEType("text/plain"),
	), func(ctx context.Context, req mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
		content, err := summary()
		if err != nil {
			return nil, err
		}
		return []mcp.ResourceContents{
			mcp.TextResourceContents{URI: "openapi://summary", MIMEType: "text/plain", Text: content},
		}, nil
	})

	s.AddResource(mcp.NewResource(
		"asyncapi://spec",
		"Full AsyncAPI specification",
		mcp.WithMIMEType("application/yaml"),
	), func(ctx context.Context, req mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
		content, err := fullAsyncSpec()
		if err != nil {
			return nil, err
		}
		return []mcp.ResourceContents{
			mcp.TextResourceContents{URI: "asyncapi://spec", MIMEType: "application/yaml", Text: content},
		}, nil
	})

	s.AddResource(mcp.NewResource(
		"asyncapi://summary",
		"Summary of AsyncAPI channels, operations and messages",
		mcp.WithMIMEType("text/plain"),
	), func(ctx context.Context, req mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
		content, err := asyncSummary()
		if err != nil {
			return nil, err
		}
		return []mcp.ResourceContents{
			mcp.TextResourceContents{URI: "asyncapi://summary", MIMEType: "text/plain", Text: content},
		}, nil
	})

	// --- Resource Templates ---

	s.AddResourceTemplate(
		mcp.NewResourceTemplate("openapi://schema/{name}", "OpenAPI Schema Detail",
			mcp.WithTemplateDescription("Detailed view of a specific OpenAPI schema with all properties, types and extensions"),
			mcp.WithTemplateMIMEType("application/json"),
		),
		func(ctx context.Context, req mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
			name := strings.TrimPrefix(req.Params.URI, "openapi://schema/")
			content, err := schemaDetail(name)
			if err != nil {
				return nil, err
			}
			return []mcp.ResourceContents{
				mcp.TextResourceContents{URI: req.Params.URI, MIMEType: "application/json", Text: content},
			}, nil
		},
	)

	s.AddResourceTemplate(
		mcp.NewResourceTemplate("openapi://endpoint/{method}/{+path}", "OpenAPI Endpoint Detail",
			mcp.WithTemplateDescription("Detailed view of a specific OpenAPI endpoint with parameters, request body and responses"),
			mcp.WithTemplateMIMEType("application/json"),
		),
		func(ctx context.Context, req mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
			method, path := parseEndpointURI(req.Params.URI)
			content, err := endpointDetail(method, path)
			if err != nil {
				return nil, err
			}
			return []mcp.ResourceContents{
				mcp.TextResourceContents{URI: req.Params.URI, MIMEType: "application/json", Text: content},
			}, nil
		},
	)

	// --- Tools ---

	s.AddTool(
		mcp.NewTool("validate",
			mcp.WithDescription("Validate an OpenAPI specification file and report any errors"),
			mcp.WithString("spec_path", mcp.Description("Path to the spec file (uses default if not specified)")),
			mcp.WithReadOnlyHintAnnotation(true),
			mcp.WithDestructiveHintAnnotation(false),
		),
		validateHandler,
	)

	s.AddTool(
		mcp.NewTool("list-examples",
			mcp.WithDescription("List available example API specifications from the examples directory"),
			mcp.WithReadOnlyHintAnnotation(true),
			mcp.WithDestructiveHintAnnotation(false),
		),
		listExamplesHandler,
	)

	s.AddTool(
		mcp.NewTool("generate",
			mcp.WithDescription("Generate a Go microservice from an OpenAPI or AsyncAPI specification using Dredger"),
			mcp.WithString("spec_path", mcp.Description("Path to the spec file"), mcp.Required()),
			mcp.WithString("output_path", mcp.Description("Output directory for generated code"), mcp.DefaultString("src")),
			mcp.WithString("project_name", mcp.Description("Name of the generated project"), mcp.DefaultString("default")),
			mcp.WithBoolean("database", mcp.Description("Include SQLite database support"), mcp.DefaultBool(false)),
			mcp.WithBoolean("frontend", mcp.Description("Include frontend/templ support"), mcp.DefaultBool(false)),
		),
		generateHandler,
	)

	// --- Prompts ---

	s.AddPrompt(
		mcp.NewPrompt("explain-spec",
			mcp.WithPromptDescription("Explain the API specification in clear, simple terms"),
			mcp.WithArgument("spec_type",
				mcp.ArgumentDescription("Type of spec: 'openapi' or 'asyncapi' (default: openapi)"),
			),
		),
		explainSpecHandler,
	)

	s.AddPrompt(
		mcp.NewPrompt("suggest-improvements",
			mcp.WithPromptDescription("Analyze the API specification and suggest improvements"),
			mcp.WithArgument("spec_type",
				mcp.ArgumentDescription("Type of spec: 'openapi' or 'asyncapi' (default: openapi)"),
			),
		),
		suggestImprovementsHandler,
	)

	log.Fatal(server.ServeStdio(s))
}
