// Code generated by ent, DO NOT EDIT.

package parcel

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/SeyramWood/bookibus/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Parcel {
	return predicate.Parcel(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Parcel {
	return predicate.Parcel(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Parcel {
	return predicate.Parcel(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Parcel {
	return predicate.Parcel(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Parcel {
	return predicate.Parcel(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Parcel {
	return predicate.Parcel(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Parcel {
	return predicate.Parcel(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Parcel {
	return predicate.Parcel(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Parcel {
	return predicate.Parcel(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Parcel {
	return predicate.Parcel(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Parcel {
	return predicate.Parcel(sql.FieldEQ(FieldUpdatedAt, v))
}

// ParcelCode applies equality check predicate on the "parcel_code" field. It's identical to ParcelCodeEQ.
func ParcelCode(v string) predicate.Parcel {
	return predicate.Parcel(sql.FieldEQ(FieldParcelCode, v))
}

// Type applies equality check predicate on the "type" field. It's identical to TypeEQ.
func Type(v string) predicate.Parcel {
	return predicate.Parcel(sql.FieldEQ(FieldType, v))
}

// SenderName applies equality check predicate on the "sender_name" field. It's identical to SenderNameEQ.
func SenderName(v string) predicate.Parcel {
	return predicate.Parcel(sql.FieldEQ(FieldSenderName, v))
}

// SenderPhone applies equality check predicate on the "sender_phone" field. It's identical to SenderPhoneEQ.
func SenderPhone(v string) predicate.Parcel {
	return predicate.Parcel(sql.FieldEQ(FieldSenderPhone, v))
}

// SenderEmail applies equality check predicate on the "sender_email" field. It's identical to SenderEmailEQ.
func SenderEmail(v string) predicate.Parcel {
	return predicate.Parcel(sql.FieldEQ(FieldSenderEmail, v))
}

// RecipientName applies equality check predicate on the "recipient_name" field. It's identical to RecipientNameEQ.
func RecipientName(v string) predicate.Parcel {
	return predicate.Parcel(sql.FieldEQ(FieldRecipientName, v))
}

// RecipientPhone applies equality check predicate on the "recipient_phone" field. It's identical to RecipientPhoneEQ.
func RecipientPhone(v string) predicate.Parcel {
	return predicate.Parcel(sql.FieldEQ(FieldRecipientPhone, v))
}

// RecipientLocation applies equality check predicate on the "recipient_location" field. It's identical to RecipientLocationEQ.
func RecipientLocation(v string) predicate.Parcel {
	return predicate.Parcel(sql.FieldEQ(FieldRecipientLocation, v))
}

// Weight applies equality check predicate on the "weight" field. It's identical to WeightEQ.
func Weight(v float32) predicate.Parcel {
	return predicate.Parcel(sql.FieldEQ(FieldWeight, v))
}

// Amount applies equality check predicate on the "amount" field. It's identical to AmountEQ.
func Amount(v float64) predicate.Parcel {
	return predicate.Parcel(sql.FieldEQ(FieldAmount, v))
}

// PaidAt applies equality check predicate on the "paid_at" field. It's identical to PaidAtEQ.
func PaidAt(v time.Time) predicate.Parcel {
	return predicate.Parcel(sql.FieldEQ(FieldPaidAt, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Parcel {
	return predicate.Parcel(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Parcel {
	return predicate.Parcel(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Parcel {
	return predicate.Parcel(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Parcel {
	return predicate.Parcel(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Parcel {
	return predicate.Parcel(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Parcel {
	return predicate.Parcel(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Parcel {
	return predicate.Parcel(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Parcel {
	return predicate.Parcel(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Parcel {
	return predicate.Parcel(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Parcel {
	return predicate.Parcel(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Parcel {
	return predicate.Parcel(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Parcel {
	return predicate.Parcel(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Parcel {
	return predicate.Parcel(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Parcel {
	return predicate.Parcel(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Parcel {
	return predicate.Parcel(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Parcel {
	return predicate.Parcel(sql.FieldLTE(FieldUpdatedAt, v))
}

// ParcelCodeEQ applies the EQ predicate on the "parcel_code" field.
func ParcelCodeEQ(v string) predicate.Parcel {
	return predicate.Parcel(sql.FieldEQ(FieldParcelCode, v))
}

// ParcelCodeNEQ applies the NEQ predicate on the "parcel_code" field.
func ParcelCodeNEQ(v string) predicate.Parcel {
	return predicate.Parcel(sql.FieldNEQ(FieldParcelCode, v))
}

// ParcelCodeIn applies the In predicate on the "parcel_code" field.
func ParcelCodeIn(vs ...string) predicate.Parcel {
	return predicate.Parcel(sql.FieldIn(FieldParcelCode, vs...))
}

// ParcelCodeNotIn applies the NotIn predicate on the "parcel_code" field.
func ParcelCodeNotIn(vs ...string) predicate.Parcel {
	return predicate.Parcel(sql.FieldNotIn(FieldParcelCode, vs...))
}

// ParcelCodeGT applies the GT predicate on the "parcel_code" field.
func ParcelCodeGT(v string) predicate.Parcel {
	return predicate.Parcel(sql.FieldGT(FieldParcelCode, v))
}

// ParcelCodeGTE applies the GTE predicate on the "parcel_code" field.
func ParcelCodeGTE(v string) predicate.Parcel {
	return predicate.Parcel(sql.FieldGTE(FieldParcelCode, v))
}

// ParcelCodeLT applies the LT predicate on the "parcel_code" field.
func ParcelCodeLT(v string) predicate.Parcel {
	return predicate.Parcel(sql.FieldLT(FieldParcelCode, v))
}

// ParcelCodeLTE applies the LTE predicate on the "parcel_code" field.
func ParcelCodeLTE(v string) predicate.Parcel {
	return predicate.Parcel(sql.FieldLTE(FieldParcelCode, v))
}

// ParcelCodeContains applies the Contains predicate on the "parcel_code" field.
func ParcelCodeContains(v string) predicate.Parcel {
	return predicate.Parcel(sql.FieldContains(FieldParcelCode, v))
}

// ParcelCodeHasPrefix applies the HasPrefix predicate on the "parcel_code" field.
func ParcelCodeHasPrefix(v string) predicate.Parcel {
	return predicate.Parcel(sql.FieldHasPrefix(FieldParcelCode, v))
}

// ParcelCodeHasSuffix applies the HasSuffix predicate on the "parcel_code" field.
func ParcelCodeHasSuffix(v string) predicate.Parcel {
	return predicate.Parcel(sql.FieldHasSuffix(FieldParcelCode, v))
}

// ParcelCodeEqualFold applies the EqualFold predicate on the "parcel_code" field.
func ParcelCodeEqualFold(v string) predicate.Parcel {
	return predicate.Parcel(sql.FieldEqualFold(FieldParcelCode, v))
}

// ParcelCodeContainsFold applies the ContainsFold predicate on the "parcel_code" field.
func ParcelCodeContainsFold(v string) predicate.Parcel {
	return predicate.Parcel(sql.FieldContainsFold(FieldParcelCode, v))
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v string) predicate.Parcel {
	return predicate.Parcel(sql.FieldEQ(FieldType, v))
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v string) predicate.Parcel {
	return predicate.Parcel(sql.FieldNEQ(FieldType, v))
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...string) predicate.Parcel {
	return predicate.Parcel(sql.FieldIn(FieldType, vs...))
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...string) predicate.Parcel {
	return predicate.Parcel(sql.FieldNotIn(FieldType, vs...))
}

// TypeGT applies the GT predicate on the "type" field.
func TypeGT(v string) predicate.Parcel {
	return predicate.Parcel(sql.FieldGT(FieldType, v))
}

// TypeGTE applies the GTE predicate on the "type" field.
func TypeGTE(v string) predicate.Parcel {
	return predicate.Parcel(sql.FieldGTE(FieldType, v))
}

// TypeLT applies the LT predicate on the "type" field.
func TypeLT(v string) predicate.Parcel {
	return predicate.Parcel(sql.FieldLT(FieldType, v))
}

// TypeLTE applies the LTE predicate on the "type" field.
func TypeLTE(v string) predicate.Parcel {
	return predicate.Parcel(sql.FieldLTE(FieldType, v))
}

// TypeContains applies the Contains predicate on the "type" field.
func TypeContains(v string) predicate.Parcel {
	return predicate.Parcel(sql.FieldContains(FieldType, v))
}

// TypeHasPrefix applies the HasPrefix predicate on the "type" field.
func TypeHasPrefix(v string) predicate.Parcel {
	return predicate.Parcel(sql.FieldHasPrefix(FieldType, v))
}

// TypeHasSuffix applies the HasSuffix predicate on the "type" field.
func TypeHasSuffix(v string) predicate.Parcel {
	return predicate.Parcel(sql.FieldHasSuffix(FieldType, v))
}

// TypeEqualFold applies the EqualFold predicate on the "type" field.
func TypeEqualFold(v string) predicate.Parcel {
	return predicate.Parcel(sql.FieldEqualFold(FieldType, v))
}

// TypeContainsFold applies the ContainsFold predicate on the "type" field.
func TypeContainsFold(v string) predicate.Parcel {
	return predicate.Parcel(sql.FieldContainsFold(FieldType, v))
}

// SenderNameEQ applies the EQ predicate on the "sender_name" field.
func SenderNameEQ(v string) predicate.Parcel {
	return predicate.Parcel(sql.FieldEQ(FieldSenderName, v))
}

// SenderNameNEQ applies the NEQ predicate on the "sender_name" field.
func SenderNameNEQ(v string) predicate.Parcel {
	return predicate.Parcel(sql.FieldNEQ(FieldSenderName, v))
}

// SenderNameIn applies the In predicate on the "sender_name" field.
func SenderNameIn(vs ...string) predicate.Parcel {
	return predicate.Parcel(sql.FieldIn(FieldSenderName, vs...))
}

// SenderNameNotIn applies the NotIn predicate on the "sender_name" field.
func SenderNameNotIn(vs ...string) predicate.Parcel {
	return predicate.Parcel(sql.FieldNotIn(FieldSenderName, vs...))
}

// SenderNameGT applies the GT predicate on the "sender_name" field.
func SenderNameGT(v string) predicate.Parcel {
	return predicate.Parcel(sql.FieldGT(FieldSenderName, v))
}

// SenderNameGTE applies the GTE predicate on the "sender_name" field.
func SenderNameGTE(v string) predicate.Parcel {
	return predicate.Parcel(sql.FieldGTE(FieldSenderName, v))
}

// SenderNameLT applies the LT predicate on the "sender_name" field.
func SenderNameLT(v string) predicate.Parcel {
	return predicate.Parcel(sql.FieldLT(FieldSenderName, v))
}

// SenderNameLTE applies the LTE predicate on the "sender_name" field.
func SenderNameLTE(v string) predicate.Parcel {
	return predicate.Parcel(sql.FieldLTE(FieldSenderName, v))
}

// SenderNameContains applies the Contains predicate on the "sender_name" field.
func SenderNameContains(v string) predicate.Parcel {
	return predicate.Parcel(sql.FieldContains(FieldSenderName, v))
}

// SenderNameHasPrefix applies the HasPrefix predicate on the "sender_name" field.
func SenderNameHasPrefix(v string) predicate.Parcel {
	return predicate.Parcel(sql.FieldHasPrefix(FieldSenderName, v))
}

// SenderNameHasSuffix applies the HasSuffix predicate on the "sender_name" field.
func SenderNameHasSuffix(v string) predicate.Parcel {
	return predicate.Parcel(sql.FieldHasSuffix(FieldSenderName, v))
}

// SenderNameEqualFold applies the EqualFold predicate on the "sender_name" field.
func SenderNameEqualFold(v string) predicate.Parcel {
	return predicate.Parcel(sql.FieldEqualFold(FieldSenderName, v))
}

// SenderNameContainsFold applies the ContainsFold predicate on the "sender_name" field.
func SenderNameContainsFold(v string) predicate.Parcel {
	return predicate.Parcel(sql.FieldContainsFold(FieldSenderName, v))
}

// SenderPhoneEQ applies the EQ predicate on the "sender_phone" field.
func SenderPhoneEQ(v string) predicate.Parcel {
	return predicate.Parcel(sql.FieldEQ(FieldSenderPhone, v))
}

// SenderPhoneNEQ applies the NEQ predicate on the "sender_phone" field.
func SenderPhoneNEQ(v string) predicate.Parcel {
	return predicate.Parcel(sql.FieldNEQ(FieldSenderPhone, v))
}

// SenderPhoneIn applies the In predicate on the "sender_phone" field.
func SenderPhoneIn(vs ...string) predicate.Parcel {
	return predicate.Parcel(sql.FieldIn(FieldSenderPhone, vs...))
}

// SenderPhoneNotIn applies the NotIn predicate on the "sender_phone" field.
func SenderPhoneNotIn(vs ...string) predicate.Parcel {
	return predicate.Parcel(sql.FieldNotIn(FieldSenderPhone, vs...))
}

// SenderPhoneGT applies the GT predicate on the "sender_phone" field.
func SenderPhoneGT(v string) predicate.Parcel {
	return predicate.Parcel(sql.FieldGT(FieldSenderPhone, v))
}

// SenderPhoneGTE applies the GTE predicate on the "sender_phone" field.
func SenderPhoneGTE(v string) predicate.Parcel {
	return predicate.Parcel(sql.FieldGTE(FieldSenderPhone, v))
}

// SenderPhoneLT applies the LT predicate on the "sender_phone" field.
func SenderPhoneLT(v string) predicate.Parcel {
	return predicate.Parcel(sql.FieldLT(FieldSenderPhone, v))
}

// SenderPhoneLTE applies the LTE predicate on the "sender_phone" field.
func SenderPhoneLTE(v string) predicate.Parcel {
	return predicate.Parcel(sql.FieldLTE(FieldSenderPhone, v))
}

// SenderPhoneContains applies the Contains predicate on the "sender_phone" field.
func SenderPhoneContains(v string) predicate.Parcel {
	return predicate.Parcel(sql.FieldContains(FieldSenderPhone, v))
}

// SenderPhoneHasPrefix applies the HasPrefix predicate on the "sender_phone" field.
func SenderPhoneHasPrefix(v string) predicate.Parcel {
	return predicate.Parcel(sql.FieldHasPrefix(FieldSenderPhone, v))
}

// SenderPhoneHasSuffix applies the HasSuffix predicate on the "sender_phone" field.
func SenderPhoneHasSuffix(v string) predicate.Parcel {
	return predicate.Parcel(sql.FieldHasSuffix(FieldSenderPhone, v))
}

// SenderPhoneEqualFold applies the EqualFold predicate on the "sender_phone" field.
func SenderPhoneEqualFold(v string) predicate.Parcel {
	return predicate.Parcel(sql.FieldEqualFold(FieldSenderPhone, v))
}

// SenderPhoneContainsFold applies the ContainsFold predicate on the "sender_phone" field.
func SenderPhoneContainsFold(v string) predicate.Parcel {
	return predicate.Parcel(sql.FieldContainsFold(FieldSenderPhone, v))
}

// SenderEmailEQ applies the EQ predicate on the "sender_email" field.
func SenderEmailEQ(v string) predicate.Parcel {
	return predicate.Parcel(sql.FieldEQ(FieldSenderEmail, v))
}

// SenderEmailNEQ applies the NEQ predicate on the "sender_email" field.
func SenderEmailNEQ(v string) predicate.Parcel {
	return predicate.Parcel(sql.FieldNEQ(FieldSenderEmail, v))
}

// SenderEmailIn applies the In predicate on the "sender_email" field.
func SenderEmailIn(vs ...string) predicate.Parcel {
	return predicate.Parcel(sql.FieldIn(FieldSenderEmail, vs...))
}

// SenderEmailNotIn applies the NotIn predicate on the "sender_email" field.
func SenderEmailNotIn(vs ...string) predicate.Parcel {
	return predicate.Parcel(sql.FieldNotIn(FieldSenderEmail, vs...))
}

// SenderEmailGT applies the GT predicate on the "sender_email" field.
func SenderEmailGT(v string) predicate.Parcel {
	return predicate.Parcel(sql.FieldGT(FieldSenderEmail, v))
}

// SenderEmailGTE applies the GTE predicate on the "sender_email" field.
func SenderEmailGTE(v string) predicate.Parcel {
	return predicate.Parcel(sql.FieldGTE(FieldSenderEmail, v))
}

// SenderEmailLT applies the LT predicate on the "sender_email" field.
func SenderEmailLT(v string) predicate.Parcel {
	return predicate.Parcel(sql.FieldLT(FieldSenderEmail, v))
}

// SenderEmailLTE applies the LTE predicate on the "sender_email" field.
func SenderEmailLTE(v string) predicate.Parcel {
	return predicate.Parcel(sql.FieldLTE(FieldSenderEmail, v))
}

// SenderEmailContains applies the Contains predicate on the "sender_email" field.
func SenderEmailContains(v string) predicate.Parcel {
	return predicate.Parcel(sql.FieldContains(FieldSenderEmail, v))
}

// SenderEmailHasPrefix applies the HasPrefix predicate on the "sender_email" field.
func SenderEmailHasPrefix(v string) predicate.Parcel {
	return predicate.Parcel(sql.FieldHasPrefix(FieldSenderEmail, v))
}

// SenderEmailHasSuffix applies the HasSuffix predicate on the "sender_email" field.
func SenderEmailHasSuffix(v string) predicate.Parcel {
	return predicate.Parcel(sql.FieldHasSuffix(FieldSenderEmail, v))
}

// SenderEmailEqualFold applies the EqualFold predicate on the "sender_email" field.
func SenderEmailEqualFold(v string) predicate.Parcel {
	return predicate.Parcel(sql.FieldEqualFold(FieldSenderEmail, v))
}

// SenderEmailContainsFold applies the ContainsFold predicate on the "sender_email" field.
func SenderEmailContainsFold(v string) predicate.Parcel {
	return predicate.Parcel(sql.FieldContainsFold(FieldSenderEmail, v))
}

// RecipientNameEQ applies the EQ predicate on the "recipient_name" field.
func RecipientNameEQ(v string) predicate.Parcel {
	return predicate.Parcel(sql.FieldEQ(FieldRecipientName, v))
}

// RecipientNameNEQ applies the NEQ predicate on the "recipient_name" field.
func RecipientNameNEQ(v string) predicate.Parcel {
	return predicate.Parcel(sql.FieldNEQ(FieldRecipientName, v))
}

// RecipientNameIn applies the In predicate on the "recipient_name" field.
func RecipientNameIn(vs ...string) predicate.Parcel {
	return predicate.Parcel(sql.FieldIn(FieldRecipientName, vs...))
}

// RecipientNameNotIn applies the NotIn predicate on the "recipient_name" field.
func RecipientNameNotIn(vs ...string) predicate.Parcel {
	return predicate.Parcel(sql.FieldNotIn(FieldRecipientName, vs...))
}

// RecipientNameGT applies the GT predicate on the "recipient_name" field.
func RecipientNameGT(v string) predicate.Parcel {
	return predicate.Parcel(sql.FieldGT(FieldRecipientName, v))
}

// RecipientNameGTE applies the GTE predicate on the "recipient_name" field.
func RecipientNameGTE(v string) predicate.Parcel {
	return predicate.Parcel(sql.FieldGTE(FieldRecipientName, v))
}

// RecipientNameLT applies the LT predicate on the "recipient_name" field.
func RecipientNameLT(v string) predicate.Parcel {
	return predicate.Parcel(sql.FieldLT(FieldRecipientName, v))
}

// RecipientNameLTE applies the LTE predicate on the "recipient_name" field.
func RecipientNameLTE(v string) predicate.Parcel {
	return predicate.Parcel(sql.FieldLTE(FieldRecipientName, v))
}

// RecipientNameContains applies the Contains predicate on the "recipient_name" field.
func RecipientNameContains(v string) predicate.Parcel {
	return predicate.Parcel(sql.FieldContains(FieldRecipientName, v))
}

// RecipientNameHasPrefix applies the HasPrefix predicate on the "recipient_name" field.
func RecipientNameHasPrefix(v string) predicate.Parcel {
	return predicate.Parcel(sql.FieldHasPrefix(FieldRecipientName, v))
}

// RecipientNameHasSuffix applies the HasSuffix predicate on the "recipient_name" field.
func RecipientNameHasSuffix(v string) predicate.Parcel {
	return predicate.Parcel(sql.FieldHasSuffix(FieldRecipientName, v))
}

// RecipientNameEqualFold applies the EqualFold predicate on the "recipient_name" field.
func RecipientNameEqualFold(v string) predicate.Parcel {
	return predicate.Parcel(sql.FieldEqualFold(FieldRecipientName, v))
}

// RecipientNameContainsFold applies the ContainsFold predicate on the "recipient_name" field.
func RecipientNameContainsFold(v string) predicate.Parcel {
	return predicate.Parcel(sql.FieldContainsFold(FieldRecipientName, v))
}

// RecipientPhoneEQ applies the EQ predicate on the "recipient_phone" field.
func RecipientPhoneEQ(v string) predicate.Parcel {
	return predicate.Parcel(sql.FieldEQ(FieldRecipientPhone, v))
}

// RecipientPhoneNEQ applies the NEQ predicate on the "recipient_phone" field.
func RecipientPhoneNEQ(v string) predicate.Parcel {
	return predicate.Parcel(sql.FieldNEQ(FieldRecipientPhone, v))
}

// RecipientPhoneIn applies the In predicate on the "recipient_phone" field.
func RecipientPhoneIn(vs ...string) predicate.Parcel {
	return predicate.Parcel(sql.FieldIn(FieldRecipientPhone, vs...))
}

// RecipientPhoneNotIn applies the NotIn predicate on the "recipient_phone" field.
func RecipientPhoneNotIn(vs ...string) predicate.Parcel {
	return predicate.Parcel(sql.FieldNotIn(FieldRecipientPhone, vs...))
}

// RecipientPhoneGT applies the GT predicate on the "recipient_phone" field.
func RecipientPhoneGT(v string) predicate.Parcel {
	return predicate.Parcel(sql.FieldGT(FieldRecipientPhone, v))
}

// RecipientPhoneGTE applies the GTE predicate on the "recipient_phone" field.
func RecipientPhoneGTE(v string) predicate.Parcel {
	return predicate.Parcel(sql.FieldGTE(FieldRecipientPhone, v))
}

// RecipientPhoneLT applies the LT predicate on the "recipient_phone" field.
func RecipientPhoneLT(v string) predicate.Parcel {
	return predicate.Parcel(sql.FieldLT(FieldRecipientPhone, v))
}

// RecipientPhoneLTE applies the LTE predicate on the "recipient_phone" field.
func RecipientPhoneLTE(v string) predicate.Parcel {
	return predicate.Parcel(sql.FieldLTE(FieldRecipientPhone, v))
}

// RecipientPhoneContains applies the Contains predicate on the "recipient_phone" field.
func RecipientPhoneContains(v string) predicate.Parcel {
	return predicate.Parcel(sql.FieldContains(FieldRecipientPhone, v))
}

// RecipientPhoneHasPrefix applies the HasPrefix predicate on the "recipient_phone" field.
func RecipientPhoneHasPrefix(v string) predicate.Parcel {
	return predicate.Parcel(sql.FieldHasPrefix(FieldRecipientPhone, v))
}

// RecipientPhoneHasSuffix applies the HasSuffix predicate on the "recipient_phone" field.
func RecipientPhoneHasSuffix(v string) predicate.Parcel {
	return predicate.Parcel(sql.FieldHasSuffix(FieldRecipientPhone, v))
}

// RecipientPhoneEqualFold applies the EqualFold predicate on the "recipient_phone" field.
func RecipientPhoneEqualFold(v string) predicate.Parcel {
	return predicate.Parcel(sql.FieldEqualFold(FieldRecipientPhone, v))
}

// RecipientPhoneContainsFold applies the ContainsFold predicate on the "recipient_phone" field.
func RecipientPhoneContainsFold(v string) predicate.Parcel {
	return predicate.Parcel(sql.FieldContainsFold(FieldRecipientPhone, v))
}

// RecipientLocationEQ applies the EQ predicate on the "recipient_location" field.
func RecipientLocationEQ(v string) predicate.Parcel {
	return predicate.Parcel(sql.FieldEQ(FieldRecipientLocation, v))
}

// RecipientLocationNEQ applies the NEQ predicate on the "recipient_location" field.
func RecipientLocationNEQ(v string) predicate.Parcel {
	return predicate.Parcel(sql.FieldNEQ(FieldRecipientLocation, v))
}

// RecipientLocationIn applies the In predicate on the "recipient_location" field.
func RecipientLocationIn(vs ...string) predicate.Parcel {
	return predicate.Parcel(sql.FieldIn(FieldRecipientLocation, vs...))
}

// RecipientLocationNotIn applies the NotIn predicate on the "recipient_location" field.
func RecipientLocationNotIn(vs ...string) predicate.Parcel {
	return predicate.Parcel(sql.FieldNotIn(FieldRecipientLocation, vs...))
}

// RecipientLocationGT applies the GT predicate on the "recipient_location" field.
func RecipientLocationGT(v string) predicate.Parcel {
	return predicate.Parcel(sql.FieldGT(FieldRecipientLocation, v))
}

// RecipientLocationGTE applies the GTE predicate on the "recipient_location" field.
func RecipientLocationGTE(v string) predicate.Parcel {
	return predicate.Parcel(sql.FieldGTE(FieldRecipientLocation, v))
}

// RecipientLocationLT applies the LT predicate on the "recipient_location" field.
func RecipientLocationLT(v string) predicate.Parcel {
	return predicate.Parcel(sql.FieldLT(FieldRecipientLocation, v))
}

// RecipientLocationLTE applies the LTE predicate on the "recipient_location" field.
func RecipientLocationLTE(v string) predicate.Parcel {
	return predicate.Parcel(sql.FieldLTE(FieldRecipientLocation, v))
}

// RecipientLocationContains applies the Contains predicate on the "recipient_location" field.
func RecipientLocationContains(v string) predicate.Parcel {
	return predicate.Parcel(sql.FieldContains(FieldRecipientLocation, v))
}

// RecipientLocationHasPrefix applies the HasPrefix predicate on the "recipient_location" field.
func RecipientLocationHasPrefix(v string) predicate.Parcel {
	return predicate.Parcel(sql.FieldHasPrefix(FieldRecipientLocation, v))
}

// RecipientLocationHasSuffix applies the HasSuffix predicate on the "recipient_location" field.
func RecipientLocationHasSuffix(v string) predicate.Parcel {
	return predicate.Parcel(sql.FieldHasSuffix(FieldRecipientLocation, v))
}

// RecipientLocationEqualFold applies the EqualFold predicate on the "recipient_location" field.
func RecipientLocationEqualFold(v string) predicate.Parcel {
	return predicate.Parcel(sql.FieldEqualFold(FieldRecipientLocation, v))
}

// RecipientLocationContainsFold applies the ContainsFold predicate on the "recipient_location" field.
func RecipientLocationContainsFold(v string) predicate.Parcel {
	return predicate.Parcel(sql.FieldContainsFold(FieldRecipientLocation, v))
}

// WeightEQ applies the EQ predicate on the "weight" field.
func WeightEQ(v float32) predicate.Parcel {
	return predicate.Parcel(sql.FieldEQ(FieldWeight, v))
}

// WeightNEQ applies the NEQ predicate on the "weight" field.
func WeightNEQ(v float32) predicate.Parcel {
	return predicate.Parcel(sql.FieldNEQ(FieldWeight, v))
}

// WeightIn applies the In predicate on the "weight" field.
func WeightIn(vs ...float32) predicate.Parcel {
	return predicate.Parcel(sql.FieldIn(FieldWeight, vs...))
}

// WeightNotIn applies the NotIn predicate on the "weight" field.
func WeightNotIn(vs ...float32) predicate.Parcel {
	return predicate.Parcel(sql.FieldNotIn(FieldWeight, vs...))
}

// WeightGT applies the GT predicate on the "weight" field.
func WeightGT(v float32) predicate.Parcel {
	return predicate.Parcel(sql.FieldGT(FieldWeight, v))
}

// WeightGTE applies the GTE predicate on the "weight" field.
func WeightGTE(v float32) predicate.Parcel {
	return predicate.Parcel(sql.FieldGTE(FieldWeight, v))
}

// WeightLT applies the LT predicate on the "weight" field.
func WeightLT(v float32) predicate.Parcel {
	return predicate.Parcel(sql.FieldLT(FieldWeight, v))
}

// WeightLTE applies the LTE predicate on the "weight" field.
func WeightLTE(v float32) predicate.Parcel {
	return predicate.Parcel(sql.FieldLTE(FieldWeight, v))
}

// WeightIsNil applies the IsNil predicate on the "weight" field.
func WeightIsNil() predicate.Parcel {
	return predicate.Parcel(sql.FieldIsNull(FieldWeight))
}

// WeightNotNil applies the NotNil predicate on the "weight" field.
func WeightNotNil() predicate.Parcel {
	return predicate.Parcel(sql.FieldNotNull(FieldWeight))
}

// AmountEQ applies the EQ predicate on the "amount" field.
func AmountEQ(v float64) predicate.Parcel {
	return predicate.Parcel(sql.FieldEQ(FieldAmount, v))
}

// AmountNEQ applies the NEQ predicate on the "amount" field.
func AmountNEQ(v float64) predicate.Parcel {
	return predicate.Parcel(sql.FieldNEQ(FieldAmount, v))
}

// AmountIn applies the In predicate on the "amount" field.
func AmountIn(vs ...float64) predicate.Parcel {
	return predicate.Parcel(sql.FieldIn(FieldAmount, vs...))
}

// AmountNotIn applies the NotIn predicate on the "amount" field.
func AmountNotIn(vs ...float64) predicate.Parcel {
	return predicate.Parcel(sql.FieldNotIn(FieldAmount, vs...))
}

// AmountGT applies the GT predicate on the "amount" field.
func AmountGT(v float64) predicate.Parcel {
	return predicate.Parcel(sql.FieldGT(FieldAmount, v))
}

// AmountGTE applies the GTE predicate on the "amount" field.
func AmountGTE(v float64) predicate.Parcel {
	return predicate.Parcel(sql.FieldGTE(FieldAmount, v))
}

// AmountLT applies the LT predicate on the "amount" field.
func AmountLT(v float64) predicate.Parcel {
	return predicate.Parcel(sql.FieldLT(FieldAmount, v))
}

// AmountLTE applies the LTE predicate on the "amount" field.
func AmountLTE(v float64) predicate.Parcel {
	return predicate.Parcel(sql.FieldLTE(FieldAmount, v))
}

// PaidAtEQ applies the EQ predicate on the "paid_at" field.
func PaidAtEQ(v time.Time) predicate.Parcel {
	return predicate.Parcel(sql.FieldEQ(FieldPaidAt, v))
}

// PaidAtNEQ applies the NEQ predicate on the "paid_at" field.
func PaidAtNEQ(v time.Time) predicate.Parcel {
	return predicate.Parcel(sql.FieldNEQ(FieldPaidAt, v))
}

// PaidAtIn applies the In predicate on the "paid_at" field.
func PaidAtIn(vs ...time.Time) predicate.Parcel {
	return predicate.Parcel(sql.FieldIn(FieldPaidAt, vs...))
}

// PaidAtNotIn applies the NotIn predicate on the "paid_at" field.
func PaidAtNotIn(vs ...time.Time) predicate.Parcel {
	return predicate.Parcel(sql.FieldNotIn(FieldPaidAt, vs...))
}

// PaidAtGT applies the GT predicate on the "paid_at" field.
func PaidAtGT(v time.Time) predicate.Parcel {
	return predicate.Parcel(sql.FieldGT(FieldPaidAt, v))
}

// PaidAtGTE applies the GTE predicate on the "paid_at" field.
func PaidAtGTE(v time.Time) predicate.Parcel {
	return predicate.Parcel(sql.FieldGTE(FieldPaidAt, v))
}

// PaidAtLT applies the LT predicate on the "paid_at" field.
func PaidAtLT(v time.Time) predicate.Parcel {
	return predicate.Parcel(sql.FieldLT(FieldPaidAt, v))
}

// PaidAtLTE applies the LTE predicate on the "paid_at" field.
func PaidAtLTE(v time.Time) predicate.Parcel {
	return predicate.Parcel(sql.FieldLTE(FieldPaidAt, v))
}

// PaidAtIsNil applies the IsNil predicate on the "paid_at" field.
func PaidAtIsNil() predicate.Parcel {
	return predicate.Parcel(sql.FieldIsNull(FieldPaidAt))
}

// PaidAtNotNil applies the NotNil predicate on the "paid_at" field.
func PaidAtNotNil() predicate.Parcel {
	return predicate.Parcel(sql.FieldNotNull(FieldPaidAt))
}

// TansTypeEQ applies the EQ predicate on the "tans_type" field.
func TansTypeEQ(v TansType) predicate.Parcel {
	return predicate.Parcel(sql.FieldEQ(FieldTansType, v))
}

// TansTypeNEQ applies the NEQ predicate on the "tans_type" field.
func TansTypeNEQ(v TansType) predicate.Parcel {
	return predicate.Parcel(sql.FieldNEQ(FieldTansType, v))
}

// TansTypeIn applies the In predicate on the "tans_type" field.
func TansTypeIn(vs ...TansType) predicate.Parcel {
	return predicate.Parcel(sql.FieldIn(FieldTansType, vs...))
}

// TansTypeNotIn applies the NotIn predicate on the "tans_type" field.
func TansTypeNotIn(vs ...TansType) predicate.Parcel {
	return predicate.Parcel(sql.FieldNotIn(FieldTansType, vs...))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v Status) predicate.Parcel {
	return predicate.Parcel(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v Status) predicate.Parcel {
	return predicate.Parcel(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...Status) predicate.Parcel {
	return predicate.Parcel(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...Status) predicate.Parcel {
	return predicate.Parcel(sql.FieldNotIn(FieldStatus, vs...))
}

// HasImages applies the HasEdge predicate on the "images" edge.
func HasImages() predicate.Parcel {
	return predicate.Parcel(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ImagesTable, ImagesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasImagesWith applies the HasEdge predicate on the "images" edge with a given conditions (other predicates).
func HasImagesWith(preds ...predicate.ParcelImage) predicate.Parcel {
	return predicate.Parcel(func(s *sql.Selector) {
		step := newImagesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasTrip applies the HasEdge predicate on the "trip" edge.
func HasTrip() predicate.Parcel {
	return predicate.Parcel(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, TripTable, TripColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTripWith applies the HasEdge predicate on the "trip" edge with a given conditions (other predicates).
func HasTripWith(preds ...predicate.Trip) predicate.Parcel {
	return predicate.Parcel(func(s *sql.Selector) {
		step := newTripStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasCompany applies the HasEdge predicate on the "company" edge.
func HasCompany() predicate.Parcel {
	return predicate.Parcel(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, CompanyTable, CompanyColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCompanyWith applies the HasEdge predicate on the "company" edge with a given conditions (other predicates).
func HasCompanyWith(preds ...predicate.Company) predicate.Parcel {
	return predicate.Parcel(func(s *sql.Selector) {
		step := newCompanyStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasDriver applies the HasEdge predicate on the "driver" edge.
func HasDriver() predicate.Parcel {
	return predicate.Parcel(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, DriverTable, DriverColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasDriverWith applies the HasEdge predicate on the "driver" edge with a given conditions (other predicates).
func HasDriverWith(preds ...predicate.CompanyUser) predicate.Parcel {
	return predicate.Parcel(func(s *sql.Selector) {
		step := newDriverStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Parcel) predicate.Parcel {
	return predicate.Parcel(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Parcel) predicate.Parcel {
	return predicate.Parcel(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Parcel) predicate.Parcel {
	return predicate.Parcel(sql.NotPredicates(p))
}
