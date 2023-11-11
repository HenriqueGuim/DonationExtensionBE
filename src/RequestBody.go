package main

type RequestBody struct {
	Amount        int    `json:"amount"`
	StripeKey     string `json:"stripeKey"`
	SuccessDomain string `json:"successDomain"`
	FailDomain    string `json:"failDomain"`
	ImgUrl        string `json:"imgUrl"`
}
