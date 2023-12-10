package services

import (
	repo "DonationBE/src/repositories"
	"github.com/gin-gonic/gin"
)

type ConfigsService struct {
	CR *repo.ConfigurationsRepo
}

func (cs *ConfigsService) SaveConfigs(ctx *gin.Context) {

	var (
		channelId              int
		stripeToken            string
		streamlabsToken        string
		streamlabsRefreshToken string
	)

	configs, _ := cs.CR.GetConfigurations(1)

	if channelId == 0 || stripeToken == "" || streamlabsToken == "" || streamlabsRefreshToken == "" {
		ctx.JSON(400, gin.H{"error": "Bad request missing parameters"})
		return
	}

	err := cs.CR.PostConfigurations(configs)

	if err != nil {
		println(err.Error())
		ctx.JSON(500, gin.H{"error": "Internal server error"})
		return
	}

	ctx.JSON(201, gin.H{"message": "Configurations saved successfully"})

}
