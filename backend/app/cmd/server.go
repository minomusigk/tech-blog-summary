package main

import (
	"github.com/gin-gonic/gin"
	"github.com/minomusigk/tech-blog-summary/backend/app/presentation/controller"
)

func main() {
	r := gin.Default()

	r.GET("/ping", controller.GetPing)
	r.Run() // listen and serve on 0.0.0.0:8080
}
