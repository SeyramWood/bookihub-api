package incident

import (
	"strings"

	"github.com/SeyramWood/app/adapters/gateways"
	"github.com/SeyramWood/app/adapters/presenters"
	requeststructs "github.com/SeyramWood/app/domain/request_structs"
	"github.com/SeyramWood/config"
	"github.com/SeyramWood/ent"
)

type service struct {
	repo        gateways.IncidentRepo
	produce     gateways.EventProducer
	storage     gateways.StorageService
	storagePath string
}

func NewService(repo gateways.IncidentRepo, storage gateways.StorageService, produce gateways.EventProducer) gateways.IncidentService {
	return &service{
		repo:        repo,
		produce:     produce,
		storage:     storage,
		storagePath: "public/incidents",
	}
}

// AddImage implements gateways.IncidentService.
func (s *service) AddImage(id int, request *requeststructs.IncidentImageRequest) (*ent.Incident, error) {
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

// Create implements gateways.IncidentService.
func (s *service) Create(companyId int, request *requeststructs.IncidentRequest) (*ent.Incident, error) {
	fileChan := make(chan map[string]any)
	errChan := make(chan error)
	go func() {
		if request.Audio != nil {
			audio, err := s.storage.UploadFile(s.storagePath, request.Audio)
			if err != nil {
				errChan <- err
				return
			}
			images, err := s.storage.UploadFiles(s.storagePath, request.Image)
			if err != nil {
				errChan <- err
				return
			}
			result := make(map[string]any)
			result["images"] = images
			result["audio"] = audio
			fileChan <- result
		}
		images, err := s.storage.UploadFiles(s.storagePath, request.Image)
		if err != nil {
			errChan <- err
			return
		}
		result := make(map[string]any)
		result["images"] = images
		fileChan <- result
	}()
	for {
		select {
		case err := <-errChan:
			return nil, err
		case files := <-fileChan:
			if request.Audio != nil {
				audio := files["audio"].(string)
				images := files["images"].([]string)
				result, err := s.repo.Insert(companyId, request, images, audio)
				if err != nil {
					go func(audio string, images []string) {
						s.storage.ExecuteTask(strings.Replace(audio, config.App().AppURL, "public", 1), "delete_file")
						for _, image := range images {
							s.storage.ExecuteTask(strings.Replace(image, config.App().AppURL, "public", 1), "delete_file")
						}
					}(result.Audio, images)
					return nil, err
				}
				return result, nil
			}
			images := files["images"].([]string)
			result, err := s.repo.Insert(companyId, request, images)
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
	}

}

// Fetch implements gateways.IncidentService.
func (s *service) Fetch(id int) (*ent.Incident, error) {
	return s.repo.Read(id)
}

// FetchAll implements gateways.IncidentService.
func (s *service) FetchAll(limit int, offset int, filter *requeststructs.IncidentFilterRequest) (*presenters.PaginationResponse, error) {
	return s.repo.ReadAll(limit, offset, filter)
}

// FetchAllByCompany implements gateways.IncidentService.
func (s *service) FetchAllByCompany(companyId int, limit int, offset int, filter *requeststructs.IncidentFilterRequest) (*presenters.PaginationResponse, error) {
	return s.repo.ReadAllByCompany(companyId, limit, offset, filter)
}

// FetchAllByDriver implements gateways.IncidentService.
func (s *service) FetchAllByDriver(driverId int, limit int, offset int, filter *requeststructs.IncidentFilterRequest) (*presenters.PaginationResponse, error) {
	return s.repo.ReadAllByDriver(driverId, limit, offset, filter)
}

// Remove implements gateways.IncidentService.
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
		go func(audio string, images []*ent.IncidentImage) {
			s.storage.ExecuteTask(strings.Replace(audio, config.App().AppURL, "public", 1), "delete_file")
			for _, image := range images {
				s.storage.ExecuteTask(strings.Replace(image.Image, config.App().AppURL, "public", 1), "delete_file")
			}
		}(result.Audio, images)
	}
	return nil
}

// RemoveImage implements gateways.IncidentService.
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

// RemoveAudio implements gateways.IncidentService.
func (s *service) RemoveAudio(id int) error {
	result, err := s.repo.Read(id)
	if err != nil {
		return err
	}
	err = s.repo.DeleteAudio(id)
	if err != nil {
		return err
	}
	go s.storage.ExecuteTask(strings.Replace(result.Audio, config.App().AppURL, "public", 1), "delete_file")
	return nil
}

// Update implements gateways.IncidentService.
func (s *service) Update(id int, request *requeststructs.IncidentUpdateRequest) (*ent.Incident, error) {
	return s.repo.Update(id, request)
}

// UpdateAudio implements gateways.IncidentService.
func (s *service) UpdateAudio(id int, request *requeststructs.IncidentAudioUpdateRequest) (string, error) {
	result, err := s.repo.Read(id)
	if err != nil {
		return "", err
	}
	audio, err := s.storage.UploadFile(s.storagePath, request.Audio)
	if err != nil {
		return "", err
	}
	_, err = s.repo.UpdateAudio(id, audio)
	if err != nil {
		go s.storage.ExecuteTask(strings.Replace(audio, config.App().AppURL, "public", 1), "delete_file")
		return "", err
	}
	go s.storage.ExecuteTask(strings.Replace(result.Audio, config.App().AppURL, "public", 1), "delete_file")
	return audio, nil
}

// UpdateImage implements gateways.IncidentService.
func (s *service) UpdateImage(id int, request *requeststructs.IncidentImageUpdateRequest) (*ent.IncidentImage, error) {
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
