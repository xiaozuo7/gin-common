package exception

import (
	"net/http"

	"com.github.gin-common/common/exceptions"
)

func UserCreateFailed() *exceptions.ApiError {
	return &exceptions.ApiError{
		Code:          "200001",
		HttpCode:      http.StatusInternalServerError,
		DefaultErrMsg: "创建用户失败",
	}
}

func UserNameDuplicate() *exceptions.ApiError {
	return &exceptions.ApiError{
		Code:          "200002",
		HttpCode:      http.StatusInternalServerError,
		DefaultErrMsg: "用户名重复",
	}
}

func UserNotFound() *exceptions.ApiError {
	return &exceptions.ApiError{
		Code:          "200003",
		HttpCode:      http.StatusNotFound,
		DefaultErrMsg: "用户不存在",
	}
}

func DeleteUserFailed() *exceptions.ApiError {
	return &exceptions.ApiError{
		Code:          "200004",
		HttpCode:      http.StatusInternalServerError,
		DefaultErrMsg: "删除用户失败",
	}
}

func ActivateUserFailed() *exceptions.ApiError {
	return &exceptions.ApiError{
		Code:          "200005",
		HttpCode:      http.StatusInternalServerError,
		DefaultErrMsg: "启用用户失败",
	}
}

func DeActivateUserFailed() *exceptions.ApiError {
	return &exceptions.ApiError{
		Code:          "200005",
		HttpCode:      http.StatusInternalServerError,
		DefaultErrMsg: "禁用用户失败",
	}
}

func UserUpdateFailed() *exceptions.ApiError {
	return &exceptions.ApiError{
		Code:          "200006",
		HttpCode:      http.StatusInternalServerError,
		DefaultErrMsg: "更新用户失败",
	}
}

func OldPassInvalid() *exceptions.ApiError {
	return &exceptions.ApiError{
		Code:          "200007",
		HttpCode:      http.StatusInternalServerError,
		DefaultErrMsg: "旧密码错误",
	}
}

func ChangePassFailed() *exceptions.ApiError {
	return &exceptions.ApiError{
		Code:          "200008",
		HttpCode:      http.StatusInternalServerError,
		DefaultErrMsg: "修改密码失败",
	}
}
