package handler

import (
	"log"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/Promise111/gin-quickstart.git/internal/utils"
	"github.com/gin-gonic/gin"
)

func Upload(c *gin.Context) {
	file, uploadErr := c.FormFile("profilePic")
	if uploadErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "Upload valid image",
			"name":    "Promise",
			"gender":  "M",
		})
		return
	}
	dst := filepath.Join("./files/", filepath.Base(file.Filename))
	saveErr := c.SaveUploadedFile(file, dst)
	if saveErr != nil {
		c.JSON(http.StatusCreated, gin.H{
			"status":   false,
			"message":  "Error: saving file failed",
			"name":     "Promise",
			"age":      30,
			"fileName": file.Filename,
			"path":     dst,
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status":   true,
		"message":  "File uploaded successfully",
		"name":     "Promise",
		"age":      30,
		"fileName": file.Filename,
		"path":     dst,
	})
}

func MultipleUpload(c *gin.Context) {
	form, uploadErr := c.MultipartForm()
	if uploadErr != nil {
		c.JSON(http.StatusBadRequest,
			gin.H{
				"status":  false,
				"message": uploadErr.Error(),
				"name":    "Promise",
				"age":     54,
			})
		return
	}
	files := form.File["files"]
	log.Println(files)

	for i, file := range files {
		extName := filepath.Ext(file.Filename)
		randomBytes, genBytesErr := utils.GenerateRandomBytes(15)
		if genBytesErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": genBytesErr.Error(),
				"name": nil,
				"age":  nil,
			})
			return
		}
		newFileName := randomBytes + strconv.Itoa(i) + extName
		dst := filepath.Join("./multiple/", filepath.Base(newFileName))
		if err := c.SaveUploadedFile(file, dst); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
				"name":    "Promise",
				"age":     26,
			})
			return
		}
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":     true,
		"message":    "Files uploaded successfully",
		"name":       "Promise",
		"age":        22,
		"fileLenght": len(files),
	})
}
