package main

import (
	"path/filepath"

	upload "evan-soft.com/bricks/gin-upload"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	api := server.Group("/api")
	config := upload.New()
	config.UploadFolder = filepath.Join("data", "attachments")
	config.StaticRoot = "attachments"
	upload.Register(api, config)

	server.Static("/attachments", "./data/attachments")
	server.Run()
}
