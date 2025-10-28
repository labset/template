package api_mcp_v1

import (
	"net/http"

	"github.com/rs/zerolog/log"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type Tool interface {
	GetTool() (*mcp.Tool, error)
	GetHandler() mcp.ToolHandler
}

type handler struct {
	wrapped *mcp.StreamableHTTPHandler
}

func newHandler() *handler {
	// Create a new MCP server with HTTP wrapped
	server := mcp.NewServer(&mcp.Implementation{
		Name:    "poc-clarity-mcp",
		Version: "1.0.0",
	}, &mcp.ServerOptions{})

	registerTools(server)

	streamableHandler := mcp.NewStreamableHTTPHandler(
		func(request *http.Request) *mcp.Server {
			return server
		},
		&mcp.StreamableHTTPOptions{},
	)

	return &handler{
		wrapped: streamableHandler,
	}
}

func registerTools(server *mcp.Server) {
	tools := []Tool{
		newHelloWorldTool(),
		// Add more tools here
	}

	for _, tool := range tools {
		wrappedTool, err := tool.GetTool()
		if err != nil {
			log.Warn().Err(err).Msg("Failed to get tool")

			continue
		}

		server.AddTool(wrappedTool, tool.GetHandler())
	}
}
