// Code generated by entc, DO NOT EDIT.

package aircraft

import (
	"airbound/internal/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Aircraft {
	return predicate.Aircraft(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Aircraft {
	return predicate.Aircraft(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Aircraft {
	return predicate.Aircraft(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Aircraft {
	return predicate.Aircraft(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Aircraft {
	return predicate.Aircraft(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Aircraft {
	return predicate.Aircraft(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Aircraft {
	return predicate.Aircraft(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Aircraft {
	return predicate.Aircraft(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Aircraft {
	return predicate.Aircraft(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// TailNumber applies equality check predicate on the "tail_number" field. It's identical to TailNumberEQ.
func TailNumber(v string) predicate.Aircraft {
	return predicate.Aircraft(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTailNumber), v))
	})
}

// Manufacturer applies equality check predicate on the "manufacturer" field. It's identical to ManufacturerEQ.
func Manufacturer(v string) predicate.Aircraft {
	return predicate.Aircraft(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldManufacturer), v))
	})
}

// Model applies equality check predicate on the "model" field. It's identical to ModelEQ.
func Model(v string) predicate.Aircraft {
	return predicate.Aircraft(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldModel), v))
	})
}

// Capacity applies equality check predicate on the "capacity" field. It's identical to CapacityEQ.
func Capacity(v int) predicate.Aircraft {
	return predicate.Aircraft(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCapacity), v))
	})
}

// Range applies equality check predicate on the "range" field. It's identical to RangeEQ.
func Range(v int) predicate.Aircraft {
	return predicate.Aircraft(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRange), v))
	})
}

// ManufacturedAt applies equality check predicate on the "manufactured_at" field. It's identical to ManufacturedAtEQ.
func ManufacturedAt(v time.Time) predicate.Aircraft {
	return predicate.Aircraft(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldManufacturedAt), v))
	})
}

// IsGrounded applies equality check predicate on the "is_grounded" field. It's identical to IsGroundedEQ.
func IsGrounded(v bool) predicate.Aircraft {
	return predicate.Aircraft(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIsGrounded), v))
	})
}

// GroundedAt applies equality check predicate on the "grounded_at" field. It's identical to GroundedAtEQ.
func GroundedAt(v time.Time) predicate.Aircraft {
	return predicate.Aircraft(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldGroundedAt), v))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Aircraft {
	return predicate.Aircraft(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Aircraft {
	return predicate.Aircraft(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// TailNumberEQ applies the EQ predicate on the "tail_number" field.
func TailNumberEQ(v string) predicate.Aircraft {
	return predicate.Aircraft(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTailNumber), v))
	})
}

// TailNumberNEQ applies the NEQ predicate on the "tail_number" field.
func TailNumberNEQ(v string) predicate.Aircraft {
	return predicate.Aircraft(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTailNumber), v))
	})
}

// TailNumberIn applies the In predicate on the "tail_number" field.
func TailNumberIn(vs ...string) predicate.Aircraft {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Aircraft(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldTailNumber), v...))
	})
}

// TailNumberNotIn applies the NotIn predicate on the "tail_number" field.
func TailNumberNotIn(vs ...string) predicate.Aircraft {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Aircraft(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldTailNumber), v...))
	})
}

// TailNumberGT applies the GT predicate on the "tail_number" field.
func TailNumberGT(v string) predicate.Aircraft {
	return predicate.Aircraft(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTailNumber), v))
	})
}

// TailNumberGTE applies the GTE predicate on the "tail_number" field.
func TailNumberGTE(v string) predicate.Aircraft {
	return predicate.Aircraft(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTailNumber), v))
	})
}

// TailNumberLT applies the LT predicate on the "tail_number" field.
func TailNumberLT(v string) predicate.Aircraft {
	return predicate.Aircraft(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTailNumber), v))
	})
}

// TailNumberLTE applies the LTE predicate on the "tail_number" field.
func TailNumberLTE(v string) predicate.Aircraft {
	return predicate.Aircraft(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTailNumber), v))
	})
}

