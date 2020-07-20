package logics

import (
	"gin_in_action/core"
	"gin_in_action/forms"
	"gin_in_action/models"
	"github.com/IcecreamLee/goutils"
	"github.com/gin-gonic/gin"
	"time"
)

func Login(form forms.LoginForm) core.ApiError {
	user := models.GetUserByName(form.Username)
	if user.ID == 0 {
		return core.ApiFailure("password error")
	}

	password := goutils.MD5Str(goutils.MD5Str(form.Password) + user.Passsalt)
	if password != user.Password {
		return core.ApiFailure("password error")
	}

	token := goutils.MD5Str(goutils.IntToString(int64(user.ID) + time.Now().UnixNano()))

	core.RedisSet("token_"+token, user.ID, time.Hour*24)

	return core.ApiSuccess(core.CodeSuccess, "success", gin.H{"token": token})
}
