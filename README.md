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

server.Static("/attachments", "./data/attachments")
server.Run()
```

## API

* Upload File : `POST {{API}}/attachment/`
* File Info : `GET {{API}}/attachment/:id`
* Delete File : `DELETE {{API}}/attachment/:id`

[postman collection](docs/postman_collection.json)