// TailNumberContains applies the Contains predicate on the "tail_number" field.
func TailNumberContains(v string) predicate.Aircraft {
	return predicate.Aircraft(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTailNumber), v))
	})
}

// TailNumberHasPrefix applies the HasPrefix predicate on the "tail_number" field.
func TailNumberHasPrefix(v string) predicate.Aircraft {
	return predicate.Aircraft(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTailNumber), v))
	})
}

// TailNumberHasSuffix applies the HasSuffix predicate on the "tail_number" field.
func TailNumberHasSuffix(v string) predicate.Aircraft {
	return predicate.Aircraft(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTailNumber), v))
	})
}

// TailNumberEqualFold applies the EqualFold predicate on the "tail_number" field.
func TailNumberEqualFold(v string) predicate.Aircraft {
	return predicate.Aircraft(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTailNumber), v))
	})
}

// TailNumberContainsFold applies the ContainsFold predicate on the "tail_number" field.
func TailNumberContainsFold(v string) predicate.Aircraft {
	return predicate.Aircraft(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTailNumber), v))
	})
}

// ManufacturerEQ applies the EQ predicate on the "manufacturer" field.
func ManufacturerEQ(v string) predicate.Aircraft {
	return predicate.Aircraft(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldManufacturer), v))
	})
}

// ManufacturerNEQ applies the NEQ predicate on the "manufacturer" field.
func ManufacturerNEQ(v string) predicate.Aircraft {
	return predicate.Aircraft(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldManufacturer), v))
	})
}

// ManufacturerIn applies the In predicate on the "manufacturer" field.
func ManufacturerIn(vs ...string) predicate.Aircraft {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Aircraft(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldManufacturer), v...))
	})
}

// ManufacturerNotIn applies the NotIn predicate on the "manufacturer" field.
func ManufacturerNotIn(vs ...string) predicate.Aircraft {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Aircraft(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldManufacturer), v...))
	})
}

// ManufacturerGT applies the GT predicate on the "manufacturer" field.
func ManufacturerGT(v string) predicate.Aircraft {
	return predicate.Aircraft(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldManufacturer), v))
	})
}

// ManufacturerGTE applies the GTE predicate on the "manufacturer" field.
func ManufacturerGTE(v string) predicate.Aircraft {
	return predicate.Aircraft(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldManufacturer), v))
	})
}

// ManufacturerLT applies the LT predicate on the "manufacturer" field.
func ManufacturerLT(v string) predicate.Aircraft {
	return predicate.Aircraft(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldManufacturer), v))
	})
}

// ManufacturerLTE applies the LTE predicate on the "manufacturer" field.
func ManufacturerLTE(v string) predicate.Aircraft {
	return predicate.Aircraft(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldManufacturer), v))
	})
}

// ManufacturerContains applies the Contains predicate on the "manufacturer" field.
func ManufacturerContains(v string) predicate.Aircraft {
	return predicate.Aircraft(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldManufacturer), v))
	})
}

// ManufacturerHasPrefix applies the HasPrefix predicate on the "manufacturer" field.
func ManufacturerHasPrefix(v string) predicate.Aircraft {
	return predicate.Aircraft(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldManufacturer), v))
	})
}

// ManufacturerHasSuffix applies the HasSuffix predicate on the "manufacturer" field.
func ManufacturerHasSuffix(v string) predicate.Aircraft {
	return predicate.Aircraft(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldManufacturer), v))
	})
}

// ManufacturerEqualFold applies the EqualFold predicate on the "manufacturer" field.
func ManufacturerEqualFold(v string) predicate.Aircraft {
	return predicate.Aircraft(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldManufacturer), v))
	})
}

// ManufacturerContainsFold applies the ContainsFold predicate on the "manufacturer" field.
func ManufacturerContainsFold(v string) predicate.Aircraft {
	return predicate.Aircraft(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldManufacturer), v))
	})
}

// ModelEQ applies the EQ predicate on the "model" field.
func ModelEQ(v string) predicate.Aircraft {
	return predicate.Aircraft(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldModel), v))
	})
}

