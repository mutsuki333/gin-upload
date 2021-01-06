package main

import (
	uploader "evan-soft.com/bricks/gin-uploader"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	api := server.Group("/api")
	uploader.Default(api)
	server.Run()
}
