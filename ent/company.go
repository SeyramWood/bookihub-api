// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/SeyramWood/bookibus/ent/company"
	"github.com/SeyramWood/bookibus/ent/schema"
)

// Company is the model entity for the Company schema.
type Company struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Phone holds the value of the "phone" field.
	Phone string `json:"phone,omitempty"`
	// Email holds the value of the "email" field.
	Email string `json:"email,omitempty"`
	// Certificate holds the value of the "certificate" field.
	Certificate string `json:"certificate,omitempty"`
	// BankAccount holds the value of the "bank_account" field.
	BankAccount *schema.BankAccount `json:"bank_account,omitempty"`
	// ContactPerson holds the value of the "contact_person" field.
	ContactPerson *schema.ContactPerson `json:"contact_person,omitempty"`
	// OnboardingStatus holds the value of the "onboarding_status" field.
	OnboardingStatus company.OnboardingStatus `json:"onboarding_status,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CompanyQuery when eager-loading is set.
	Edges        CompanyEdges `json:"edges"`
	selectValues sql.SelectValues
}

// CompanyEdges holds the relations/edges for other nodes in the graph.
type CompanyEdges struct {
	// Profile holds the value of the profile edge.
	Profile []*CompanyUser `json:"profile,omitempty"`
	// Terminals holds the value of the terminals edge.
	Terminals []*Terminal `json:"terminals,omitempty"`
	// Vehicles holds the value of the vehicles edge.
	Vehicles []*Vehicle `json:"vehicles,omitempty"`
	// Routes holds the value of the routes edge.
	Routes []*Route `json:"routes,omitempty"`
	// Trips holds the value of the trips edge.
	Trips []*Trip `json:"trips,omitempty"`
	// Bookings holds the value of the bookings edge.
	Bookings []*Booking `json:"bookings,omitempty"`
	// Incidents holds the value of the incidents edge.
	Incidents []*Incident `json:"incidents,omitempty"`
	// Parcels holds the value of the parcels edge.
	Parcels []*Parcel `json:"parcels,omitempty"`
	// Transactions holds the value of the transactions edge.
	Transactions []*Transaction `json:"transactions,omitempty"`
	// Notifications holds the value of the notifications edge.
	Notifications []*Notification `json:"notifications,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [10]bool
}

// ProfileOrErr returns the Profile value or an error if the edge
// was not loaded in eager-loading.
func (e CompanyEdges) ProfileOrErr() ([]*CompanyUser, error) {
	if e.loadedTypes[0] {
		return e.Profile, nil
	}
	return nil, &NotLoadedError{edge: "profile"}
}

// TerminalsOrErr returns the Terminals value or an error if the edge
// was not loaded in eager-loading.
func (e CompanyEdges) TerminalsOrErr() ([]*Terminal, error) {
	if e.loadedTypes[1] {
		return e.Terminals, nil
	}
	return nil, &NotLoadedError{edge: "terminals"}
}

// VehiclesOrErr returns the Vehicles value or an error if the edge
// was not loaded in eager-loading.
func (e CompanyEdges) VehiclesOrErr() ([]*Vehicle, error) {
	if e.loadedTypes[2] {
		return e.Vehicles, nil
	}
	return nil, &NotLoadedError{edge: "vehicles"}
}

// RoutesOrErr returns the Routes value or an error if the edge
// was not loaded in eager-loading.
func (e CompanyEdges) RoutesOrErr() ([]*Route, error) {
	if e.loadedTypes[3] {
		return e.Routes, nil
	}
	return nil, &NotLoadedError{edge: "routes"}
}

// TripsOrErr returns the Trips value or an error if the edge
// was not loaded in eager-loading.
func (e CompanyEdges) TripsOrErr() ([]*Trip, error) {
	if e.loadedTypes[4] {
		return e.Trips, nil
	}
	return nil, &NotLoadedError{edge: "trips"}
}

// BookingsOrErr returns the Bookings value or an error if the edge
// was not loaded in eager-loading.
func (e CompanyEdges) BookingsOrErr() ([]*Booking, error) {
	if e.loadedTypes[5] {
		return e.Bookings, nil
	}
	return nil, &NotLoadedError{edge: "bookings"}
}

// IncidentsOrErr returns the Incidents value or an error if the edge
// was not loaded in eager-loading.
func (e CompanyEdges) IncidentsOrErr() ([]*Incident, error) {
	if e.loadedTypes[6] {
		return e.Incidents, nil
	}
	return nil, &NotLoadedError{edge: "incidents"}
}

// ParcelsOrErr returns the Parcels value or an error if the edge
// was not loaded in eager-loading.
func (e CompanyEdges) ParcelsOrErr() ([]*Parcel, error) {
	if e.loadedTypes[7] {
		return e.Parcels, nil
	}
	return nil, &NotLoadedError{edge: "parcels"}
}

// TransactionsOrErr returns the Transactions value or an error if the edge
// was not loaded in eager-loading.
func (e CompanyEdges) TransactionsOrErr() ([]*Transaction, error) {
	if e.loadedTypes[8] {
		return e.Transactions, nil
	}
	return nil, &NotLoadedError{edge: "transactions"}
}

