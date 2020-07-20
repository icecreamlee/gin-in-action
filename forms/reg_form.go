package forms

type RegForm struct {
	Username   string `form:"username" binding:"required,min=2,max=20"`
	Password   string `form:"password" binding:"required,min=6,max=20"`
	RePassword string `form:"repassword" binding:"required,min=6,max=20,eqcsfield=Password"`
}
