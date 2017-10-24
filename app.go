package main

import (
	// "database/sql"
	// "net/http"

	"github.com/gin-gonic/gin"
)

type App struct {
	Router *gin.Engine
}

func (a *App) Initialize() {

	a.Router = gin.Default()
	a.registerRoutes()
}

func (a *App) registerRoutes() {
	a.Router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to go API",
		})
	})
}
