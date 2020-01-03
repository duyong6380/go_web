package main

import "router"

func main(){
    route := router.Init()
    route.Run(":8000") 
}
