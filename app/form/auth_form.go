package form

type LoginForm struct {
	UserName string `binding:"required" json:"username"`
	Password string `binding:"required" json:"password"`
}
