package presenters

import (
	"github.com/gofiber/fiber/v2"

	"github.com/SeyramWood/ent"
)

type (
	CustomerResponseData struct {
		ID         int    `json:"id"`
		LastName   string `json:"lastName"`
		OtherName  string `json:"otherName"`
		Phone      string `json:"phone"`
		OtherPhone string `json:"otherPhone"`
		CreatedAt  any    `json:"createdAt"`
		UpdatedAt  any    `json:"updatedAt"`
	}
)

func CustomerResponse(data *ent.Customer) *fiber.Map {
	return SuccessResponse(&CustomerResponseData{
		ID:         data.ID,
		LastName:   data.LastName,
		OtherName:  data.OtherName,
		Phone:      data.Phone,
		OtherPhone: data.OtherPhone,
		CreatedAt:  data.CreatedAt,
		UpdatedAt:  data.UpdatedAt,
	})
}
func CustomersResponse(data *PaginationResponse) *fiber.Map {
	var response []*CustomerResponseData
	for _, u := range data.Data.([]*ent.Customer) {
		response = append(response, &CustomerResponseData{
			ID:         u.ID,
			LastName:   u.LastName,
			OtherName:  u.OtherName,
			Phone:      u.Phone,
			OtherPhone: u.OtherPhone,
			CreatedAt:  u.CreatedAt,
			UpdatedAt:  u.UpdatedAt,
		})
	}
	data.Data = response
	return SuccessResponse(data)
}
