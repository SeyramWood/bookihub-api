package presenters

import (
	"github.com/gofiber/fiber/v2"

	"github.com/SeyramWood/ent"
)

type (
	BookingPassengerResponseData struct {
		ID       int     `json:"id"`
		FullName string  `json:"fullName"`
		Amount   float64 `json:"amount"`
		Maturity string  `json:"maturity"`
		Gender   string  `json:"gender"`
	}
	BookingLuggagesResponseData struct {
		ID       int     `json:"id"`
		Baggage  string  `json:"baggage"`
		Quantity int     `json:"quantity"`
		Amount   float64 `json:"amount"`
	}
	BookingContactResponseData struct {
		ID       int    `json:"id"`
		FullName string `json:"fullName"`
		Email    string `json:"email"`
		Phone    string `json:"phone"`
	}
	BookingResponseData struct {
		ID              int                             `json:"id"`
		TripID          string                          `json:"tripId"`
		BoardingPoint   string                          `json:"boardingPoint"`
		VAT             float64                         `json:"vat"`
		SMSFee          float64                         `json:"smsFee"`
		Amount          float64                         `json:"amount"`
		RefundAmount    float64                         `json:"refundAmount"`
		TransactionType string                          `json:"transType"`
		SMSNotification bool                            `json:"smsNotification"`
		Status          string                          `json:"status"`
		RefundedAt      any                             `json:"refundedAt"`
		Passengers      []*BookingPassengerResponseData `json:"passengers"`
		Luggages        []*BookingLuggagesResponseData  `json:"luggages"`
		Contact         *BookingContactResponseData     `json:"contact"`
		Trip            *TripResponseData               `json:"trip"`
		CreatedAt       any                             `json:"createdAt"`
		UpdatedAt       any                             `json:"updatedAt"`
	}
)

