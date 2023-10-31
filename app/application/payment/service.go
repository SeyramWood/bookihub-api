package payment

import (
	"github.com/SeyramWood/app/adapters/gateways"
	"github.com/SeyramWood/config"
)

func NewPaymentService() gateways.PaymentService {
	switch config.Payment().Gateway {
	case "paystack":
		return newPaystackService()
	default:
		return nil
	}
}
