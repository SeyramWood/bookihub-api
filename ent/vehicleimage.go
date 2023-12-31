// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/SeyramWood/bookibus/ent/vehicle"
	"github.com/SeyramWood/bookibus/ent/vehicleimage"
)

// VehicleImage is the model entity for the VehicleImage schema.
type VehicleImage struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Image holds the value of the "image" field.
	Image string `json:"image,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the VehicleImageQuery when eager-loading is set.
	Edges          VehicleImageEdges `json:"edges"`
	vehicle_images *int
	selectValues   sql.SelectValues
}

// VehicleImageEdges holds the relations/edges for other nodes in the graph.
type VehicleImageEdges struct {
	// Vehicle holds the value of the vehicle edge.
	Vehicle *Vehicle `json:"vehicle,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// VehicleOrErr returns the Vehicle value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e VehicleImageEdges) VehicleOrErr() (*Vehicle, error) {
	if e.loadedTypes[0] {
		if e.Vehicle == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: vehicle.Label}
		}
		return e.Vehicle, nil
	}
	return nil, &NotLoadedError{edge: "vehicle"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*VehicleImage) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case vehicleimage.FieldID:
			values[i] = new(sql.NullInt64)
		case vehicleimage.FieldImage:
			values[i] = new(sql.NullString)
		case vehicleimage.FieldCreatedAt, vehicleimage.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case vehicleimage.ForeignKeys[0]: // vehicle_images
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the VehicleImage fields.
func (vi *VehicleImage) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case vehicleimage.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			vi.ID = int(value.Int64)
		case vehicleimage.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				vi.CreatedAt = value.Time
			}
		case vehicleimage.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				vi.UpdatedAt = value.Time
			}
		case vehicleimage.FieldImage:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field image", values[i])
			} else if value.Valid {
				vi.Image = value.String
			}
		case vehicleimage.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field vehicle_images", value)
			} else if value.Valid {
				vi.vehicle_images = new(int)
				*vi.vehicle_images = int(value.Int64)
			}
		default:
			vi.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the VehicleImage.
// This includes values selected through modifiers, order, etc.
func (vi *VehicleImage) Value(name string) (ent.Value, error) {
	return vi.selectValues.Get(name)
}

// QueryVehicle queries the "vehicle" edge of the VehicleImage entity.
func (vi *VehicleImage) QueryVehicle() *VehicleQuery {
	return NewVehicleImageClient(vi.config).QueryVehicle(vi)
}

// Update returns a builder for updating this VehicleImage.
// Note that you need to call VehicleImage.Unwrap() before calling this method if this VehicleImage
// was returned from a transaction, and the transaction was committed or rolled back.
func (vi *VehicleImage) Update() *VehicleImageUpdateOne {
	return NewVehicleImageClient(vi.config).UpdateOne(vi)
}

// Unwrap unwraps the VehicleImage entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (vi *VehicleImage) Unwrap() *VehicleImage {
	_tx, ok := vi.config.driver.(*txDriver)
	if !ok {
		panic("ent: VehicleImage is not a transactional entity")
	}
	vi.config.driver = _tx.drv
	return vi
}

// String implements the fmt.Stringer.
func (vi *VehicleImage) String() string {
	var builder strings.Builder
	builder.WriteString("VehicleImage(")
	builder.WriteString(fmt.Sprintf("id=%v, ", vi.ID))
	builder.WriteString("created_at=")
	builder.WriteString(vi.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(vi.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("image=")
	builder.WriteString(vi.Image)
	builder.WriteByte(')')
	return builder.String()
}

// VehicleImages is a parsable slice of VehicleImage.
type VehicleImages []*VehicleImage
