package main

import (
	"net/http"
	"platform/backend/config"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

const (
	SessionName   = "labset.clarity"
	SessionMaxAge = 600 // 10 minutes in seconds
)

func setupRouter(cfg config.Config, deps *dependencies) *gin.Engine {
	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{cfg.Frontend.URL},
		AllowMethods: []string{"GET", "POST", "OPTIONS"},
		AllowHeaders: []string{
			"Origin",
			"Content-Length",
			"Content-Type",
			"Connect-Protocol-Version",
		},
		AllowCredentials: true,
	}))

	// Setup session store
	sessionStore := cookie.NewStore(
		[]byte(cfg.Session.Secret),
	)
	sessionStore.Options(sessions.Options{
		Path:     "/",
		Domain:   cfg.Session.Domain,
		MaxAge:   SessionMaxAge,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	})

	router.Use(sessions.Sessions(SessionName, sessionStore))

	// Setup routes

	return router
}
