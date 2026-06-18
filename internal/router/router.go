package router

import (
	"github.com/Promise111/gin-quickstart.git/internal/handler"
	"github.com/gin-gonic/gin"
)

func New() *gin.Engine {
	r := gin.Default()
	r.MaxMultipartMemory = 8 << 20 // 8 MiB

	r.GET("/:name", handler.GetByName)
	r.POST("/user", handler.CreateUser)
	r.PUT("/upload", handler.Upload)
	r.PUT("/multiple-upload", handler.MultipleUpload)

	return r
}
