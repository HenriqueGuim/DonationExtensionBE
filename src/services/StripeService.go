package services

import (
	"DonationBE/src/models"
	repo "DonationBE/src/repositories"
	"DonationBE/src/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/checkout/session"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
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

	s, err := utils.CreateCheckoutSession(int64(requestBody.Amount*100), requestBody.SuccessDomain, requestBody.FailDomain, requestBody.ImgUrl)

	if err != nil {
		println(err.Error())
		ctx.JSON(500, gin.H{"error": "Internal server error"})
		return
	}

	go ss.checkPayment(*s, configs, requestBody)

	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")

	responseBody := gin.H{"id": s.ID, "url": s.URL}

	ctx.JSON(200, responseBody)
}

func (ss *StripeService) checkPayment(stripeSession stripe.CheckoutSession, configs models.Configs, donoReq models.DonoRequestBody) {
	for i := 0; i < 30; i++ {
		newSession, err := session.Get(stripeSession.ID, &stripe.CheckoutSessionParams{})

		if err != nil {
			println(err.Error())
			return
		}

		if newSession.PaymentStatus != "paid" {
			println("Payment done")
			ss.streamlabsRegistDono(configs, donoReq)
			return
		}

		println("Payment not done yet")
		time.Sleep(1 * time.Minute)
	}
}

func (ss *StripeService) streamlabsRegistDono(configs models.Configs, donoReq models.DonoRequestBody) {
	url := os.Getenv("STREAMLABS_URL") + "/donations"

	payload := strings.NewReader(
		`{"name": "` + donoReq.UserName + `",
"message": "` + donoReq.Message + `",
"identifier": "` + donoReq.UserName + `",
"amount": ` + `"` + fmt.Sprintf("%.2f", donoReq.Amount) + `"` + `,
"currency": "eur"
}`)

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("Authorization", "Bearer "+configs.StreamlabsToken)

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		println(err.Error())
		return
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			println(err.Error())
		}
	}(res.Body)
	body, _ := io.ReadAll(res.Body)

	if res.StatusCode != 200 {
		println("Error doing request")
		println(string(body))
		return
	}

	fmt.Println(string(body))
}
