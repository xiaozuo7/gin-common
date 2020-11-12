package exception

import (
	"net/http"

	"com.github.gin-common/common/exceptions"
)

func BookCreateFailed() *exceptions.ApiError {
	return &exceptions.ApiError{
		Code:          "200001",
		HttpCode:      http.StatusInternalServerError,
		DefaultErrMsg: "创建用户失败",
	}
}
