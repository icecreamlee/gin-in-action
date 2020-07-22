package controllers

import (
	"gin_in_action/core"
	"gin_in_action/forms"
	"gin_in_action/logics"
	"gin_in_action/models"
	"github.com/IcecreamLee/goutils"
	"github.com/gin-gonic/gin"
)

var loginUser models.User

// Index 首页接口
func Index(c *gin.Context) {
	products := models.GetRecentProducts()
	user := models.GetUser(loginUser.ID)
	c.JSON(200, core.ApiSuccess(core.CodeSuccess, "success", gin.H{"user": user, "products": products}))
}

// Reg 注册接口
func Reg(c *gin.Context) {
	var regForm forms.RegForm
	err := c.ShouldBindJSON(&regForm)
	if err != nil {
		c.JSON(200, core.ApiFailure(err.Error()))
		return
	}
	c.JSON(200, logics.Reg(regForm))
}

// Login 登录接口
func Login(c *gin.Context) {
	var loginForm forms.LoginForm
	err := c.BindJSON(&loginForm)
	if err != nil {
		c.JSON(200, core.ApiFailure(err.Error()))
		return
	}
	c.JSON(200, logics.Login(loginForm))
}

// Logout 退出登录
func Logout(c *gin.Context) {
	c.JSON(200, logics.Logout(c.GetHeader("token")))
}

// CheckLogin 检查登录
func CheckLogin(c *gin.Context) {
	token := c.GetHeader("token")
	loginUser.ID = uint(goutils.ToInt(core.RedisGet("token_" + token)))
	loginUser = models.GetUser(loginUser.ID)
	if loginUser.ID == 0 {
		c.JSON(200, core.ApiFailure(core.CodeLoginFirst, "login first"))
		c.Abort()
		return
	}
}
