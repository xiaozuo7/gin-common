package service

import (
	"com.github.gin-common/app/model"
)

type UserService interface {
	// 创建用户
	CreateUser(user *model.User, password string) (*model.User, error)
	// 更新用户
	UpdateUser(id uint, updateInfo model.User) (*model.User, error)
	// 删除用户
	DeleteUser(id uint) error
	// 激活用户
	ActivateUser(id uint) (*model.User, error)
	// 禁用用户
	DeactivateUser(id uint) (*model.User, error)
	// 通过ID获取用户信息
	GetUserInfoById(id uint) (*model.User, error)
	// 通过ID修改用户密码
	ChangePassword(id uint, oldPass string, newPass string) error
	// 通过用户名获取用户信息
	GetUserInfoByUserName(username string) (*model.User, error)
}
