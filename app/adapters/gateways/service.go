package gateways

import (
	"mime/multipart"
	"time"

	"github.com/SeyramWood/bookibus/app/adapters/presenters"
	requeststructs "github.com/SeyramWood/bookibus/app/domain/request_structs"
	"github.com/SeyramWood/bookibus/ent"
)

type (
	StorageService interface {
		UploadFile(dir string, f *multipart.FileHeader) (string, error)
		UploadResizeImage(dir string, f *multipart.FileHeader, width, height int) (string, error)
		UploadFiles(dir string, files []*multipart.FileHeader) ([]string, error)
		Disk(disk string) StorageService
		ExecuteTask(data any, taskType string)
		Listen()
		Done()
		Close()
	}
	CacheService interface {
		Set(key string, value any, ttl time.Duration) error
		Get(key string, obj any) error
		Exist(key string) bool
		Delete(key string) error
	}
	EventProducer interface {
		Queue(queue string, payload any)
	}
	OTPVerifier interface {
		VerifyOTP(otp string) bool
	}
	PaymentService interface {
		Pay(request any) (any, error)
		Verify(reference string) (*requeststructs.PaymentReferenceResponse, error)
	}
	AuthService interface {
		OTPVerifier
		Login(request *requeststructs.UserLoginRequest) (*presenters.AuthTokenData, error)
		RefreshToken(refreshToken string) (*presenters.AuthTokenData, error)
		Logout() error
		SendPasswordResetCode(request *requeststructs.UsernameRequest) (string, error)
		UpdatePassword(sessionID int, request *requeststructs.UpdatePasswordRequest) (*ent.User, error)
		ResetPassword(request *requeststructs.ResetPasswordRequest) (*ent.User, error)
	}
	BookibusUserService interface {
		FetchAll(limit, offset int) (*presenters.PaginationResponse, error)
		Fetch(id int) (*ent.BookibusUser, error)
		Create(request *requeststructs.BookibusUserRequest) (*ent.BookibusUser, error)
		Update(id int, request *requeststructs.BookibusUserUpdateRequest) (*ent.BookibusUser, error)
		Remove(id int) error
	}
	CompanyUserService interface {
		FetchAll(limit, offset int) (*presenters.PaginationResponse, error)
		FetchAllByCompany(companyId, limit, offset int) (*presenters.PaginationResponse, error)
		Fetch(id int) (*ent.CompanyUser, error)
		Create(request *requeststructs.CompanyUserRequest) (*ent.CompanyUser, error)
		Update(id int, request *requeststructs.CompanyUserUpdateRequest) (*ent.CompanyUser, error)
		Remove(id int) error
	}
	CustomerService interface {
		FetchAll(limit, offset int) (*presenters.PaginationResponse, error)
		Fetch(id int) (*ent.Customer, error)
		Create(request *requeststructs.CustomerRequest) (*ent.Customer, error)
		Update(id int, request *requeststructs.CustomerUpdateRequest) (*ent.Customer, error)
		Remove(id int) error
	}
	CompanyService interface {
		FetchAll(limit, offset int) (*presenters.PaginationResponse, error)
		Fetch(id int) (*ent.Company, error)
		Create(request *requeststructs.CompanyRequest) (*ent.Company, error)
		Update(id int, request *requeststructs.CompanyUpdateRequest) (*ent.Company, error)
		Remove(id int) error
	}
	TerminalService interface {
		FetchAll(limit, offset int) (*presenters.PaginationResponse, error)
		FetchAllByCompany(companyId, limit, offset int) (*presenters.PaginationResponse, error)
		Fetch(id int) (*ent.Terminal, error)
		Create(companyId int, request *requeststructs.TerminalRequest) (*ent.Terminal, error)
		Update(id int, request *requeststructs.TerminalRequest) (*ent.Terminal, error)
		Remove(id int) error
	}
	VehicleService interface {
		FetchAll(limit, offset int) (*presenters.PaginationResponse, error)
		FetchAllByCompany(companyId, limit, offset int) (*presenters.PaginationResponse, error)
		Fetch(id int) (*ent.Vehicle, error)
		Create(companyId int, request *requeststructs.VehicleRequest) (*ent.Vehicle, error)
		AddImage(id int, request *requeststructs.VehicleImageRequest) (*ent.Vehicle, error)
		UpdateImage(id int, request *requeststructs.VehicleImageUpdateRequest) (*ent.VehicleImage, error)
		Update(id int, request *requeststructs.VehicleUpdateRequest) (*ent.Vehicle, error)
		Remove(id int) error
		RemoveImage(id int) error
	}
	RouteService interface {
		FetchAll(limit, offset int) (*presenters.PaginationResponse, error)
		FetchAllDistinct(limit, offset int) (*presenters.PaginationResponse, error)
		FetchAllByCompany(companyId, limit, offset int) (*presenters.PaginationResponse, error)
		Fetch(id int) (*ent.Route, error)
		Create(companyId int, request *requeststructs.RouteRequest) (*ent.Route, error)
		AddRouteStop(id int, request *requeststructs.RouteStopRequest) (*ent.RouteStop, error)
		Update(id int, request *requeststructs.RouteUpdateRequest) (*ent.Route, error)
		UpdateStop(id int, request *requeststructs.RouteStopRequest) (*ent.RouteStop, error)
		Remove(id int) error
		RemoveStop(id int) error
	}
	TripService interface {
		FetchAll(limit, offset int, filter *requeststructs.TripFilterRequest) (*presenters.PaginationResponse, error)
		FetchAllByCompany(companyId, limit, offset int, filter *requeststructs.TripFilterRequest) (*presenters.PaginationResponse, error)
		FetchAllByDriver(driverId, limit, offset int, filter *requeststructs.TripFilterRequest) (*presenters.PaginationResponse, error)
		FetchAllCustomer(limit, offset int, filter *requeststructs.CustomerTripFilterRequest) (*presenters.PaginationResponse, error)
		FetchAllPopular(limit, offset int) (*presenters.PaginationResponse, error)
		Fetch(id int) (*ent.Trip, error)
		Create(companyId int, request *requeststructs.TripRequest) (*ent.Trip, error)
		Update(id int, request *requeststructs.TripUpdateRequest) (*ent.Trip, error)
		UpdateStatus(id int, status string) (*ent.Trip, error)
		UpdateInspection(id int, request *requeststructs.TripInspectionStatusRequest) (*ent.Trip, error)
		UpdateSchedule(id int, status bool) (*ent.Trip, error)
		Remove(id int) error
	}
	BookingService interface {
		FetchAll(limit, offset int, filter *requeststructs.BookingFilterRequest) (*presenters.PaginationResponse, error)
		FetchAllByCompany(companyId, limit, offset int, filter *requeststructs.BookingFilterRequest) (*presenters.PaginationResponse, error)
		FetchAllCustomer(limit, offset int, filter *requeststructs.BookingFilterRequest, customerId ...int) (*presenters.PaginationResponse, error)
		Fetch(id int) (*ent.Booking, error)
		Create(request *requeststructs.BookingRequest, transType string) (*ent.Booking, error)
		SaveToCache(reference string, request *requeststructs.BookingRequest) error
		Update(id int, request *requeststructs.BookingUpdateRequest) (*ent.Booking, error)
		CancelBooking(id int, request *requeststructs.BookingCancelRequest) (*ent.Booking, error)
		Remove(id int) error
	}
	ParcelService interface {
		FetchAll(limit, offset int, filter *requeststructs.ParcelFilterRequest) (*presenters.PaginationResponse, error)
		FetchAllByCompany(companyId, limit, offset int, filter *requeststructs.ParcelFilterRequest) (*presenters.PaginationResponse, error)
		FetchAllByDriver(driverId, limit, offset int, filter *requeststructs.ParcelFilterRequest) (*presenters.PaginationResponse, error)
		FetchByCode(code string) (*ent.Parcel, error)
		Fetch(id int) (*ent.Parcel, error)
		Create(companyId int, request *requeststructs.ParcelRequest, transType string) (*ent.Parcel, error)
		AddImage(id int, request *requeststructs.ParcelImageRequest) (*ent.Parcel, error)
		Update(id int, request *requeststructs.ParcelUpdateRequest) (*ent.Parcel, error)
		UpdateStatus(id int, request *requeststructs.ParcelDeliveredRequest) (*ent.Parcel, error)
		UpdateImage(id int, request *requeststructs.ParcelImageUpdateRequest) (*ent.ParcelImage, error)
		Remove(id int) error
		RemoveImage(id int) error
	}
	IncidentService interface {
		FetchAll(limit, offset int, filter *requeststructs.IncidentFilterRequest) (*presenters.PaginationResponse, error)
		FetchAllByCompany(companyId, limit, offset int, filter *requeststructs.IncidentFilterRequest) (*presenters.PaginationResponse, error)
		FetchAllByDriver(driverId, limit, offset int, filter *requeststructs.IncidentFilterRequest) (*presenters.PaginationResponse, error)
		Fetch(id int) (*ent.Incident, error)
		Create(companyId int, request *requeststructs.IncidentRequest) (*ent.Incident, error)
		AddImage(id int, request *requeststructs.IncidentImageRequest) (*ent.Incident, error)
		Update(id int, request *requeststructs.IncidentUpdateRequest) (*ent.Incident, error)
		UpdateStatus(id int, status string) (*ent.Incident, error)
		UpdateAudio(id int, request *requeststructs.IncidentAudioUpdateRequest) (string, error)
		UpdateImage(id int, request *requeststructs.IncidentImageUpdateRequest) (*ent.IncidentImage, error)
		Remove(id int) error
		RemoveImage(id int) error
		RemoveAudio(id int) error
	}
)
