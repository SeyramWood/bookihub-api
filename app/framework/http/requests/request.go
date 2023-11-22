package requests

import (
	"net/http"
	"strconv"

	"github.com/SeyramWood/valid"

	requeststructs "github.com/SeyramWood/bookibus/app/domain/request_structs"
	"github.com/SeyramWood/bookibus/config"
)

func validConfig() *valid.Config {
	return &valid.Config{
		DB: &valid.Database{
			Host: config.DB().Host,
			Port: func() int {
				dbPort, _ := strconv.Atoi(config.DB().Port)
				if dbPort > 0 {
					return dbPort
				}
				return 3306
			}(),
			Name:     config.DB().Name,
			Username: config.DB().Username,
			Password: config.DB().Password,
			Driver:   config.DB().Driver,
		},
	}
}

func ValidateAuthUser(next http.Handler) http.Handler {
	return valid.New(validConfig()).RequestStruct(new(requeststructs.UserLoginRequest)).ValidateRequest(next)
}
func ValidateUsername(next http.Handler) http.Handler {
	return valid.New(validConfig()).RequestStruct(new(requeststructs.UsernameRequest)).ValidateRequest(next)
}
func ValidateAvatarUpdate(next http.Handler) http.Handler {
	return valid.New(validConfig()).RequestStruct(new(requeststructs.AvatarUpdateRequest)).ValidateRequest(next)
}
func ValidateBookibusUser(next http.Handler) http.Handler {
	return valid.New(validConfig()).RequestStruct(new(requeststructs.BookibusUserRequest)).ValidateRequest(next)
}
func ValidateBookibusUserUpdate(next http.Handler) http.Handler {
	return valid.New(validConfig()).RequestStruct(new(requeststructs.BookibusUserUpdateRequest)).ValidateRequest(next)
}
func ValidateCompanyUser(next http.Handler) http.Handler {
	return valid.New(validConfig()).RequestStruct(new(requeststructs.CompanyUserRequest)).ValidateRequest(next)
}
func ValidateCompanyUserUpdate(next http.Handler) http.Handler {
	return valid.New(validConfig()).RequestStruct(new(requeststructs.CompanyUserUpdateRequest)).ValidateRequest(next)
}

func ValidateCustomer(next http.Handler) http.Handler {
	return valid.New(validConfig()).RequestStruct(new(requeststructs.CustomerRequest)).ValidateRequest(next)
}
func ValidateCustomerUpdate(next http.Handler) http.Handler {
	return valid.New(validConfig()).RequestStruct(new(requeststructs.CustomerUpdateRequest)).ValidateRequest(next)
}

func ValidateCompany(next http.Handler) http.Handler {
	return valid.New(validConfig()).RequestStruct(new(requeststructs.CompanyRequest)).ValidateRequest(next)
}
func ValidateCompanyUpdate(next http.Handler) http.Handler {
	return valid.New(validConfig()).RequestStruct(new(requeststructs.CompanyUpdateRequest)).ValidateRequest(next)
}

func ValidateTerminal(next http.Handler) http.Handler {
	return valid.New(validConfig()).RequestStruct(new(requeststructs.TerminalRequest)).ValidateRequest(next)
}

func ValidateVehicle(next http.Handler) http.Handler {
	return valid.New(validConfig()).RequestStruct(new(requeststructs.VehicleRequest)).ValidateRequest(next)
}
func ValidateVehicleUpdate(next http.Handler) http.Handler {
	return valid.New(validConfig()).RequestStruct(new(requeststructs.VehicleUpdateRequest)).ValidateRequest(next)
}
func ValidateVehicleImage(next http.Handler) http.Handler {
	return valid.New(validConfig()).RequestStruct(new(requeststructs.VehicleImageRequest)).ValidateRequest(next)
}
func ValidateVehicleImageUpdate(next http.Handler) http.Handler {
	return valid.New(validConfig()).RequestStruct(new(requeststructs.VehicleImageUpdateRequest)).ValidateRequest(next)
}

func ValidateRoute(next http.Handler) http.Handler {
	return valid.New(validConfig()).RequestStruct(new(requeststructs.RouteRequest)).ValidateRequest(next)
}
func ValidateRouteUpdate(next http.Handler) http.Handler {
	return valid.New(validConfig()).RequestStruct(new(requeststructs.RouteUpdateRequest)).ValidateRequest(next)
}
func ValidateRouteStop(next http.Handler) http.Handler {
	return valid.New(validConfig()).RequestStruct(new(requeststructs.RouteStopRequest)).ValidateRequest(next)
}

func ValidateTrip(next http.Handler) http.Handler {
	return valid.New(validConfig()).RequestStruct(new(requeststructs.TripRequest)).ValidateRequest(next)
}
func ValidateTripUpdate(next http.Handler) http.Handler {
	return valid.New(validConfig()).RequestStruct(new(requeststructs.TripUpdateRequest)).ValidateRequest(next)
}

func ValidateBooking(next http.Handler) http.Handler {
	return valid.New(validConfig()).RequestStruct(new(requeststructs.BookingRequest)).ValidateRequest(next)
}
func ValidateBookingUpdate(next http.Handler) http.Handler {
	return valid.New(validConfig()).RequestStruct(new(requeststructs.BookingUpdateRequest)).ValidateRequest(next)
}
func ValidateBookingCancel(next http.Handler) http.Handler {
	return valid.New(validConfig()).RequestStruct(new(requeststructs.BookingCancelRequest)).ValidateRequest(next)
}

func ValidateParcel(next http.Handler) http.Handler {
	return valid.New(validConfig()).RequestStruct(new(requeststructs.ParcelRequest)).ValidateRequest(next)
}
func ValidateParcelImage(next http.Handler) http.Handler {
	return valid.New(validConfig()).RequestStruct(new(requeststructs.ParcelImageRequest)).ValidateRequest(next)
}
func ValidateParcelUpdate(next http.Handler) http.Handler {
	return valid.New(validConfig()).RequestStruct(new(requeststructs.ParcelUpdateRequest)).ValidateRequest(next)
}
func ValidateParcelImageUpdate(next http.Handler) http.Handler {
	return valid.New(validConfig()).RequestStruct(new(requeststructs.ParcelImageUpdateRequest)).ValidateRequest(next)
}
func ValidateParcelDeliveredUpdate(next http.Handler) http.Handler {
	return valid.New(validConfig()).RequestStruct(new(requeststructs.ParcelDeliveredRequest)).ValidateRequest(next)
}

func ValidateIncident(next http.Handler) http.Handler {
	return valid.New(validConfig()).RequestStruct(new(requeststructs.IncidentRequest)).ValidateRequest(next)
}
func ValidateIncidentUpdate(next http.Handler) http.Handler {
	return valid.New(validConfig()).RequestStruct(new(requeststructs.IncidentUpdateRequest)).ValidateRequest(next)
}
func ValidateIncidentImage(next http.Handler) http.Handler {
	return valid.New(validConfig()).RequestStruct(new(requeststructs.IncidentImageRequest)).ValidateRequest(next)
}
func ValidateIncidentImageUpdate(next http.Handler) http.Handler {
	return valid.New(validConfig()).RequestStruct(new(requeststructs.IncidentImageUpdateRequest)).ValidateRequest(next)
}
func ValidateIncidentAudioUpdate(next http.Handler) http.Handler {
	return valid.New(validConfig()).RequestStruct(new(requeststructs.IncidentAudioUpdateRequest)).ValidateRequest(next)
}
