package presenters

import (
	"github.com/gofiber/fiber/v2"

	"github.com/SeyramWood/bookibus/ent"
)

type (
	IncidentResponseData struct {
		ID          int                          `json:"id"`
		Time        any                          `json:"time"`
		Location    string                       `json:"location"`
		Description string                       `json:"description"`
		Audio       string                       `json:"audio"`
		Type        string                       `json:"type"`
		Images      []*IncidentImageResponseData `json:"images"`
		Trip        *TripResponseData            `json:"trip"`
		Status      string                       `json:"status"`
		CreatedAt   any                          `json:"createdAt,omitempty"`
		UpdatedAt   any                          `json:"updatedAt,omitempty"`
	}
	IncidentImageResponseData struct {
		ID    int    `json:"id"`
		Image string `json:"image"`
	}
)

func IncidentResponse(data *ent.Incident) *fiber.Map {
	return SuccessResponse(&IncidentResponseData{
		ID:          data.ID,
		Time:        parseNullDatetime(data.Time),
		Location:    data.Location,
		Description: data.Description,
		Audio:       data.Audio,
		Type:        data.Type,
		Images: func() []*IncidentImageResponseData {
			if images, err := data.Edges.ImagesOrErr(); err == nil && len(images) > 0 {
				response := make([]*IncidentImageResponseData, 0, len(images))
				for _, image := range images {
					response = append(response, &IncidentImageResponseData{ID: image.ID, Image: image.Image})
				}
				return response
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
					Terminal:         &TripTerminalResponseData{},
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
		Status:    string(data.Status),
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	})
}
func IncidentsResponse(data *PaginationResponse) *fiber.Map {
	var response []*IncidentResponseData
	for _, i := range data.Data.([]*ent.Incident) {
		response = append(response, &IncidentResponseData{
			ID:          i.ID,
			Time:        parseNullDatetime(i.Time),
			Location:    i.Location,
			Description: i.Description,
			Audio:       i.Audio,
			Type:        i.Type,
			Images: func() []*IncidentImageResponseData {
				if images, err := i.Edges.ImagesOrErr(); err == nil && len(images) > 0 {
					response := make([]*IncidentImageResponseData, 0, len(images))
					for _, image := range images {
						response = append(response, &IncidentImageResponseData{ID: image.ID, Image: image.Image})
					}
					return response
				}
				return nil
			}(),
			Trip: func() *TripResponseData {
				if t, err := i.Edges.TripOrErr(); err == nil {
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
						Terminal:         &TripTerminalResponseData{},
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
			Status:    string(i.Status),
			CreatedAt: i.CreatedAt,
			UpdatedAt: i.UpdatedAt,
		})
	}
	data.Data = response
	return SuccessResponse(data)
}
func IncidentImageResponse(data *ent.IncidentImage) *fiber.Map {
	return SuccessResponse(&ParcelImageResponseData{
		ID:    data.ID,
		Image: data.Image,
	})
}
