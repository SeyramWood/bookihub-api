package bookibus_user

import (
	"strings"

	"github.com/SeyramWood/bookibus/app/adapters/gateways"
	"github.com/SeyramWood/bookibus/app/adapters/presenters"
	"github.com/SeyramWood/bookibus/app/application"
	"github.com/SeyramWood/bookibus/app/domain"
	requeststructs "github.com/SeyramWood/bookibus/app/domain/request_structs"
	"github.com/SeyramWood/bookibus/config"
	"github.com/SeyramWood/bookibus/ent"
)

type service struct {
	repo     gateways.BookibusUserRepo
	producer gateways.EventProducer
	storage  gateways.StorageService
}

func NewService(repo gateways.BookibusUserRepo, produce gateways.EventProducer, storage gateways.StorageService) gateways.BookibusUserService {
	return &service{
		repo:     repo,
		producer: produce,
		storage:  storage,
	}
}

// Create implements gateways.BookibusUserService.
func (s *service) Create(request *requeststructs.BookibusUserRequest) (*ent.BookibusUser, error) {
	password := application.RandomString(12)
	result, err := s.repo.Insert(request, password)
	if err != nil {
		return nil, err
	}
	s.producer.Queue("notification:email", domain.MailerMessage{
		To:      request.Username,
		Subject: "NEW USER ACCOUNT - BookiRide",
		Data: map[string]string{
			"username": request.Username,
			"password": password,
			"url":      config.App().AppBookiBusURL,
		},
		Template: "newuser",
	})
	return result, nil
}

// Fetch implements gateways.BookibusUserService.
func (s *service) Fetch(id int) (*ent.BookibusUser, error) {
	return s.repo.Read(id)
}

// FetchAll implements gateways.BookibusUserService.
func (s *service) FetchAll(limit int, offset int) (*presenters.PaginationResponse, error) {
	return s.repo.ReadAll(limit, offset)
}

// Remove implements gateways.BookibusUserService.
func (s *service) Remove(id int) error {
	user, err := s.repo.Read(id)
	if err != nil {
		return err
	}
	if user.Edges.Profile.Avatar != "" {
		go s.storage.ExecuteTask(strings.Replace(user.Edges.Profile.Avatar, config.App().AppURL, "public", 1), "delete_file")
	}
	return s.repo.Delete(id)
}

// Update implements gateways.BookibusUserService.
func (s *service) Update(id int, request *requeststructs.BookibusUserUpdateRequest) (*ent.BookibusUser, error) {
	return s.repo.Update(id, request)
}
