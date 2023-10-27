package requeststructs

type (
	BookingRequest struct {
		BoardingPointID string                     `json:"boardingPointId" validate:"required"`
		Passenger       []*BookingPassengerRequest `json:"passenger" validate:"required"`
		Luggage         []*BookingLuggagesRequest  `json:"luggage"`
		Contact         *BookingContactRequest     `json:"contact" validate:"required"`
		SMSNotification bool                       `json:"smsNotification"`
		TransactionType string                     `json:"transactionType" validate:"string"`
		VAT             float64                    `json:"vat"`
		SMSFee          float64                    `json:"smsFee"`
		Amount          float64                    `json:"amount" validate:"required"`
		TripID          int                        `json:"tripId" validate:"required"`
		CompanyID       int                        `json:"companyId" validate:"required"`
		CustomerID      int                        `json:"customerId"`
	}
	BookingUpdateRequest struct {
		BoardingPointID string                           `json:"boardingPointId" validate:"required"`
		Passenger       []*BookingPassengerUpdateRequest `json:"passenger" validate:"required"`
		Luggage         []*BookingLuggagesUpdateRequest  `json:"luggage"`
		Contact         *BookingContactUpdateRequest     `json:"contact" validate:"required"`
		SMSNotification bool                             `json:"smsNotification"`
		TransactionType string                           `json:"transactionType" validate:"string"`
		VAT             float64                          `json:"vat"`
		SMSFee          float64                          `json:"smsFee"`
		Amount          float64                          `json:"amount" validate:"required"`
	}
	BookingPassengerRequest struct {
		FullName string  `json:"fullName" validate:"required|ascii"`
		Amount   float64 `json:"amount" validate:"required"`
		Maturity string  `json:"maturity" validate:"required|string"`
		Gender   string  `json:"gender" validate:"required|string"`
	}
	BookingLuggagesRequest struct {
		Baggage  string  `json:"baggage"`
		Quantity int     `json:"quantity"`
		Amount   float64 `json:"amount"`
	}
	BookingContactRequest struct {
		FullName string `json:"fullName" validate:"required|ascii"`
		Email    string `json:"email" validate:"required|email"`
		Phone    string `json:"phone" validate:"required|phone"`
	}
	BookingPassengerUpdateRequest struct {
		ID       int     `json:"id" validate:"required"`
		FullName string  `json:"fullName" validate:"required|ascii"`
		Maturity string  `json:"maturity" validate:"required|string"`
		Gender   string  `json:"gender" validate:"required|string"`
		Amount   float64 `json:"amount" validate:"required"`
	}
	BookingLuggagesUpdateRequest struct {
		ID       int     `json:"id" validate:"required"`
		Baggage  string  `json:"baggage"`
		Quantity int     `json:"quantity"`
		Amount   float64 `json:"amount"`
	}
	BookingContactUpdateRequest struct {
		ID       int    `json:"id" validate:"required"`
		FullName string `json:"fullName" validate:"required|ascii"`
		Email    string `json:"email" validate:"required|email"`
		Phone    string `json:"phone" validate:"required|phone"`
	}
	BookingCancelRequest struct {
		Amount float64 `json:"amount" validate:"required"`
	}
	BookingFilterRequest struct {
		BookingNumber string
		FullName      string
		Active        bool
		Completed     bool
		Canceled      bool
	}
)
