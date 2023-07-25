package src

import (
	"io"
	"os"
	"sport_news/config"
	"sport_news/src/api/category"
	user "sport_news/src/api/user"
	"sport_news/src/tools"

	"github.com/gin-gonic/gin"
)

func Init_app() *gin.Engine {
	set_mode(config.ENV.GIN_MODE)
	router := gin.New()
	router.Use(gin.Logger())
	router.SetTrustedProxies(nil)
	router.Static("/public", config.ENV.DIR)
	router.Use(tools.Cors)
	set_routers(router)
	return router
}

func set_mode(mode string) {
	switch mode {
	case "debug":
		gin.SetMode(gin.DebugMode)
	case "release":
		gin.SetMode(gin.ReleaseMode)
		gin.DisableConsoleColor()
		file, _ := os.Create("gin.log")
		gin.DefaultWriter = io.MultiWriter(file)
	}

}

func set_routers(router *gin.Engine) {
	create_router(router, "user", user.Controller)
	create_router(router, "category", category.Controller)

}

func create_router(
	router *gin.Engine,
	name string,
	controller func(router *gin.RouterGroup),
) {
	group_name := router.Group(name)
	controller(group_name)
}
