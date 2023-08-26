package main

import (
	"log"

	"github.com/booklib/util"
	"github.com/gin-gonic/gin"
)

func main() {

	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	// fs := http.FileServer(http.Dir("./static"))
	// http.Handle("/", fs)

	log.Println("Application running in environment:",
		config.RuntimeSetup, config.AppPort)

	var router *gin.Engine = gin.Default()
	router.Static("/", "./static")
	router.Run(config.ServerAddress + ":" + config.AppPort)

}
