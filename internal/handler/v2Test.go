package handler

import (
	// "embed"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetDocs(engine *gin.Engine) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.File("assets/docs/walking with God.docx")
	}
}

func GetDocsDownload(engine *gin.Engine) gin.HandlerFunc {
	return func(c *gin.Context) {
		// var embeddedAssets embed.FS
		c.FileAttachment("assets/docs/walking with God.docx", "walking-doc")
	}
}

func SomeDataFromReader(engine *gin.Engine) gin.HandlerFunc {
	return func(c *gin.Context) {
		response, err := http.Get("https://raw.githubusercontent.com/gin-gonic/logo/master/color.png")

		if err != nil || response.StatusCode != http.StatusOK {
			c.Status(http.StatusServiceUnavailable)
			return
		}
		defer response.Body.Close()

		reader := response.Body
		contentLength := response.ContentLength
		contentType := response.Header.Get("Content-Type")

		extraHeaders := map[string]string{
			"Content-Disposition": `attachment; filename="gopher.png"`,
		}

		c.DataFromReader(http.StatusOK, contentLength, contentType, reader, extraHeaders)
	}
}

func GetHTML1(engine *gin.Engine) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "fileOne.html", gin.H{"title": "Gin Quick-Start"})
	}
}

func TestV2(engine *gin.Engine) gin.HandlerFunc {
	return func(c *gin.Context) {
		example := c.MustGet("example").(string)

		// it would print: "12345"
		log.Println(example)
	}
}

func Panic(engine *gin.Engine) gin.HandlerFunc {
	return func(c *gin.Context) {
		// panic with a string -- the custom middleware could save this to a database or report it to the user
		panic("foo")
	}
}
