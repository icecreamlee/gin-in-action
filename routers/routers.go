package routers

import (
	"gin_in_action/configs"
	"gin_in_action/controllers"
	"gin_in_action/core"
	"github.com/gin-gonic/gin"
	"time"
)

func SetRouters(r *gin.Engine) {
	if configs.RateLimit > 0 {
		r.Use(RateLimiter())
	}

	r.POST("/reg", controllers.Reg)
	r.POST("/login", controllers.Login)
	r.GET("/logout", controllers.Logout)

	r.GET("/", controllers.Index)
	r.GET("/index", controllers.CheckLogin, controllers.Index)
	r.GET("/product", controllers.CheckLogin, controllers.Product)
	r.GET("/products", controllers.CheckLogin, controllers.ProductList)
	r.POST("/add-product", controllers.CheckLogin, controllers.AddProduct)
	r.POST("/update-product", controllers.CheckLogin, controllers.UpdateProduct)
}

// RateLimiter 限流器中间件
func RateLimiter() gin.HandlerFunc {
	// 限流器
	rateLimiter := core.NewTokenBucket(time.Second/time.Duration(configs.RateLimit), configs.RateLimit)
	return func(context *gin.Context) {
		if rateLimiter.Take() {
			context.Next()
		} else {
			context.String(503, "503 Service Unavailable")
			context.Abort()
		}
	}
}
