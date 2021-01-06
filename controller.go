package upload

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func controllers(base *gin.RouterGroup) {
	api := base.Group("/attachment")
	api.POST("/", uploader.upload)
	api.GET("/:id", uploader.Info)
	api.DELETE("/:id", uploader.delete)
}

func (u *Uploader) upload(c *gin.Context) {
	file, err := u.Upload(c)
	if err != nil {
		abort(c, err)
	} else {
		c.JSON(http.StatusOK, file)
	}
}

//Info file
func (u *Uploader) Info(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		abort(c, err)
		return
	}
	file := &File{}
	err = u.DB.First(file, id).Error
	if err != nil {
		abort(c, err)
		return
	}
	file.URL = url("/", u.StaticRoot, file.Path)
	c.JSON(http.StatusOK, file)
}

func (u *Uploader) delete(c *gin.Context) {
	file, err := u.Delete(c.Param("id"))
	if err != nil {
		abort(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"name":    file.Filename,
	})
}
