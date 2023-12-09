package main

import (
	"DonationBE/src/repositories"
	"DonationBE/src/services"
	"github.com/gin-gonic/gin"
)

var (
	configRepo        = repositories.NewConfigurationsRepo()
	stripeService     = services.StripeService{CR: &configRepo}
	streamlabsService = services.StreamlabsService{}
	configsService    = services.ConfigsService{CR: &configRepo}
)

func StripeCheckoutController(router *gin.Engine) {
	router.POST("/createCheckout", stripeService.CreateCheckout)
}

func StreamLabsRedirectAutorize(router *gin.Engine) {
	router.GET("/streamlabs/authorize", streamlabsService.AuthorizeStreamLabs)
}

func StreamLabsGetTokens(router *gin.Engine) {
	router.GET("/streamlabs/token", streamlabsService.GetTokens)
}

func SaveConfigs(router *gin.Engine) {
	router.POST("/saveConfigs", configsService.SaveConfigs)
}
