package router

import (
	"github.com/gin-gonic/gin"
	"auth"
	_ "index"
)

var router *gin.Engine = gin.Default()

func RegistRouter(r ... map[string] gin.HandlerFunc) {
	for _,r_map := range(r) {
		for key ,value := range(r_map) {
	        router.GET(key,value)
	    }
	}
    
}

func Init() *gin.Engine {
	r := auth.Init()
	r_index := index.Init()
	RegistRouter(r,r_index)
	return router
}