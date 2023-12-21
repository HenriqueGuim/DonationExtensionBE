package services

import (
	"DonationBE/src/models"
	repo "DonationBE/src/repositories"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strings"
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

	err = saveConfigsOnTwitch(configs)
	if err != nil {
		println(err.Error())
		ctx.JSON(500, gin.H{"error": "Internal server error"})
		return
	}

	ctx.JSON(201, gin.H{"message": "Configurations saved successfully"})

}

func saveConfigsOnTwitch(configs models.Configs) error {
	twitchUrl := fmt.Sprintf(`https://api.twitch.tv/extensions/%s/0.0.1/required_configuration?channel_id=%s`,
		os.Getenv("TWITCH_CLIENTID"), string(rune(configs.ChannelId)))

	body := strings.NewReader(`
{
	"required_configuration": "EBSConfigured"
}
`)

	req, err := http.NewRequest("PUT", twitchUrl, body)

	if err != nil {

		return err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Client-Id", os.Getenv("TWITCH_CLIENTID"))
	req.Header.Add("Authorization", "Bearer "+configs.TwitchToken)

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return err
	}

	println(res.Status)

	return nil
}
