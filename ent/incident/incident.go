// Code generated by ent, DO NOT EDIT.

package incident

import (
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the incident type in the database.
	Label = "incident"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldTime holds the string denoting the time field in the database.
	FieldTime = "time"
	// FieldLocation holds the string denoting the location field in the database.
	FieldLocation = "location"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldAudio holds the string denoting the audio field in the database.
	FieldAudio = "audio"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// EdgeImages holds the string denoting the images edge name in mutations.
	EdgeImages = "images"
	// EdgeTrip holds the string denoting the trip edge name in mutations.
	EdgeTrip = "trip"
	// EdgeCompany holds the string denoting the company edge name in mutations.
	EdgeCompany = "company"
	// EdgeDriver holds the string denoting the driver edge name in mutations.
	EdgeDriver = "driver"
	// Table holds the table name of the incident in the database.
	Table = "incidents"
	// ImagesTable is the table that holds the images relation/edge.
	ImagesTable = "incident_images"
	// ImagesInverseTable is the table name for the IncidentImage entity.
	// It exists in this package in order to avoid circular dependency with the "incidentimage" package.
	ImagesInverseTable = "incident_images"
	// ImagesColumn is the table column denoting the images relation/edge.
	ImagesColumn = "incident_images"
	// TripTable is the table that holds the trip relation/edge.
	TripTable = "incidents"
	// TripInverseTable is the table name for the Trip entity.
	// It exists in this package in order to avoid circular dependency with the "trip" package.
	TripInverseTable = "trips"
	// TripColumn is the table column denoting the trip relation/edge.
	TripColumn = "trip_incidents"
	// CompanyTable is the table that holds the company relation/edge.
	CompanyTable = "incidents"
	// CompanyInverseTable is the table name for the Company entity.
	// It exists in this package in order to avoid circular dependency with the "company" package.
	CompanyInverseTable = "companies"
	// CompanyColumn is the table column denoting the company relation/edge.
	CompanyColumn = "company_incidents"
	// DriverTable is the table that holds the driver relation/edge.
	DriverTable = "incidents"
	// DriverInverseTable is the table name for the CompanyUser entity.
	// It exists in this package in order to avoid circular dependency with the "companyuser" package.
	DriverInverseTable = "company_users"
	// DriverColumn is the table column denoting the driver relation/edge.
	DriverColumn = "company_user_incidents"
)

// Columns holds all SQL columns for incident fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldTime,
	FieldLocation,
	FieldDescription,
	FieldType,
	FieldAudio,
	FieldStatus,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "incidents"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"company_incidents",
	"company_user_incidents",
	"trip_incidents",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// LocationValidator is a validator for the "location" field. It is called by the builders before save.
	LocationValidator func(string) error
	// DescriptionValidator is a validator for the "description" field. It is called by the builders before save.
	DescriptionValidator func(string) error
	// TypeValidator is a validator for the "type" field. It is called by the builders before save.
	TypeValidator func(string) error
)

// Status defines the type for the "status" enum field.
type Status string

// StatusPending is the default value of the Status enum.
const DefaultStatus = StatusPending

// Status values.
const (
	StatusPending    Status = "pending"
	StatusInProgress Status = "in-progress"
	StatusResolved   Status = "resolved"
)

func (s Status) String() string {
	return string(s)
}

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s Status) error {
	switch s {
	case StatusPending, StatusInProgress, StatusResolved:
		return nil
	default:
		return fmt.Errorf("incident: invalid enum value for status field: %q", s)
	}
}

// OrderOption defines the ordering options for the Incident queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByTime orders the results by the time field.
func ByTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTime, opts...).ToFunc()
}

// ByLocation orders the results by the location field.
func ByLocation(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLocation, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// ByType orders the results by the type field.
func ByType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldType, opts...).ToFunc()
}

// ByAudio orders the results by the audio field.
func ByAudio(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAudio, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByImagesCount orders the results by images count.
func ByImagesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newImagesStep(), opts...)
	}
}

// ByImages orders the results by images terms.
func ByImages(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newImagesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByTripField orders the results by trip field.
func ByTripField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newTripStep(), sql.OrderByField(field, opts...))
	}
}

// ByCompanyField orders the results by company field.
func ByCompanyField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCompanyStep(), sql.OrderByField(field, opts...))
	}
}

// ByDriverField orders the results by driver field.
func ByDriverField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newDriverStep(), sql.OrderByField(field, opts...))
	}
}
func newImagesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ImagesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ImagesTable, ImagesColumn),
	)
}
func newTripStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(TripInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, TripTable, TripColumn),
	)
}
func newCompanyStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CompanyInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, CompanyTable, CompanyColumn),
	)
}
func newDriverStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(DriverInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, DriverTable, DriverColumn),
	)
}
