package routers

import (
	"reflect"
	"runtime"
	"strings"

	"com.github.gin-common/common/controllers"

	"github.com/gin-gonic/gin"
)

type RouteDesc struct {
	Method     string
	MiddleWare []controllers.MiddlewareFunc
	Controller []controllers.ControllerFunc
}

type GinRouterInterface interface {
	GroupName() string
	GroupConfig() map[string][]RouteDesc
	GroupMiddleware() []controllers.MiddlewareFunc
}

func handleRouter(engine *gin.Engine, router GinRouterInterface) {
	groupName := router.GroupName()
	groupConfigs := router.GroupConfig()
	groupMiddlewares := router.GroupMiddleware()
	group := handleGroupName(engine, groupName)
	middlewareFlag := handleGroupMiddleware(group, groupMiddlewares...)
	handleGroupConfigs(group, groupConfigs, middlewareFlag)
}

func handleGroupMiddleware(group *gin.RouterGroup, middlewares ...controllers.MiddlewareFunc) map[string]bool {
	var middlewareHandlers []gin.HandlerFunc
	var middlewareFlag = make(map[string]bool)
	for _, middlewareFunc := range middlewares {
		funcName := runtime.FuncForPC(reflect.ValueOf(middlewareFunc).Pointer()).Name()
		if _, ok := middlewareFlag[funcName]; !ok {
			middlewareFlag[funcName] = true
			middlewareHandlers = append(middlewareHandlers, controllers.MiddlewareHandler(middlewareFunc))
		}
	}
	group.Use(middlewareHandlers...)
	return middlewareFlag
}

func handleGroupName(engine *gin.Engine, groupName string) *gin.RouterGroup {
	trimName := strings.Trim(groupName, " ")
	if trimName == "" {
		return &engine.RouterGroup
	}
	return engine.Group(trimName)
}

type RequestFunc func(group *gin.RouterGroup, relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes

func getFuncByMethod(method string) RequestFunc {
	upperMethod := strings.ToUpper(method)
	return func(group *gin.RouterGroup, relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
		return group.Handle(upperMethod, relativePath, handlers...)
	}
}

func handleGroupConfigs(group *gin.RouterGroup, groupConfigs map[string][]RouteDesc, alreadyUseMiddlewareFlag map[string]bool) {
	for relativePath, routeDescribes := range groupConfigs {
		for _, routeDesc := range routeDescribes {
			method := routeDesc.Method

			var controllerHandlers []gin.HandlerFunc
			var middlewareFlag = make(map[string]bool)

			for _, middlewareFunc := range routeDesc.MiddleWare {
				funcName := runtime.FuncForPC(reflect.ValueOf(middlewareFunc).Pointer()).Name()
				if _, ok := middlewareFlag[funcName]; !ok {
					if _, alreadyUse := alreadyUseMiddlewareFlag[funcName]; !alreadyUse {
						middlewareFlag[funcName] = true
						controllerHandlers = append(controllerHandlers, controllers.MiddlewareHandler(middlewareFunc))
					}
				}
			}

			var controllerFlag = make(map[string]bool)
			for _, controllerFunc := range routeDesc.Controller {
				funcName := runtime.FuncForPC(reflect.ValueOf(controllerFunc).Pointer()).Name()
				if _, ok := controllerFlag[funcName]; !ok {
					controllerFlag[funcName] = true
					controllerHandlers = append(controllerHandlers, controllers.ControllerHandler(controllerFunc))
				}
			}

			process := getFuncByMethod(method)
			process(group, relativePath, controllerHandlers...)
		}
	}
}

func CombineRouters(engine *gin.Engine, routers ...GinRouterInterface) {
	for _, router := range routers {
		handleRouter(engine, router)
	}
}
