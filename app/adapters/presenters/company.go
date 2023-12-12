package presenters

import (
	"github.com/gofiber/fiber/v2"

	"github.com/SeyramWood/bookibus/ent"
	"github.com/SeyramWood/bookibus/ent/parcelimage"
	"github.com/SeyramWood/bookibus/ent/schema"
)

type (
	CompanyUserResponseData struct {
		ID         int    `json:"id"`
		LastName   string `json:"lastName"`
		OtherName  string `json:"otherName"`
		Phone      string `json:"phone"`
		OtherPhone string `json:"otherPhone"`
		Role       string `json:"role,omitempty"`
		CreatedAt  any    `json:"createdAt,omitempty"`
		UpdatedAt  any    `json:"updatedAt,omitempty"`
	}
	AdminCompanyResponseData struct {
		ID            int                   `json:"id"`
		Name          string                `json:"name"`
		Phone         string                `json:"phone"`
		Email         string                `json:"email"`
		Certificate   string                `json:"certificate"`
		BankAccount   *schema.BankAccount   `json:"bankAccount"`
		ContactPerson *schema.ContactPerson `json:"contactPerson"`
		Status        string                `json:"status"`
		CreatedAt     any                   `json:"createdAt,omitempty"`
		UpdatedAt     any                   `json:"updatedAt,omitempty"`
	}
	CompanyResponseData struct {
		ID        int    `json:"id"`
		Name      string `json:"name"`
		Phone     string `json:"phone"`
		Email     string `json:"email"`
		CreatedAt any    `json:"createdAt,omitempty"`
		UpdatedAt any    `json:"updatedAt,omitempty"`
	}
	TerminalResponseData struct {
		ID   int    `json:"id,omitempty"`
		Name string `json:"name"`
	}
	TripTerminalResponseData struct {
		From *TerminalResponseData `json:"from"`
		To   *TerminalResponseData `json:"to"`
	}
	VehicleImageResponseData struct {
		ID    int    `json:"id"`
		Image string `json:"image"`
	}
	VehicleResponseData struct {
		ID                 int                         `json:"id"`
		RegistrationNumber string                      `json:"registrationNumber"`
		Model              string                      `json:"model"`
		Seat               int                         `json:"seat"`
		Images             []*VehicleImageResponseData `json:"images"`
		CreatedAt          any                         `json:"createdAt,omitempty"`
		UpdatedAt          any                         `json:"updatedAt,omitempty"`
	}
	RouteStopResponseData struct {
		ID        int     `json:"id"`
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
	}
	RouteResponseData struct {
		ID            int                      `json:"id,omitempty"`
		From          string                   `json:"from"`
		To            string                   `json:"to"`
		FromLatitude  float64                  `json:"fromLatitude,omitempty"`
		FromLongitude float64                  `json:"fromLongitude,omitempty"`
		ToLatitude    float64                  `json:"toLatitude,omitempty"`
		ToLongitude   float64                  `json:"toLongitude,omitempty"`
		Stops         []*RouteStopResponseData `json:"stops,omitempty"`
		CreatedAt     any                      `json:"createdAt,omitempty"`
		UpdatedAt     any                      `json:"updatedAt,omitempty"`
	}
	TripInspectionStatus struct {
		Exterior           bool `json:"exterior"`
		Interior           bool `json:"interior"`
		EngineCompartment  bool `json:"engineCompartment"`
		BrakeAndSteering   bool `json:"brakeAndSteering"`
		EmergencyEquipment bool `json:"emergencyEquipment"`
		FuelAndFluid       bool `json:"fuelAndFluid"`
	}
	TripResponseData struct {
		ID               int                       `json:"id"`
		DepartureDate    any                       `json:"departureDate"`
		ArrivalDate      any                       `json:"arrivalDate"`
		ReturnDate       any                       `json:"returnDate"`
		Type             string                    `json:"type"`
		InspectionStatus *TripInspectionStatus     `json:"inspectionStatus"`
		Status           string                    `json:"status"`
		Scheduled        bool                      `json:"scheduled"`
		SeatLeft         int                       `json:"seatLeft"`
		Rate             float64                   `json:"rate"`
		Discount         float32                   `json:"discount"`
		Terminal         *TripTerminalResponseData `json:"terminal"`
		Vehicle          *VehicleResponseData      `json:"vehicle"`
		Route            *RouteResponseData        `json:"route"`
		Driver           *CompanyUserResponseData  `json:"driver"`
		Company          *CompanyResponseData      `json:"company"`
		Bookings         []*BookingResponseData    `json:"bookings,omitempty"`
		Delivery         []*ParcelResponseData     `json:"delivery,omitempty"`
		Incident         []*IncidentResponseData   `json:"incident,omitempty"`
		CreatedAt        any                       `json:"createdAt,omitempty"`
		UpdatedAt        any                       `json:"updatedAt,omitempty"`
	}
)

