package forms

type LoginForm struct {
	Username string `json:"username" form:"username" binding:"required,min=6,max=20"`
	Password string `json:"password" form:"password" binding:"required,min=6,max=20"`
}
