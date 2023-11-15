package requeststructs

import "mime/multipart"

type (
	CompanyUserRequest struct {
		LastName   string `json:"lastName" validate:"required"`
		OtherName  string `json:"otherName" validate:"required"`
		Phone      string `json:"phone" validate:"required|phone"`
		OtherPhone string `json:"otherPhone" validate:"phone"`
		Role       string `json:"role" validate:"required|string"`
		CompanyID  int    `json:"companyId" validate:"required"`
		Username   string `json:"username" validate:"required|email|unique:users.username"`
	}
	CompanyUserUpdateRequest struct {
		LastName   string `json:"lastName" validate:"required"`
		OtherName  string `json:"otherName" validate:"required"`
		Phone      string `json:"phone" validate:"required|phone"`
		OtherPhone string `json:"otherPhone" validate:"phone"`
	}
	CompanyRequest struct {
		CompanyName     string `json:"companyName" validate:"required|unique:companies.name"`
		CompanyPhone    string `json:"companyPhone" validate:"required|phone|unique:companies.phone"`
		CompanyEmail    string `json:"companyEmail" validate:"required|email|unique:companies.email"`
		Username        string `json:"username" validate:"required|email|unique:users"`
		Password        string `json:"password" validate:"required|min:8"`
		ConfirmPassword string `json:"confirmPassword" validate:"required|min:8|match:password"`
		Terms           bool   `json:"terms" validate:"required|bool"`
	}
	CompanyUpdateRequest struct {
		CompanyName       string `json:"companyName" validate:"required"`
		CompanyPhone      string `json:"companyPhone" validate:"required|phone"`
		CompanyOtherPhone string `json:"companyOtherPhone" validate:"phone"`
		CompanyEmail      string `json:"companyEmail" validate:"required|email"`
	}

	VehicleRequest struct {
		RegistrationNumber string                  `json:"registrationNumber" validate:"required|ascii|unique:vehicles.registration_number"`
		Model              string                  `json:"model" validate:"required|ascii"`
		Seat               int                     `json:"seat" validate:"required|int"`
		Image              []*multipart.FileHeader `json:"image" form:"image" validate:"required|slice:max:3|image|size:1MB"`
	}
	VehicleUpdateRequest struct {
		RegistrationNumber string `json:"registrationNumber" validate:"required|ascii|unique:vehicles.registration_number"`
		Model              string `json:"model" validate:"required|ascii"`
		Seat               int    `json:"seat" validate:"required|int"`
	}
	VehicleFormattedRequest struct {
		RegistrationNumber string   `json:"registrationNumber" validate:"required|ascii|unique:vehicles.registration_number"`
		Model              string   `json:"model" validate:"required|ascii"`
		Seat               int      `json:"seat" validate:"required"`
		Image              []string `json:"image" form:"image" validate:"required|image|size:1MB"`
	}
	VehicleImageRequest struct {
		Image []*multipart.FileHeader `json:"image" form:"image" validate:"required|slice:max:3|image|size:1MB"`
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
		FromLatitude  float64             `json:"fromLatitude" validate:"required|float"`
		FromLongitude float64             `json:"fromLongitude" validate:"required|float"`
		ToLatitude    float64             `json:"toLatitude" validate:"required|float"`
		ToLongitude   float64             `json:"toLongitude" validate:"required|float"`
		Stops         []*RouteStopRequest `json:"stops" validate:""`
		Rate          float64             `json:"rate" validate:"required|float"`
		Discount      float32             `json:"discount" validate:"float|min:1|max:100"`
	}
	RouteUpdateRequest struct {
		From          string  `json:"from" validate:"required|ascii"`
		To            string  `json:"to" validate:"required|ascii"`
		FromLatitude  float64 `json:"fromLatitude" validate:"required|float"`
		FromLongitude float64 `json:"fromLongitude" validate:"required|float"`
		ToLatitude    float64 `json:"toLatitude" validate:"required|float"`
		ToLongitude   float64 `json:"toLongitude" validate:"required|float"`
		Rate          float64 `json:"rate" validate:"required|float"`
		Discount      float32 `json:"discount" validate:"min:1|max:100"`
	}
	TripRequest struct {
		FromTerminalID int    `json:"fromTerminalId" validate:"required|int"`
		ToTerminalID   int    `json:"toTerminalId" validate:"required|int"`
		VehicleID      int    `json:"vehicleId" validate:"required|int"`
		RouteID        int    `json:"routeId" validate:"required|int"`
		DriverID       int    `json:"driverId" validate:"required|int"`
		DepartureDate  string `json:"departureDate" validate:"required|rfc3339"`
		ArrivalDate    string `json:"arrivalDate" validate:"required|rfc3339"`
		ReturnDate     string `json:"returnDate" validate:"rfc3339"`
		TripType       string `json:"tripType" validate:"required|string"`
		Schedule       bool   `json:"schedule" validate:""`
	}
	TripUpdateRequest struct {
		FromTerminalID int    `json:"fromTerminalId" validate:"required|int"`
		ToTerminalID   int    `json:"toTerminalId" validate:"required|int"`
		VehicleID      int    `json:"vehicleId" validate:"required|int"`
		DriverID       int    `json:"driverId" validate:"required|int"`
		DepartureDate  string `json:"departureDate" validate:"required|rfc3339"`
		ArrivalDate    string `json:"arrivalDate" validate:"required|rfc3339"`
		ReturnDate     string `json:"returnDate" validate:"rfc3339"`
		TripType       string `json:"tripType" validate:"required|string"`
	}
	TripInspectionStatusRequest struct {
		Exterior           bool `json:"exterior"`
		Interior           bool `json:"interior"`
		EngineCompartment  bool `json:"engineCompartment"`
		BrakeAndSteering   bool `json:"brakeAndSteering"`
		EmergencyEquipment bool `json:"emergencyEquipment"`
		FuelAndFluid       bool `json:"fuelAndFluid"`
	}

	TerminalRequest struct {
		Name string `json:"name" validate:"required|ascii"`
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
