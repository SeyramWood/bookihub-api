package requeststructs

import "mime/multipart"

type (
	IncidentRequest struct {
		Time        string                  `json:"time" validate:"required"`
		Location    string                  `json:"location" validate:"required"`
		Description string                  `json:"description" validate:"required|ascii"`
		Image       []*multipart.FileHeader `json:"image" form:"image" validate:"required|max:5|image|size:2MB"`
		Audio       *multipart.FileHeader   `json:"voiceNote" form:"voiceNote" validate:"mimes:mp3,wav,m4a,aac,voc,flac,amr,au|size:10MB"`
		TripID      int                     `json:"tripId" validate:"required"`
		DriverID    int                     `json:"driverId" validate:"required"`
	}
	IncidentUpdateRequest struct {
		Location    string `json:"location" validate:"required"`
		Description string `json:"description" validate:"required|ascii"`
	}
	IncidentImageRequest struct {
		Image []*multipart.FileHeader `json:"image" form:"image" validate:"required|max:5|image|size:2MB"`
	}
	IncidentImageUpdateRequest struct {
		Image *multipart.FileHeader `json:"image" form:"image" validate:"required|min:1|image|size:2MB"`
	}
	IncidentAudioUpdateRequest struct {
		Audio *multipart.FileHeader `json:"voiceNote" form:"voiceNote" validate:"mimes:mp3,wav,m4a,aac,voc,flac,amr,au|size:10MB"`
	}
	IncidentFilterRequest struct {
		Datetime string
	}
)
