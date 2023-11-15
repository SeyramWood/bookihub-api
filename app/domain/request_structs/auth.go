package requeststructs

type (
	UserLoginRequest struct {
		Username string `json:"username" validate:"required|username"`
		Password string `json:"password" validate:"required|min:8"`
		UserType string `json:"userType" validate:"required|string"`
	}
	UsernameRequest struct {
		Username string `json:"username" validate:"required|username"`
	}
	ResetPasswordRequest struct {
		Username           string `json:"username" validate:"required|username"`
		NewPassword        string `json:"newPassword" validate:"required|min:8"`
		ConfirmNewPassword string `json:"confirmNewPassword" validate:"required|min:8|match:newPassword"`
	}
	UpdatePasswordRequest struct {
		CurrentPassword    string `json:"currentPassword" validate:"required|min:8"`
		NewPassword        string `json:"newPassword" validate:"required|min:8"`
		ConfirmNewPassword string `json:"confirmNewPassword" validate:"required|min:8|match:newPassword"`
	}
)
