package requeststructs

import "mime/multipart"

type (
	CompanyUserRequest struct {
		LastName   string `json:"lastName" validate:"required"`
		OtherName  string `json:"otherName" validate:"required"`
		Phone      string `json:"phone" validate:"required|phone_with_code"`
		OtherPhone string `json:"otherPhone" validate:"phone_with_code"`
		Role       string `json:"role" validate:"required|string"`
		CompanyID  int    `json:"companyId" validate:"required"`
		Username   string `json:"username" validate:"required|email|unique:users"`
	}
	CompanyUserUpdateRequest struct {
		LastName   string `json:"lastName" validate:"required"`
		OtherName  string `json:"otherName" validate:"required"`
		Phone      string `json:"phone" validate:"required|phone_with_code"`
		OtherPhone string `json:"otherPhone" validate:"phone_with_code"`
	}
	CompanyRequest struct {
		CompanyName     string `json:"companyName" validate:"required|unique:companies.name"`
		CompanyPhone    string `json:"companyPhone" validate:"required|phone_with_code|unique:companies.phone"`
		CompanyEmail    string `json:"companyEmail" validate:"required|email|unique:companies.email"`
		Username        string `json:"username" validate:"required|email|unique:users"`
		Password        string `json:"password" validate:"required|min:8"`
		ConfirmPassword string `json:"confirmPassword" validate:"required|min:8|match:Password"`
		Terms           bool   `json:"terms" validate:"required|bool"`
	}
	CompanyUpdateRequest struct {
		CompanyName       string `json:"companyName" validate:"required|unique:companies"`
		CompanyPhone      string `json:"companyPhone" validate:"required|phone_with_code"`
		CompanyOtherPhone string `json:"companyOtherPhone" validate:"phone_with_code"`
		CompanyEmail      string `json:"companyEmail" validate:"required|email"`
	}

	VehicleRequest struct {
		RegistrationNumber string                  `json:"registrationNumber" validate:"required|ascii|unique:vehicles.registration_number"`
		Model              string                  `json:"model" validate:"required|ascii"`
		Seat               int                     `json:"seat" validate:"required"`
		Image              []*multipart.FileHeader `json:"image" form:"image" validate:"required|max:3|image|size:1MB"`
	}
	VehicleUpdateRequest struct {
		RegistrationNumber string `json:"registrationNumber" validate:"required|ascii|unique:vehicles.registration_number"`
		Model              string `json:"model" validate:"required|ascii"`
		Seat               int    `json:"seat" validate:"required"`
	}
	VehicleFormattedRequest struct {
		RegistrationNumber string   `json:"registrationNumber" validate:"required|ascii|unique:vehicles.registration_number"`
		Model              string   `json:"model" validate:"required|ascii"`
		Seat               int      `json:"seat" validate:"required"`
		Image              []string `json:"image" form:"image" validate:"required|image|size:1MB"`
	}
	VehicleImageRequest struct {
		Image []*multipart.FileHeader `json:"image" form:"image" validate:"required|max:3|image|size:1MB"`
	}
	VehicleImageUpdateRequest struct {
		Image *multipart.FileHeader `json:"image" form:"image" validate:"required|image|size:1MB"`
	}
	RouteStopRequest struct {
		Latitude  float64 `json:"latitude" validate:"required"`
		Longitude float64 `json:"longitude" validate:"required"`
	}
	RouteRequest struct {
		From          string              `json:"from" validate:"required|ascii"`
		To            string              `json:"to" validate:"required|ascii"`
		FromLatitude  float64             `json:"fromLatitude" validate:"required"`
		FromLongitude float64             `json:"fromLongitude" validate:"required"`
		ToLatitude    float64             `json:"toLatitude" validate:"required"`
		ToLongitude   float64             `json:"toLongitude" validate:"required"`
		Stops         []*RouteStopRequest `json:"stops"`
		Rate          float64             `json:"rate" validate:"required"`
		Discount      float32             `json:"discount" validate:"min:1|max:100"`
	}
	RouteUpdateRequest struct {
		From          string  `json:"from" validate:"required|ascii"`
		To            string  `json:"to" validate:"required|ascii"`
		FromLatitude  float64 `json:"fromLatitude" validate:"required"`
		FromLongitude float64 `json:"fromLongitude" validate:"required"`
		ToLatitude    float64 `json:"toLatitude" validate:"required"`
		ToLongitude   float64 `json:"toLongitude" validate:"required"`
		Rate          float64 `json:"rate" validate:"required"`
		Discount      float32 `json:"discount" validate:"min:1|max:100"`
	}
	TripRequest struct {
		VehicleID     int      `json:"vehicleId" validate:"required"`
		RouteID       int      `json:"routeId" validate:"required"`
		DriverID      int      `json:"driverId" validate:"required"`
		DepartureDate string   `json:"departureDate" validate:"required|ascii"`
		ArrivalDate   string   `json:"arrivalDate" validate:"required|ascii"`
		ReturnDate    string   `json:"returnDate" validate:"ascii"`
		TripType      string   `json:"tripType" validate:"required|string"`
		Schedule      bool     `json:"schedule"`
		BoardingPoint []string `json:"boardingPoint" validate:"required"`
	}
	TripUpdateRequest struct {
		VehicleID     int                  `json:"vehicleId" validate:"required"`
		DriverID      int                  `json:"driverId" validate:"required"`
		DepartureDate string               `json:"departureDate" validate:"required|ascii"`
		ArrivalDate   string               `json:"arrivalDate" validate:"required|ascii"`
		ReturnDate    string               `json:"returnDate" validate:"ascii"`
		TripType      string               `json:"tripType" validate:"required|string"`
		BoardingPoint []*TripBoardingPoint `json:"boardingPoint" validate:"required"`
	}
	TripInspectionStatusRequest struct {
		Exterior           bool `json:"exterior"`
		Interior           bool `json:"interior"`
		EngineCompartment  bool `json:"engineCompartment"`
		BrakeAndSteering   bool `json:"brakeAndSteering"`
		EmergencyEquipment bool `json:"emergencyEquipment"`
		FuelAndFluid       bool `json:"fuelAndFluid"`
	}
	TripBoardingPoint struct {
		ID       string `json:"id"`
		Location string `json:"location"`
	}
	TripNewBoardingPoint struct {
		BoardingPoint []string `json:"boardingPoint" validate:"required"`
	}
	TripFilterRequest struct {
		Today     bool
		Scheduled bool
		Completed bool
	}
	CustomerTripFilterRequest struct {
		CompanyID     int    `json:"companyId"`
		TripType      string `json:"tripType"`
		From          string `json:"from"`
		To            string `json:"to"`
		DepartureDate string `json:"departureDate"`
		ReturnDate    string `json:"returnDate"`
		Passengers    int    `json:"passenger"`
	}
)
