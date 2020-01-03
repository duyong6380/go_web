package main

import ( 
    "github.com/gin-gonic/gin"
    "auth"
    "index"
    "fmt"
) 

var module_name []string = []string{"index","auth"}

func regist_router(r *gin.Engine,router map[string] gin.HandlerFunc) {
	for key ,value := range(router) {
		r.GET(key,value)
	}
}

func set_module_router(r *gin.Engine) {
	for module := range module_name {
		module_router := module.InitAuthRouter()
    	RegistRouter(r,module_router)
	}
}

func main(){ 
    router := gin.Default()
    set_module_router(router)
    router.Run(":8000") 
}
