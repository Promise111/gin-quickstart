package handler

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func LongAsync(engine *gin.Engine) gin.HandlerFunc {
	return func(c *gin.Context) {
		// create context copy to be used inside goroutine
		cCp := c.Copy()
		go func() {
			// simulate a long standing task with time.Sleep(). 5 seconds
			time.Sleep(5 * time.Second)

			// note that I am using the copied c.Context "cCp", O di very IMPORTANT
			log.Println("Done! in path " + cCp.Request.URL.Path)
		}()
	}
}

func LongSync(engine *gin.Engine) gin.HandlerFunc {
	return func(c *gin.Context) {
		// simulate a long standing task with time.Sleep(). 5 seconds
		time.Sleep(5 * time.Second)

		// Since we are not using a goroutine, we do not have to copy c.Context like we did in the handler above
		log.Println("Done! in path " + c.Request.URL.Path)
	}
}
