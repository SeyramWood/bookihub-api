package requeststructs

type (
	BookingRequest struct {
		Reference       string                     `json:"reference" validate:"required"`
		Passenger       []*BookingPassengerRequest `json:"passenger" validate:"required"`
		Luggage         []*BookingLuggagesRequest  `json:"luggage" validate:"slice:max:2"`
		Contact         *BookingContactRequest     `json:"contact" validate:"required"`
		SMSNotification bool                       `json:"smsNotification" validate:""`
		VAT             float64                    `json:"vat" validate:"float"`
		Amount          float64                    `json:"amount" validate:"required|float"`
		TripID          int                        `json:"tripId" validate:"required|int"`
		CompanyID       int                        `json:"companyId" validate:"required|int"`
		CustomerID      int                        `json:"customerId" validate:"int"`
	}
	BookingUpdateRequest struct {
		Passenger       []*BookingPassengerUpdateRequest `json:"passenger" validate:"required"`
		Luggage         []*BookingLuggagesUpdateRequest  `json:"luggage" validate:"slice:max:2"`
		Contact         *BookingContactUpdateRequest     `json:"contact" validate:"required"`
		SMSNotification bool                             `json:"smsNotification" validate:""`
		TransactionType string                           `json:"transactionType" validate:"string"`
		VAT             float64                          `json:"vat" validate:"float"`
		Amount          float64                          `json:"amount" validate:"required|float"`
	}
	BookingPassengerRequest struct {
		FullName string  `json:"fullName" validate:"required|ascii"`
		Amount   float64 `json:"amount" validate:"required|float"`
		Maturity string  `json:"maturity" validate:"required|string"`
		Gender   string  `json:"gender" validate:"required|string"`
	}
	BookingLuggagesRequest struct {
		Baggage  string  `json:"baggage" validate:"required|ascii"`
		Quantity int     `json:"quantity" validate:"required|int"`
		Amount   float64 `json:"amount" validate:"required|float"`
	}
	BookingContactRequest struct {
		FullName string `json:"fullName" validate:"required|ascii"`
		Email    string `json:"email" validate:"email"`
		Phone    string `json:"phone" validate:"required|phone_with_code"`
	}
	BookingPassengerUpdateRequest struct {
		ID       int     `json:"id" validate:"required|int"`
		FullName string  `json:"fullName" validate:"required|ascii"`
		Maturity string  `json:"maturity" validate:"required|string"`
		Gender   string  `json:"gender" validate:"required|string"`
		Amount   float64 `json:"amount" validate:"required|float"`
	}
	BookingLuggagesUpdateRequest struct {
		ID       int     `json:"id" validate:"required"`
		Baggage  string  `json:"baggage" validate:"ascii"`
		Quantity int     `json:"quantity" validate:"int"`
		Amount   float64 `json:"amount" validate:"float"`
	}
	BookingContactUpdateRequest struct {
		ID       int    `json:"id" validate:"required"`
		FullName string `json:"fullName" validate:"required|ascii"`
		Email    string `json:"email" validate:"required|email"`
		Phone    string `json:"phone" validate:"required|phone_with_code"`
	}
	BookingCancelRequest struct {
		Amount float64 `json:"amount" validate:"required|float"`
	}
	BookingFilterRequest struct {
		BookingNumber string
		Phone         string
		Active        bool
		Completed     bool
		Canceled      bool
	}
)
