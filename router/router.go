package router

import "github.com/gin-gonic/gin"

type Router interface {
	Load(router *gin.Engine)
}
