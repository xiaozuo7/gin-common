package model

import (
	"com.github.gin-common/common/models"
	"com.github.gin-common/util"
)

type User struct {
	models.BaseModel
	models.SoftDeleteModel
	models.ActivateModel
	Username string `gorm:"unique;not null;size:256" json:"username"`
	Name     string `gorm:"size:256;not null;default:''" json:"name"`
	Password string `gorm:"size:256;not null;" json:"-"`
	Email    string `gorm:"size:256" json:"email"`
}

func (user *User) SetPass(pwd string) error {
	// 修改密码
	encryptPass, err := util.HashAndSalt([]byte(pwd))
	if err != nil {
		return err
	}
	user.Password = encryptPass
	return nil
}

func (user *User) CheckPass(pwd string) bool {
	return util.CompareHash(user.Password, []byte(pwd))
}
