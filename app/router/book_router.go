package router

import (
	"com.github.gin-common/common/controllers"
	"com.github.gin-common/common/routers"
	"com.github.gin-common/wires"
	"net/http"
)

type BookRouter struct{}

func (router BookRouter) GroupName() string {
	return "/book"
}

func (router BookRouter) GroupConfig() map[string][]routers.RouteDesc {
	return map[string][]routers.RouteDesc{
		"": {
			{Method: http.MethodPost, Controller: []controllers.ControllerFunc{wires.CreateBookController}},
		},
	}
}

func (router BookRouter) GroupMiddleware() []controllers.MiddlewareFunc {
	return []controllers.MiddlewareFunc{
	}
}
