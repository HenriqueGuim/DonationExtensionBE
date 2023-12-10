package services

import (
	"DonationBE/src/models"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strings"
)

type StreamlabsService struct {
}

func (*StreamlabsService) AuthorizeStreamLabs(ctx *gin.Context) {
	responseType := "code"
	clientId := os.Getenv("STREAMLABS_CLIENT_ID")
	redirectUri := os.Getenv("REDIRECT_URL") + "/streamlabs/token"
	scope := "donations.create"
	url := os.Getenv("STREAMLABS_URL") + "/api/v2.0/authorize"

	url += "?client_id=" + clientId
	url += "&redirect_uri=" + redirectUri
	url += "&scope=" + scope
	url += "&response_type=" + responseType
	url += "&state=123456"

	ctx.Redirect(301, url)

}

func (*StreamlabsService) GetTokens(ctx *gin.Context) {
	code := ctx.Query("code")

	url := os.Getenv("STREAMLABS_URL") + "/token"
	clientId := os.Getenv("STREAMLABS_CLIENT_ID")
	clientSecret := os.Getenv("STREAMLABS_CLIENT_SECRET")
	redirectUri := os.Getenv("STREAMLABS_URL") + "/streamlabs/token"

	body := strings.NewReader(`
{
	"grant_type": "authorization_code",
	"client_id": "` + clientId + `",
	"client_secret": "` + clientSecret + `",
	"redirect_uri": "` + redirectUri + `",
	"code": "` + code + `"
}
`)

	req, err := http.NewRequest("POST", url, body)

	if err != nil {
		println(err.Error())
		println("Error creating request")
		ctx.JSON(500, gin.H{"error": "Internal server error"})
		return
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	res, err := http.DefaultClient.Do(req)

	if err != nil || res.StatusCode != 200 {
		println(err.Error())
		println("Error doing request")
		ctx.JSON(500, gin.H{"error": "Internal server error"})
		return
	}

	println("status of request: " + res.Status)

	var resBody models.StreamLabsTokensRequestBody

	err = json.NewDecoder(res.Body).Decode(&resBody)

	if err != nil {
		println(err.Error())
		println("Error decoding body")
		ctx.JSON(500, gin.H{"error": "Internal server error"})
		return
	}

	ctx.Data(200, "text/html", []byte(getTokensHTML(resBody)))
}

func getTokensHTML(tokens models.StreamLabsTokensRequestBody) string {
	return `
<html>
<head>
<title>Streamlabs tokens</title>
</head>
<body>
<h1>Streamlabs tokens</h1>
<p>Access token: ` + tokens.AccessToken + `</p>
<p>Refresh token: ` + tokens.RefreshToken + `</p>
</body>
</html>`
}
