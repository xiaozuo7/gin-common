package router

type BookRouter struct{}

func (router BookRouter) GroupName() string {
	return "/book"
}

//func (router BookRouter) GroupConfig() map[string][]routers.RouteDesc {
//	return map[string][]routers.RouteDesc{
//		"": {
//			{Method: http.MethodPost, Controller: []controllers.ControllerFunc{wires.CreateUserController}},
//		},
//	}
//}

//func (router BookRouter) GroupMiddleware() []controllers.MiddlewareFunc {
//	return []controllers.MiddlewareFunc{
//		wires.AuthMiddleware,
//		wires.DeactivatedAbortMiddleware,
//	}
//}
