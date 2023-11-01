// Code generated by ent, DO NOT EDIT.

package booking

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/SeyramWood/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Booking {
	return predicate.Booking(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Booking {
	return predicate.Booking(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Booking {
	return predicate.Booking(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Booking {
	return predicate.Booking(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Booking {
	return predicate.Booking(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Booking {
	return predicate.Booking(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Booking {
	return predicate.Booking(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Booking {
	return predicate.Booking(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Booking {
	return predicate.Booking(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Booking {
	return predicate.Booking(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Booking {
	return predicate.Booking(sql.FieldEQ(FieldUpdatedAt, v))
}

// Reference applies equality check predicate on the "reference" field. It's identical to ReferenceEQ.
func Reference(v string) predicate.Booking {
	return predicate.Booking(sql.FieldEQ(FieldReference, v))
}

// BookingNumber applies equality check predicate on the "booking_number" field. It's identical to BookingNumberEQ.
func BookingNumber(v string) predicate.Booking {
	return predicate.Booking(sql.FieldEQ(FieldBookingNumber, v))
}

// Vat applies equality check predicate on the "vat" field. It's identical to VatEQ.
func Vat(v float64) predicate.Booking {
	return predicate.Booking(sql.FieldEQ(FieldVat, v))
}

// SmsFee applies equality check predicate on the "sms_fee" field. It's identical to SmsFeeEQ.
func SmsFee(v float64) predicate.Booking {
	return predicate.Booking(sql.FieldEQ(FieldSmsFee, v))
}

// Amount applies equality check predicate on the "amount" field. It's identical to AmountEQ.
func Amount(v float64) predicate.Booking {
	return predicate.Booking(sql.FieldEQ(FieldAmount, v))
}

// RefundAmount applies equality check predicate on the "refund_amount" field. It's identical to RefundAmountEQ.
func RefundAmount(v float64) predicate.Booking {
	return predicate.Booking(sql.FieldEQ(FieldRefundAmount, v))
}

// PaidAt applies equality check predicate on the "paid_at" field. It's identical to PaidAtEQ.
func PaidAt(v time.Time) predicate.Booking {
	return predicate.Booking(sql.FieldEQ(FieldPaidAt, v))
}

// RefundAt applies equality check predicate on the "refund_at" field. It's identical to RefundAtEQ.
func RefundAt(v time.Time) predicate.Booking {
	return predicate.Booking(sql.FieldEQ(FieldRefundAt, v))
}

// SmsNotification applies equality check predicate on the "sms_notification" field. It's identical to SmsNotificationEQ.
func SmsNotification(v bool) predicate.Booking {
	return predicate.Booking(sql.FieldEQ(FieldSmsNotification, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Booking {
	return predicate.Booking(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Booking {
	return predicate.Booking(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Booking {
	return predicate.Booking(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Booking {
	return predicate.Booking(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Booking {
	return predicate.Booking(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Booking {
	return predicate.Booking(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Booking {
	return predicate.Booking(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Booking {
	return predicate.Booking(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Booking {
	return predicate.Booking(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Booking {
	return predicate.Booking(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Booking {
	return predicate.Booking(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Booking {
	return predicate.Booking(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Booking {
	return predicate.Booking(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Booking {
	return predicate.Booking(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Booking {
	return predicate.Booking(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Booking {
	return predicate.Booking(sql.FieldLTE(FieldUpdatedAt, v))
}

// ReferenceEQ applies the EQ predicate on the "reference" field.
func ReferenceEQ(v string) predicate.Booking {
	return predicate.Booking(sql.FieldEQ(FieldReference, v))
}

// ReferenceNEQ applies the NEQ predicate on the "reference" field.
func ReferenceNEQ(v string) predicate.Booking {
	return predicate.Booking(sql.FieldNEQ(FieldReference, v))
}

// ReferenceIn applies the In predicate on the "reference" field.
func ReferenceIn(vs ...string) predicate.Booking {
	return predicate.Booking(sql.FieldIn(FieldReference, vs...))
}

// ReferenceNotIn applies the NotIn predicate on the "reference" field.
func ReferenceNotIn(vs ...string) predicate.Booking {
	return predicate.Booking(sql.FieldNotIn(FieldReference, vs...))
}

// ReferenceGT applies the GT predicate on the "reference" field.
func ReferenceGT(v string) predicate.Booking {
	return predicate.Booking(sql.FieldGT(FieldReference, v))
}

// ReferenceGTE applies the GTE predicate on the "reference" field.
func ReferenceGTE(v string) predicate.Booking {
	return predicate.Booking(sql.FieldGTE(FieldReference, v))
}

// ReferenceLT applies the LT predicate on the "reference" field.
func ReferenceLT(v string) predicate.Booking {
	return predicate.Booking(sql.FieldLT(FieldReference, v))
}

// ReferenceLTE applies the LTE predicate on the "reference" field.
func ReferenceLTE(v string) predicate.Booking {
	return predicate.Booking(sql.FieldLTE(FieldReference, v))
}

// ReferenceContains applies the Contains predicate on the "reference" field.
func ReferenceContains(v string) predicate.Booking {
	return predicate.Booking(sql.FieldContains(FieldReference, v))
}

// ReferenceHasPrefix applies the HasPrefix predicate on the "reference" field.
func ReferenceHasPrefix(v string) predicate.Booking {
	return predicate.Booking(sql.FieldHasPrefix(FieldReference, v))
}

// ReferenceHasSuffix applies the HasSuffix predicate on the "reference" field.
func ReferenceHasSuffix(v string) predicate.Booking {
	return predicate.Booking(sql.FieldHasSuffix(FieldReference, v))
}

// ReferenceIsNil applies the IsNil predicate on the "reference" field.
func ReferenceIsNil() predicate.Booking {
	return predicate.Booking(sql.FieldIsNull(FieldReference))
}

// ReferenceNotNil applies the NotNil predicate on the "reference" field.
func ReferenceNotNil() predicate.Booking {
	return predicate.Booking(sql.FieldNotNull(FieldReference))
}

// ReferenceEqualFold applies the EqualFold predicate on the "reference" field.
func ReferenceEqualFold(v string) predicate.Booking {
	return predicate.Booking(sql.FieldEqualFold(FieldReference, v))
}

// ReferenceContainsFold applies the ContainsFold predicate on the "reference" field.
func ReferenceContainsFold(v string) predicate.Booking {
	return predicate.Booking(sql.FieldContainsFold(FieldReference, v))
}

// BookingNumberEQ applies the EQ predicate on the "booking_number" field.
func BookingNumberEQ(v string) predicate.Booking {
	return predicate.Booking(sql.FieldEQ(FieldBookingNumber, v))
}

// BookingNumberNEQ applies the NEQ predicate on the "booking_number" field.
func BookingNumberNEQ(v string) predicate.Booking {
	return predicate.Booking(sql.FieldNEQ(FieldBookingNumber, v))
}

// BookingNumberIn applies the In predicate on the "booking_number" field.
func BookingNumberIn(vs ...string) predicate.Booking {
	return predicate.Booking(sql.FieldIn(FieldBookingNumber, vs...))
}

// BookingNumberNotIn applies the NotIn predicate on the "booking_number" field.
func BookingNumberNotIn(vs ...string) predicate.Booking {
	return predicate.Booking(sql.FieldNotIn(FieldBookingNumber, vs...))
}

// BookingNumberGT applies the GT predicate on the "booking_number" field.
func BookingNumberGT(v string) predicate.Booking {
	return predicate.Booking(sql.FieldGT(FieldBookingNumber, v))
}

// BookingNumberGTE applies the GTE predicate on the "booking_number" field.
func BookingNumberGTE(v string) predicate.Booking {
	return predicate.Booking(sql.FieldGTE(FieldBookingNumber, v))
}

// BookingNumberLT applies the LT predicate on the "booking_number" field.
func BookingNumberLT(v string) predicate.Booking {
	return predicate.Booking(sql.FieldLT(FieldBookingNumber, v))
}

// BookingNumberLTE applies the LTE predicate on the "booking_number" field.
func BookingNumberLTE(v string) predicate.Booking {
	return predicate.Booking(sql.FieldLTE(FieldBookingNumber, v))
}

// BookingNumberContains applies the Contains predicate on the "booking_number" field.
func BookingNumberContains(v string) predicate.Booking {
	return predicate.Booking(sql.FieldContains(FieldBookingNumber, v))
}

// BookingNumberHasPrefix applies the HasPrefix predicate on the "booking_number" field.
func BookingNumberHasPrefix(v string) predicate.Booking {
	return predicate.Booking(sql.FieldHasPrefix(FieldBookingNumber, v))
}

// BookingNumberHasSuffix applies the HasSuffix predicate on the "booking_number" field.
func BookingNumberHasSuffix(v string) predicate.Booking {
	return predicate.Booking(sql.FieldHasSuffix(FieldBookingNumber, v))
}

// BookingNumberEqualFold applies the EqualFold predicate on the "booking_number" field.
func BookingNumberEqualFold(v string) predicate.Booking {
	return predicate.Booking(sql.FieldEqualFold(FieldBookingNumber, v))
}

// BookingNumberContainsFold applies the ContainsFold predicate on the "booking_number" field.
func BookingNumberContainsFold(v string) predicate.Booking {
	return predicate.Booking(sql.FieldContainsFold(FieldBookingNumber, v))
}

// VatEQ applies the EQ predicate on the "vat" field.
func VatEQ(v float64) predicate.Booking {
	return predicate.Booking(sql.FieldEQ(FieldVat, v))
}

// VatNEQ applies the NEQ predicate on the "vat" field.
func VatNEQ(v float64) predicate.Booking {
	return predicate.Booking(sql.FieldNEQ(FieldVat, v))
}

// VatIn applies the In predicate on the "vat" field.
func VatIn(vs ...float64) predicate.Booking {
	return predicate.Booking(sql.FieldIn(FieldVat, vs...))
}

// VatNotIn applies the NotIn predicate on the "vat" field.
func VatNotIn(vs ...float64) predicate.Booking {
	return predicate.Booking(sql.FieldNotIn(FieldVat, vs...))
}

// VatGT applies the GT predicate on the "vat" field.
func VatGT(v float64) predicate.Booking {
	return predicate.Booking(sql.FieldGT(FieldVat, v))
}

// VatGTE applies the GTE predicate on the "vat" field.
func VatGTE(v float64) predicate.Booking {
	return predicate.Booking(sql.FieldGTE(FieldVat, v))
}

// VatLT applies the LT predicate on the "vat" field.
func VatLT(v float64) predicate.Booking {
	return predicate.Booking(sql.FieldLT(FieldVat, v))
}

// VatLTE applies the LTE predicate on the "vat" field.
func VatLTE(v float64) predicate.Booking {
	return predicate.Booking(sql.FieldLTE(FieldVat, v))
}

// SmsFeeEQ applies the EQ predicate on the "sms_fee" field.
func SmsFeeEQ(v float64) predicate.Booking {
	return predicate.Booking(sql.FieldEQ(FieldSmsFee, v))
}

// SmsFeeNEQ applies the NEQ predicate on the "sms_fee" field.
func SmsFeeNEQ(v float64) predicate.Booking {
	return predicate.Booking(sql.FieldNEQ(FieldSmsFee, v))
}

// SmsFeeIn applies the In predicate on the "sms_fee" field.
func SmsFeeIn(vs ...float64) predicate.Booking {
	return predicate.Booking(sql.FieldIn(FieldSmsFee, vs...))
}

// SmsFeeNotIn applies the NotIn predicate on the "sms_fee" field.
func SmsFeeNotIn(vs ...float64) predicate.Booking {
	return predicate.Booking(sql.FieldNotIn(FieldSmsFee, vs...))
}

// SmsFeeGT applies the GT predicate on the "sms_fee" field.
func SmsFeeGT(v float64) predicate.Booking {
	return predicate.Booking(sql.FieldGT(FieldSmsFee, v))
}

// SmsFeeGTE applies the GTE predicate on the "sms_fee" field.
func SmsFeeGTE(v float64) predicate.Booking {
	return predicate.Booking(sql.FieldGTE(FieldSmsFee, v))
}

// SmsFeeLT applies the LT predicate on the "sms_fee" field.
func SmsFeeLT(v float64) predicate.Booking {
	return predicate.Booking(sql.FieldLT(FieldSmsFee, v))
}

// SmsFeeLTE applies the LTE predicate on the "sms_fee" field.
func SmsFeeLTE(v float64) predicate.Booking {
	return predicate.Booking(sql.FieldLTE(FieldSmsFee, v))
}

// AmountEQ applies the EQ predicate on the "amount" field.
func AmountEQ(v float64) predicate.Booking {
	return predicate.Booking(sql.FieldEQ(FieldAmount, v))
}

// AmountNEQ applies the NEQ predicate on the "amount" field.
func AmountNEQ(v float64) predicate.Booking {
	return predicate.Booking(sql.FieldNEQ(FieldAmount, v))
}

// AmountIn applies the In predicate on the "amount" field.
func AmountIn(vs ...float64) predicate.Booking {
	return predicate.Booking(sql.FieldIn(FieldAmount, vs...))
}

// AmountNotIn applies the NotIn predicate on the "amount" field.
func AmountNotIn(vs ...float64) predicate.Booking {
	return predicate.Booking(sql.FieldNotIn(FieldAmount, vs...))
}

// AmountGT applies the GT predicate on the "amount" field.
func AmountGT(v float64) predicate.Booking {
	return predicate.Booking(sql.FieldGT(FieldAmount, v))
}

// AmountGTE applies the GTE predicate on the "amount" field.
func AmountGTE(v float64) predicate.Booking {
	return predicate.Booking(sql.FieldGTE(FieldAmount, v))
}

// AmountLT applies the LT predicate on the "amount" field.
func AmountLT(v float64) predicate.Booking {
	return predicate.Booking(sql.FieldLT(FieldAmount, v))
}

// AmountLTE applies the LTE predicate on the "amount" field.
func AmountLTE(v float64) predicate.Booking {
	return predicate.Booking(sql.FieldLTE(FieldAmount, v))
}

// RefundAmountEQ applies the EQ predicate on the "refund_amount" field.
func RefundAmountEQ(v float64) predicate.Booking {
	return predicate.Booking(sql.FieldEQ(FieldRefundAmount, v))
}

// RefundAmountNEQ applies the NEQ predicate on the "refund_amount" field.
func RefundAmountNEQ(v float64) predicate.Booking {
	return predicate.Booking(sql.FieldNEQ(FieldRefundAmount, v))
}

// RefundAmountIn applies the In predicate on the "refund_amount" field.
func RefundAmountIn(vs ...float64) predicate.Booking {
	return predicate.Booking(sql.FieldIn(FieldRefundAmount, vs...))
}

// RefundAmountNotIn applies the NotIn predicate on the "refund_amount" field.
func RefundAmountNotIn(vs ...float64) predicate.Booking {
	return predicate.Booking(sql.FieldNotIn(FieldRefundAmount, vs...))
}

// RefundAmountGT applies the GT predicate on the "refund_amount" field.
func RefundAmountGT(v float64) predicate.Booking {
	return predicate.Booking(sql.FieldGT(FieldRefundAmount, v))
}

// RefundAmountGTE applies the GTE predicate on the "refund_amount" field.
func RefundAmountGTE(v float64) predicate.Booking {
	return predicate.Booking(sql.FieldGTE(FieldRefundAmount, v))
}

// RefundAmountLT applies the LT predicate on the "refund_amount" field.
func RefundAmountLT(v float64) predicate.Booking {
	return predicate.Booking(sql.FieldLT(FieldRefundAmount, v))
}

// RefundAmountLTE applies the LTE predicate on the "refund_amount" field.
func RefundAmountLTE(v float64) predicate.Booking {
	return predicate.Booking(sql.FieldLTE(FieldRefundAmount, v))
}

// RefundAmountIsNil applies the IsNil predicate on the "refund_amount" field.
func RefundAmountIsNil() predicate.Booking {
	return predicate.Booking(sql.FieldIsNull(FieldRefundAmount))
}

// RefundAmountNotNil applies the NotNil predicate on the "refund_amount" field.
func RefundAmountNotNil() predicate.Booking {
	return predicate.Booking(sql.FieldNotNull(FieldRefundAmount))
}

// PaidAtEQ applies the EQ predicate on the "paid_at" field.
func PaidAtEQ(v time.Time) predicate.Booking {
	return predicate.Booking(sql.FieldEQ(FieldPaidAt, v))
}

// PaidAtNEQ applies the NEQ predicate on the "paid_at" field.
func PaidAtNEQ(v time.Time) predicate.Booking {
	return predicate.Booking(sql.FieldNEQ(FieldPaidAt, v))
}

// PaidAtIn applies the In predicate on the "paid_at" field.
func PaidAtIn(vs ...time.Time) predicate.Booking {
	return predicate.Booking(sql.FieldIn(FieldPaidAt, vs...))
}

// PaidAtNotIn applies the NotIn predicate on the "paid_at" field.
func PaidAtNotIn(vs ...time.Time) predicate.Booking {
	return predicate.Booking(sql.FieldNotIn(FieldPaidAt, vs...))
}

// PaidAtGT applies the GT predicate on the "paid_at" field.
func PaidAtGT(v time.Time) predicate.Booking {
	return predicate.Booking(sql.FieldGT(FieldPaidAt, v))
}

// PaidAtGTE applies the GTE predicate on the "paid_at" field.
func PaidAtGTE(v time.Time) predicate.Booking {
	return predicate.Booking(sql.FieldGTE(FieldPaidAt, v))
}

// PaidAtLT applies the LT predicate on the "paid_at" field.
func PaidAtLT(v time.Time) predicate.Booking {
	return predicate.Booking(sql.FieldLT(FieldPaidAt, v))
}

// PaidAtLTE applies the LTE predicate on the "paid_at" field.
func PaidAtLTE(v time.Time) predicate.Booking {
	return predicate.Booking(sql.FieldLTE(FieldPaidAt, v))
}

// PaidAtIsNil applies the IsNil predicate on the "paid_at" field.
func PaidAtIsNil() predicate.Booking {
	return predicate.Booking(sql.FieldIsNull(FieldPaidAt))
}

// PaidAtNotNil applies the NotNil predicate on the "paid_at" field.
func PaidAtNotNil() predicate.Booking {
	return predicate.Booking(sql.FieldNotNull(FieldPaidAt))
}

// RefundAtEQ applies the EQ predicate on the "refund_at" field.
func RefundAtEQ(v time.Time) predicate.Booking {
	return predicate.Booking(sql.FieldEQ(FieldRefundAt, v))
}

// RefundAtNEQ applies the NEQ predicate on the "refund_at" field.
func RefundAtNEQ(v time.Time) predicate.Booking {
	return predicate.Booking(sql.FieldNEQ(FieldRefundAt, v))
}

// RefundAtIn applies the In predicate on the "refund_at" field.
func RefundAtIn(vs ...time.Time) predicate.Booking {
	return predicate.Booking(sql.FieldIn(FieldRefundAt, vs...))
}

// RefundAtNotIn applies the NotIn predicate on the "refund_at" field.
func RefundAtNotIn(vs ...time.Time) predicate.Booking {
	return predicate.Booking(sql.FieldNotIn(FieldRefundAt, vs...))
}

// RefundAtGT applies the GT predicate on the "refund_at" field.
func RefundAtGT(v time.Time) predicate.Booking {
	return predicate.Booking(sql.FieldGT(FieldRefundAt, v))
}

// RefundAtGTE applies the GTE predicate on the "refund_at" field.
func RefundAtGTE(v time.Time) predicate.Booking {
	return predicate.Booking(sql.FieldGTE(FieldRefundAt, v))
}

// RefundAtLT applies the LT predicate on the "refund_at" field.
func RefundAtLT(v time.Time) predicate.Booking {
	return predicate.Booking(sql.FieldLT(FieldRefundAt, v))
}

// RefundAtLTE applies the LTE predicate on the "refund_at" field.
func RefundAtLTE(v time.Time) predicate.Booking {
	return predicate.Booking(sql.FieldLTE(FieldRefundAt, v))
}

// RefundAtIsNil applies the IsNil predicate on the "refund_at" field.
func RefundAtIsNil() predicate.Booking {
	return predicate.Booking(sql.FieldIsNull(FieldRefundAt))
}

// RefundAtNotNil applies the NotNil predicate on the "refund_at" field.
func RefundAtNotNil() predicate.Booking {
	return predicate.Booking(sql.FieldNotNull(FieldRefundAt))
}

// TansTypeEQ applies the EQ predicate on the "tans_type" field.
func TansTypeEQ(v TansType) predicate.Booking {
	return predicate.Booking(sql.FieldEQ(FieldTansType, v))
}

// TansTypeNEQ applies the NEQ predicate on the "tans_type" field.
func TansTypeNEQ(v TansType) predicate.Booking {
	return predicate.Booking(sql.FieldNEQ(FieldTansType, v))
}

// TansTypeIn applies the In predicate on the "tans_type" field.
func TansTypeIn(vs ...TansType) predicate.Booking {
	return predicate.Booking(sql.FieldIn(FieldTansType, vs...))
}

// TansTypeNotIn applies the NotIn predicate on the "tans_type" field.
func TansTypeNotIn(vs ...TansType) predicate.Booking {
	return predicate.Booking(sql.FieldNotIn(FieldTansType, vs...))
}

// SmsNotificationEQ applies the EQ predicate on the "sms_notification" field.
func SmsNotificationEQ(v bool) predicate.Booking {
	return predicate.Booking(sql.FieldEQ(FieldSmsNotification, v))
}

// SmsNotificationNEQ applies the NEQ predicate on the "sms_notification" field.
func SmsNotificationNEQ(v bool) predicate.Booking {
	return predicate.Booking(sql.FieldNEQ(FieldSmsNotification, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v Status) predicate.Booking {
	return predicate.Booking(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v Status) predicate.Booking {
	return predicate.Booking(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...Status) predicate.Booking {
	return predicate.Booking(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...Status) predicate.Booking {
	return predicate.Booking(sql.FieldNotIn(FieldStatus, vs...))
}

// HasPassengers applies the HasEdge predicate on the "passengers" edge.
func HasPassengers() predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, PassengersTable, PassengersColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPassengersWith applies the HasEdge predicate on the "passengers" edge with a given conditions (other predicates).
func HasPassengersWith(preds ...predicate.Passenger) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		step := newPassengersStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasLuggages applies the HasEdge predicate on the "luggages" edge.
func HasLuggages() predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, LuggagesTable, LuggagesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasLuggagesWith applies the HasEdge predicate on the "luggages" edge with a given conditions (other predicates).
func HasLuggagesWith(preds ...predicate.CustomerLuggage) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		step := newLuggagesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasContact applies the HasEdge predicate on the "contact" edge.
func HasContact() predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, ContactTable, ContactColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasContactWith applies the HasEdge predicate on the "contact" edge with a given conditions (other predicates).
func HasContactWith(preds ...predicate.CustomerContact) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		step := newContactStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasTrip applies the HasEdge predicate on the "trip" edge.
func HasTrip() predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, TripTable, TripColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTripWith applies the HasEdge predicate on the "trip" edge with a given conditions (other predicates).
func HasTripWith(preds ...predicate.Trip) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		step := newTripStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasCompany applies the HasEdge predicate on the "company" edge.
func HasCompany() predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, CompanyTable, CompanyColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCompanyWith applies the HasEdge predicate on the "company" edge with a given conditions (other predicates).
func HasCompanyWith(preds ...predicate.Company) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		step := newCompanyStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasCustomer applies the HasEdge predicate on the "customer" edge.
func HasCustomer() predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, CustomerTable, CustomerColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCustomerWith applies the HasEdge predicate on the "customer" edge with a given conditions (other predicates).
func HasCustomerWith(preds ...predicate.Customer) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		step := newCustomerStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Booking) predicate.Booking {
	return predicate.Booking(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Booking) predicate.Booking {
	return predicate.Booking(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Booking) predicate.Booking {
	return predicate.Booking(sql.NotPredicates(p))
}
