package presenters

import (
	"strings"

	"github.com/gofiber/fiber/v2"

	"github.com/SeyramWood/bookibus/ent"
)

type (
	AuthSession struct {
		ID            int    `json:"id"`
		CompanyID     int    `json:"companyId,omitempty"`
		CompanyStatus string `json:"companyStatus,omitempty"`
		DisplayName   string `json:"displayName"`
		Username      string `json:"username"`
		Avatar        string `json:"avatar,omitempty"`
		Role          string `json:"role,omitempty"`
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
			Avatar:      data.Avatar,
			Role:        string(profile.Role),
		}
	}
	if profile, err := data.Edges.CompanyUserOrErr(); err == nil {
		return &AuthSession{
			ID:            data.ID,
			CompanyID:     profile.Edges.Company.ID,
			CompanyStatus: string(profile.Edges.Company.OnboardingStatus),
			DisplayName:   strings.Split(profile.OtherName, " ")[0],
			Username:      data.Username,
			Avatar:        data.Avatar,
			Role:          string(profile.Role),
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
