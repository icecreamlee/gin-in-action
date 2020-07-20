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

func Index(c *gin.Context) {
	c.JSON(200, core.ApiSuccess(
		core.CodeSuccess, "success", struct {
			ID   uint   `gorm:"primary_key",json:"id"`
			Name string `json:"name"`
		}{loginUser.ID, loginUser.Name},
	))
}

func Reg(c *gin.Context) {
	var regForm forms.RegForm
	err := c.ShouldBindJSON(&regForm)
	if err != nil {
		c.JSON(200, core.ApiFailure(err.Error()))
		return
	}
	c.JSON(200, logics.Reg(regForm))
}

func Login(c *gin.Context) {
	var loginForm forms.LoginForm
	err := c.BindJSON(&loginForm)
	if err != nil {
		c.JSON(200, core.ApiFailure(err.Error()))
		return
	}
	c.JSON(200, logics.Login(loginForm))
}

// ManagePermissionValidation 验证管理相关页面的密码不正确则显示404页
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
