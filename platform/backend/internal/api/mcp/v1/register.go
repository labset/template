package api_mcp_v1

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type Dependencies struct {
	// Add any dependencies here if needed in the future
}

func Register(mcpGroup *gin.RouterGroup, _ Dependencies) {
	log.Info().Msg("Registering MCP server routes")

	mcpHandler := newHandler()

	mcpGroup.Any("/*any", func(c *gin.Context) {
		mcpHandler.wrapped.ServeHTTP(c.Writer, c.Request)
	})
}