// NotificationsOrErr returns the Notifications value or an error if the edge
// was not loaded in eager-loading.
func (e CompanyEdges) NotificationsOrErr() ([]*Notification, error) {
	if e.loadedTypes[9] {
		return e.Notifications, nil
	}
	return nil, &NotLoadedError{edge: "notifications"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Company) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case company.FieldBankAccount, company.FieldContactPerson:
			values[i] = new([]byte)
		case company.FieldID:
			values[i] = new(sql.NullInt64)
		case company.FieldName, company.FieldPhone, company.FieldEmail, company.FieldCertificate, company.FieldOnboardingStatus:
			values[i] = new(sql.NullString)
		case company.FieldCreatedAt, company.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Company fields.
func (c *Company) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case company.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			c.ID = int(value.Int64)
		case company.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				c.CreatedAt = value.Time
			}
		case company.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				c.UpdatedAt = value.Time
			}
		case company.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				c.Name = value.String
			}
		case company.FieldPhone:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field phone", values[i])
			} else if value.Valid {
				c.Phone = value.String
			}
		case company.FieldEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email", values[i])
			} else if value.Valid {
				c.Email = value.String
			}
		case company.FieldCertificate:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field certificate", values[i])
			} else if value.Valid {
				c.Certificate = value.String
			}
		case company.FieldBankAccount:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field bank_account", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &c.BankAccount); err != nil {
					return fmt.Errorf("unmarshal field bank_account: %w", err)
				}
			}
		case company.FieldContactPerson:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field contact_person", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &c.ContactPerson); err != nil {
					return fmt.Errorf("unmarshal field contact_person: %w", err)
				}
			}
		case company.FieldOnboardingStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field onboarding_status", values[i])
			} else if value.Valid {
				c.OnboardingStatus = company.OnboardingStatus(value.String)
			}
		default:
			c.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Company.
// This includes values selected through modifiers, order, etc.
func (c *Company) Value(name string) (ent.Value, error) {
	return c.selectValues.Get(name)
}

// QueryProfile queries the "profile" edge of the Company entity.
func (c *Company) QueryProfile() *CompanyUserQuery {
	return NewCompanyClient(c.config).QueryProfile(c)
}

// QueryTerminals queries the "terminals" edge of the Company entity.
func (c *Company) QueryTerminals() *TerminalQuery {
	return NewCompanyClient(c.config).QueryTerminals(c)
}

// QueryVehicles queries the "vehicles" edge of the Company entity.
func (c *Company) QueryVehicles() *VehicleQuery {
	return NewCompanyClient(c.config).QueryVehicles(c)
}

// QueryRoutes queries the "routes" edge of the Company entity.
func (c *Company) QueryRoutes() *RouteQuery {
	return NewCompanyClient(c.config).QueryRoutes(c)
}

// QueryTrips queries the "trips" edge of the Company entity.
func (c *Company) QueryTrips() *TripQuery {
	return NewCompanyClient(c.config).QueryTrips(c)
}

// QueryBookings queries the "bookings" edge of the Company entity.
func (c *Company) QueryBookings() *BookingQuery {
	return NewCompanyClient(c.config).QueryBookings(c)
}

// QueryIncidents queries the "incidents" edge of the Company entity.
func (c *Company) QueryIncidents() *IncidentQuery {
	return NewCompanyClient(c.config).QueryIncidents(c)
}

// QueryParcels queries the "parcels" edge of the Company entity.
func (c *Company) QueryParcels() *ParcelQuery {
	return NewCompanyClient(c.config).QueryParcels(c)
}

// QueryTransactions queries the "transactions" edge of the Company entity.
func (c *Company) QueryTransactions() *TransactionQuery {
	return NewCompanyClient(c.config).QueryTransactions(c)
}

// QueryNotifications queries the "notifications" edge of the Company entity.
func (c *Company) QueryNotifications() *NotificationQuery {
	return NewCompanyClient(c.config).QueryNotifications(c)
}

// Update returns a builder for updating this Company.
// Note that you need to call Company.Unwrap() before calling this method if this Company
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Company) Update() *CompanyUpdateOne {
	return NewCompanyClient(c.config).UpdateOne(c)
}

// Unwrap unwraps the Company entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Company) Unwrap() *Company {
	_tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Company is not a transactional entity")
	}
	c.config.driver = _tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Company) String() string {
	var builder strings.Builder
	builder.WriteString("Company(")
	builder.WriteString(fmt.Sprintf("id=%v, ", c.ID))
	builder.WriteString("created_at=")
	builder.WriteString(c.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(c.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(c.Name)
	builder.WriteString(", ")
	builder.WriteString("phone=")
	builder.WriteString(c.Phone)
	builder.WriteString(", ")
	builder.WriteString("email=")
	builder.WriteString(c.Email)
	builder.WriteString(", ")
	builder.WriteString("certificate=")
	builder.WriteString(c.Certificate)
	builder.WriteString(", ")
	builder.WriteString("bank_account=")
	builder.WriteString(fmt.Sprintf("%v", c.BankAccount))
	builder.WriteString(", ")
	builder.WriteString("contact_person=")
	builder.WriteString(fmt.Sprintf("%v", c.ContactPerson))
	builder.WriteString(", ")
	builder.WriteString("onboarding_status=")
	builder.WriteString(fmt.Sprintf("%v", c.OnboardingStatus))
	builder.WriteByte(')')
	return builder.String()
}

// Companies is a parsable slice of Company.
type Companies []*Company
