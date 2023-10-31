package parcel

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/SeyramWood/app/adapters/gateways"
	"github.com/SeyramWood/app/adapters/presenters"
	"github.com/SeyramWood/app/application/payment"
	"github.com/SeyramWood/app/domain"
	requeststructs "github.com/SeyramWood/app/domain/request_structs"
	"github.com/SeyramWood/config"
	"github.com/SeyramWood/ent"
)

var (
	ErrWrongPackageCode error = errors.New("wrong package code")
)

type service struct {
	repo        gateways.ParcelRepo
	payment     gateways.PaymentService
	producer    gateways.EventProducer
	storage     gateways.StorageService
	storagePath string
}

func NewService(repo gateways.ParcelRepo, storage gateways.StorageService, payment gateways.PaymentService, produce gateways.EventProducer) gateways.ParcelService {
	return &service{
		repo:        repo,
		payment:     payment,
		producer:    produce,
		storage:     storage,
		storagePath: "public/parcels",
	}
}

// AddImage implements gateways.ParcelService.
func (s *service) AddImage(id int, request *requeststructs.ParcelImageRequest) (*ent.Parcel, error) {
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

// Create implements gateways.ParcelService.
func (s *service) Create(companyId int, request *requeststructs.ParcelRequest, transType string) (*ent.Parcel, error) {
	images, err := s.storage.UploadFiles(s.storagePath, request.Image)
	if err != nil {
		return nil, err
	}
	if transType == "cash" {
		result, err := s.repo.Insert(companyId, request, &requeststructs.PaymentReferenceResponse{
			PaidAt:    time.Now().Format(time.RFC3339),
			TransType: strings.ToLower(transType),
		}, images)
		if err != nil {
			go func(images []string) {
				for _, image := range images {
					s.storage.ExecuteTask(strings.Replace(image, config.App().AppURL, "public", 1), "delete_file")
				}
			}(images)
			return nil, err
		}
		s.producer.Queue("notification:sms", domain.SMSPayload{
			Message:    fmt.Sprintf("Thanks for your trusting us. Kindly mention the package code to only the recipient.\n Package Code: %s", result.ParcelCode),
			Recipients: []string{request.SenderPhone},
		})
		return result, nil
	}
	resp, err := s.payment.Verify(request.Reference)
	if err != nil {
		return nil, err
	}
	if resp.Status && resp.Message == payment.VerificationSuccessful {
		result, err := s.repo.Insert(companyId, request, resp, images)
		if err != nil {
			go func(images []string) {
				for _, image := range images {
					s.storage.ExecuteTask(strings.Replace(image, config.App().AppURL, "public", 1), "delete_file")
				}
			}(images)
			return nil, err
		}
		s.producer.Queue("notification:sms", domain.SMSPayload{
			Message:    fmt.Sprintf("Thanks for your trusting us. Kindly mention the package code to only the recipient.\n Package Code: %s", result.ParcelCode),
			Recipients: []string{request.SenderPhone},
		})
		return result, nil
	}
	return nil, errors.New(strings.ToLower(resp.Message))
}

// Fetch implements gateways.ParcelService.
func (s *service) Fetch(id int) (*ent.Parcel, error) {
	return s.repo.Read(id)
}

// FetchAll implements gateways.ParcelService.
func (s *service) FetchAll(limit int, offset int, filter *requeststructs.ParcelFilterRequest) (*presenters.PaginationResponse, error) {
	return s.repo.ReadAll(limit, offset, filter)
}

// FetchAllByCompany implements gateways.ParcelService.
func (s *service) FetchAllByCompany(companyId int, limit int, offset int, filter *requeststructs.ParcelFilterRequest) (*presenters.PaginationResponse, error) {
	return s.repo.ReadAllByCompany(companyId, limit, offset, filter)
}

// FetchAllByDriver implements gateways.ParcelService.
func (s *service) FetchAllByDriver(driverId int, limit int, offset int, filter *requeststructs.ParcelFilterRequest) (*presenters.PaginationResponse, error) {
	return s.repo.ReadAllByDriver(driverId, limit, offset, filter)
}

// Remove implements gateways.ParcelService.
func (s *service) Remove(id int) error {
	result, err := s.repo.Read(id)
	if err != nil {
		return err
	}
	err = s.repo.Delete(id)
	if err != nil {
		return err
	}
	if images, err := result.Edges.ImagesOrErr(); err != nil && len(images) > 0 {
		go func(images []*ent.ParcelImage) {
			for _, image := range images {
				s.storage.ExecuteTask(strings.Replace(image.Image, config.App().AppURL, "public", 1), "delete_file")
			}
		}(images)
	}
	return nil
}

// RemoveImage implements gateways.ParcelService.
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

// Update implements gateways.ParcelService.
func (s *service) Update(id int, request *requeststructs.ParcelUpdateRequest) (*ent.Parcel, error) {
	return s.repo.Update(id, request)
}

// UpdateImage implements gateways.ParcelService.
func (s *service) UpdateImage(id int, request *requeststructs.ParcelImageUpdateRequest) (*ent.ParcelImage, error) {
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

// UpdateStatus implements gateways.ParcelService.
func (s *service) UpdateStatus(id int, request *requeststructs.ParcelDeliveredRequest) (*ent.Parcel, error) {
	result, err := s.repo.Read(id)
	if err != nil {
		return nil, err
	}
	if result.ParcelCode != request.PackageCode {
		return nil, ErrWrongPackageCode
	}
	images, err := s.storage.UploadFiles(s.storagePath, request.Image)
	if err != nil {
		return nil, err
	}
	res, err := s.repo.UpdateStatus(id, images)
	if err != nil {
		go func(images []string) {
			for _, image := range images {
				s.storage.ExecuteTask(strings.Replace(image, config.App().AppURL, "public", 1), "delete_file")
			}
		}(images)
		return nil, err
	}
	s.producer.Queue("notification:sms", domain.SMSPayload{
		Message:    fmt.Sprintf("Your package has been delivered successfully.\n Package Code: %s", result.ParcelCode),
		Recipients: []string{result.SenderPhone},
	})
	return res, nil
}
