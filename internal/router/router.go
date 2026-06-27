package router

import (
	"github.com/Promise111/gin-quickstart.git/internal/handler"
	"github.com/Promise111/gin-quickstart.git/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func New() *gin.Engine {
	r := gin.Default()
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("bookabledate", handler.BookableDate)
	}
	r.MaxMultipartMemory = 8 << 20 // 8 MiB

	api := r.Group(utils.APIPrefix)

	{
		v1 := api.Group(utils.V1Prefix)
		v1.GET("/final", handler.FinalHandler)
		v1.GET("/test", handler.TestHandler(r))
		v1.GET("/:name", handler.GetByName)
		v1.POST("/user", handler.CreateUser)
		v1.PUT("/upload", handler.Upload)
		v1.PUT("/multiple-upload", handler.MultipleUpload)
		v1.POST("/login", handler.LoginHandler(r))
		v1.GET("/booking", handler.GetBookableDate(r))
		v1.GET("/testing", handler.GetStartPage(r))
		v1.POST("/testing", handler.GetStartPage(r))
		v1.POST("/person", handler.GetPerson(r))
	}

	return r
}
