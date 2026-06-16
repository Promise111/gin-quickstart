package main

import (
	"log"
	"maps"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

type Metadata struct {
	Ids         map[string]string
	CourseNames map[string]string
}

func main() {
	log.SetPrefix("Quickstart: ")
	log.SetFlags(0)
	router := gin.Default()
	router.MaxMultipartMemory = 8 << 20 // 8 MiB

	router.GET("/:name", func(c *gin.Context) {
		var name string = c.Param("name")
		c.JSON(http.StatusOK, gin.H{
			"status":  true,
			"message": "User data fetched successfully",
			"name":    name,
		})
	})

	router.POST("/user", func(c *gin.Context) {
		age := c.DefaultPostForm("age", "18")
		gender := c.DefaultQuery("gender", "M")
		name := c.PostForm("name")

		idsQuery := c.QueryMap("ids")
		idMap := make(map[string]string)
		maps.Copy(idMap, idsQuery)

		courseNameMap := make(map[string]string)
		courseNames := c.PostFormMap("names")
		maps.Copy(courseNameMap, courseNames)
		meta := Metadata{
			Ids:         idMap,
			CourseNames: courseNameMap,
		}
		log.Println(meta)

		c.JSON(http.StatusCreated, gin.H{
			"status":  true,
			"message": "User created successfully",
			"age":     age,
			"gender":  strings.ToUpper(gender),
			"meta":    meta,
			"name":    name,
		})
	})

	router.PUT("/upload", func(c *gin.Context) {
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
	})

	router.Run()
}
