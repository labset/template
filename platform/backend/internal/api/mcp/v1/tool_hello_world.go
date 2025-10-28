package api_mcp_v1

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/google/jsonschema-go/jsonschema"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/rs/zerolog/log"
)

type helloWorldTool struct {
}

func newHelloWorldTool() Tool {
	return &helloWorldTool{}
}

type HelloInput struct {
	Name string `json:"name" jsonschema:"the name to say hello to"`
}

type HelloOutput struct {
	Message string `json:"message" jsonschema:"the hello message"`
}

func (h *helloWorldTool) GetTool() (*mcp.Tool, error) {
	inputSchema, err := jsonschema.For[HelloInput](&jsonschema.ForOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to generate input schema: %w", err)
	}

	outputSchema, err := jsonschema.For[HelloOutput](&jsonschema.ForOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to generate output schema: %w", err)
	}

	return &mcp.Tool{
		Name:         "hello-world-tool",
		Description:  "A simple Hello World tool",
		InputSchema:  inputSchema,
		OutputSchema: outputSchema,
	}, nil
}

var (
	errNameEmpty = errors.New("name parameter cannot be empty")
)

func (h *helloWorldTool) GetHandler() mcp.ToolHandler {
	return func(ctx context.Context, request *mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		// Marshal the arguments to JSON
		argsJSON, err := json.Marshal(request.Params.Arguments)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal arguments: %w", err)
		}

		// Unmarshal into typed HelloInput struct
		var input HelloInput
		if err = json.Unmarshal(argsJSON, &input); err != nil {
			return nil, fmt.Errorf("failed to unmarshal input: %w", err)
		}

		log.Info().Str("name", input.Name).Msg("Processing hello world tool")

		// Validate name is not empty
		if input.Name == "" {
			return nil, errNameEmpty
		}

		// Create greeting message using typed output
		output := HelloOutput{
			Message: fmt.Sprintf("Hello, %s! Welcome to POC Clarity MCP Server.", input.Name),
		}

		log.Info().Str("message", output.Message).Msg("Generated greeting")

		// Marshal output to JSON for structured response
		outputJSON, err := json.Marshal(output)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal output: %w", err)
		}

		// Return the result with text content
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				&mcp.TextContent{
					Text: string(outputJSON),
				},
			},
		}, nil
	}
}
