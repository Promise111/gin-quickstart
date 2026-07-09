package handler

import (
	// "embed"
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
