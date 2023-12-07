package main

import (
	"DonationBE/src/services"
	"github.com/gin-gonic/gin"
)

func StripeCheckoutController(router *gin.Engine) {
	router.POST("/createCheckout", services.CreateCheckout)
}

func StreamLabsRedirectAutorize(router *gin.Engine) {
	router.GET("/streamlabs/authorize", services.AuthorizeStreamLabs)
}

func StreamLabsGetTokens(router *gin.Engine) {
	router.GET("/streamlabs/token", services.GetTokens)
}
