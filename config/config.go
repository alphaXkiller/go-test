package config

import (
	"fmt"
	// "net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"rr/research-tool/routes"
)

type App struct {
	Router *gin.Engine
	DB     *gorm.DB
}

func (a *App) Initialize() {
	db, err := gorm.Open("mysql", "root:root@/research_tool?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic(fmt.Sprintf("database connection error: %v \n", err))
	}

	a.DB = db
	a.Router = gin.Default()
	a.registerRoutes()
}

func (a *App) registerRoutes() {
	a.Router.GET("/", routes.Index)
}
