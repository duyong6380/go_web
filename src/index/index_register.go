package index

import (
	"github.com/gin-gonic/gin"
)

/*
	本模块的是根模块实现
*/

// 初始化模块的路由
func Init() map[string] gin.HandlerFunc {
	var router map[string] gin.HandlerFunc
	router = make(map[string] gin.HandlerFunc)
	router["/"] = Index
	return router
}