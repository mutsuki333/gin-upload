# gin-upload

File related operation APIs

## Example

```go
server := gin.Default()
api := server.Group("/api")
config := upload.New()
config.UploadFolder = filepath.Join("data", "attachments")
config.StaticRoot = "attachments"
upload.Register(api, config)

//Custom Logic
api.POST("/MyUploads", func(c *gin.Context) {
    // some logic or validation
    file, err := upload.Upload(c)
    if err != nil {
        c.AbortWithStatusJSON(500, gin.H{
            "message": err.Error(),
        })
    } else {
        c.JSON(http.StatusOK, file)
    }
})

server.Static("/attachments", "./data/attachments")
server.Run()
```

## URL Endpoints

* Upload File : `POST {{API}}/attachment/`
* File Info : `GET {{API}}/attachment/:id`
* Delete File : `DELETE {{API}}/attachment/:id`

[postman collection](docs/postman_collection.json)

## Upload API

