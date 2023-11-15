package api

import (
	"encoding/json"

	"github.com/SeyramWood/bookibus/app/adapters/presenters"
	"github.com/gofiber/fiber/v2"
)

func responseJSON(c *fiber.Ctx, data []byte, statusCode int) error {
	resData := fiber.Map{}
	err := json.Unmarshal(data, &resData)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}
	return c.Status(statusCode).JSON(resData)
}
func fiberMapToByte(data *fiber.Map) []byte {
	byteData, err := json.Marshal(data)
	if err != nil {
		return nil
	}
	return byteData
}

func getSession(res []byte, session *presenters.AuthSession) error {
	resData := fiber.Map{}
	err := json.Unmarshal(res, &resData)
	if err != nil {
		return err
	}
	bytesData, err := json.Marshal(resData["data"])
	if err != nil {
		return err
	}
	err = json.Unmarshal(bytesData, session)
	if err != nil {
		return err
	}
	return nil
}
func toBytes(data any) ([]byte, error) {
	bytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}
