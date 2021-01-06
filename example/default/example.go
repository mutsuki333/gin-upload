package main

import (
	upload "evan-soft.com/bricks/gin-upload"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	api := server.Group("/api")
	upload.Default(api)
	server.Run()
}
