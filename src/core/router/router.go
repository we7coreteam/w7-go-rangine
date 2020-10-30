package router

import "github.com/gin-gonic/gin"

/**
定义标准路由接口
定义路由时需要继承该接口
*/
type Router interface {
	Load(router *gin.Engine)
}