// ModelNEQ applies the NEQ predicate on the "model" field.
func ModelNEQ(v string) predicate.Aircraft {
	return predicate.Aircraft(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldModel), v))
	})
}

// ModelIn applies the In predicate on the "model" field.
func ModelIn(vs ...string) predicate.Aircraft {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Aircraft(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldModel), v...))
	})
}

// ModelNotIn applies the NotIn predicate on the "model" field.
func ModelNotIn(vs ...string) predicate.Aircraft {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Aircraft(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldModel), v...))
	})
}

// ModelGT applies the GT predicate on the "model" field.
func ModelGT(v string) predicate.Aircraft {
	return predicate.Aircraft(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldModel), v))
	})
}

// ModelGTE applies the GTE predicate on the "model" field.
func ModelGTE(v string) predicate.Aircraft {
	return predicate.Aircraft(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldModel), v))
	})
}

// ModelLT applies the LT predicate on the "model" field.
func ModelLT(v string) predicate.Aircraft {
	return predicate.Aircraft(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldModel), v))
	})
}

// ModelLTE applies the LTE predicate on the "model" field.
func ModelLTE(v string) predicate.Aircraft {
	return predicate.Aircraft(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldModel), v))
	})
}

// ModelContains applies the Contains predicate on the "model" field.
func ModelContains(v string) predicate.Aircraft {
	return predicate.Aircraft(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldModel), v))
	})
}

// ModelHasPrefix applies the HasPrefix predicate on the "model" field.
func ModelHasPrefix(v string) predicate.Aircraft {
	return predicate.Aircraft(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldModel), v))
	})
}

// ModelHasSuffix applies the HasSuffix predicate on the "model" field.
func ModelHasSuffix(v string) predicate.Aircraft {
	return predicate.Aircraft(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldModel), v))
	})
}

// ModelEqualFold applies the EqualFold predicate on the "model" field.
func ModelEqualFold(v string) predicate.Aircraft {
	return predicate.Aircraft(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldModel), v))
	})
}

// ModelContainsFold applies the ContainsFold predicate on the "model" field.
func ModelContainsFold(v string) predicate.Aircraft {
	return predicate.Aircraft(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldModel), v))
	})
}

// CapacityEQ applies the EQ predicate on the "capacity" field.
func CapacityEQ(v int) predicate.Aircraft {
	return predicate.Aircraft(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCapacity), v))
	})
}

// CapacityNEQ applies the NEQ predicate on the "capacity" field.
func CapacityNEQ(v int) predicate.Aircraft {
	return predicate.Aircraft(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCapacity), v))
	})
}

// CapacityIn applies the In predicate on the "capacity" field.
func CapacityIn(vs ...int) predicate.Aircraft {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Aircraft(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCapacity), v...))
	})
}

// CapacityNotIn applies the NotIn predicate on the "capacity" field.
func CapacityNotIn(vs ...int) predicate.Aircraft {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Aircraft(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCapacity), v...))
	})
}

// CapacityGT applies the GT predicate on the "capacity" field.
func CapacityGT(v int) predicate.Aircraft {
	return predicate.Aircraft(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCapacity), v))
	})
}

// CapacityGTE applies the GTE predicate on the "capacity" field.
func CapacityGTE(v int) predicate.Aircraft {
	return predicate.Aircraft(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCapacity), v))
	})
}

// CapacityLT applies the LT predicate on the "capacity" field.
func CapacityLT(v int) predicate.Aircraft {
	return predicate.Aircraft(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCapacity), v))
	})
}

// CapacityLTE applies the LTE predicate on the "capacity" field.
func CapacityLTE(v int) predicate.Aircraft {
	return predicate.Aircraft(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCapacity), v))
	})
}

// RangeEQ applies the EQ predicate on the "range" field.
func RangeEQ(v int) predicate.Aircraft {
	return predicate.Aircraft(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRange), v))
	})
}

// RangeNEQ applies the NEQ predicate on the "range" field.
func RangeNEQ(v int) predicate.Aircraft {
	return predicate.Aircraft(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldRange), v))
	})
}

