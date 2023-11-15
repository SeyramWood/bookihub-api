package requeststructs

import "mime/multipart"

type (
	IncidentRequest struct {
		Time        string                  `json:"time" validate:"required|rfc3339"`
		Location    string                  `json:"location" validate:"required|ascii"`
		Description string                  `json:"description" validate:"required|ascii"`
		Type        string                  `json:"type" validate:"required|string"`
		Image       []*multipart.FileHeader `json:"image" form:"image" validate:"required|slice:max:5|image|size:2MB"`
		Audio       *multipart.FileHeader   `json:"voiceNote" form:"voiceNote" validate:"mimes:mp3,wav,m4a,aac,voc,flac,amr,au|size:10mb"`
		TripID      int                     `json:"tripId" validate:"required|int"`
		DriverID    int                     `json:"driverId" validate:"required|int"`
	}
	IncidentUpdateRequest struct {
		Location    string `json:"location" validate:"required"`
		Description string `json:"description" validate:"required|ascii"`
		Type        string `json:"type" validate:"required|string"`
	}
	IncidentImageRequest struct {
		Image []*multipart.FileHeader `json:"image" form:"image" validate:"required|slice:max:5|image|size:2MB"`
	}
	IncidentImageUpdateRequest struct {
		Image *multipart.FileHeader `json:"image" form:"image" validate:"required|image|size:2MB"`
	}
	IncidentAudioUpdateRequest struct {
		Audio *multipart.FileHeader `json:"voiceNote" form:"voiceNote" validate:"mimes:mp3,wav,m4a,aac,voc,flac,amr,au|size:10MB"`
	}
	IncidentFilterRequest struct {
		Datetime string
	}
)
