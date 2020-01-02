package main

import ( 
    "github.com/gin-gonic/gin"
) 



func main(){ 
    router := gin.Default() 
    RegistRouter(router)
    router.Run(":8000") 
}
