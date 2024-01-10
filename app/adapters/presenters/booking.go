package presenters

import (
	"sync"

	"github.com/gofiber/fiber/v2"

	"github.com/SeyramWood/bookibus/ent"
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
		SMSNotification bool                            `json:"smsNotification"`
		Status          string                          `json:"status"`
		Passengers      []*BookingPassengerResponseData `json:"passengers"`
		Luggages        []*BookingLuggagesResponseData  `json:"luggages"`
		Contact         *BookingContactResponseData     `json:"contact"`
		Transaction     *TransactionResponseData        `json:"transaction"`
		Trip            *TripResponseData               `json:"trip"`
		CreatedAt       any                             `json:"createdAt"`
		UpdatedAt       any                             `json:"updatedAt"`
	}
	BookingTicketResponseData struct {
		TripID      string                      `json:"tripId"`
		Passengers  int                         `json:"passengers"`
		Luggages    int                         `json:"luggages"`
		Contact     *BookingContactResponseData `json:"contact"`
		Transaction *TransactionResponseData    `json:"transaction"`
		Trip        *TripResponseData           `json:"trip"`
	}

	BookingPassengerDetailResponseData struct {
		ID            int                             `json:"id"`
		BookingNumber string                          `json:"bookingNumber"`
		Passengers    []*BookingPassengerResponseData `json:"passengers"`
		Luggages      []*BookingLuggagesResponseData  `json:"luggages"`
		Contact       *BookingContactResponseData     `json:"contact"`
		Transaction   *TransactionResponseData        `json:"transaction"`
		CreatedAt     any                             `json:"createdAt"`
	}
)

