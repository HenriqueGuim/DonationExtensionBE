package models

type DonoRequestBody struct {
	Amount        int    `json:"amount"`
	UserName      string `json:"userName"`
	Message       string `json:"message"`
	ImgUrl        string `json:"imgUrl"`
	SuccessDomain string `json:"successDomain"`
	FailDomain    string `json:"failDomain"`
	StripeToken   string `json:"stripeToken"`
}
