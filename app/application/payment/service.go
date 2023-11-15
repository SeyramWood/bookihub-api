package payment

import (
	"github.com/SeyramWood/bookibus/app/adapters/gateways"
	"github.com/SeyramWood/bookibus/config"
)

func NewPaymentService() gateways.PaymentService {
	switch config.Payment().Gateway {
	case "paystack":
		return newPaystackService()
	default:
		return nil
	}
}
