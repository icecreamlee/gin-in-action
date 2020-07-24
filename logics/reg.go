package logics

import (
	"fmt"
	"gin_in_action/core"
	"gin_in_action/forms"
	"gin_in_action/models"
	"github.com/IcecreamLee/goutils"
	"math/rand"
	"strconv"
)

func Reg(form forms.RegForm) core.ApiError {
	user := models.GetUserByName(form.Username)
	fmt.Printf("%+v\n", user)
	if user.ID > 0 {
		return core.ApiError{Msg: "username exist"}
	}

	passsalt := strconv.Itoa(rand.Intn(899999) + 100000)
	user = models.User{
		Name:     form.Username,
		Password: goutils.MD5Str(goutils.MD5Str(form.Password) + passsalt),
		Passsalt: passsalt,
		IsDelete: 0,
	}

	if err := models.DB.Create(&user).Error; err != nil {
		return core.ApiFailure(core.CodeFailure, "failed to create user", err.Error())
	}
	return core.ApiSuccess()
}