func CompanyUserResponse(data *ent.CompanyUser) *fiber.Map {
	return SuccessResponse(&CompanyUserResponseData{
		ID:         data.ID,
		LastName:   data.LastName,
		OtherName:  data.OtherName,
		Phone:      data.Phone,
		OtherPhone: data.OtherPhone,
		Role:       string(data.UserRole),
		CreatedAt:  data.CreatedAt,
		UpdatedAt:  data.UpdatedAt,
	})
}
func CompanyUsersResponse(data *PaginationResponse) *fiber.Map {
	var response []*CompanyUserResponseData
	for _, u := range data.Data.([]*ent.CompanyUser) {
		response = append(response, &CompanyUserResponseData{
			ID:         u.ID,
			LastName:   u.LastName,
			OtherName:  u.OtherName,
			Phone:      u.Phone,
			OtherPhone: u.OtherPhone,
			Role:       string(u.UserRole),
			CreatedAt:  u.CreatedAt,
			UpdatedAt:  u.UpdatedAt,
		})
	}
	data.Data = response
	return SuccessResponse(data)
}
func CompanyResponse(data *ent.Company) *fiber.Map {
	return SuccessResponse(&AdminCompanyResponseData{
		ID:            data.ID,
		Name:          data.Name,
		Phone:         data.Phone,
		Email:         data.Email,
		Certificate:   data.Certificate,
		BankAccount:   data.BankAccount,
		ContactPerson: data.ContactPerson,
		Status:        string(data.OnboardingStatus),
		CreatedAt:     data.CreatedAt,
		UpdatedAt:     data.UpdatedAt,
	})
}
func CompaniesResponse(data *PaginationResponse) *fiber.Map {
	var response []*AdminCompanyResponseData
	for _, c := range data.Data.([]*ent.Company) {
		response = append(response, &AdminCompanyResponseData{
			ID:            c.ID,
			Name:          c.Name,
			Phone:         c.Phone,
			Email:         c.Email,
			Certificate:   c.Certificate,
			BankAccount:   c.BankAccount,
			ContactPerson: c.ContactPerson,
			Status:        string(c.OnboardingStatus),
			CreatedAt:     c.CreatedAt,
			UpdatedAt:     c.UpdatedAt,
		})
	}
	data.Data = response
	return SuccessResponse(data)
}

func TerminalResponse(data *ent.Terminal) *fiber.Map {
	return SuccessResponse(&TerminalResponseData{
		ID:   data.ID,
		Name: data.Name,
	})
}
func TerminalsResponse(data *PaginationResponse) *fiber.Map {
	var response []*TerminalResponseData
	for _, c := range data.Data.([]*ent.Terminal) {
		response = append(response, &TerminalResponseData{
			ID:   c.ID,
			Name: c.Name,
		})
	}
	data.Data = response
	return SuccessResponse(data)
}

func VehicleImageResponse(data *ent.VehicleImage) *fiber.Map {
	return SuccessResponse(&VehicleImageResponseData{
		ID:    data.ID,
		Image: data.Image,
	})
}
func VehicleResponse(data *ent.Vehicle) *fiber.Map {
	return SuccessResponse(&VehicleResponseData{
		ID:                 data.ID,
		RegistrationNumber: data.RegistrationNumber,
		Model:              data.Model,
		Seat:               data.Seat,
		Images: func() []*VehicleImageResponseData {
			if images, err := data.Edges.ImagesOrErr(); err == nil {
				response := make([]*VehicleImageResponseData, 0, len(images))
				for _, image := range images {
					response = append(response, &VehicleImageResponseData{
						ID:    image.ID,
						Image: image.Image,
					})
				}
				return response
			}
			return nil
		}(),
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	})
}
func VehiclesResponse(data *PaginationResponse) *fiber.Map {
	var response []*VehicleResponseData
	for _, v := range data.Data.([]*ent.Vehicle) {
		response = append(response, &VehicleResponseData{
			ID:                 v.ID,
			RegistrationNumber: v.RegistrationNumber,
			Model:              v.Model,
			Seat:               v.Seat,
			Images: func() []*VehicleImageResponseData {
				if images, err := v.Edges.ImagesOrErr(); err == nil {
					response := make([]*VehicleImageResponseData, 0, len(images))
					for _, image := range images {
						response = append(response, &VehicleImageResponseData{
							ID:    image.ID,
							Image: image.Image,
						})
					}
					return response
				}
				return nil
			}(),
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		})
	}
	data.Data = response
	return SuccessResponse(data)
}

