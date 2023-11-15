package requeststructs

import "mime/multipart"

type (
	ParcelRequest struct {
		SenderName        string                  `json:"senderName" validate:"required|string"`
		SenderPhone       string                  `json:"senderPhone" validate:"required|phone"`
		SenderEmail       string                  `json:"senderEmail" validate:"email"`
		RecipientName     string                  `json:"recipientName" validate:"required|string"`
		RecipientPhone    string                  `json:"recipientPhone" validate:"required|phone"`
		RecipientLocation string                  `json:"recipientLocation" validate:"required|ascii"`
		Type              string                  `json:"type" validate:"required|string"`
		Image             []*multipart.FileHeader `json:"image" form:"image" validate:"required|slice:min:2|slice:max:3|image|size:2MB"`
		Weight            float32                 `json:"weight" validate:"required|float"`
		Amount            float64                 `json:"amount" validate:"required|float"`
		Reference         string                  `json:"reference" validate:"required"`
		TripID            int                     `json:"tripId" validate:"required"`
		DriverID          int                     `json:"driverId" validate:"required"`
	}
	ParcelUpdateRequest struct {
		SenderName        string `json:"senderName" validate:"required|string"`
		SenderPhone       string `json:"senderPhone" validate:"required|phone"`
		SenderEmail       string `json:"senderEmail" validate:"required|email"`
		RecipientName     string `json:"recipientName" validate:"required|string"`
		RecipientPhone    string `json:"recipientPhone" validate:"required|phone"`
		RecipientLocation string `json:"recipientLocation" validate:"required|ascii"`
		Type              string `json:"type" validate:"required|string"`
	}
	ParcelDeliveredRequest struct {
		PackageCode string                  `json:"packageCode" validate:"required"`
		Image       []*multipart.FileHeader `json:"image" form:"image" validate:"required|slice:min:2|slice:max:2|image|size:2MB"`
	}
	ParcelImageRequest struct {
		Image []*multipart.FileHeader `json:"image" form:"image" validate:"required|slice:max:3|image|size:2MB"`
	}
	ParcelImageUpdateRequest struct {
		Image *multipart.FileHeader `json:"image" form:"image" validate:"required|slice:min:1|image|size:2MB"`
	}
	ParcelFilterRequest struct {
		Status string
	}
)
