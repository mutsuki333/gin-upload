package main

import (
	"github.com/gin-gonic/gin"
	upload "github.com/mutsuki333/gin-upload"
)

func main() {
	server := gin.Default()
	api := server.Group("/api")
	upload.Default(api)
	server.Run()
}