// RangeIn applies the In predicate on the "range" field.
func RangeIn(vs ...int) predicate.Aircraft {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Aircraft(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldRange), v...))
	})
}

// RangeNotIn applies the NotIn predicate on the "range" field.
func RangeNotIn(vs ...int) predicate.Aircraft {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Aircraft(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldRange), v...))
	})
}

// RangeGT applies the GT predicate on the "range" field.
func RangeGT(v int) predicate.Aircraft {
	return predicate.Aircraft(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldRange), v))
	})
}

// RangeGTE applies the GTE predicate on the "range" field.
func RangeGTE(v int) predicate.Aircraft {
	return predicate.Aircraft(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldRange), v))
	})
}

// RangeLT applies the LT predicate on the "range" field.
func RangeLT(v int) predicate.Aircraft {
	return predicate.Aircraft(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldRange), v))
	})
}

// RangeLTE applies the LTE predicate on the "range" field.
func RangeLTE(v int) predicate.Aircraft {
	return predicate.Aircraft(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldRange), v))
	})
}

// ManufacturedAtEQ applies the EQ predicate on the "manufactured_at" field.
func ManufacturedAtEQ(v time.Time) predicate.Aircraft {
	return predicate.Aircraft(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldManufacturedAt), v))
	})
}

// ManufacturedAtNEQ applies the NEQ predicate on the "manufactured_at" field.
func ManufacturedAtNEQ(v time.Time) predicate.Aircraft {
	return predicate.Aircraft(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldManufacturedAt), v))
	})
}

// ManufacturedAtIn applies the In predicate on the "manufactured_at" field.
func ManufacturedAtIn(vs ...time.Time) predicate.Aircraft {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Aircraft(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldManufacturedAt), v...))
	})
}

// ManufacturedAtNotIn applies the NotIn predicate on the "manufactured_at" field.
func ManufacturedAtNotIn(vs ...time.Time) predicate.Aircraft {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Aircraft(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldManufacturedAt), v...))
	})
}

// ManufacturedAtGT applies the GT predicate on the "manufactured_at" field.
func ManufacturedAtGT(v time.Time) predicate.Aircraft {
	return predicate.Aircraft(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldManufacturedAt), v))
	})
}

// ManufacturedAtGTE applies the GTE predicate on the "manufactured_at" field.
func ManufacturedAtGTE(v time.Time) predicate.Aircraft {
	return predicate.Aircraft(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldManufacturedAt), v))
	})
}

// ManufacturedAtLT applies the LT predicate on the "manufactured_at" field.
func ManufacturedAtLT(v time.Time) predicate.Aircraft {
	return predicate.Aircraft(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldManufacturedAt), v))
	})
}

// ManufacturedAtLTE applies the LTE predicate on the "manufactured_at" field.
func ManufacturedAtLTE(v time.Time) predicate.Aircraft {
	return predicate.Aircraft(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldManufacturedAt), v))
	})
}

// IsGroundedEQ applies the EQ predicate on the "is_grounded" field.
func IsGroundedEQ(v bool) predicate.Aircraft {
	return predicate.Aircraft(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIsGrounded), v))
	})
}

// IsGroundedNEQ applies the NEQ predicate on the "is_grounded" field.
func IsGroundedNEQ(v bool) predicate.Aircraft {
	return predicate.Aircraft(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldIsGrounded), v))
	})
}

// GroundedAtEQ applies the EQ predicate on the "grounded_at" field.
func GroundedAtEQ(v time.Time) predicate.Aircraft {
	return predicate.Aircraft(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldGroundedAt), v))
	})
}

// GroundedAtNEQ applies the NEQ predicate on the "grounded_at" field.
func GroundedAtNEQ(v time.Time) predicate.Aircraft {
	return predicate.Aircraft(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldGroundedAt), v))
	})
}

// GroundedAtIn applies the In predicate on the "grounded_at" field.
func GroundedAtIn(vs ...time.Time) predicate.Aircraft {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Aircraft(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldGroundedAt), v...))
	})
}