func BookingTicketResponse(data *ent.Booking) *BookingTicketResponseData {
	return &BookingTicketResponseData{
		TripID: data.BookingNumber,
		Passengers: func() int {
			if passengers, err := data.Edges.PassengersOrErr(); err == nil && len(passengers) > 0 {
				return len(passengers)
			}
			return 0
		}(),
		Luggages: func() int {
			if luggages, err := data.Edges.LuggagesOrErr(); err == nil && len(luggages) > 0 {
				qty := 0
				for _, luggage := range luggages {
					qty += luggage.Quantity
				}
				return qty
			}
			return 0
		}(),
		Contact: func() *BookingContactResponseData {
			if c, err := data.Edges.ContactOrErr(); err == nil {
				return &BookingContactResponseData{ID: c.ID, FullName: c.FullName, Email: c.Email, Phone: c.Phone}
			}
			if c, err := data.Edges.CustomerOrErr(); err == nil {
				return &BookingContactResponseData{ID: c.ID, FullName: c.OtherName + " " + c.LastName, Email: c.Edges.Profile.Username, Phone: c.Phone}
			}
			return nil
		}(),
		Transaction: func() *TransactionResponseData {
			if t, err := data.Edges.TransactionOrErr(); err == nil && t != nil {
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
		Trip: func() *TripResponseData {
			if t, err := data.Edges.TripOrErr(); err == nil {
				return &TripResponseData{
					ID:               t.ID,
					DepartureDate:    t.DepartureDate.Local().Format("Jan 02, 2006 15:04PM"),
					ArrivalDate:      t.ArrivalDate,
					ReturnDate:       parseNullDatetime(t.ReturnDate),
					Type:             string(t.Type),
					InspectionStatus: &TripInspectionStatus{Exterior: t.ExteriorInspected, Interior: t.InteriorInspected, EngineCompartment: t.EngineCompartmentInspected, BrakeAndSteering: t.BrakeAndSteeringInspected, EmergencyEquipment: t.EmergencyEquipmentInspected, FuelAndFluid: t.FuelAndFluidsInspected},
					Status:           string(t.Status),
					Scheduled:        t.Scheduled,
					SeatLeft:         t.SeatLeft,
					Rate:             t.Rate,
					Discount:         t.Discount,
					Terminal: &TripTerminalResponseData{
						From: func() *TerminalResponseData {
							if tr, err := t.Edges.FromTerminalOrErr(); err == nil && tr != nil {
								return &TerminalResponseData{
									ID:        tr.ID,
									Address:   tr.Address,
									Latitude:  tr.Latitude,
									Longitude: tr.Longitude,
								}
							}
							return nil
						}(),
						To: func() *TerminalResponseData {
							if tr, err := t.Edges.ToTerminalOrErr(); err == nil && tr != nil {
								return &TerminalResponseData{
									ID:        tr.ID,
									Address:   tr.Address,
									Latitude:  tr.Latitude,
									Longitude: tr.Longitude,
								}
							}
							return nil
						}(),
					},
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
							return &RouteResponseData{
								ID:   r.ID,
								From: r.FromLocation,
								To:   r.ToLocation,
								Stops: func() []*RouteStopResponseData {
									if stops, err := t.Edges.StopsOrErr(); err == nil && len(stops) > 0 {
										response := make([]*RouteStopResponseData, 0, len(stops))
										for _, s := range stops {
											response = append(response, &RouteStopResponseData{
												ID:        s.ID,
												Address:   s.Address,
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
					CreatedAt: nil,
					UpdatedAt: data,
				}
			}
			return nil
		}(),
	}
}

func BookingResponse(data *ent.Booking) *fiber.Map {
	return SuccessResponse(&BookingResponseData{
		ID:              data.ID,
		TripID:          data.BookingNumber,
		SMSNotification: data.SmsNotification,
		Status:          string(data.Status),
		Passengers: func() []*BookingPassengerResponseData {
			if passengers, err := data.Edges.PassengersOrErr(); err == nil && len(passengers) > 0 {
				response := make([]*BookingPassengerResponseData, 0, len(passengers))
				for _, passenger := range passengers {
					response = append(response, &BookingPassengerResponseData{ID: passenger.ID, FullName: passenger.FullName, Amount: passenger.Amount, Maturity: string(passenger.Maturity), Gender: string(passenger.Gender)})
				}
				return response
			}
			return nil
		}(),
		Luggages: func() []*BookingLuggagesResponseData {
			if luggages, err := data.Edges.LuggagesOrErr(); err == nil && len(luggages) > 0 {
				response := make([]*BookingLuggagesResponseData, 0, len(luggages))
				for _, luggage := range luggages {
					response = append(response, &BookingLuggagesResponseData{ID: luggage.ID, Baggage: string(luggage.Baggage), Quantity: luggage.Quantity, Amount: luggage.Amount})
				}
				return response
			}
			return nil
		}(),
		Contact: func() *BookingContactResponseData {
			if c, err := data.Edges.ContactOrErr(); err == nil {
				return &BookingContactResponseData{ID: c.ID, FullName: c.FullName, Email: c.Email, Phone: c.Phone}
			}
			if c, err := data.Edges.CustomerOrErr(); err == nil {
				return &BookingContactResponseData{ID: c.ID, FullName: c.OtherName + " " + c.LastName, Email: c.Edges.Profile.Username, Phone: c.Phone}
			}
			return nil
		}(),
		Transaction: func() *TransactionResponseData {
			if t, err := data.Edges.TransactionOrErr(); err == nil && t != nil {
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
		Trip: func() *TripResponseData {
			if t, err := data.Edges.TripOrErr(); err == nil {
				return &TripResponseData{
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
					Terminal: &TripTerminalResponseData{
						From: func() *TerminalResponseData {
							if tr, err := t.Edges.FromTerminalOrErr(); err == nil && tr != nil {
								return &TerminalResponseData{
									ID:        tr.ID,
									Address:   tr.Address,
									Latitude:  tr.Latitude,
									Longitude: tr.Longitude,
								}
							}
							return nil
						}(),
						To: func() *TerminalResponseData {
							if tr, err := t.Edges.ToTerminalOrErr(); err == nil && tr != nil {
								return &TerminalResponseData{
									ID:        tr.ID,
									Address:   tr.Address,
									Latitude:  tr.Latitude,
									Longitude: tr.Longitude,
								}
							}
							return nil
						}(),
					},
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
							return &RouteResponseData{
								ID:   r.ID,
								From: r.FromLocation,
								To:   r.ToLocation,
								Stops: func() []*RouteStopResponseData {
									if stops, err := t.Edges.StopsOrErr(); err == nil && len(stops) > 0 {
										response := make([]*RouteStopResponseData, 0, len(stops))
										for _, s := range stops {
											response = append(response, &RouteStopResponseData{
												ID:        s.ID,
												Address:   s.Address,
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
				}
			}
			return nil
		}(),
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	})
}
func BookingPassengerDetailsResponse(data []*ent.Booking) *fiber.Map {
	wg := &sync.WaitGroup{}
	var response []*BookingPassengerDetailResponseData
	wg.Add(1)
	go func(data []*ent.Booking) {
		defer wg.Done()
		var res []*BookingPassengerDetailResponseData
		for _, b := range data {
			res = append(res, &BookingPassengerDetailResponseData{
				ID:            b.ID,
				BookingNumber: b.BookingNumber,
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
				CreatedAt: b.CreatedAt,
			})
		}
		response = res
	}(data)
	wg.Wait()
	return SuccessResponse(response)
}
func BookingsResponse(data *PaginationResponse) *fiber.Map {
	var response []*BookingResponseData
	for _, b := range data.Data.([]*ent.Booking) {
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
			Trip: func() *TripResponseData {
				if t, err := b.Edges.TripOrErr(); err == nil {
					return &TripResponseData{
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
						Terminal: &TripTerminalResponseData{
							From: func() *TerminalResponseData {
								if tr, err := t.Edges.FromTerminalOrErr(); err == nil && tr != nil {
									return &TerminalResponseData{
										ID:        tr.ID,
										Address:   tr.Address,
										Latitude:  tr.Latitude,
										Longitude: tr.Longitude,
									}
								}
								return nil
							}(),
							To: func() *TerminalResponseData {
								if tr, err := t.Edges.ToTerminalOrErr(); err == nil && tr != nil {
									return &TerminalResponseData{
										ID:        tr.ID,
										Address:   tr.Address,
										Latitude:  tr.Latitude,
										Longitude: tr.Longitude,
									}
								}
								return nil
							}(),
						},
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
								return &RouteResponseData{
									ID:   r.ID,
									From: r.FromLocation,
									To:   r.ToLocation,
									Stops: func() []*RouteStopResponseData {
										if stops, err := t.Edges.StopsOrErr(); err == nil && len(stops) > 0 {
											response := make([]*RouteStopResponseData, 0, len(stops))
											for _, s := range stops {
												response = append(response, &RouteStopResponseData{
													ID:        s.ID,
													Address:   s.Address,
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
