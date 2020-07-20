package main

import (
	"gin_in_action/configs"
	"gin_in_action/core"
	"gin_in_action/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(configs.ENV)
	router := gin.Default()
	routers.SetRouters(router)
	router.Static("/static", core.RootPath+"statics")
	//router.StaticFile("/favicon.ico", core.RootPath+"web/statics/favicon.ico")
	//router.LoadHTMLGlob(core.RootPath + "web/views/*")
	router.Run(":" + configs.Port) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
