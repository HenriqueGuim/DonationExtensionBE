package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true

	obtainCheckoutController(router)

	router.Use(cors.New(config))
	err := router.Run(":8080")
	if err != nil {
		log.Fatal("Router couldn't start")
	}

}
