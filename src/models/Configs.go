package models

type Configs struct {
	ChannelId              int    `json:"channelId"`
	StripeToken            string `json:"stripeToken"`
	StreamlabsToken        string `json:"streamlabsToken"`
	StreamlabsRefreshToken string `json:"streamlabsRefreshToken"`
}
