package presenters

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

type (
	PaginationResponse struct {
		Count int `json:"count,omitempty"`
		Data  any `json:"data"`
	}
)

func ErrorResponse(e error) *fiber.Map {
	return &fiber.Map{
		"status": false,
		"error":  e.Error(),
	}
}

func SuccessResponse(data any) *fiber.Map {

	if data == nil {
		return &fiber.Map{
			"status": true,
			"data":   nil,
		}
	}

	d, ok := data.(*PaginationResponse)
	if ok && d.Count == 0 {
		return &fiber.Map{
			"status": true,
			"data":   nil,
		}
	}
	return &fiber.Map{
		"status": true,
		"data":   data,
	}
}

func UnprocessableEntityResponse(errors any) *fiber.Map {
	return &fiber.Map{
		"status": false,
		"errors": errors,
	}
}

func MessageResponse(data any) *fiber.Map {
	return &fiber.Map{
		"status":  true,
		"message": data,
	}
}

func parseNullDatetime(datetime time.Time) any {
	if datetime.Format("2006-00-00T00:00:00Z") == "0001-00-00T00:00:00Z" {
		return nil
	}
	return datetime
}
