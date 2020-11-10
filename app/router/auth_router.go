package router

import (
	"net/http"

	"com.github.gin-common/common/controllers"

	"com.github.gin-common/wires"

	"com.github.gin-common/common/routers"
)

type AuthRouter struct{}

func (router AuthRouter) GroupName() string {
	return ""
}

func (router AuthRouter) GroupConfig() map[string][]routers.RouteDesc {
	return map[string][]routers.RouteDesc{
		"/login": {
			routers.RouteDesc{Method: http.MethodPost, Controller: []controllers.ControllerFunc{wires.LoginController}},
		},
		"/logout": {
			routers.RouteDesc{Method: http.MethodPost, MiddleWare: []controllers.MiddlewareFunc{wires.AuthMiddleware},
				Controller: []controllers.ControllerFunc{wires.LogoutController}},
		},
		"/current_user": {
			routers.RouteDesc{Method: http.MethodGet, MiddleWare: []controllers.MiddlewareFunc{wires.AuthMiddleware, wires.DeactivatedAbortMiddleware},
				Controller: []controllers.ControllerFunc{wires.CurrentUserController}},
		},
	}
}

func (router AuthRouter) GroupMiddleware() []controllers.MiddlewareFunc {
	return []controllers.MiddlewareFunc{}
}