func BookingResponse(data *ent.Booking) *fiber.Map {
	return SuccessResponse(&BookingResponseData{
		ID:     data.ID,
		TripID: data.BookingNumber,
		BoardingPoint: func() string {
			if t, err := data.Edges.TripOrErr(); err == nil {
				for _, bp := range t.BoardingPoints {
					if bp.ID == data.BoardingPoint {
						return bp.Location
					}
				}
			}
			return ""
		}(),
		VAT:             data.Vat,
		SMSFee:          data.SmsFee,
		Amount:          data.Amount,
		RefundAmount:    data.RefundAmount,
		TransactionType: string(data.TansType),
		SMSNotification: data.SmsNotification,
		Status:          string(data.Status),
		RefundedAt:      parseNullDatetime(data.RefundAt),
		Passengers: func() []*BookingPassengerResponseData {
			if passengers, err := data.Edges.PassengersOrErr(); err == nil {
				if len(passengers) == 0 {
					return nil
				}
				response := make([]*BookingPassengerResponseData, 0, len(passengers))
				for _, passenger := range passengers {
					response = append(response, &BookingPassengerResponseData{
						ID:       passenger.ID,
						FullName: passenger.FullName,
						Amount:   passenger.Amount,
						Maturity: string(passenger.Maturity),
						Gender:   string(passenger.Gender),
					})
				}
				return response
			}
			return nil
		}(),
		Luggages: func() []*BookingLuggagesResponseData {
			if luggages, err := data.Edges.LuggagesOrErr(); err == nil {
				if len(luggages) == 0 {
					return nil
				}
				response := make([]*BookingLuggagesResponseData, 0, len(luggages))
				for _, luggage := range luggages {
					response = append(response, &BookingLuggagesResponseData{
						ID:       luggage.ID,
						Baggage:  string(luggage.Baggage),
						Quantity: luggage.Quantity,
						Amount:   luggage.Amount,
					})
				}
				return response
			}
			return nil
		}(),
		Contact: func() *BookingContactResponseData {
			if c, err := data.Edges.ContactOrErr(); err == nil {
				return &BookingContactResponseData{
					ID:       c.ID,
					FullName: c.FullName,
					Email:    c.Email,
					Phone:    c.Phone,
				}
			}
			if c, err := data.Edges.CustomerOrErr(); err == nil {
				return &BookingContactResponseData{
					ID:       c.ID,
					FullName: c.OtherName + " " + c.LastName,
					Email:    c.Edges.Profile.Username,
					Phone:    c.Phone,
				}
			}
			return nil
		}(),
		Trip: func() *TripResponseData {
			if t, err := data.Edges.TripOrErr(); err == nil {
				return &TripResponseData{
					ID:            t.ID,
					DepartureDate: t.DepartureDate,
					ArrivalDate:   t.ArrivalDate,
					ReturnDate:    parseNullDatetime(t.ReturnDate),
					Type:          string(t.Type),
					InspectionStatus: &TripInspectionStatus{
						Exterior:           t.ExteriorInspected,
						Interior:           t.InteriorInspected,
						EngineCompartment:  t.EngineCompartmentInspected,
						BrakeAndSteering:   t.BrakeAndSteeringInspected,
						EmergencyEquipment: t.EmergencyEquipmentInspected,
						FuelAndFluid:       t.FuelAndFluidsInspected,
					},
					Status:        string(t.Status),
					Scheduled:     t.Scheduled,
					SeatLeft:      t.SeatLeft,
					BoardingPoint: t.BoardingPoints,
					Vehicle: func() *VehicleResponseData {
						if v, err := t.Edges.VehicleOrErr(); err == nil {
							return &VehicleResponseData{
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
							}
						}
						return nil
					}(),
					Route: func() *RouteResponseData {
						if r, err := t.Edges.RouteOrErr(); err == nil {
							return &RouteResponseData{
								ID:            r.ID,
								From:          r.FromLocation,
								To:            r.ToLocation,
								FromLatitude:  r.FromLatitude,
								FromLongitude: r.FromLongitude,
								ToLatitude:    r.ToLatitude,
								ToLongitude:   r.ToLongitude,
								Rate:          r.Rate,
								Discount:      r.Discount,
								Stops: func() []*RouteStopResponseData {
									if stops, err := r.Edges.StopsOrErr(); err == nil {
										response := make([]*RouteStopResponseData, 0, len(stops))
										for _, s := range stops {
											response = append(response, &RouteStopResponseData{
												ID:        s.ID,
												Latitude:  s.Latitude,
												Longitude: s.Longitude,
											})
										}
										return response
									}
									return nil
								}(),
							}
						}
						return nil
					}(),
					Driver: func() *CompanyUserResponseData {
						if d, err := t.Edges.DriverOrErr(); err == nil {
							return &CompanyUserResponseData{
								ID:         d.ID,
								LastName:   d.LastName,
								OtherName:  d.OtherName,
								Phone:      d.Phone,
								OtherPhone: d.OtherPhone,
							}
						}
						return nil
					}(),
					Company: func() *CompanyResponseData {
						if c, err := t.Edges.CompanyOrErr(); err == nil {
							return &CompanyResponseData{
								ID:    c.ID,
								Name:  c.Name,
								Phone: c.Phone,
								Email: c.Email,
							}
						}
						return nil
					}(),
				}
			}
			return nil
		}(),
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	})
}
func BookingsResponse(data *PaginationResponse) *fiber.Map {
	var response []*BookingResponseData
	for _, b := range data.Data.([]*ent.Booking) {
		response = append(response, &BookingResponseData{
			ID:     b.ID,
			TripID: b.BookingNumber,
			BoardingPoint: func() string {
				if t, err := b.Edges.TripOrErr(); err == nil {
					for _, bp := range t.BoardingPoints {
						if bp.ID == b.BoardingPoint {
							return bp.Location
						}
					}
				}
				return ""
			}(),
			VAT:             b.Vat,
			SMSFee:          b.SmsFee,
			Amount:          b.Amount,
			RefundAmount:    b.RefundAmount,
			TransactionType: string(b.TansType),
			SMSNotification: b.SmsNotification,
			Status:          string(b.Status),
			RefundedAt:      parseNullDatetime(b.RefundAt),
			Passengers: func() []*BookingPassengerResponseData {
				if passengers, err := b.Edges.PassengersOrErr(); err == nil {
					if len(passengers) == 0 {
						return nil
					}
					response := make([]*BookingPassengerResponseData, 0, len(passengers))
					for _, passenger := range passengers {
						response = append(response, &BookingPassengerResponseData{
							ID:       passenger.ID,
							FullName: passenger.FullName,
							Amount:   passenger.Amount,
							Maturity: string(passenger.Maturity),
							Gender:   string(passenger.Gender),
						})
					}
					return response
				}
				return nil
			}(),
			Luggages: func() []*BookingLuggagesResponseData {
				if luggages, err := b.Edges.LuggagesOrErr(); err == nil {
					if len(luggages) == 0 {
						return nil
					}
					response := make([]*BookingLuggagesResponseData, 0, len(luggages))
					for _, luggage := range luggages {
						response = append(response, &BookingLuggagesResponseData{
							ID:       luggage.ID,
							Baggage:  string(luggage.Baggage),
							Quantity: luggage.Quantity,
							Amount:   luggage.Amount,
						})
					}
					return response
				}
				return nil
			}(),
			Contact: func() *BookingContactResponseData {
				if c, err := b.Edges.ContactOrErr(); err == nil {
					return &BookingContactResponseData{
						ID:       c.ID,
						FullName: c.FullName,
						Email:    c.Email,
						Phone:    c.Phone,
					}
				}
				if c, err := b.Edges.CustomerOrErr(); err == nil {
					return &BookingContactResponseData{
						ID:       c.ID,
						FullName: c.OtherName + " " + c.LastName,
						Email:    c.Edges.Profile.Username,
						Phone:    c.Phone,
					}
				}
				return nil
			}(),
			Trip: func() *TripResponseData {
				if t, err := b.Edges.TripOrErr(); err == nil {
					return &TripResponseData{
						ID:            t.ID,
						DepartureDate: t.DepartureDate,
						ArrivalDate:   t.ArrivalDate,
						ReturnDate:    parseNullDatetime(t.ReturnDate),
						Type:          string(t.Type),
						InspectionStatus: &TripInspectionStatus{
							Exterior:           t.ExteriorInspected,
							Interior:           t.InteriorInspected,
							EngineCompartment:  t.EngineCompartmentInspected,
							BrakeAndSteering:   t.BrakeAndSteeringInspected,
							EmergencyEquipment: t.EmergencyEquipmentInspected,
							FuelAndFluid:       t.FuelAndFluidsInspected,
						},
						Status:        string(t.Status),
						Scheduled:     t.Scheduled,
						SeatLeft:      t.SeatLeft,
						BoardingPoint: t.BoardingPoints,
						Vehicle: func() *VehicleResponseData {
							if v, err := t.Edges.VehicleOrErr(); err == nil {
								return &VehicleResponseData{
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
								}
							}
							return nil
						}(),
						Route: func() *RouteResponseData {
							if r, err := t.Edges.RouteOrErr(); err == nil {
								return &RouteResponseData{
									ID:            r.ID,
									From:          r.FromLocation,
									To:            r.ToLocation,
									FromLatitude:  r.FromLatitude,
									FromLongitude: r.FromLongitude,
									ToLatitude:    r.ToLatitude,
									ToLongitude:   r.ToLongitude,
									Rate:          r.Rate,
									Discount:      r.Discount,
									Stops: func() []*RouteStopResponseData {
										if stops, err := r.Edges.StopsOrErr(); err == nil {
											response := make([]*RouteStopResponseData, 0, len(stops))
											for _, s := range stops {
												response = append(response, &RouteStopResponseData{
													ID:        s.ID,
													Latitude:  s.Latitude,
													Longitude: s.Longitude,
												})
											}
											return response
										}
										return nil
									}(),
								}
							}
							return nil
						}(),
						Driver: func() *CompanyUserResponseData {
							if d, err := t.Edges.DriverOrErr(); err == nil {
								return &CompanyUserResponseData{
									ID:         d.ID,
									LastName:   d.LastName,
									OtherName:  d.OtherName,
									Phone:      d.Phone,
									OtherPhone: d.OtherPhone,
								}
							}
							return nil
						}(),
						Company: func() *CompanyResponseData {
							if c, err := t.Edges.CompanyOrErr(); err == nil {
								return &CompanyResponseData{
									ID:    c.ID,
									Name:  c.Name,
									Phone: c.Phone,
									Email: c.Email,
								}
							}
							return nil
						}(),
					}
				}
				return nil
			}(),
			CreatedAt: b.CreatedAt,
			UpdatedAt: b.UpdatedAt,
		})
	}
	data.Data = response
	return SuccessResponse(data)
}
