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
)
