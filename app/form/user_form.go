package form

type CreateUserForm struct {
	UserName string `binding:"required" json:"username"`
	Password string `binding:"required" json:"password"`
	Name     string `json:"name"`
	Email    string `binding:"omitempty,email" json:"email"`
}

type UpdateUserForm struct {
	UserName string `binding:"required_without_all=Name Email" json:"username"`
	Name     string `binding:"required_without_all=UserName Email" json:"name"`
	Email    string `binding:"required_without_all=UserName Name,omitempty,email" json:"email"`
}

type ChangePassForm struct {
	OldPassword string `binding:"required,nefield=NewPassword" json:"old_password"`
	NewPassword string `binding:"required" json:"new_password"`
}
