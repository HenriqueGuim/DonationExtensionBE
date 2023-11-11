package main

import (
	"github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/checkout/session"
)

func createCheckoutSession(amount int, successDomain string, failDomain string, imgUrl string) (*stripe.CheckoutSession, error) {
	params := &stripe.CheckoutSessionParams{
		LineItems: []*stripe.CheckoutSessionLineItemParams{
			&stripe.CheckoutSessionLineItemParams{
				PriceData: &stripe.CheckoutSessionLineItemPriceDataParams{
					Currency: stripe.String(string(stripe.CurrencyEUR)),
					ProductData: &stripe.CheckoutSessionLineItemPriceDataProductDataParams{
						Name: stripe.String("Donation"),
						Images: []*string{
							stripe.String(imgUrl),
						},
					},
					UnitAmount: stripe.Int64(int64(amount)),
				},
				Quantity: stripe.Int64(1),
			},
		},
		Mode:       stripe.String(string(stripe.CheckoutSessionModePayment)),
		SuccessURL: stripe.String(successDomain),
		CancelURL:  stripe.String(failDomain),
	}

	s, err := session.New(params)

	return s, err

}
