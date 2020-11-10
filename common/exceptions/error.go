package exceptions

import (
	"errors"
	"net/http"
	"reflect"
	"runtime"
	"sync"

	"com.github.gin-common/common/resp"
	"com.github.gin-common/common/validator_trans"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
	"github.com/go-playground/validator/v10"
)

type ApiError struct {
	Code          string
	HttpCode      int
	Err           error
	DefaultErrMsg string
	Data          gin.H
}

type ApiErrorDefFunc func() *ApiError
type ApiErrorModifyFunc func(e *ApiError)

func NewError(errDefFunc ApiErrorDefFunc, errModifyFuncList ...ApiErrorModifyFunc) ApiErrorDefFunc {
	return func() *ApiError {
		e := errDefFunc()
		for _, modifyFunc := range errModifyFuncList {
			modifyFunc(e)
		}
		return e
	}
}

func WithError(err error) ApiErrorModifyFunc {
	return func(e *ApiError) {
		e.Err = err
	}
}

func WithValidateError(err error) ApiErrorModifyFunc {
	return func(e *ApiError) {
		if validateErrors, ok := err.(validator.ValidationErrors); ok {
			data := gin.H{}
			for _, fe := range validateErrors {
				errMsg := fe.Translate(validator_trans.Trans)
				data[fe.Field()] = errMsg
			}
			e.Err = errors.New("参数错误")
			e.Data = data
		} else {
			e.Err = err
		}
	}
}

// 存放所有预定义的异常
var DefinedErrorsMap = map[string]*ApiError{}
var mu sync.RWMutex

func GetDefinedErrors(defFunc ApiErrorDefFunc) *ApiError {
	// lazy init *ApiErr object Goroutine-safe
	// 可以并发读，不能并发写
	funcName := runtime.FuncForPC(reflect.ValueOf(defFunc).Pointer()).Name()
	mu.RLock()
	if DefinedErrorsMap[funcName] != nil {
		defer mu.RUnlock()
		return DefinedErrorsMap[funcName]
	}
	mu.RUnlock()

	if DefinedErrorsMap[funcName] == nil {
		mu.Lock()
		defer mu.Unlock()
		if DefinedErrorsMap[funcName] == nil {
			DefinedErrorsMap[funcName] = defFunc()
		}
	}
	return DefinedErrorsMap[funcName]
}

func (e *ApiError) RenderJSONFunc() resp.RenderFunc {
	var msg string
	if e.Err == nil {
		msg = e.DefaultErrMsg
	} else {
		msg = e.Err.Error()
	}
	return func(c *gin.Context) {
		c.Render(e.HttpCode, render.JSON{Data: resp.Response{
			Code: e.Code,
			Msg:  msg,
			Data: e.Data,
		}})
	}
}

func (e *ApiError) Error() string {
	if e.Err == nil {
		return e.DefaultErrMsg
	} else {
		return e.Err.Error()
	}
}

func ServerError() *ApiError {
	return &ApiError{
		Code:          "777777",
		HttpCode:      http.StatusInternalServerError,
		DefaultErrMsg: "未知错误",
	}
}

func Forbidden() *ApiError {
	return &ApiError{
		Code:          "100403",
		HttpCode:      http.StatusForbidden,
		DefaultErrMsg: "没有权限",
	}
}

func Unauthorized() *ApiError {
	return &ApiError{
		Code:          "100401",
		HttpCode:      http.StatusUnauthorized,
		DefaultErrMsg: "验证失败",
	}
}

func BadRequest() *ApiError {
	return &ApiError{
		Code:          "100400",
		HttpCode:      http.StatusBadRequest,
		DefaultErrMsg: "参数错误",
	}
}

func Timeout() *ApiError {
	return &ApiError{
		Code:          "100501",
		HttpCode:      http.StatusInternalServerError,
		DefaultErrMsg: "超时",
	}
}
