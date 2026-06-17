package main

import (
	"log"
	"maps"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/Promise111/gin-quickstart.git/internal/utils"
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
		gender := c.Query("gender")
		c.JSON(http.StatusOK, gin.H{
			"status":  true,
			"message": "User data fetched successfully",
			"name":    name,
			"gender":  strings.ToUpper(gender),
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

	router.PUT("/multiple-upload", func(c *gin.Context) {
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
			log.Println(extName)
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
	})

	router.Run()
}