// GroundedAtNotIn applies the NotIn predicate on the "grounded_at" field.
func GroundedAtNotIn(vs ...time.Time) predicate.Aircraft {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Aircraft(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldGroundedAt), v...))
	})
}

// GroundedAtGT applies the GT predicate on the "grounded_at" field.
func GroundedAtGT(v time.Time) predicate.Aircraft {
	return predicate.Aircraft(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldGroundedAt), v))
	})
}

// GroundedAtGTE applies the GTE predicate on the "grounded_at" field.
func GroundedAtGTE(v time.Time) predicate.Aircraft {
	return predicate.Aircraft(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldGroundedAt), v))
	})
}

// GroundedAtLT applies the LT predicate on the "grounded_at" field.
func GroundedAtLT(v time.Time) predicate.Aircraft {
	return predicate.Aircraft(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldGroundedAt), v))
	})
}

// GroundedAtLTE applies the LTE predicate on the "grounded_at" field.
func GroundedAtLTE(v time.Time) predicate.Aircraft {
	return predicate.Aircraft(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldGroundedAt), v))
	})
}

// GroundedAtIsNil applies the IsNil predicate on the "grounded_at" field.
func GroundedAtIsNil() predicate.Aircraft {
	return predicate.Aircraft(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldGroundedAt)))
	})
}

// GroundedAtNotNil applies the NotNil predicate on the "grounded_at" field.
func GroundedAtNotNil() predicate.Aircraft {
	return predicate.Aircraft(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldGroundedAt)))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Aircraft {
	return predicate.Aircraft(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Aircraft {
	return predicate.Aircraft(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Aircraft {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Aircraft(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Aircraft {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Aircraft(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Aircraft {
	return predicate.Aircraft(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Aircraft {
	return predicate.Aircraft(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Aircraft {
	return predicate.Aircraft(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Aircraft {
	return predicate.Aircraft(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Aircraft {
	return predicate.Aircraft(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Aircraft {
	return predicate.Aircraft(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Aircraft {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Aircraft(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Aircraft {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Aircraft(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Aircraft {
	return predicate.Aircraft(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Aircraft {
	return predicate.Aircraft(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Aircraft {
	return predicate.Aircraft(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Aircraft {
	return predicate.Aircraft(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// HasAirline applies the HasEdge predicate on the "airline" edge.
func HasAirline() predicate.Aircraft {
	return predicate.Aircraft(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(AirlineTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, AirlineTable, AirlineColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAirlineWith applies the HasEdge predicate on the "airline" edge with a given conditions (other predicates).
func HasAirlineWith(preds ...predicate.Airline) predicate.Aircraft {
	return predicate.Aircraft(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(AirlineInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, AirlineTable, AirlineColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasFlightInstance applies the HasEdge predicate on the "flight_instance" edge.
func HasFlightInstance() predicate.Aircraft {
	return predicate.Aircraft(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(FlightInstanceTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, FlightInstanceTable, FlightInstanceColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasFlightInstanceWith applies the HasEdge predicate on the "flight_instance" edge with a given conditions (other predicates).
func HasFlightInstanceWith(preds ...predicate.FlightInstance) predicate.Aircraft {
	return predicate.Aircraft(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(FlightInstanceInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, FlightInstanceTable, FlightInstanceColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasSeats applies the HasEdge predicate on the "seats" edge.
func HasSeats() predicate.Aircraft {
	return predicate.Aircraft(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(SeatsTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, SeatsTable, SeatsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSeatsWith applies the HasEdge predicate on the "seats" edge with a given conditions (other predicates).
func HasSeatsWith(preds ...predicate.Seat) predicate.Aircraft {
	return predicate.Aircraft(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(SeatsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, SeatsTable, SeatsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Aircraft) predicate.Aircraft {
	return predicate.Aircraft(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Aircraft) predicate.Aircraft {
	return predicate.Aircraft(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Aircraft) predicate.Aircraft {
	return predicate.Aircraft(func(s *sql.Selector) {
		p(s.Not())
	})
}
