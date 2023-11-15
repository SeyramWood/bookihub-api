package presenters

import (
	"github.com/gofiber/fiber/v2"

	"github.com/SeyramWood/bookibus/ent"
	"github.com/SeyramWood/bookibus/ent/parcelimage"
)

type (
	ParcelResponseData struct {
		ID                int                        `json:"id"`
		ParcelCode        string                     `json:"packageCode"`
		SenderName        string                     `json:"senderName"`
		SenderPhone       string                     `json:"senderPhone"`
		SenderEmail       string                     `json:"senderEmail"`
		RecipientName     string                     `json:"recipientName"`
		RecipientPhone    string                     `json:"recipientPhone"`
		RecipientLocation string                     `json:"recipientLocation"`
		Weight            float32                    `json:"weight"`
		Amount            float64                    `json:"amount"`
		TransType         string                     `json:"transType"`
		Status            string                     `json:"status"`
		Type              string                     `json:"type"`
		ParcelImages      []*ParcelImageResponseData `json:"packageImages"`
		RecipientImages   []*ParcelImageResponseData `json:"recipientImages"`
		Trip              *TripResponseData          `json:"trip"`
		CreatedAt         any                        `json:"createdAt,omitempty"`
		UpdatedAt         any                        `json:"updatedAt,omitempty"`
	}
	ParcelImageResponseData struct {
		ID    int    `json:"id"`
		Image string `json:"image"`
	}
)

func ParcelResponse(data *ent.Parcel) *fiber.Map {
	return SuccessResponse(&ParcelResponseData{
		ID:                data.ID,
		ParcelCode:        data.ParcelCode,
		SenderName:        data.SenderName,
		SenderPhone:       data.SenderPhone,
		SenderEmail:       data.SenderEmail,
		RecipientName:     data.RecipientName,
		RecipientPhone:    data.RecipientPhone,
		RecipientLocation: data.RecipientLocation,
		Weight:            data.Weight,
		Amount:            data.Amount,
		TransType:         string(data.TansType),
		Status:            string(data.Status),
		Type:              data.Type,
		ParcelImages: func() []*ParcelImageResponseData {
			if images, err := data.Edges.ImagesOrErr(); err == nil && len(images) > 0 {
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
			if images, err := data.Edges.ImagesOrErr(); err == nil && len(images) > 0 {
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
		Trip: func() *TripResponseData {
			if t, err := data.Edges.TripOrErr(); err == nil {
				return &TripResponseData{ID: t.ID, DepartureDate: t.DepartureDate, ArrivalDate: t.ArrivalDate, ReturnDate: parseNullDatetime(t.ReturnDate), Type: string(t.Type), InspectionStatus: &TripInspectionStatus{Exterior: t.ExteriorInspected, Interior: t.InteriorInspected, EngineCompartment: t.EngineCompartmentInspected, BrakeAndSteering: t.BrakeAndSteeringInspected, EmergencyEquipment: t.EmergencyEquipmentInspected, FuelAndFluid: t.FuelAndFluidsInspected}, Status: string(t.Status), Scheduled: t.Scheduled, SeatLeft: t.SeatLeft, Vehicle: func() *VehicleResponseData {
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
				}(), Route: func() *RouteResponseData {
					if r, err := t.Edges.RouteOrErr(); err == nil {
						return &RouteResponseData{ID: r.ID, From: r.FromLocation, To: r.ToLocation, FromLatitude: r.FromLatitude, FromLongitude: r.FromLongitude, ToLatitude: r.ToLatitude, ToLongitude: r.ToLongitude, Rate: r.Rate, Discount: r.Discount, Stops: func() []*RouteStopResponseData {
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
				}(), Driver: func() *CompanyUserResponseData {
					if d, err := t.Edges.DriverOrErr(); err == nil {
						return &CompanyUserResponseData{ID: d.ID, LastName: d.LastName, OtherName: d.OtherName, Phone: d.Phone, OtherPhone: d.OtherPhone}
					}
					return nil
				}(), Company: func() *CompanyResponseData {
					if c, err := t.Edges.CompanyOrErr(); err == nil {
						return &CompanyResponseData{ID: c.ID, Name: c.Name, Phone: c.Phone, Email: c.Email}
					}
					return nil
				}()}
			}
			return nil
		}(),
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	})
}
func ParcelsResponse(data *PaginationResponse) *fiber.Map {
	var response []*ParcelResponseData
	for _, p := range data.Data.([]*ent.Parcel) {
		response = append(response, &ParcelResponseData{
			ID:                p.ID,
			ParcelCode:        p.ParcelCode,
			SenderName:        p.SenderName,
			SenderPhone:       p.SenderPhone,
			SenderEmail:       p.SenderEmail,
			RecipientName:     p.RecipientName,
			RecipientPhone:    p.RecipientPhone,
			RecipientLocation: p.RecipientLocation,
			Weight:            p.Weight,
			Amount:            p.Amount,
			TransType:         string(p.TansType),
			Status:            string(p.Status),
			Type:              p.Type,
			ParcelImages: func() []*ParcelImageResponseData {
				if images, err := p.Edges.ImagesOrErr(); err == nil && len(images) > 0 {
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
				if images, err := p.Edges.ImagesOrErr(); err == nil && len(images) > 0 {
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
			Trip: func() *TripResponseData {
				if t, err := p.Edges.TripOrErr(); err == nil {
					return &TripResponseData{ID: t.ID, DepartureDate: t.DepartureDate, ArrivalDate: t.ArrivalDate, ReturnDate: parseNullDatetime(t.ReturnDate), Type: string(t.Type), InspectionStatus: &TripInspectionStatus{Exterior: t.ExteriorInspected, Interior: t.InteriorInspected, EngineCompartment: t.EngineCompartmentInspected, BrakeAndSteering: t.BrakeAndSteeringInspected, EmergencyEquipment: t.EmergencyEquipmentInspected, FuelAndFluid: t.FuelAndFluidsInspected}, Status: string(t.Status), Scheduled: t.Scheduled, SeatLeft: t.SeatLeft, Vehicle: func() *VehicleResponseData {
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
					}(), Route: func() *RouteResponseData {
						if r, err := t.Edges.RouteOrErr(); err == nil {
							return &RouteResponseData{ID: r.ID, From: r.FromLocation, To: r.ToLocation, FromLatitude: r.FromLatitude, FromLongitude: r.FromLongitude, ToLatitude: r.ToLatitude, ToLongitude: r.ToLongitude, Rate: r.Rate, Discount: r.Discount, Stops: func() []*RouteStopResponseData {
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
					}(), Driver: func() *CompanyUserResponseData {
						if d, err := t.Edges.DriverOrErr(); err == nil {
							return &CompanyUserResponseData{ID: d.ID, LastName: d.LastName, OtherName: d.OtherName, Phone: d.Phone, OtherPhone: d.OtherPhone}
						}
						return nil
					}(), Company: func() *CompanyResponseData {
						if c, err := t.Edges.CompanyOrErr(); err == nil {
							return &CompanyResponseData{ID: c.ID, Name: c.Name, Phone: c.Phone, Email: c.Email}
						}
						return nil
					}()}
				}
				return nil
			}(),
			CreatedAt: p.CreatedAt,
			UpdatedAt: p.UpdatedAt,
		})
	}
	data.Data = response
	return SuccessResponse(data)
}
func ParcelImageResponse(data *ent.ParcelImage) *fiber.Map {
	return SuccessResponse(&ParcelImageResponseData{
		ID:    data.ID,
		Image: data.Image,
	})
}
