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
		Username   string `json:"username" validate:"required|email|unique:users.username"`
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
		Username        string `json:"username" validate:"required|email|unique:users.username"`
		Password        string `json:"password" validate:"required|min:8"`
		ConfirmPassword string `json:"confirmPassword" validate:"required|min:8|match:password"`
		Terms           bool   `json:"terms" validate:"required|bool"`
	}
	CompanyUpdateRequest struct {
		CompanyName  string `json:"companyName" validate:"required"`
		CompanyPhone string `json:"companyPhone" validate:"required|phone_with_code"`
		CompanyEmail string `json:"companyEmail" validate:"required|email"`
	}
	CompanyOnboardingRequest struct {
		BankAccount   *CompanyBankAccountRequest   `json:"bankAccount" validate:"-"`
		ContactPerson *CompanyContactPersonRequest `json:"contactPerson" validate:"-"`
	}
	BookiOnboardingRequest struct {
		CompanyName     string                       `json:"companyName" validate:"required|unique:companies.name"`
		CompanyPhone    string                       `json:"companyPhone" validate:"required|phone_with_code|unique:companies.phone"`
		CompanyEmail    string                       `json:"companyEmail" validate:"required|email|unique:companies.email"`
		BankAccount     *CompanyBankAccountRequest   `json:"bankAccount" validate:"-"`
		ContactPerson   *CompanyContactPersonRequest `json:"contactPerson" validate:"-"`
		ManagerUsername string                       `json:"managerUsername" validate:"required|email|unique:users.username"`
	}
	CompanyBankAccountRequest struct {
		AccountName   string `json:"accountName" validate:"required|ascii"`
		AccountNumber string `json:"accountNumber" validate:"required|ascii"`
		Bank          string `json:"bank" validate:"required|ascii"`
		Branch        string `json:"branch" validate:"required|ascii"`
	}
	CompanyMomoAccountRequest struct {
		AccountName string `json:"accountName" validate:"required|ascii"`
		PhoneNumber string `json:"phoneNumber" validate:"required|phone_with_code"`
		Provider    string `json:"provider" validate:"required|string"`
	}
	CompanyContactPersonRequest struct {
		Name     string `json:"name" validate:"ascii"`
		Position string `json:"position" validate:"ascii"`
		Phone    string `json:"phone" validate:"phone_with_code"`
		Email    string `json:"email" validate:"email"`
	}
	CompanyContactPersonUpdateRequest struct {
		Name     string `json:"name" validate:"required|ascii"`
		Position string `json:"position" validate:"required|ascii"`
		Phone    string `json:"phone" validate:"required|phone_with_code"`
		Email    string `json:"email" validate:"email"`
	}
	CompanyCertificateUpdateRequest struct {
		BusinessCertificate *multipart.FileHeader `json:"businessCertificate" form:"businessCertificate" validate:"required|image|size:3MB"`
	}
	CompanyLogoUpdateRequest struct {
		BusinessLogo *multipart.FileHeader `json:"businessLogo" form:"businessLogo" validate:"required|image|size:2MB"`
	}
	VehicleRequest struct {
		RegistrationNumber string                  `json:"registrationNumber" validate:"required|ascii"`
		Model              string                  `json:"model" validate:"required|ascii"`
		Seat               int                     `json:"seat" validate:"required|int"`
		Image              []*multipart.FileHeader `json:"image" form:"image" validate:"required|slice:max:3|image|size:3MB"`
	}
	VehicleUpdateRequest struct {
		RegistrationNumber string `json:"registrationNumber" validate:"required|ascii"`
		Model              string `json:"model" validate:"required|ascii"`
		Seat               int    `json:"seat" validate:"required|int"`
	}
	VehicleFormattedRequest struct {
		RegistrationNumber string   `json:"registrationNumber" validate:"required|ascii"`
		Model              string   `json:"model" validate:"required|ascii"`
		Seat               int      `json:"seat" validate:"required"`
		Image              []string `json:"image" form:"image" validate:"required|image|size:3MB"`
	}
	VehicleImageRequest struct {
		Image []*multipart.FileHeader `json:"image" form:"image" validate:"required|slice:max:3|image|size:3MB"`
	}
	VehicleImageUpdateRequest struct {
		Image *multipart.FileHeader `json:"image" form:"image" validate:"required|image|size:3MB"`
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
	}
	RouteUpdateRequest struct {
		From          string  `json:"from" validate:"required|ascii"`
		To            string  `json:"to" validate:"required|ascii"`
		FromLatitude  float64 `json:"fromLatitude" validate:"required|float"`
		FromLongitude float64 `json:"fromLongitude" validate:"required|float"`
		ToLatitude    float64 `json:"toLatitude" validate:"required|float"`
		ToLongitude   float64 `json:"toLongitude" validate:"required|float"`
	}
	TripRequest struct {
		FromTerminalID int     `json:"fromTerminalId" validate:"required|int"`
		ToTerminalID   int     `json:"toTerminalId" validate:"required|int"`
		VehicleID      int     `json:"vehicleId" validate:"required|int"`
		RouteID        int     `json:"routeId" validate:"required|int"`
		DriverID       int     `json:"driverId" validate:"required|int"`
		DepartureDate  string  `json:"departureDate" validate:"required|rfc3339"`
		ArrivalDate    string  `json:"arrivalDate" validate:"required|rfc3339"`
		ReturnDate     string  `json:"returnDate" validate:"rfc3339"`
		TripType       string  `json:"tripType" validate:"required|string"`
		Schedule       bool    `json:"schedule" validate:""`
		Rate           float64 `json:"rate" validate:"required|float"`
		Discount       float32 `json:"discount" validate:"min:1|max:100"`
	}
	TripUpdateRequest struct {
		FromTerminalID int     `json:"fromTerminalId" validate:"required|int"`
		ToTerminalID   int     `json:"toTerminalId" validate:"required|int"`
		VehicleID      int     `json:"vehicleId" validate:"required|int"`
		RouteID        int     `json:"routeId" validate:"required|int"`
		DriverID       int     `json:"driverId" validate:"required|int"`
		DepartureDate  string  `json:"departureDate" validate:"required|rfc3339"`
		ArrivalDate    string  `json:"arrivalDate" validate:"required|rfc3339"`
		ReturnDate     string  `json:"returnDate" validate:"rfc3339"`
		TripType       string  `json:"tripType" validate:"required|string"`
		Rate           float64 `json:"rate" validate:"required|float"`
		Discount       float32 `json:"discount" validate:"min:1|max:100"`
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
		From       string
		To         string
		Datetime   string
		Today      bool
		Scheduled  bool
		Completed  bool
		Passengers int
		TimeRange  string
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
