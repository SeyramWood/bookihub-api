// Code generated by ent, DO NOT EDIT.

package notification

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the notification type in the database.
	Label = "notification"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldEvent holds the string denoting the event field in the database.
	FieldEvent = "event"
	// FieldActivity holds the string denoting the activity field in the database.
	FieldActivity = "activity"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldSubjectType holds the string denoting the subject_type field in the database.
	FieldSubjectType = "subject_type"
	// FieldSubjectID holds the string denoting the subject_id field in the database.
	FieldSubjectID = "subject_id"
	// FieldCreatorType holds the string denoting the creator_type field in the database.
	FieldCreatorType = "creator_type"
	// FieldCreatorID holds the string denoting the creator_id field in the database.
	FieldCreatorID = "creator_id"
	// FieldCustomerReadAt holds the string denoting the customer_read_at field in the database.
	FieldCustomerReadAt = "customer_read_at"
	// FieldBookibusReadAt holds the string denoting the bookibus_read_at field in the database.
	FieldBookibusReadAt = "bookibus_read_at"
	// FieldCompanyReadAt holds the string denoting the company_read_at field in the database.
	FieldCompanyReadAt = "company_read_at"
	// FieldData holds the string denoting the data field in the database.
	FieldData = "data"
	// EdgeBookibusUser holds the string denoting the bookibus_user edge name in mutations.
	EdgeBookibusUser = "bookibus_user"
	// EdgeCompanyUser holds the string denoting the company_user edge name in mutations.
	EdgeCompanyUser = "company_user"
	// EdgeCustomer holds the string denoting the customer edge name in mutations.
	EdgeCustomer = "customer"
	// EdgeCompany holds the string denoting the company edge name in mutations.
	EdgeCompany = "company"
	// Table holds the table name of the notification in the database.
	Table = "notifications"
	// BookibusUserTable is the table that holds the bookibus_user relation/edge. The primary key declared below.
	BookibusUserTable = "bookibus_user_notifications"
	// BookibusUserInverseTable is the table name for the BookibusUser entity.
	// It exists in this package in order to avoid circular dependency with the "bookibususer" package.
	BookibusUserInverseTable = "bookibus_users"
	// CompanyUserTable is the table that holds the company_user relation/edge. The primary key declared below.
	CompanyUserTable = "company_user_notifications"
	// CompanyUserInverseTable is the table name for the CompanyUser entity.
	// It exists in this package in order to avoid circular dependency with the "companyuser" package.
	CompanyUserInverseTable = "company_users"
	// CustomerTable is the table that holds the customer relation/edge. The primary key declared below.
	CustomerTable = "customer_notifications"
	// CustomerInverseTable is the table name for the Customer entity.
	// It exists in this package in order to avoid circular dependency with the "customer" package.
	CustomerInverseTable = "customers"
	// CompanyTable is the table that holds the company relation/edge.
	CompanyTable = "notifications"
	// CompanyInverseTable is the table name for the Company entity.
	// It exists in this package in order to avoid circular dependency with the "company" package.
	CompanyInverseTable = "companies"
	// CompanyColumn is the table column denoting the company relation/edge.
	CompanyColumn = "company_notifications"
)

// Columns holds all SQL columns for notification fields.
var Columns = []string{
	FieldID,
	FieldEvent,
	FieldActivity,
	FieldDescription,
	FieldSubjectType,
	FieldSubjectID,
	FieldCreatorType,
	FieldCreatorID,
	FieldCustomerReadAt,
	FieldBookibusReadAt,
	FieldCompanyReadAt,
	FieldData,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "notifications"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"company_notifications",
}

var (
	// BookibusUserPrimaryKey and BookibusUserColumn2 are the table columns denoting the
	// primary key for the bookibus_user relation (M2M).
	BookibusUserPrimaryKey = []string{"bookibus_user_id", "notification_id"}
	// CompanyUserPrimaryKey and CompanyUserColumn2 are the table columns denoting the
	// primary key for the company_user relation (M2M).
	CompanyUserPrimaryKey = []string{"company_user_id", "notification_id"}
	// CustomerPrimaryKey and CustomerColumn2 are the table columns denoting the
	// primary key for the customer relation (M2M).
	CustomerPrimaryKey = []string{"customer_id", "notification_id"}
)

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
	// EventValidator is a validator for the "event" field. It is called by the builders before save.
	EventValidator func(string) error
	// ActivityValidator is a validator for the "activity" field. It is called by the builders before save.
	ActivityValidator func(string) error
	// DescriptionValidator is a validator for the "description" field. It is called by the builders before save.
	DescriptionValidator func(string) error
	// SubjectTypeValidator is a validator for the "subject_type" field. It is called by the builders before save.
	SubjectTypeValidator func(string) error
	// CreatorTypeValidator is a validator for the "creator_type" field. It is called by the builders before save.
	CreatorTypeValidator func(string) error
)

// OrderOption defines the ordering options for the Notification queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByEvent orders the results by the event field.
func ByEvent(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEvent, opts...).ToFunc()
}

// ByActivity orders the results by the activity field.
func ByActivity(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldActivity, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// BySubjectType orders the results by the subject_type field.
func BySubjectType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSubjectType, opts...).ToFunc()
}

// BySubjectID orders the results by the subject_id field.
func BySubjectID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSubjectID, opts...).ToFunc()
}

// ByCreatorType orders the results by the creator_type field.
func ByCreatorType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatorType, opts...).ToFunc()
}

// ByCreatorID orders the results by the creator_id field.
func ByCreatorID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatorID, opts...).ToFunc()
}

// ByCustomerReadAt orders the results by the customer_read_at field.
func ByCustomerReadAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCustomerReadAt, opts...).ToFunc()
}

// ByBookibusUserCount orders the results by bookibus_user count.
func ByBookibusUserCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newBookibusUserStep(), opts...)
	}
}

// ByBookibusUser orders the results by bookibus_user terms.
func ByBookibusUser(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newBookibusUserStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByCompanyUserCount orders the results by company_user count.
func ByCompanyUserCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newCompanyUserStep(), opts...)
	}
}

// ByCompanyUser orders the results by company_user terms.
func ByCompanyUser(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCompanyUserStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByCustomerCount orders the results by customer count.
func ByCustomerCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newCustomerStep(), opts...)
	}
}

// ByCustomer orders the results by customer terms.
func ByCustomer(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCustomerStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByCompanyField orders the results by company field.
func ByCompanyField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCompanyStep(), sql.OrderByField(field, opts...))
	}
}
func newBookibusUserStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(BookibusUserInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, BookibusUserTable, BookibusUserPrimaryKey...),
	)
}
func newCompanyUserStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CompanyUserInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, CompanyUserTable, CompanyUserPrimaryKey...),
	)
}
func newCustomerStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CustomerInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, CustomerTable, CustomerPrimaryKey...),
	)
}
func newCompanyStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CompanyInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, CompanyTable, CompanyColumn),
	)
}
