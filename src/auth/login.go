package auth

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func Login(c *gin.Context) {
    username := c.Query("username")
    password := c.Query("password")
    if username == "hello" && password == "world" {
       c.String(http.StatusOK,"login success")
    } else {
        c.String(http.StatusOK,"login failed")
    }
}


func Logout(c *gin.Context) {
	c.String(http.StatusOK,"logout success")
}