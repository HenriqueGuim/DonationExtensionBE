package services

import (
	"DonationBE/src/models"
	repo "DonationBE/src/repositories"
	"github.com/gin-gonic/gin"
)

type ConfigsService struct {
	CR *repo.ConfigurationsRepo
}

func (cs *ConfigsService) SaveConfigs(ctx *gin.Context) {

	var configs models.Configs

	if err := ctx.BindJSON(&configs); err != nil {
		println(err.Error())
		ctx.JSON(400, gin.H{"error": "Bad request binding json"})
		return
	}

	if configs.ChannelId == 0 || configs.StripeToken == "" || configs.StreamlabsToken == "" || configs.StreamlabsRefreshToken == "" {
		ctx.JSON(400, gin.H{"error": "Bad request missing parameters"})
		return
	}

	err := cs.CR.PostConfigurations(configs)

	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")

	if err != nil {
		println(err.Error())
		ctx.JSON(500, gin.H{"error": "Internal server error"})
		return
	}

	ctx.JSON(201, gin.H{"message": "Configurations saved successfully"})

}
