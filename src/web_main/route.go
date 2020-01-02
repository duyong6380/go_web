package main


import (
	"auth"
	"github.com/gin-gonic/gin"
	"net/http"
)

func index(c *gin.Context) {
    c.String(http.StatusOK, "Hello World")
}

func RegistRouter(r *gin.Engine) {
    r.GET("/", index)
    r.GET("/auth/login",auth.Login)
}
