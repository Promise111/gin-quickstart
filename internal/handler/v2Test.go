package handler

import (
	// "embed"
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
