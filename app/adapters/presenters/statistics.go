package presenters

type (
	AdminRevenue struct {
		Total    float64 `json:"total"`
		Trip     float64 `json:"trip"`
		Delivery float64 `json:"delivery"`
	}
	AdminOverview struct {
		Month  string  `json:"x"`
		Amount float64 `json:"y"`
	}
	AdminBestSelling struct {
		ID      int     `json:"id"`
		Company string  `json:"company"`
		Product string  `json:"product"`
		Amount  float64 `json:"amount"`
	}
	AdminCompanyOverview struct {
		Revenue  *AdminRevenue `json:"revenue"`
		Staff    int           `json:"staff"`
		Terminal int           `json:"terminal"`
		Fleet    int           `json:"fleet"`
		Customer int           `json:"customer"`
	}
	CompanyTripOverview struct {
		Trip        int `json:"trip"`
		Customer    int `json:"customer"`
		Package     int `json:"package"`
		NewCustomer int `json:"newCustomer"`
		Incident    int `json:"incident"`
	}
	CompanyRevenueOverview struct {
		Day     string  `json:"day"`
		Trip    float64 `json:"trip"`
		Package float64 `json:"package"`
	}
	CompanyMonthRevenue struct {
		CurrentMonth float64 `json:"currentMonth"`
		LastMonth    float64 `json:"lastMonth"`
	}
	CompanyIncidentOverview struct {
		Total      int `json:"total"`
		Pending    int `json:"pending"`
		InProgress int `json:"inProgress"`
		Resolved   int `json:"resolved"`
		Accident   int `json:"accident"`
		Mechanical int `json:"mechanical"`
	}
)