func RouteStopResponse(data *ent.RouteStop) *fiber.Map {
	return SuccessResponse(&RouteStopResponseData{
		ID:        data.ID,
		Latitude:  data.Latitude,
		Longitude: data.Longitude,
	})
}
func RouteResponse(data *ent.Route) *fiber.Map {
	return SuccessResponse(&RouteResponseData{
		ID:            data.ID,
		From:          data.FromLocation,
		To:            data.ToLocation,
		FromLatitude:  data.FromLatitude,
		FromLongitude: data.FromLongitude,
		ToLatitude:    data.ToLatitude,
		ToLongitude:   data.ToLongitude,
		Stops: func() []*RouteStopResponseData {
			if stops, err := data.Edges.StopsOrErr(); err == nil {
				response := make([]*RouteStopResponseData, 0, len(stops))
				for _, s := range stops {
					response = append(response, &RouteStopResponseData{ID: s.ID, Latitude: s.Latitude, Longitude: s.Longitude})
				}
				return response
			}
			return nil
		}(),
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	})
}
func RoutesResponse(data *PaginationResponse) *fiber.Map {
	var response []*RouteResponseData
	for _, r := range data.Data.([]*ent.Route) {
		response = append(response, &RouteResponseData{
			ID:            r.ID,
			From:          r.FromLocation,
			To:            r.ToLocation,
			FromLatitude:  r.FromLatitude,
			FromLongitude: r.FromLongitude,
			ToLatitude:    r.ToLatitude,
			ToLongitude:   r.ToLongitude,
			Stops: func() []*RouteStopResponseData {
				if stops, err := r.Edges.StopsOrErr(); err == nil {
					response := make([]*RouteStopResponseData, 0, len(stops))
					for _, r := range stops {
						response = append(response, &RouteStopResponseData{
							ID:        r.ID,
							Latitude:  r.Latitude,
							Longitude: r.Longitude,
						})
					}
					return response
				}
				return nil
			}(),
			CreatedAt: r.CreatedAt,
			UpdatedAt: r.UpdatedAt,
		})
	}
	data.Data = response
	return SuccessResponse(data)
}
func DistinctRoutesResponse(data *PaginationResponse) *fiber.Map {
	var response []*RouteResponseData
	for _, r := range data.Data.([]*ent.Route) {
		response = append(response, &RouteResponseData{
			From: r.FromLocation,
			To:   r.ToLocation,
		})
	}
	data.Data = response
	return SuccessResponse(data)
}

