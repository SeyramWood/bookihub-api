package requeststructs

type (
	BookibusUserRequest struct {
		LastName   string `json:"lastName" validate:"required"`
		OtherName  string `json:"otherName" validate:"required"`
		Phone      string `json:"phone" validate:"required|phone_with_code"`
		OtherPhone string `json:"otherPhone" validate:"phone_with_code"`
		Role       string `json:"role" validate:"required|string"`
		Username   string `json:"username" validate:"required|email|unique:users"`
	}
	BookibusUserUpdateRequest struct {
		LastName   string `json:"lastName" validate:"required"`
		OtherName  string `json:"otherName" validate:"required"`
		Phone      string `json:"phone" validate:"required|phone_with_code"`
		OtherPhone string `json:"otherPhone" validate:"phone_with_code"`
	}

	TransactionChargeRequest struct {
		PaymentGatewayServiceFee float32 `json:"paymentGatewayServiceFee" validate:"required|float|max:100"`
		TripServiceFee           float32 `json:"tripServiceFee" validate:"required|float|max:100"`
		ParcelServiceFee         float32 `json:"deliveryServiceFee" validate:"required|float|max:100"`
		TripCancellationFee      float32 `json:"tripCancellationFee" validate:"required|float|max:100"`
	}
)
