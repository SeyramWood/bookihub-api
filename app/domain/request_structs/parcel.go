package requeststructs

import "mime/multipart"

type (
	ParcelRequest struct {
		SenderName        string                  `json:"senderName" validate:"required|string"`
		SenderPhone       string                  `json:"senderPhone" validate:"required|phone"`
		RecipientName     string                  `json:"recipientName" validate:"required|string"`
		RecipientPhone    string                  `json:"recipientPhone" validate:"required|phone"`
		RecipientLocation string                  `json:"recipientLocation" validate:"required|ascii"`
		Image             []*multipart.FileHeader `json:"image" form:"image" validate:"required|min:2|max:3|image|size:2MB"`
		Amount            float64                 `json:"amount" validate:"required"`
		Reference         string                  `json:"reference" validate:"required"`
		TripID            int                     `json:"tripId" validate:"required"`
		DriverID          int                     `json:"driverId" validate:"required"`
	}
	ParcelUpdateRequest struct {
		SenderName        string `json:"senderName" validate:"required|string"`
		SenderPhone       string `json:"senderPhone" validate:"required|phone"`
		RecipientName     string `json:"recipientName" validate:"required|string"`
		RecipientPhone    string `json:"recipientPhone" validate:"required|phone"`
		RecipientLocation string `json:"recipientLocation" validate:"required|ascii"`
	}
	ParcelDeliveredRequest struct {
		PackageCode string                  `json:"packageCode" validate:"required"`
		Image       []*multipart.FileHeader `json:"image" form:"image" validate:"required|min:2|max:2|image|size:2MB"`
	}
	ParcelImageRequest struct {
		Image []*multipart.FileHeader `json:"image" form:"image" validate:"required|max:3|image|size:2MB"`
	}
	ParcelImageUpdateRequest struct {
		Image *multipart.FileHeader `json:"image" form:"image" validate:"required|min:1|image|size:2MB"`
	}
	ParcelFilterRequest struct {
		Status string
	}
)
