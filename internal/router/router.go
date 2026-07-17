package router

import (
	"github.com/Promise111/gin-quickstart.git/internal/handler"
	"github.com/Promise111/gin-quickstart.git/internal/middleware"
	"github.com/Promise111/gin-quickstart.git/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func New() *gin.Engine {
	// r := gin.Default()
	r := gin.New()
	// register custom Logger middleware
	r.Use(middleware.Logger())

	// Recovery middleware recovers from any panics and returns 500 if there was one
	r.Use(middleware.Recovery())

	// Attach the error-handling middleware
	r.Use(middleware.ErrorHandler())

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("bookabledate", handler.BookableDate)
	}
	r.MaxMultipartMemory = 8 << 20 // 8 MiB

	// r.Static("/assets", "./assets")
	r.StaticFS("/assets", gin.Dir("./assets", true))

	r.LoadHTMLGlob("internal/templates/*")
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
		v1.GET("/search", handler.Search(r))
		v1.GET("/:name/:id", handler.BindURI(r))
		v1.GET("/testText", handler.TextUmarshal(r))
		v1.GET("/testBind", handler.UnmarshalParam(r))
		v1.GET("/header", handler.GetHeaders(r))
		v1.GET("/checkboxes", handler.GetCheckBoxes(r))
		v1.POST("/checkboxes", handler.PostCheckBoxes(r))
		v1.POST("/login-urlencoded", handler.LoginUrlEncoded(r))
	}

	{
		v2 := api.Group(utils.V2Prefix)
		v2.GET("/getb", handler.GetB(r))
		v2.GET("/getc", handler.GetC(r))
		v2.GET("/getd", handler.GetD(r))
		v2.POST("/bind", handler.BindMultipleStruct(r))
		v2.GET("/list", handler.ListCustomerBindind(r))
		v2.GET("/walking-god", handler.GetDocs(r))
		v2.GET("/download-doc", handler.GetDocsDownload(r))
		v2.GET("/someDataFromReader", handler.SomeDataFromReader(r))
		v2.GET("/html", handler.GetHTML1(r))
		v2.GET("/test", handler.TestV2(r))
		v2.GET("/panic", handler.Panic(r))
		v2.GET("/error", handler.HandleError(r))
	}

	{
		authorized := r.Group("/admin", middleware.BasicAuth)
		// /admin/secrets endpoint
		// hit "localhost:8080/admin/secrets
		authorized.GET("/secrets", handler.HandleSecrets(r))
	}

	return r
}
