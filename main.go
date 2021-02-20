package main

import (
	"dqh-test/router"
	"github.com/gin-gonic/gin"
)

func main() {

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	router.RegisterRouter(r)

	r.Run(":8080")

}

