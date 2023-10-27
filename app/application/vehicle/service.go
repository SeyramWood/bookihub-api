package vehicle

import (
	"strings"

	"github.com/SeyramWood/app/adapters/gateways"
	"github.com/SeyramWood/app/adapters/presenters"
	requeststructs "github.com/SeyramWood/app/domain/request_structs"
	"github.com/SeyramWood/config"
	"github.com/SeyramWood/ent"
)

type service struct {
	repo        gateways.VehicleRepo
	storage     gateways.StorageService
	storagePath string
}

func NewService(repo gateways.VehicleRepo, storage gateways.StorageService) gateways.VehicleService {
	return &service{
		repo:        repo,
		storage:     storage,
		storagePath: "public/vehicles",
	}
}

// AddImage implements gateways.VehicleService.
func (s *service) AddImage(id int, request *requeststructs.VehicleImageRequest) (*ent.Vehicle, error) {
	images, err := s.storage.UploadFiles(s.storagePath, request.Image)
	if err != nil {
		return nil, err
	}
	result, err := s.repo.InsertImage(id, images)
	if err != nil {
		go func(images []string) {
			for _, image := range images {
				s.storage.ExecuteTask(strings.Replace(image, config.App().AppURL, "public", 1), "delete_file")
			}
		}(images)
		return nil, err
	}
	return result, nil
}

// Create implements gateways.VehicleService.
func (s *service) Create(companyId int, request *requeststructs.VehicleRequest) (*ent.Vehicle, error) {
	images, err := s.storage.UploadFiles(s.storagePath, request.Image)
	if err != nil {
		return nil, err
	}
	result, err := s.repo.Insert(companyId, &requeststructs.VehicleFormattedRequest{
		RegistrationNumber: request.RegistrationNumber,
		Model:              request.Model,
		Seat:               request.Seat,
		Image:              images,
	})
	if err != nil {
		go func(images []string) {
			for _, image := range images {
				s.storage.ExecuteTask(strings.Replace(image, config.App().AppURL, "public", 1), "delete_file")
			}
		}(images)
		return nil, err
	}
	return result, nil
}

// Fetch implements gateways.VehicleService.
func (s *service) Fetch(id int) (*ent.Vehicle, error) {
	return s.repo.Read(id)
}

// FetchAll implements gateways.VehicleService.
func (s *service) FetchAll(limit int, offset int) (*presenters.PaginationResponse, error) {
	return s.repo.ReadAll(limit, offset)
}

// FetchAllByCompany implements gateways.VehicleService.
func (s *service) FetchAllByCompany(companyId int, limit int, offset int) (*presenters.PaginationResponse, error) {
	return s.repo.ReadAllByCompany(companyId, limit, offset)
}

// Remove implements gateways.VehicleService.
func (s *service) Remove(id int) error {
	return s.repo.Delete(id)
}

// RemoveImage implements gateways.VehicleService.
func (s *service) RemoveImage(id int) error {
	img, err := s.repo.ReadImage(id)
	if err != nil {
		return err
	}
	if err := s.repo.DeleteImage(id); err != nil {
		return err
	}
	go s.storage.ExecuteTask(strings.Replace(img.Image, config.App().AppURL, "public", 1), "delete_file")
	return nil
}

// Update implements gateways.VehicleService.
func (s *service) Update(id int, request *requeststructs.VehicleUpdateRequest) (*ent.Vehicle, error) {
	return s.repo.Update(id, request)
}

// UpdateImage implements gateways.VehicleService.
func (s *service) UpdateImage(id int, request *requeststructs.VehicleImageUpdateRequest) (*ent.VehicleImage, error) {
	img, err := s.repo.ReadImage(id)
	if err != nil {
		return nil, err
	}
	image, err := s.storage.UploadFile(s.storagePath, request.Image)
	if err != nil {
		return nil, err
	}
	result, err := s.repo.UpdateImage(id, image)
	if err != nil {
		go s.storage.ExecuteTask(strings.Replace(image, config.App().AppURL, "public", 1), "delete_file")
		return nil, err
	}
	go s.storage.ExecuteTask(strings.Replace(img.Image, config.App().AppURL, "public", 1), "delete_file")
	return result, nil
}
