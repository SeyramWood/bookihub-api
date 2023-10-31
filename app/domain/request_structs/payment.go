package requeststructs

type (
	PaymentReferenceResponse struct {
		Status                     bool
		Message, PaidAt, TransType string
	}
)