func TripResponse(data *ent.Trip) *fiber.Map {
	return SuccessResponse(&TripResponseData{
		ID:               data.ID,
		DepartureDate:    data.DepartureDate,
		ArrivalDate:      data.ArrivalDate,
		ReturnDate:       parseNullDatetime(data.ReturnDate),
		Type:             string(data.Type),
		InspectionStatus: &TripInspectionStatus{Exterior: data.ExteriorInspected, Interior: data.InteriorInspected, EngineCompartment: data.EngineCompartmentInspected, BrakeAndSteering: data.BrakeAndSteeringInspected, EmergencyEquipment: data.EmergencyEquipmentInspected, FuelAndFluid: data.FuelAndFluidsInspected},
		Status:           string(data.Status),
		Scheduled:        data.Scheduled,
		SeatLeft:         data.SeatLeft,
		Rate:             data.Rate,
		Discount:         data.Discount,
		Terminal: func() *TripTerminalResponseData {
			if from, err := data.Edges.FromTerminalOrErr(); err == nil {
				if to, err := data.Edges.ToTerminalOrErr(); err == nil {
					return &TripTerminalResponseData{From: &TerminalResponseData{ID: from.ID, Name: from.Name}, To: &TerminalResponseData{ID: to.ID, Name: to.Name}}
				}
			}
			return nil
		}(),
		Vehicle: func() *VehicleResponseData {
			if v, err := data.Edges.VehicleOrErr(); err == nil {
				return &VehicleResponseData{ID: v.ID, RegistrationNumber: v.RegistrationNumber, Model: v.Model, Seat: v.Seat, Images: func() []*VehicleImageResponseData {
					if images, err := v.Edges.ImagesOrErr(); err == nil && len(images) > 0 {
						response := make([]*VehicleImageResponseData, 0, len(images))
						for _, image := range images {
							response = append(response, &VehicleImageResponseData{ID: image.ID, Image: image.Image})
						}
						return response
					}
					return nil
				}()}
			}
			return nil
		}(),
		Route: func() *RouteResponseData {
			if r, err := data.Edges.RouteOrErr(); err == nil {
				return &RouteResponseData{ID: r.ID, From: r.FromLocation, To: r.ToLocation, FromLatitude: r.FromLatitude, FromLongitude: r.FromLongitude, ToLatitude: r.ToLatitude, ToLongitude: r.ToLongitude, Stops: func() []*RouteStopResponseData {
					if stops, err := r.Edges.StopsOrErr(); err == nil && len(stops) > 0 {
						response := make([]*RouteStopResponseData, 0, len(stops))
						for _, s := range stops {
							response = append(response, &RouteStopResponseData{ID: s.ID, Latitude: s.Latitude, Longitude: s.Longitude})
						}
						return response
					}
					return nil
				}()}
			}
			return nil
		}(),
		Driver: func() *CompanyUserResponseData {
			if d, err := data.Edges.DriverOrErr(); err == nil {
				return &CompanyUserResponseData{ID: d.ID, LastName: d.LastName, OtherName: d.OtherName, Phone: d.Phone, OtherPhone: d.OtherPhone}
			}
			return nil
		}(),
		Company: func() *CompanyResponseData {
			if c, err := data.Edges.CompanyOrErr(); err == nil {
				return &CompanyResponseData{ID: c.ID, Name: c.Name, Phone: c.Phone, Email: c.Email}
			}
			return nil
		}(),
		Bookings: func() []*BookingResponseData {
			if bookings, err := data.Edges.BookingsOrErr(); err == nil && len(bookings) > 0 {
				response := make([]*BookingResponseData, 0, len(bookings))
				for _, b := range bookings {
					response = append(response, &BookingResponseData{
						ID:              b.ID,
						TripID:          b.BookingNumber,
						SMSNotification: b.SmsNotification,
						Status:          string(b.Status),
						Passengers: func() []*BookingPassengerResponseData {
							if passengers, err := b.Edges.PassengersOrErr(); err == nil && len(passengers) > 0 {
								response := make([]*BookingPassengerResponseData, 0, len(passengers))
								for _, passenger := range passengers {
									response = append(response, &BookingPassengerResponseData{ID: passenger.ID, FullName: passenger.FullName, Amount: passenger.Amount, Maturity: string(passenger.Maturity), Gender: string(passenger.Gender)})
								}
								return response
							}
							return nil
						}(),
						Luggages: func() []*BookingLuggagesResponseData {
							if luggages, err := b.Edges.LuggagesOrErr(); err == nil && len(luggages) > 0 {
								response := make([]*BookingLuggagesResponseData, 0, len(luggages))
								for _, luggage := range luggages {
									response = append(response, &BookingLuggagesResponseData{ID: luggage.ID, Baggage: string(luggage.Baggage), Quantity: luggage.Quantity, Amount: luggage.Amount})
								}
								return response
							}
							return nil
						}(),
						Contact: func() *BookingContactResponseData {
							if c, err := b.Edges.ContactOrErr(); err == nil {
								return &BookingContactResponseData{ID: c.ID, FullName: c.FullName, Email: c.Email, Phone: c.Phone}
							}
							if c, err := b.Edges.CustomerOrErr(); err == nil {
								return &BookingContactResponseData{ID: c.ID, FullName: c.OtherName + " " + c.LastName, Email: c.Edges.Profile.Username, Phone: c.Phone}
							}
							return nil
						}(),
						Transaction: func() *TransactionResponseData {
							if t, err := b.Edges.TransactionOrErr(); err == nil && t != nil {
								return &TransactionResponseData{
									Reference:       t.Reference,
									TransactionType: string(t.Channel),
									Amount:          t.Amount,
									VAT:             t.Vat,
									TransactionFee:  t.TransactionFee,
									CancellationFee: t.CancellationFee,
									PaidAt:          parseNullDatetime(t.PaidAt),
									CanceledAt:      parseNullDatetime(t.CanceledAt),
								}
							}
							return nil
						}(),
					})
				}
				return response
			}
			return nil
		}(),
		Delivery: func() []*ParcelResponseData {
			if parcels, err := data.Edges.ParcelsOrErr(); err == nil && len(parcels) > 0 {
				response := make([]*ParcelResponseData, 0, len(parcels))
				for _, parcel := range parcels {
					response = append(response, &ParcelResponseData{
						ID:                parcel.ID,
						ParcelCode:        parcel.ParcelCode,
						SenderName:        parcel.SenderName,
						SenderPhone:       parcel.SenderPhone,
						SenderEmail:       parcel.SenderEmail,
						RecipientName:     parcel.RecipientName,
						RecipientPhone:    parcel.RecipientPhone,
						RecipientLocation: parcel.RecipientLocation,
						Weight:            parcel.Weight,
						Status:            string(parcel.Status),
						Type:              parcel.Type,
						ParcelImages: func() []*ParcelImageResponseData {
							if images, err := parcel.Edges.ImagesOrErr(); err == nil && len(images) > 0 {
								response := make([]*ParcelImageResponseData, 0, len(images))
								for _, image := range images {
									if image.Kind != parcelimage.KindParcel {
										continue
									}
									response = append(response, &ParcelImageResponseData{ID: image.ID, Image: image.Image})
								}
								return response
							}
							return nil
						}(),
						RecipientImages: func() []*ParcelImageResponseData {
							if images, err := parcel.Edges.ImagesOrErr(); err == nil && len(images) > 0 {
								response := make([]*ParcelImageResponseData, 0, len(images))
								for _, image := range images {
									if image.Kind != parcelimage.KindRecipient {
										continue
									}
									response = append(response, &ParcelImageResponseData{ID: image.ID, Image: image.Image})
								}
								return response
							}
							return nil
						}(),
						Transaction: func() *TransactionResponseData {
							if t, err := parcel.Edges.TransactionOrErr(); err == nil && t != nil {
								return &TransactionResponseData{
									Reference:       t.Reference,
									TransactionType: string(t.Channel),
									Amount:          t.Amount,
									TransactionFee:  t.TransactionFee,
									PaidAt:          parseNullDatetime(t.PaidAt),
								}
							}
							return nil
						}(),
					})
				}
				return response
			}
			return nil
		}(),
		Incident: func() []*IncidentResponseData {
			if incidents, err := data.Edges.IncidentsOrErr(); err != nil && len(incidents) > 0 {
				response := make([]*IncidentResponseData, 0, len(incidents))
				for _, incident := range incidents {
					response = append(response, &IncidentResponseData{ID: incident.ID, Time: parseNullDatetime(incident.Time), Location: incident.Location, Description: incident.Description, Audio: incident.Audio, Images: func() []*IncidentImageResponseData {
						if images, err := incident.Edges.ImagesOrErr(); err == nil && len(images) > 0 {
							response := make([]*IncidentImageResponseData, 0, len(images))
							for _, image := range images {
								response = append(response, &IncidentImageResponseData{ID: image.ID, Image: image.Image})
							}
							return response
						}
						return nil
					}()})
					return response
				}
			}
			return nil
		}(),
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	})
}
func TripsResponse(data *PaginationResponse) *fiber.Map {
	var response []*TripResponseData
	for _, t := range data.Data.([]*ent.Trip) {
		response = append(response, &TripResponseData{
			ID:               t.ID,
			DepartureDate:    t.DepartureDate,
			ArrivalDate:      t.ArrivalDate,
			ReturnDate:       parseNullDatetime(t.ReturnDate),
			Type:             string(t.Type),
			InspectionStatus: &TripInspectionStatus{Exterior: t.ExteriorInspected, Interior: t.InteriorInspected, EngineCompartment: t.EngineCompartmentInspected, BrakeAndSteering: t.BrakeAndSteeringInspected, EmergencyEquipment: t.EmergencyEquipmentInspected, FuelAndFluid: t.FuelAndFluidsInspected},
			Status:           string(t.Status),
			Scheduled:        t.Scheduled,
			SeatLeft:         t.SeatLeft,
			Rate:             t.Rate,
			Discount:         t.Discount,
			Terminal: func() *TripTerminalResponseData {
				if from, err := t.Edges.FromTerminalOrErr(); err == nil {
					if to, err := t.Edges.ToTerminalOrErr(); err == nil {
						return &TripTerminalResponseData{From: &TerminalResponseData{ID: from.ID, Name: from.Name}, To: &TerminalResponseData{ID: to.ID, Name: to.Name}}
					}
				}
				return nil
			}(),
			Vehicle: func() *VehicleResponseData {
				if v, err := t.Edges.VehicleOrErr(); err == nil {
					return &VehicleResponseData{ID: v.ID, RegistrationNumber: v.RegistrationNumber, Model: v.Model, Seat: v.Seat, Images: func() []*VehicleImageResponseData {
						if images, err := v.Edges.ImagesOrErr(); err == nil && len(images) > 0 {
							response := make([]*VehicleImageResponseData, 0, len(images))
							for _, image := range images {
								response = append(response, &VehicleImageResponseData{ID: image.ID, Image: image.Image})
							}
							return response
						}
						return nil
					}()}
				}
				return nil
			}(),
			Route: func() *RouteResponseData {
				if r, err := t.Edges.RouteOrErr(); err == nil {
					return &RouteResponseData{ID: r.ID, From: r.FromLocation, To: r.ToLocation, FromLatitude: r.FromLatitude, FromLongitude: r.FromLongitude, ToLatitude: r.ToLatitude, ToLongitude: r.ToLongitude, Stops: func() []*RouteStopResponseData {
						if stops, err := r.Edges.StopsOrErr(); err == nil && len(stops) > 0 {
							response := make([]*RouteStopResponseData, 0, len(stops))
							for _, s := range stops {
								response = append(response, &RouteStopResponseData{ID: s.ID, Latitude: s.Latitude, Longitude: s.Longitude})
							}
							return response
						}
						return nil
					}()}
				}
				return nil
			}(),
			Driver: func() *CompanyUserResponseData {
				if d, err := t.Edges.DriverOrErr(); err == nil {
					return &CompanyUserResponseData{ID: d.ID, LastName: d.LastName, OtherName: d.OtherName, Phone: d.Phone, OtherPhone: d.OtherPhone}
				}
				return nil
			}(),
			Company: func() *CompanyResponseData {
				if c, err := t.Edges.CompanyOrErr(); err == nil {
					return &CompanyResponseData{ID: c.ID, Name: c.Name, Phone: c.Phone, Email: c.Email}
				}
				return nil
			}(),
			Bookings: func() []*BookingResponseData {
				if bookings, err := t.Edges.BookingsOrErr(); err == nil && len(bookings) > 0 {
					response := make([]*BookingResponseData, 0, len(bookings))
					for _, b := range bookings {
						response = append(response, &BookingResponseData{
							ID:              b.ID,
							TripID:          b.BookingNumber,
							SMSNotification: b.SmsNotification,
							Status:          string(b.Status),
							Passengers: func() []*BookingPassengerResponseData {
								if passengers, err := b.Edges.PassengersOrErr(); err == nil && len(passengers) > 0 {
									response := make([]*BookingPassengerResponseData, 0, len(passengers))
									for _, passenger := range passengers {
										response = append(response, &BookingPassengerResponseData{ID: passenger.ID, FullName: passenger.FullName, Amount: passenger.Amount, Maturity: string(passenger.Maturity), Gender: string(passenger.Gender)})
									}
									return response
								}
								return nil
							}(),
							Luggages: func() []*BookingLuggagesResponseData {
								if luggages, err := b.Edges.LuggagesOrErr(); err == nil && len(luggages) > 0 {
									response := make([]*BookingLuggagesResponseData, 0, len(luggages))
									for _, luggage := range luggages {
										response = append(response, &BookingLuggagesResponseData{ID: luggage.ID, Baggage: string(luggage.Baggage), Quantity: luggage.Quantity, Amount: luggage.Amount})
									}
									return response
								}
								return nil
							}(),
							Contact: func() *BookingContactResponseData {
								if c, err := b.Edges.ContactOrErr(); err == nil {
									return &BookingContactResponseData{ID: c.ID, FullName: c.FullName, Email: c.Email, Phone: c.Phone}
								}
								if c, err := b.Edges.CustomerOrErr(); err == nil {
									return &BookingContactResponseData{ID: c.ID, FullName: c.OtherName + " " + c.LastName, Email: c.Edges.Profile.Username, Phone: c.Phone}
								}
								return nil
							}(),
							Transaction: func() *TransactionResponseData {
								if t, err := b.Edges.TransactionOrErr(); err == nil && t != nil {
									return &TransactionResponseData{
										Reference:       t.Reference,
										TransactionType: string(t.Channel),
										Amount:          t.Amount,
										VAT:             t.Vat,
										TransactionFee:  t.TransactionFee,
										CancellationFee: t.CancellationFee,
										PaidAt:          parseNullDatetime(t.PaidAt),
										CanceledAt:      parseNullDatetime(t.CanceledAt),
									}
								}
								return nil
							}(),
						})
					}
					return response
				}
				return nil
			}(),
			Delivery: func() []*ParcelResponseData {
				if parcels, err := t.Edges.ParcelsOrErr(); err == nil && len(parcels) > 0 {
					response := make([]*ParcelResponseData, 0, len(parcels))
					for _, parcel := range parcels {
						response = append(response, &ParcelResponseData{
							ID:                parcel.ID,
							ParcelCode:        parcel.ParcelCode,
							SenderName:        parcel.SenderName,
							SenderPhone:       parcel.SenderPhone,
							SenderEmail:       parcel.SenderEmail,
							RecipientName:     parcel.RecipientName,
							RecipientPhone:    parcel.RecipientPhone,
							RecipientLocation: parcel.RecipientLocation,
							Weight:            parcel.Weight,
							Status:            string(parcel.Status),
							Type:              parcel.Type,
							ParcelImages: func() []*ParcelImageResponseData {
								if images, err := parcel.Edges.ImagesOrErr(); err == nil && len(images) > 0 {
									response := make([]*ParcelImageResponseData, 0, len(images))
									for _, image := range images {
										if image.Kind != parcelimage.KindParcel {
											continue
										}
										response = append(response, &ParcelImageResponseData{ID: image.ID, Image: image.Image})
									}
									return response
								}
								return nil
							}(),
							RecipientImages: func() []*ParcelImageResponseData {
								if images, err := parcel.Edges.ImagesOrErr(); err == nil && len(images) > 0 {
									response := make([]*ParcelImageResponseData, 0, len(images))
									for _, image := range images {
										if image.Kind != parcelimage.KindRecipient {
											continue
										}
										response = append(response, &ParcelImageResponseData{ID: image.ID, Image: image.Image})
									}
									return response
								}
								return nil
							}(),
							Transaction: func() *TransactionResponseData {
								if t, err := parcel.Edges.TransactionOrErr(); err == nil && t != nil {
									return &TransactionResponseData{
										Reference:       t.Reference,
										TransactionType: string(t.Channel),
										Amount:          t.Amount,
										TransactionFee:  t.TransactionFee,
										PaidAt:          parseNullDatetime(t.PaidAt),
									}
								}
								return nil
							}(),
						})
					}
					return response
				}
				return nil
			}(),
			Incident: func() []*IncidentResponseData {
				if incidents, err := t.Edges.IncidentsOrErr(); err != nil && len(incidents) > 0 {
					response := make([]*IncidentResponseData, 0, len(incidents))
					for _, incident := range incidents {
						response = append(response, &IncidentResponseData{ID: incident.ID, Time: parseNullDatetime(incident.Time), Location: incident.Location, Description: incident.Description, Audio: incident.Audio, Images: func() []*IncidentImageResponseData {
							if images, err := incident.Edges.ImagesOrErr(); err == nil && len(images) > 0 {
								response := make([]*IncidentImageResponseData, 0, len(images))
								for _, image := range images {
									response = append(response, &IncidentImageResponseData{ID: image.ID, Image: image.Image})
								}
								return response
							}
							return nil
						}()})
						return response
					}
				}
				return nil
			}(),
			CreatedAt: t.CreatedAt,
			UpdatedAt: t.UpdatedAt,
		})
	}

	data.Data = response
	return SuccessResponse(data)
}
