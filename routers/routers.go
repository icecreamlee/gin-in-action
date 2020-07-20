package routers

import (
	"gin_in_action/controllers"
	"github.com/gin-gonic/gin"
)

func SetRouters(r *gin.Engine) {
	r.GET("/", controllers.Index)
	r.GET("/index", controllers.CheckLogin, controllers.Index)
	r.POST("/login", controllers.Login)
	r.POST("/reg", controllers.Reg)
}
