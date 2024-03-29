package models

type DonoRequestBody struct {
	Amount        float32 `json:"amount"`
	UserName      string  `json:"userName"`
	Message       string  `json:"message"`
	ImgUrl        string  `json:"imgUrl"`
	SuccessDomain string  `json:"successDomain"`
	FailDomain    string  `json:"failDomain"`
	ChannelId     int     `json:"channelId"`
}
