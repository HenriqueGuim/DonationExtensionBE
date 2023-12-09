package services

import (
	"DonationBE/src/models"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

type StreamlabsService struct {
}

func (*StreamlabsService) AuthorizeStreamLabs(ctx *gin.Context) {
	responseType := "code"
	clientId := "6e62955e-2c61-4fe5-acfc-3b9637e33469"
	redirectUri := "http://localhost:8080/streamlabs/token"
	scope := "donations.create"
	url := "https://streamlabs.com/api/v2.0/authorize"

	url += "?client_id=" + clientId
	url += "&redirect_uri=" + redirectUri
	url += "&scope=" + scope
	url += "&response_type=" + responseType
	url += "&state=123456"

	ctx.Redirect(301, url)

}

func (*StreamlabsService) GetTokens(ctx *gin.Context) {
	code := ctx.Query("code")

	url := "https://streamlabs.com/api/v2.0/token"
	clientId := "6e62955e-2c61-4fe5-acfc-3b9637e33469"
	clientSecret := "cbooySyWt2Zn5LMemQCFyr3eadfVsIxbJSPea8vN"
	redirectUri := "http://localhost:8080/"

	url += "?grant_type=authorization_code"
	url += "&client_id=" + clientId
	url += "&client_secret=" + clientSecret
	url += "&redirect_uri=" + redirectUri
	url += "&code=" + code

	req, err := http.NewRequest("POST", url, nil)

	if err != nil {
		println(err.Error())
		println("Error creating request")
		ctx.JSON(500, gin.H{"error": "Internal server error"})
		return
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		println(err.Error())
		println("Error doing request")
		ctx.JSON(500, gin.H{"error": "Internal server error"})
		return
	}

	println(res.Status)

	var resBody models.StreamLabsTokensRequestBody

	err = json.NewDecoder(res.Body).Decode(&resBody)

	if err != nil {
		println(err.Error())
		println("Error decoding body")
		ctx.JSON(500, gin.H{"error": "Internal server error"})
		return
	}

	ctx.JSON(200, resBody)
}
