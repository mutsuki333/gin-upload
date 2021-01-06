# gin-uploader

File related operation APIs

## Example

```go
server := gin.Default()
api := server.Group("/api")
config := uploader.New()
config.UploadFolder = filepath.Join("data", "attachments")
config.StaticRoot = "attachments"
uploader.Register(api, config)

server.Static("/attachments", "./data/attachments")
server.Run()
```

## API

[postman collection](docs/postman_collection.json)