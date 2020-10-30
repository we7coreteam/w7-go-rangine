package core

import (
	"github.com/gin-gonic/gin"
)

type App struct {
	address   string
	Host      string
	Port      int
	Route     []RouterHandler
	GinRouter *gin.Engine
}

type Router interface {
	Load(router *gin.Engine)
}

func New() *App {
	return &App{
		GinRouter: gin.New(),
	}
}

func (mySelf *App) SetRoute(routeList []RouterHandler) {
	mySelf.Route = routeList
}

func (mySelf *App) Run(address string) {
	// 加载路由
	for _, callback := range mySelf.Route {
		callback(mySelf.GinRouter)
	}

	// 启动服务
	mySelf.GinRouter.Run(address)
}

type RouterHandler func(ginEngine *gin.Engine)
