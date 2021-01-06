package main

import (
	"path/filepath"

	uploader "evan-soft.com/bricks/gin-uploader"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	api := server.Group("/api")
	config := uploader.New()
	config.UploadFolder = filepath.Join("data", "attachments")
	config.StaticRoot = "attachments"
	uploader.Register(api, config)

	server.Static("/attachments", "./data/attachments")
	server.Run()
}
