package router

import (
	"net/http"

	"com.github.gin-common/common/controllers"

	"com.github.gin-common/wires"

	"com.github.gin-common/common/routers"
)

type UserRouter struct{}

func (router UserRouter) GroupName() string {
	return "/user"
}

func (router UserRouter) GroupConfig() map[string][]routers.RouteDesc {
	return map[string][]routers.RouteDesc{
		"": {
			{Method: http.MethodPost, Controller: []controllers.ControllerFunc{wires.CreateUserController}},
		},
		"/:userID": {
			{Method: http.MethodGet, Controller: []controllers.ControllerFunc{wires.GetUserInfoController}},
			{Method: http.MethodPut, Controller: []controllers.ControllerFunc{wires.UpdateUserController}},
			{Method: http.MethodDelete, Controller: []controllers.ControllerFunc{wires.DeleteUserController}},
		},
		"/:userID/activate": {
			{Method: http.MethodPatch, Controller: []controllers.ControllerFunc{wires.ActivateUserController}},
		},
		"/:userID/deactivate": {
			{Method: http.MethodPatch, Controller: []controllers.ControllerFunc{wires.DeActivateUserController}},
		},
		"/:userID/changePass": {
			{Method: http.MethodPost, Controller: []controllers.ControllerFunc{wires.ChangePasswordController}},
		},
	}
}

func (router UserRouter) GroupMiddleware() []controllers.MiddlewareFunc {
	return []controllers.MiddlewareFunc{
		wires.AuthMiddleware,
		wires.DeactivatedAbortMiddleware,
	}
}
