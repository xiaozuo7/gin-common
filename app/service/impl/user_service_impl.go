package impl

import (
	"errors"
	"time"

	"com.github.gin-common/common/models"

	"com.github.gin-common/app/exception"
	"com.github.gin-common/app/model"
	"com.github.gin-common/common/exceptions"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

type UserServiceImpl struct {
	session *gorm.DB
}

func (service *UserServiceImpl) Init(session *gorm.DB) {
	service.session = session
}

func (service *UserServiceImpl) CreateUser(user *model.User, password string) (*model.User, error) {
	// 创建用户
	e := user.SetPass(password)
	if e != nil {
		return nil, exceptions.GetDefinedErrors(exception.UserCreateFailed)
	}
	user.ActivateAt = time.Now()
	result := service.session.Create(user)
	if result.Error != nil {
		if err, ok := result.Error.(*mysql.MySQLError); ok {
			if err.Number == uint16(1062) {
				return nil, exceptions.GetDefinedErrors(exception.UserNameDuplicate)
			}
		}
		return nil, exceptions.GetDefinedErrors(exception.UserCreateFailed)
	}

	return user, nil
}

func (service *UserServiceImpl) UpdateUser(id uint, updateInfo model.User) (*model.User, error) {
	// 更新用户

	var user, err = service.GetUserInfoById(id)
	if err != nil {
		return nil, err
	}
	if result := service.session.Model(user).Updates(updateInfo); result.Error != nil {
		return nil, exceptions.GetDefinedErrors(exception.UserUpdateFailed)
	}
	return user, nil
}

func (service *UserServiceImpl) DeleteUser(id uint) error {
	// 删除用户
	user, err := service.GetUserInfoById(id)
	if err != nil {
		return err
	}
	if result := models.Delete(service.session, id, user); result.Error != nil {
		return exceptions.GetDefinedErrors(exception.DeleteUserFailed)
	}
	return nil
}

func (service *UserServiceImpl) ActivateUser(id uint) (*model.User, error) {
	// 启用用户
	user, err := service.GetUserInfoById(id)
	if err != nil {
		return nil, err
	}
	if result := models.Activate(service.session, id, user); result.Error != nil {
		return nil, exceptions.GetDefinedErrors(exception.ActivateUserFailed)
	}
	return user, nil
}

func (service *UserServiceImpl) DeactivateUser(id uint) (*model.User, error) {
	// 禁用用户
	user, err := service.GetUserInfoById(id)
	if err != nil {
		return nil, err
	}
	if result := models.Deactivate(service.session, id, user); result.Error != nil {
		return nil, exceptions.GetDefinedErrors(exception.DeActivateUserFailed)
	}
	return user, nil
}

func (service *UserServiceImpl) GetUserInfoById(id uint) (*model.User, error) {
	user := &model.User{}
	if result := service.session.Where("ID=?", id).Take(user); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, exceptions.GetDefinedErrors(exception.UserNotFound)
		}
		return nil, result.Error
	}
	return user, nil
}

func (service *UserServiceImpl) GetUserInfoByUserName(username string) (*model.User, error) {
	// 通过用户名获取用户
	user := &model.User{}
	if result := service.session.Where("username=?", username).Take(user); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, exceptions.GetDefinedErrors(exception.UserNotFound)
		}
		return nil, result.Error
	}
	return user, nil
}

func (service *UserServiceImpl) ChangePassword(id uint, oldPass string, newPass string) error {
	// 修改密码
	user, err := service.GetUserInfoById(id)
	if err != nil {
		return err
	}
	if user.CheckPass(oldPass) {
		if err := user.SetPass(newPass); err != nil {
			return exceptions.GetDefinedErrors(exception.ChangePassFailed)
		}
	} else {
		return exceptions.GetDefinedErrors(exception.OldPassInvalid)
	}
	service.session.Save(user)
	return nil
}
