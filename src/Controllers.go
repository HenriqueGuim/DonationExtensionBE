package main

import "github.com/gin-gonic/gin"

func obtainCheckoutController(router *gin.Engine) {
	router.POST("/createCheckout", createCheckout)
}
