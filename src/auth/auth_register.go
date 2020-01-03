package auth

import "github.com/gin-gonic/gin"

/*
	本模块的是登陆认证模块
*/

// 认证模块的url根路径
var auth_root string = "/auth/"

// 初始化模块的路由
func InitRouter() map[string] gin.HandlerFunc {
	var router map[string] gin.HandlerFunc
	router = make(map[string] gin.HandlerFunc)
	router[auth_root + "login"] = Login
	router[auth_root + "logout"] = Logout
	return router
}

