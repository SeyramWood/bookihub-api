package presenters

type (
	AdminRevenue struct {
		Total    float64 `json:"total"`
		Trip     float64 `json:"trip"`
		Delivery float64 `json:"delivery"`
	}
	AdminOverview struct {
		Month string  `json:"x"`
		Data  float64 `json:"y"`
	}
	AdminBestSelling struct {
		Company string  `json:"company"`
		Product string  `json:"product"`
		Amount  float64 `json:"amount"`
	}
)
