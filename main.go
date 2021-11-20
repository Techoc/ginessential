package main

import (
	"github.com/gin-gonic/gin"
	"github.com/techoc/ginessential/common"
)

func main() {
	common.InitDB()
	r := gin.Default()
	r = CollectRoute(r)
	panic(r.Run()) // listen and serve on 0.0.0.0:8080
}