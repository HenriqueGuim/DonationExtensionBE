package main

import (
	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go/v76"
)

func createCheckout(ctx *gin.Context) {

	var requestBody RequestBody

	if err := ctx.BindJSON(&requestBody); err != nil {
		println(err.Error())
		ctx.JSON(400, gin.H{"error": "Bad request"})
		return
	}

	stripe.Key = requestBody.StripeKey

	s, err := createCheckoutSession(requestBody.Amount*100, requestBody.SuccessDomain, requestBody.FailDomain, requestBody.ImgUrl)

	if err != nil {
		println(err.Error())
		ctx.JSON(500, gin.H{"error": "Internal server error"})
		return
	}

	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")

	responseBody := gin.H{"id": s.ID, "url": s.URL}

	ctx.JSON(200, responseBody)
}
