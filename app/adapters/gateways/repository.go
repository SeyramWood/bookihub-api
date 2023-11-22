package gateways

import (
	"github.com/SeyramWood/bookibus/app/adapters/presenters"
	requeststructs "github.com/SeyramWood/bookibus/app/domain/request_structs"
	"github.com/SeyramWood/bookibus/ent"
)

type (
	AuthRepo interface {
		ReadByID(id int) (*ent.User, error)
		ReadByUsername(username string) (*ent.User, error)
		UpdatePassword(sessionID int, request *requeststructs.UpdatePasswordRequest) (*ent.User, error)
		UpdateAvatar(userID int, avatar string) error
		ResetPassword(request *requeststructs.ResetPasswordRequest) (*ent.User, error)
	}
	BookibusUserRepo interface {
		ReadAll(limit, offset int) (*presenters.PaginationResponse, error)
		Read(id int) (*ent.BookibusUser, error)
		Insert(request *requeststructs.BookibusUserRequest, password string) (*ent.BookibusUser, error)
		Update(id int, request *requeststructs.BookibusUserUpdateRequest) (*ent.BookibusUser, error)
		Delete(id int) error
	}
	CompanyUserRepo interface {
		ReadAll(limit, offset int) (*presenters.PaginationResponse, error)
		FetchAllByCompany(companyId, limit, offset int) (*presenters.PaginationResponse, error)
		Read(id int) (*ent.CompanyUser, error)
		Insert(request *requeststructs.CompanyUserRequest, password string) (*ent.CompanyUser, error)
		Update(id int, request *requeststructs.CompanyUserUpdateRequest) (*ent.CompanyUser, error)
		Delete(id int) error
	}
	CustomerRepo interface {
		ReadAll(limit, offset int) (*presenters.PaginationResponse, error)
		Read(id int) (*ent.Customer, error)
		Insert(request *requeststructs.CustomerRequest) (*ent.Customer, error)
		Update(id int, request *requeststructs.CustomerUpdateRequest) (*ent.Customer, error)
		Delete(id int) error
	}
	CompanyRepo interface {
		ReadAll(limit, offset int) (*presenters.PaginationResponse, error)
		Read(id int) (*ent.Company, error)
		Insert(request *requeststructs.CompanyRequest) (*ent.Company, error)
		Update(id int, request *requeststructs.CompanyUpdateRequest) (*ent.Company, error)
		Delete(id int) error
	}
	TerminalRepo interface {
		ReadAll(limit, offset int) (*presenters.PaginationResponse, error)
		ReadAllByCompany(companyId, limit, offset int) (*presenters.PaginationResponse, error)
		Read(id int) (*ent.Terminal, error)
		Insert(companyId int, request *requeststructs.TerminalRequest) (*ent.Terminal, error)
		Update(id int, request *requeststructs.TerminalRequest) (*ent.Terminal, error)
		Delete(id int) error
	}
	VehicleRepo interface {
		ReadAll(limit, offset int) (*presenters.PaginationResponse, error)
		ReadAllByCompany(companyId, limit, offset int) (*presenters.PaginationResponse, error)
		Read(id int) (*ent.Vehicle, error)
		ReadImage(id int) (*ent.VehicleImage, error)
		Insert(companyId int, request *requeststructs.VehicleFormattedRequest) (*ent.Vehicle, error)
		InsertImage(id int, request []string) (*ent.Vehicle, error)
		UpdateImage(id int, request string) (*ent.VehicleImage, error)
		Update(id int, request *requeststructs.VehicleUpdateRequest) (*ent.Vehicle, error)
		Delete(id int) error
		DeleteImage(id int) error
	}
	RouteRepo interface {
		ReadAll(limit, offset int) (*presenters.PaginationResponse, error)
		ReadAllDistinct(limit, offset int) (*presenters.PaginationResponse, error)
		ReadAllByCompany(companyId, limit, offset int) (*presenters.PaginationResponse, error)
		Read(id int) (*ent.Route, error)
		Insert(companyId int, request *requeststructs.RouteRequest) (*ent.Route, error)
		InsertRouteStop(id int, request *requeststructs.RouteStopRequest) (*ent.RouteStop, error)
		Update(id int, request *requeststructs.RouteUpdateRequest) (*ent.Route, error)
		UpdateStop(id int, request *requeststructs.RouteStopRequest) (*ent.RouteStop, error)
		Delete(id int) error
		DeleteStop(id int) error
	}
	TripRepo interface {
		ReadAll(limit, offset int, filter *requeststructs.TripFilterRequest) (*presenters.PaginationResponse, error)
		ReadAllByCompany(companyId, limit, offset int, filter *requeststructs.TripFilterRequest) (*presenters.PaginationResponse, error)
		ReadAllByDriver(driverId, limit, offset int, filter *requeststructs.TripFilterRequest) (*presenters.PaginationResponse, error)
		ReadAllCustomer(limit, offset int, filter *requeststructs.CustomerTripFilterRequest) (*presenters.PaginationResponse, error)
		ReadAllPopular(limit, offset int) (*presenters.PaginationResponse, error)
		ReadAllSearch(searchKey string, limit, offset int, filter *requeststructs.TripFilterRequest) (*presenters.PaginationResponse, error)
		ReadAllSearchByCompany(searchKey string, companyId, limit, offset int, filter *requeststructs.TripFilterRequest) (*presenters.PaginationResponse, error)
		Read(id int) (*ent.Trip, error)
		Insert(companyId int, request *requeststructs.TripRequest) (*ent.Trip, error)
		Update(id int, request *requeststructs.TripUpdateRequest) (*ent.Trip, error)
		UpdateStatus(id int, status string) (*ent.Trip, error)
		UpdateInspection(id int, request *requeststructs.TripInspectionStatusRequest) (*ent.Trip, error)
		UpdateSchedule(id int, status bool) (*ent.Trip, error)
		Delete(id int) error
	}
	BookingRepo interface {
		ReadAll(limit, offset int, filter *requeststructs.BookingFilterRequest) (*presenters.PaginationResponse, error)
		ReadAllByCompany(companyId, limit, offset int, filter *requeststructs.BookingFilterRequest) (*presenters.PaginationResponse, error)
		ReadAllCustomer(limit, offset int, filter *requeststructs.BookingFilterRequest, customerId ...int) (*presenters.PaginationResponse, error)
		Read(id int) (*ent.Booking, error)
		Insert(request *requeststructs.BookingRequest, refResponse *requeststructs.PaymentReferenceResponse) (*ent.Booking, error)
		Update(id int, request *requeststructs.BookingUpdateRequest) (*ent.Booking, error)
		CancelBooking(id int, request *requeststructs.BookingCancelRequest) (*ent.Booking, error)
		Delete(id int) error
	}

	ParcelRepo interface {
		ReadAll(limit, offset int, filter *requeststructs.ParcelFilterRequest) (*presenters.PaginationResponse, error)
		ReadAllByCompany(companyId, limit, offset int, filter *requeststructs.ParcelFilterRequest) (*presenters.PaginationResponse, error)
		ReadAllByDriver(driverId, limit, offset int, filter *requeststructs.ParcelFilterRequest) (*presenters.PaginationResponse, error)
		ReadByCode(code string) (*ent.Parcel, error)
		Read(id int) (*ent.Parcel, error)
		ReadImage(id int) (*ent.ParcelImage, error)
		Insert(companyId int, request *requeststructs.ParcelRequest, refResponse *requeststructs.PaymentReferenceResponse, images []string) (*ent.Parcel, error)
		InsertImage(id int, request []string) (*ent.Parcel, error)
		Update(id int, request *requeststructs.ParcelUpdateRequest) (*ent.Parcel, error)
		UpdateStatus(id int, images []string) (*ent.Parcel, error)
		UpdateImage(id int, request string) (*ent.ParcelImage, error)
		Delete(id int) error
		DeleteImage(id int) error
	}
	IncidentRepo interface {
		ReadAll(limit, offset int, filter *requeststructs.IncidentFilterRequest) (*presenters.PaginationResponse, error)
		ReadAllByCompany(companyId, limit, offset int, filter *requeststructs.IncidentFilterRequest) (*presenters.PaginationResponse, error)
		ReadAllByDriver(driverId, limit, offset int, filter *requeststructs.IncidentFilterRequest) (*presenters.PaginationResponse, error)
		Read(id int) (*ent.Incident, error)
		ReadImage(id int) (*ent.IncidentImage, error)
		Insert(companyId int, request *requeststructs.IncidentRequest, images []string, audio ...string) (*ent.Incident, error)
		InsertImage(id int, request []string) (*ent.Incident, error)
		Update(id int, request *requeststructs.IncidentUpdateRequest) (*ent.Incident, error)
		UpdateStatus(id int, status string) (*ent.Incident, error)
		UpdateAudio(id int, request string) (*ent.Incident, error)
		UpdateImage(id int, request string) (*ent.IncidentImage, error)
		Delete(id int) error
		DeleteImage(id int) error
		DeleteAudio(id int) error
	}
)
