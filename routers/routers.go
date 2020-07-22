package routers

import (
	"gin_in_action/controllers"
	"github.com/gin-gonic/gin"
)

func SetRouters(r *gin.Engine) {
	r.POST("/reg", controllers.Reg)
	r.POST("/login", controllers.Login)
	r.GET("/logout", controllers.Logout)

	r.GET("/", controllers.Index)
	r.GET("/index", controllers.CheckLogin, controllers.Index)
	r.GET("/product", controllers.CheckLogin, controllers.Product)
	r.GET("/products", controllers.CheckLogin, controllers.ProductList)
}
