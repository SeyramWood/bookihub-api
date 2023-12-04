package presenters

type (
	TransactionResponseData struct {
		ID              int     `json:"id,omitempty"`
		Reference       string  `json:"reference"`
		TransactionType string  `json:"transType"`
		Amount          float64 `json:"amount"`
		VAT             float64 `json:"vat,omitempty"`
		TransactionFee  float64 `json:"transactionFee"`
		CancellationFee float64 `json:"cancellationFee,omitempty"`
		PaidAt          any     `json:"paidAt"`
		CanceledAt      any     `json:"canceledAt,omitempty"`
		CreatedAt       any     `json:"createdAt,omitempty"`
		UpdatedAt       any     `json:"updatedAt,omitempty"`
	}
)
