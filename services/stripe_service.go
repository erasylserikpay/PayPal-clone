package services

import (
    "log"
    "paypal-clone/models"
    "github.com/stripe/stripe-go/v72"
    "github.com/stripe/stripe-go/v72/charge"
)

func ProcessPayment(payment models.StripePayment) (*stripe.Charge, error) {
    stripe.Key = "your_stripe_secret_key"

    params := &stripe.ChargeParams{
        Amount:   stripe.Int64(payment.Amount),
        Currency: stripe.String(payment.Currency),
        Source:   &stripe.SourceParams{Token: stripe.String(payment.Source)},
    }

    ch, err := charge.New(params)
    if err != nil {
        log.Println("Stripe charge error:", err)
        return nil, err
    }

    return ch, nil
}
