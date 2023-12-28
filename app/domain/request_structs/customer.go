package requeststructs

type (
	CustomerRequest struct {
		LastName        string `json:"lastName" validate:"required"`
		OtherName       string `json:"otherName" validate:"required"`
		Phone           string `json:"phone" validate:"required|phone_with_code"`
		Username        string `json:"username" validate:"required|email|unique:users.username"`
		Password        string `json:"password" validate:"required|min:8"`
		ConfirmPassword string `json:"confirmPassword" validate:"required|min:8|match:password"`
		Terms           bool   `json:"terms" validate:"required|bool"`
	}

	CustomerUpdateRequest struct {
		LastName  string `json:"lastName" validate:"required"`
		OtherName string `json:"otherName" validate:"required"`
		Phone     string `json:"phone" validate:"required|phone_with_code"`
	}
)
