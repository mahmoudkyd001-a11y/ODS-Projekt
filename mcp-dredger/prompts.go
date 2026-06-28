package main

import (
	"context"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
)

func explainSpecHandler(ctx context.Context, req mcp.GetPromptRequest) (*mcp.GetPromptResult, error) {
	specType := req.Params.Arguments["spec_type"]
	if specType == "" {
		specType = "openapi"
	}

	var summaryText string
	var err error

	switch specType {
	case "asyncapi":
		summaryText, err = asyncSummary()
	default:
		summaryText, err = summary()
	}

	if err != nil {
		return nil, fmt.Errorf("failed to load spec summary: %w", err)
	}

	return &mcp.GetPromptResult{
		Description: "Explain the API specification",
		Messages: []mcp.PromptMessage{
			{
				Role: mcp.RoleUser,
				Content: mcp.NewTextContent(fmt.Sprintf(
					"Here is a summary of the %s specification:\n\n%s\n\nPlease explain this API in clear, simple terms. "+
						"Describe what the API does, what endpoints/operations are available, and what data models it uses. "+
						"If there are any notable design patterns or conventions, mention those too.",
					specType, summaryText,
				)),
			},
		},
	}, nil
}

func suggestImprovementsHandler(ctx context.Context, req mcp.GetPromptRequest) (*mcp.GetPromptResult, error) {
	specType := req.Params.Arguments["spec_type"]
	if specType == "" {
		specType = "openapi"
	}

	var specContent string
	var err error

	switch specType {
	case "asyncapi":
		specContent, err = fullAsyncSpec()
	default:
		specContent, err = fullSpec()
	}

	if err != nil {
		return nil, fmt.Errorf("failed to load spec: %w", err)
	}

	return &mcp.GetPromptResult{
		Description: "Suggest improvements for the API specification",
		Messages: []mcp.PromptMessage{
			{
				Role: mcp.RoleUser,
				Content: mcp.NewTextContent(fmt.Sprintf(
					"Here is the full %s specification:\n\n%s\n\n"+
						"Please analyze this API specification and suggest improvements. Consider:\n"+
						"- Missing or incomplete descriptions\n"+
						"- Inconsistent naming conventions\n"+
						"- Missing error responses\n"+
						"- Missing validation rules\n"+
						"- Security considerations\n"+
						"- Best practices for %s specifications\n\n"+
						"Provide specific, actionable suggestions.",
					specType, specContent, specType,
				)),
			},
		},
	}, nil
}
