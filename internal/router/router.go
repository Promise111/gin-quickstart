package router

import (
	"github.com/Promise111/gin-quickstart.git/internal/handler"
	"github.com/Promise111/gin-quickstart.git/internal/utils"
	"github.com/gin-gonic/gin"
)

func New() *gin.Engine {
	r := gin.Default()
	r.MaxMultipartMemory = 8 << 20 // 8 MiB

	api := r.Group(utils.APIPrefix)

	{
		v1 := api.Group(utils.V1Prefix)
		v1.GET("/:name", handler.GetByName)
		v1.POST("/user", handler.CreateUser)
		v1.PUT("/upload", handler.Upload)
		v1.PUT("/multiple-upload", handler.MultipleUpload)
	}

	return r
}
