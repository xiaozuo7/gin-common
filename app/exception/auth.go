package exception

import (
	"net/http"

	"com.github.gin-common/common/exceptions"
)

func UserNameOrPassInvalid() *exceptions.ApiError {
	return &exceptions.ApiError{
		Code:          "300001",
		HttpCode:      http.StatusInternalServerError,
		DefaultErrMsg: "用户名或密码错误",
	}
}

func LoginFailed() *exceptions.ApiError {
	return &exceptions.ApiError{
		Code:          "300002",
		HttpCode:      http.StatusInternalServerError,
		DefaultErrMsg: "登录失败",
	}
}

func TokenInvalid() *exceptions.ApiError {
	return &exceptions.ApiError{
		Code:          "300003",
		HttpCode:      http.StatusUnauthorized,
		DefaultErrMsg: "token无效",
	}
}

func TokenExpired() *exceptions.ApiError {
	return &exceptions.ApiError{
		Code:          "300004",
		HttpCode:      http.StatusUnauthorized,
		DefaultErrMsg: "token过期",
	}
}

func UserDeactivated() *exceptions.ApiError {
	return &exceptions.ApiError{
		Code:          "300005",
		HttpCode:      http.StatusUnauthorized,
		DefaultErrMsg: "用户被禁用",
	}
}
