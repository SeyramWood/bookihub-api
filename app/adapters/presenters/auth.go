package presenters

import (
	"strings"

	"github.com/gofiber/fiber/v2"

	"github.com/SeyramWood/bookibus/ent"
)

type (
	AuthSession struct {
		ID          int    `json:"id"`
		CompanyID   int    `json:"companyId,omitempty"`
		DisplayName string `json:"displayName"`
		Username    string `json:"username"`
		Role        string `json:"role,omitempty"`
	}
	AuthTokenData struct {
		AccessToken  string `json:"accessToken"`
		RefreshToken string `json:"refreshToken,omitempty"`
	}
)

func AuthSessionResponse(data *ent.User) *fiber.Map {
	return SuccessResponse(FormatSession(data))
}

func FormatSession(data *ent.User) *AuthSession {
	if profile, err := data.Edges.BookibusUserOrErr(); err == nil {
		return &AuthSession{
			ID:          data.ID,
			DisplayName: strings.Split(profile.OtherName, " ")[0],
			Username:    data.Username,
			Role:        string(profile.Role),
		}
	}
	if profile, err := data.Edges.CompanyUserOrErr(); err == nil {
		return &AuthSession{
			ID:          data.ID,
			CompanyID:   profile.Edges.Company.ID,
			DisplayName: strings.Split(profile.OtherName, " ")[0],
			Username:    data.Username,
			Role:        string(profile.Role),
		}
	}
	if profile, err := data.Edges.CustomerOrErr(); err == nil {
		return &AuthSession{
			ID:          data.ID,
			DisplayName: strings.Split(profile.OtherName, " ")[0],
			Username:    data.Username,
		}
	}
	return nil
}
