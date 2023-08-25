package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {

	// fs := http.FileServer(http.Dir("./static"))
	// http.Handle("/", fs)

	log.Println("Application running in environment:", os.Getenv("RUNTIME_SETUP"), os.Getenv("PORT"))

	var router *gin.Engine = gin.Default()
	router.Static("/", "./static")
	router.Run()

}
