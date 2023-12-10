package services

import (
	"DonationBE/src/models"
	repo "DonationBE/src/repositories"
	"DonationBE/src/utils"
	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go/v76"
)

type StripeService struct {
	CR *repo.ConfigurationsRepo
}

func (ss *StripeService) CreateCheckout(ctx *gin.Context) {

	var requestBody models.DonoRequestBody

	if err := ctx.BindJSON(&requestBody); err != nil {
		println(err.Error())
		ctx.JSON(400, gin.H{"error": "Bad request"})
		return
	}

	configs, err := ss.CR.GetConfigurations(requestBody.ChannelId)

	if err != nil {
		println(err.Error())
		ctx.JSON(500, gin.H{"error": "Internal server error"})
		return
	}

	stripe.Key = configs.StripeToken

	s, err := utils.CreateCheckoutSession(requestBody.Amount*100, requestBody.SuccessDomain, requestBody.FailDomain, requestBody.ImgUrl)

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
