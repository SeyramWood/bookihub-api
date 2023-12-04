package domain

import (
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/SeyramWood/bookibus/app/application"
)

type (
	MailerMessage struct {
		From        string
		FromName    string
		To          string
		Subject     string
		Attachments []string
		Data        any
		DataMap     map[string]any
		Template    string
	}
	SMSPayload struct {
		Message    string   `json:"message"`
		Recipients []string `json:"recipients"`
	}
)

func ComputeTransactionFee(gatewayFee, serviceFee float32, amount float64) (fee float64) {
	return math.Round(((float64(gatewayFee+serviceFee)/100)*amount)*100) / 100
}
func ComputeCancellationFee(cancellationPercentage float32, amount float64) (fee float64) {
	return math.Round(((float64(cancellationPercentage)/100)*amount)*100) / 100
}
func ComputeUniqueCode(delimiter string) (code string) {
	t := time.Now()
	st := strings.Split(fmt.Sprintf("%d", t.Year()), "")[2:]
	code = fmt.Sprintf("%s%s%s%sBR", application.OTP(4), st[0], st[1], delimiter)
	return
}
