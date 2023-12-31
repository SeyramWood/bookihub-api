package auth

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/SeyramWood/bookibus/app/adapters/gateways"
	"github.com/SeyramWood/bookibus/app/adapters/presenters"
	"github.com/SeyramWood/bookibus/app/application"
	"github.com/SeyramWood/bookibus/app/domain"
	requeststructs "github.com/SeyramWood/bookibus/app/domain/request_structs"
	"github.com/SeyramWood/bookibus/config"
	"github.com/SeyramWood/bookibus/ent"
	"github.com/SeyramWood/bookibus/ent/user"
	"github.com/SeyramWood/bookibus/utils/jwt"
)

const (
	ACCESS_TOKEN_EXPIRY  = time.Second * 60 * 20            // 15 minutes
	REFRESH_TOKEN_EXPIRY = time.Second * 60 * 60 * 24 * 366 // 1year
)

var ErrBadRequest = errors.New("bad request")

type service struct {
	repo        gateways.AuthRepo
	cache       gateways.CacheService
	producer    gateways.EventProducer
	storage     gateways.StorageService
	storagePath string
	jwt         *jwt.JWT
}

func NewService(repo gateways.AuthRepo, cacheSrv gateways.CacheService, jwt *jwt.JWT, producer gateways.EventProducer, storage gateways.StorageService) gateways.AuthService {
	return &service{
		repo:        repo,
		cache:       cacheSrv,
		producer:    producer,
		storage:     storage,
		storagePath: "public/user/avatar",
		jwt:         jwt,
	}
}

// Login implements gateways.AuthService.
func (s *service) Login(request *requeststructs.UserLoginRequest) (*presenters.AuthTokenData, error) {
	result, err := s.repo.ReadByUsername(request.Username)
	if err != nil {
		return nil, err
	}
	if !s.hashCheck(result.Password, request.Password) || !(result.Type == user.Type(request.UserType)) {
		return nil, ErrBadRequest
	}
	session := presenters.FormatSession(result)
	accessTokenExpiry, accessToken, err := s.jwt.GenerateToken(ACCESS_TOKEN_EXPIRY, session)
	if err != nil {
		return nil, err
	}
	refreshTokenExpiry, refreshToken, err := s.jwt.GenerateToken(REFRESH_TOKEN_EXPIRY, strconv.Itoa(session.ID))
	if err != nil {
		return nil, err
	}
	if result.Type == user.TypeCompany && result.Edges.CompanyUser.Role == "driver" {
		return &presenters.AuthTokenData{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		}, nil
	}
	return &presenters.AuthTokenData{
		AccessToken:        accessToken,
		RefreshToken:       refreshToken,
		AccessTokenExpiry:  accessTokenExpiry,
		RefreshTokenExpiry: refreshTokenExpiry,
		User:               session,
	}, nil
}

// Logout implements gateways.AuthService.
func (s *service) Logout() error {
	panic("unimplemented")
}

// RefreshToken implements gateways.AuthService.
func (s *service) RefreshToken(refreshToken string) (*presenters.AuthTokenData, error) {
	claims, err := s.jwt.ValidateToken(refreshToken)
	if err != nil {
		return nil, err
	}
	sessionId, _ := strconv.Atoi(claims["session"].(string))
	result, err := s.repo.ReadByID(sessionId)
	if err != nil {
		return nil, err
	}
	session := presenters.FormatSession(result)
	accessTokenExpiry, accessToken, err := s.jwt.GenerateToken(ACCESS_TOKEN_EXPIRY, session)
	if err != nil {
		return nil, err
	}
	refreshTokenExpiry, refreshToken, err := s.jwt.GenerateToken(REFRESH_TOKEN_EXPIRY, strconv.Itoa(session.ID))
	if err != nil {
		return nil, err
	}

	if result.Type == user.TypeCompany && result.Edges.CompanyUser.Role == "driver" {
		return &presenters.AuthTokenData{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		}, nil
	}

	return &presenters.AuthTokenData{
		AccessToken:        accessToken,
		RefreshToken:       refreshToken,
		AccessTokenExpiry:  accessTokenExpiry,
		RefreshTokenExpiry: refreshTokenExpiry,
		User:               session,
	}, nil
}

// ResetPassword implements gateways.AuthService.
func (s *service) ResetPassword(request *requeststructs.ResetPasswordRequest) (*ent.User, error) {
	return s.repo.ResetPassword(request)
}

// SendPasswordResetCode implements gateways.AuthService.
func (s *service) SendPasswordResetCode(request *requeststructs.UsernameRequest) (string, error) {
	if _, err := s.repo.ReadByUsername(request.Username); err != nil {
		return "", err
	}
	code := application.OTP(6)
	if application.UsernameType(request.Username, "email") {
		s.producer.Queue("notification:email", domain.MailerMessage{
			To:       request.Username,
			Subject:  "PASSWORD RESET - BookiRide",
			Data:     code,
			Template: "resetpassword",
		})
	}
	if application.UsernameType(request.Username, "phone") {
		s.producer.Queue(
			"notification:sms",
			&domain.SMSPayload{
				Message: fmt.Sprintf(
					"You are a step away from completing your password reset! Enter the reset code to proceed. %s",
					code,
				),
				Recipients: []string{fmt.Sprintf("+233%s", request.Username)},
			},
		)
	}
	go func(s *service) {
		_ = s.cache.Set(code, code, time.Minute*15)
	}(s)

	return code, nil
}

// UpdateAvatar implements gateways.AuthService.
func (s *service) UpdateAvatar(userID int, request *requeststructs.AvatarUpdateRequest) (string, error) {
	user, err := s.repo.ReadByID(userID)
	if err != nil {
		return "", err
	}
	avatar, err := s.storage.UploadFile(s.storagePath, request.Avatar)
	if err != nil {
		return "", err
	}
	if err := s.repo.UpdateAvatar(userID, avatar); err != nil {
		go s.storage.ExecuteTask(strings.Replace(avatar, config.App().AppURL, "public", 1), "delete_file")
		return "", err
	}
	go s.storage.ExecuteTask(strings.Replace(user.Avatar, config.App().AppURL, "public", 1), "delete_file")
	return avatar, nil
}

// UpdatePassword implements gateways.AuthService.
func (s *service) UpdatePassword(sessionID int, request *requeststructs.UpdatePasswordRequest) (*ent.User, error) {
	result, err := s.repo.ReadByID(sessionID)
	if err != nil || !s.hashCheck(result.Password, request.CurrentPassword) {
		return nil, ErrBadRequest
	}
	return s.repo.UpdatePassword(sessionID, request)
}

// VerifyOTP implements gateways.AuthService.
func (s *service) VerifyOTP(otp string) bool {
	if !s.cache.Exist(otp) {
		return false
	}
	s.cache.Delete(otp)
	return true
}

func (s *service) hashCheck(hash []byte, plain string) bool {
	if err := bcrypt.CompareHashAndPassword(hash, []byte(plain)); err != nil {
		return false
	}
	return true
}
