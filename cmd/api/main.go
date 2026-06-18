package main

import (
	"log"

	"github.com/Promise111/gin-quickstart.git/internal/router"
)

func main() {
	log.SetPrefix("Quickstart: ")
	log.SetFlags(0)

	r := router.New()
	r.Run()
}
