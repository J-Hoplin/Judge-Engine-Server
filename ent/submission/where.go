// Code generated by ent, DO NOT EDIT.

package submission

import (
	"judge-engine/ent/predicate"
	"judge-engine/ent/schema/enum"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Submission {
	return predicate.Submission(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Submission {
	return predicate.Submission(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Submission {
	return predicate.Submission(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Submission {
	return predicate.Submission(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Submission {
	return predicate.Submission(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Submission {
	return predicate.Submission(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Submission {
	return predicate.Submission(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Submission {
	return predicate.Submission(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Submission {
	return predicate.Submission(sql.FieldLTE(FieldID, id))
}

// Code applies equality check predicate on the "code" field. It's identical to CodeEQ.
func Code(v string) predicate.Submission {
	return predicate.Submission(sql.FieldEQ(FieldCode, v))
}

// CodeLength applies equality check predicate on the "codeLength" field. It's identical to CodeLengthEQ.
func CodeLength(v int) predicate.Submission {
	return predicate.Submission(sql.FieldEQ(FieldCodeLength, v))
}

// Memory applies equality check predicate on the "memory" field. It's identical to MemoryEQ.
func Memory(v int) predicate.Submission {
	return predicate.Submission(sql.FieldEQ(FieldMemory, v))
}

// CreatedAt applies equality check predicate on the "createdAt" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Submission {
	return predicate.Submission(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updatedAt" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Submission {
	return predicate.Submission(sql.FieldEQ(FieldUpdatedAt, v))
}

// CodeEQ applies the EQ predicate on the "code" field.
func CodeEQ(v string) predicate.Submission {
	return predicate.Submission(sql.FieldEQ(FieldCode, v))
}

// CodeNEQ applies the NEQ predicate on the "code" field.
func CodeNEQ(v string) predicate.Submission {
	return predicate.Submission(sql.FieldNEQ(FieldCode, v))
}

// CodeIn applies the In predicate on the "code" field.
func CodeIn(vs ...string) predicate.Submission {
	return predicate.Submission(sql.FieldIn(FieldCode, vs...))
}

// CodeNotIn applies the NotIn predicate on the "code" field.
func CodeNotIn(vs ...string) predicate.Submission {
	return predicate.Submission(sql.FieldNotIn(FieldCode, vs...))
}

// CodeGT applies the GT predicate on the "code" field.
func CodeGT(v string) predicate.Submission {
	return predicate.Submission(sql.FieldGT(FieldCode, v))
}

// CodeGTE applies the GTE predicate on the "code" field.
func CodeGTE(v string) predicate.Submission {
	return predicate.Submission(sql.FieldGTE(FieldCode, v))
}

// CodeLT applies the LT predicate on the "code" field.
func CodeLT(v string) predicate.Submission {
	return predicate.Submission(sql.FieldLT(FieldCode, v))
}

// CodeLTE applies the LTE predicate on the "code" field.
func CodeLTE(v string) predicate.Submission {
	return predicate.Submission(sql.FieldLTE(FieldCode, v))
}

// CodeContains applies the Contains predicate on the "code" field.
func CodeContains(v string) predicate.Submission {
	return predicate.Submission(sql.FieldContains(FieldCode, v))
}

// CodeHasPrefix applies the HasPrefix predicate on the "code" field.
func CodeHasPrefix(v string) predicate.Submission {
	return predicate.Submission(sql.FieldHasPrefix(FieldCode, v))
}

// CodeHasSuffix applies the HasSuffix predicate on the "code" field.
func CodeHasSuffix(v string) predicate.Submission {
	return predicate.Submission(sql.FieldHasSuffix(FieldCode, v))
}

// CodeEqualFold applies the EqualFold predicate on the "code" field.
func CodeEqualFold(v string) predicate.Submission {
	return predicate.Submission(sql.FieldEqualFold(FieldCode, v))
}

// CodeContainsFold applies the ContainsFold predicate on the "code" field.
func CodeContainsFold(v string) predicate.Submission {
	return predicate.Submission(sql.FieldContainsFold(FieldCode, v))
}

// CodeLengthEQ applies the EQ predicate on the "codeLength" field.
func CodeLengthEQ(v int) predicate.Submission {
	return predicate.Submission(sql.FieldEQ(FieldCodeLength, v))
}

// CodeLengthNEQ applies the NEQ predicate on the "codeLength" field.
func CodeLengthNEQ(v int) predicate.Submission {
	return predicate.Submission(sql.FieldNEQ(FieldCodeLength, v))
}

// CodeLengthIn applies the In predicate on the "codeLength" field.
func CodeLengthIn(vs ...int) predicate.Submission {
	return predicate.Submission(sql.FieldIn(FieldCodeLength, vs...))
}

// CodeLengthNotIn applies the NotIn predicate on the "codeLength" field.
func CodeLengthNotIn(vs ...int) predicate.Submission {
	return predicate.Submission(sql.FieldNotIn(FieldCodeLength, vs...))
}

// CodeLengthGT applies the GT predicate on the "codeLength" field.
func CodeLengthGT(v int) predicate.Submission {
	return predicate.Submission(sql.FieldGT(FieldCodeLength, v))
}

// CodeLengthGTE applies the GTE predicate on the "codeLength" field.
func CodeLengthGTE(v int) predicate.Submission {
	return predicate.Submission(sql.FieldGTE(FieldCodeLength, v))
}

// CodeLengthLT applies the LT predicate on the "codeLength" field.
func CodeLengthLT(v int) predicate.Submission {
	return predicate.Submission(sql.FieldLT(FieldCodeLength, v))
}

// CodeLengthLTE applies the LTE predicate on the "codeLength" field.
func CodeLengthLTE(v int) predicate.Submission {
	return predicate.Submission(sql.FieldLTE(FieldCodeLength, v))
}

// MemoryEQ applies the EQ predicate on the "memory" field.
func MemoryEQ(v int) predicate.Submission {
	return predicate.Submission(sql.FieldEQ(FieldMemory, v))
}

// MemoryNEQ applies the NEQ predicate on the "memory" field.
func MemoryNEQ(v int) predicate.Submission {
	return predicate.Submission(sql.FieldNEQ(FieldMemory, v))
}

// MemoryIn applies the In predicate on the "memory" field.
func MemoryIn(vs ...int) predicate.Submission {
	return predicate.Submission(sql.FieldIn(FieldMemory, vs...))
}

// MemoryNotIn applies the NotIn predicate on the "memory" field.
func MemoryNotIn(vs ...int) predicate.Submission {
	return predicate.Submission(sql.FieldNotIn(FieldMemory, vs...))
}

// MemoryGT applies the GT predicate on the "memory" field.
func MemoryGT(v int) predicate.Submission {
	return predicate.Submission(sql.FieldGT(FieldMemory, v))
}

// MemoryGTE applies the GTE predicate on the "memory" field.
func MemoryGTE(v int) predicate.Submission {
	return predicate.Submission(sql.FieldGTE(FieldMemory, v))
}

// MemoryLT applies the LT predicate on the "memory" field.
func MemoryLT(v int) predicate.Submission {
	return predicate.Submission(sql.FieldLT(FieldMemory, v))
}

// MemoryLTE applies the LTE predicate on the "memory" field.
func MemoryLTE(v int) predicate.Submission {
	return predicate.Submission(sql.FieldLTE(FieldMemory, v))
}

// ResponseEQ applies the EQ predicate on the "response" field.
func ResponseEQ(v enum.ResponseType) predicate.Submission {
	vc := v
	return predicate.Submission(sql.FieldEQ(FieldResponse, vc))
}

// ResponseNEQ applies the NEQ predicate on the "response" field.
func ResponseNEQ(v enum.ResponseType) predicate.Submission {
	vc := v
	return predicate.Submission(sql.FieldNEQ(FieldResponse, vc))
}

// ResponseIn applies the In predicate on the "response" field.
func ResponseIn(vs ...enum.ResponseType) predicate.Submission {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Submission(sql.FieldIn(FieldResponse, v...))
}

// ResponseNotIn applies the NotIn predicate on the "response" field.
func ResponseNotIn(vs ...enum.ResponseType) predicate.Submission {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Submission(sql.FieldNotIn(FieldResponse, v...))
}

// CreatedAtEQ applies the EQ predicate on the "createdAt" field.
func CreatedAtEQ(v time.Time) predicate.Submission {
	return predicate.Submission(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "createdAt" field.
func CreatedAtNEQ(v time.Time) predicate.Submission {
	return predicate.Submission(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "createdAt" field.
func CreatedAtIn(vs ...time.Time) predicate.Submission {
	return predicate.Submission(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "createdAt" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Submission {
	return predicate.Submission(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "createdAt" field.
func CreatedAtGT(v time.Time) predicate.Submission {
	return predicate.Submission(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "createdAt" field.
func CreatedAtGTE(v time.Time) predicate.Submission {
	return predicate.Submission(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "createdAt" field.
func CreatedAtLT(v time.Time) predicate.Submission {
	return predicate.Submission(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "createdAt" field.
func CreatedAtLTE(v time.Time) predicate.Submission {
	return predicate.Submission(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updatedAt" field.
func UpdatedAtEQ(v time.Time) predicate.Submission {
	return predicate.Submission(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updatedAt" field.
func UpdatedAtNEQ(v time.Time) predicate.Submission {
	return predicate.Submission(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updatedAt" field.
func UpdatedAtIn(vs ...time.Time) predicate.Submission {
	return predicate.Submission(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updatedAt" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Submission {
	return predicate.Submission(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updatedAt" field.
func UpdatedAtGT(v time.Time) predicate.Submission {
	return predicate.Submission(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updatedAt" field.
func UpdatedAtGTE(v time.Time) predicate.Submission {
	return predicate.Submission(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updatedAt" field.
func UpdatedAtLT(v time.Time) predicate.Submission {
	return predicate.Submission(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updatedAt" field.
func UpdatedAtLTE(v time.Time) predicate.Submission {
	return predicate.Submission(sql.FieldLTE(FieldUpdatedAt, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Submission) predicate.Submission {
	return predicate.Submission(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Submission) predicate.Submission {
	return predicate.Submission(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Submission) predicate.Submission {
	return predicate.Submission(sql.NotPredicates(p))
}
