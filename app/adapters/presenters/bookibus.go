package presenters

import (
	"github.com/gofiber/fiber/v2"

	"github.com/SeyramWood/bookibus/ent"
)

type (
	BookibusUserResponseData struct {
		ID         int    `json:"id"`
		LastName   string `json:"lastName"`
		OtherName  string `json:"otherName"`
		Phone      string `json:"phone"`
		OtherPhone string `json:"otherPhone"`
		Role       string `json:"role"`
		CreatedAt  any    `json:"createdAt"`
		UpdatedAt  any    `json:"updatedAt"`
	}
	ChargeResponseData struct {
		ID                       int     `json:"id"`
		PaymentGatewayServiceFee float32 `json:"paymentGatewayServiceFee"`
		TripServiceFee           float32 `json:"tripServiceFee"`
		ParcelServiceFee         float32 `json:"deliveryServiceFee"`
		TripCancellationFee      float32 `json:"tripCancellationFee"`
	}
)

func BookibusUserResponse(data *ent.BookibusUser) *fiber.Map {
	return SuccessResponse(&BookibusUserResponseData{
		ID:         data.ID,
		LastName:   data.LastName,
		OtherName:  data.OtherName,
		Phone:      data.Phone,
		OtherPhone: data.OtherPhone,
		Role:       string(data.Role),
		CreatedAt:  data.CreatedAt,
		UpdatedAt:  data.UpdatedAt,
	})
}
func BookibusUsersResponse(data *PaginationResponse) *fiber.Map {
	var response []*BookibusUserResponseData
	for _, u := range data.Data.([]*ent.BookibusUser) {
		response = append(response, &BookibusUserResponseData{
			ID:         u.ID,
			LastName:   u.LastName,
			OtherName:  u.OtherName,
			Phone:      u.Phone,
			OtherPhone: u.OtherPhone,
			Role:       string(u.Role),
			CreatedAt:  u.CreatedAt,
			UpdatedAt:  u.UpdatedAt,
		})
	}
	data.Data = response
	return SuccessResponse(data)
}

func ChargeResponse(data *ent.Configuration) *fiber.Map {
	if data == nil {
		return SuccessResponse(nil)
	}
	return SuccessResponse(&ChargeResponseData{
		ID:                       data.ID,
		PaymentGatewayServiceFee: data.Charge.PaymentGatewayServiceFee,
		TripServiceFee:           data.Charge.TripServiceFee,
		ParcelServiceFee:         data.Charge.ParcelServiceFee,
		TripCancellationFee:      data.Charge.TripCancellationFee,
	})
}
