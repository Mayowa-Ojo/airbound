// Code generated by entc, DO NOT EDIT.

package passenger

import (
	"airbound/internal/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Passenger {
	return predicate.Passenger(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Passenger {
	return predicate.Passenger(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Passenger {
	return predicate.Passenger(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Passenger {
	return predicate.Passenger(func(s *sql.Selector) {
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
func IDNotIn(ids ...uuid.UUID) predicate.Passenger {
	return predicate.Passenger(func(s *sql.Selector) {
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
func IDGT(id uuid.UUID) predicate.Passenger {
	return predicate.Passenger(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Passenger {
	return predicate.Passenger(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Passenger {
	return predicate.Passenger(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Passenger {
	return predicate.Passenger(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Firstname applies equality check predicate on the "firstname" field. It's identical to FirstnameEQ.
func Firstname(v string) predicate.Passenger {
	return predicate.Passenger(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFirstname), v))
	})
}

// Lastname applies equality check predicate on the "lastname" field. It's identical to LastnameEQ.
func Lastname(v string) predicate.Passenger {
	return predicate.Passenger(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLastname), v))
	})
}

// Age applies equality check predicate on the "age" field. It's identical to AgeEQ.
func Age(v int) predicate.Passenger {
	return predicate.Passenger(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAge), v))
	})
}

// PassportNumber applies equality check predicate on the "passport_number" field. It's identical to PassportNumberEQ.
func PassportNumber(v string) predicate.Passenger {
	return predicate.Passenger(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPassportNumber), v))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Passenger {
	return predicate.Passenger(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Passenger {
	return predicate.Passenger(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// FirstnameEQ applies the EQ predicate on the "firstname" field.
func FirstnameEQ(v string) predicate.Passenger {
	return predicate.Passenger(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFirstname), v))
	})
}

// FirstnameNEQ applies the NEQ predicate on the "firstname" field.
func FirstnameNEQ(v string) predicate.Passenger {
	return predicate.Passenger(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldFirstname), v))
	})
}

// FirstnameIn applies the In predicate on the "firstname" field.
func FirstnameIn(vs ...string) predicate.Passenger {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Passenger(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldFirstname), v...))
	})
}

// FirstnameNotIn applies the NotIn predicate on the "firstname" field.
func FirstnameNotIn(vs ...string) predicate.Passenger {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Passenger(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldFirstname), v...))
	})
}

// FirstnameGT applies the GT predicate on the "firstname" field.
func FirstnameGT(v string) predicate.Passenger {
	return predicate.Passenger(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldFirstname), v))
	})
}

// FirstnameGTE applies the GTE predicate on the "firstname" field.
func FirstnameGTE(v string) predicate.Passenger {
	return predicate.Passenger(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldFirstname), v))
	})
}

// FirstnameLT applies the LT predicate on the "firstname" field.
func FirstnameLT(v string) predicate.Passenger {
	return predicate.Passenger(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldFirstname), v))
	})
}

// FirstnameLTE applies the LTE predicate on the "firstname" field.
func FirstnameLTE(v string) predicate.Passenger {
	return predicate.Passenger(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldFirstname), v))
	})
}

// FirstnameContains applies the Contains predicate on the "firstname" field.
func FirstnameContains(v string) predicate.Passenger {
	return predicate.Passenger(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldFirstname), v))
	})
}

// FirstnameHasPrefix applies the HasPrefix predicate on the "firstname" field.
func FirstnameHasPrefix(v string) predicate.Passenger {
	return predicate.Passenger(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldFirstname), v))
	})
}

// FirstnameHasSuffix applies the HasSuffix predicate on the "firstname" field.
func FirstnameHasSuffix(v string) predicate.Passenger {
	return predicate.Passenger(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldFirstname), v))
	})
}

// FirstnameEqualFold applies the EqualFold predicate on the "firstname" field.
func FirstnameEqualFold(v string) predicate.Passenger {
	return predicate.Passenger(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldFirstname), v))
	})
}

// FirstnameContainsFold applies the ContainsFold predicate on the "firstname" field.
func FirstnameContainsFold(v string) predicate.Passenger {
	return predicate.Passenger(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldFirstname), v))
	})
}

// LastnameEQ applies the EQ predicate on the "lastname" field.
func LastnameEQ(v string) predicate.Passenger {
	return predicate.Passenger(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLastname), v))
	})
}

// LastnameNEQ applies the NEQ predicate on the "lastname" field.
func LastnameNEQ(v string) predicate.Passenger {
	return predicate.Passenger(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLastname), v))
	})
}

// LastnameIn applies the In predicate on the "lastname" field.
func LastnameIn(vs ...string) predicate.Passenger {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Passenger(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldLastname), v...))
	})
}

// LastnameNotIn applies the NotIn predicate on the "lastname" field.
func LastnameNotIn(vs ...string) predicate.Passenger {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Passenger(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldLastname), v...))
	})
}

// LastnameGT applies the GT predicate on the "lastname" field.
func LastnameGT(v string) predicate.Passenger {
	return predicate.Passenger(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldLastname), v))
	})
}

// LastnameGTE applies the GTE predicate on the "lastname" field.
func LastnameGTE(v string) predicate.Passenger {
	return predicate.Passenger(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldLastname), v))
	})
}

// LastnameLT applies the LT predicate on the "lastname" field.
func LastnameLT(v string) predicate.Passenger {
	return predicate.Passenger(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldLastname), v))
	})
}

// LastnameLTE applies the LTE predicate on the "lastname" field.
func LastnameLTE(v string) predicate.Passenger {
	return predicate.Passenger(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldLastname), v))
	})
}

// LastnameContains applies the Contains predicate on the "lastname" field.
func LastnameContains(v string) predicate.Passenger {
	return predicate.Passenger(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldLastname), v))
	})
}

// LastnameHasPrefix applies the HasPrefix predicate on the "lastname" field.
func LastnameHasPrefix(v string) predicate.Passenger {
	return predicate.Passenger(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldLastname), v))
	})
}

// LastnameHasSuffix applies the HasSuffix predicate on the "lastname" field.
func LastnameHasSuffix(v string) predicate.Passenger {
	return predicate.Passenger(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldLastname), v))
	})
}

// LastnameEqualFold applies the EqualFold predicate on the "lastname" field.
func LastnameEqualFold(v string) predicate.Passenger {
	return predicate.Passenger(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldLastname), v))
	})
}

// LastnameContainsFold applies the ContainsFold predicate on the "lastname" field.
func LastnameContainsFold(v string) predicate.Passenger {
	return predicate.Passenger(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldLastname), v))
	})
}

// AgeEQ applies the EQ predicate on the "age" field.
func AgeEQ(v int) predicate.Passenger {
	return predicate.Passenger(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAge), v))
	})
}

// AgeNEQ applies the NEQ predicate on the "age" field.
func AgeNEQ(v int) predicate.Passenger {
	return predicate.Passenger(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAge), v))
	})
}

// AgeIn applies the In predicate on the "age" field.
func AgeIn(vs ...int) predicate.Passenger {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Passenger(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldAge), v...))
	})
}

// AgeNotIn applies the NotIn predicate on the "age" field.
func AgeNotIn(vs ...int) predicate.Passenger {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Passenger(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldAge), v...))
	})
}

// AgeGT applies the GT predicate on the "age" field.
func AgeGT(v int) predicate.Passenger {
	return predicate.Passenger(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAge), v))
	})
}

// AgeGTE applies the GTE predicate on the "age" field.
func AgeGTE(v int) predicate.Passenger {
	return predicate.Passenger(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAge), v))
	})
}

// AgeLT applies the LT predicate on the "age" field.
func AgeLT(v int) predicate.Passenger {
	return predicate.Passenger(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAge), v))
	})
}

// AgeLTE applies the LTE predicate on the "age" field.
func AgeLTE(v int) predicate.Passenger {
	return predicate.Passenger(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAge), v))
	})
}

// PassportNumberEQ applies the EQ predicate on the "passport_number" field.
func PassportNumberEQ(v string) predicate.Passenger {
	return predicate.Passenger(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPassportNumber), v))
	})
}

// PassportNumberNEQ applies the NEQ predicate on the "passport_number" field.
func PassportNumberNEQ(v string) predicate.Passenger {
	return predicate.Passenger(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPassportNumber), v))
	})
}

// PassportNumberIn applies the In predicate on the "passport_number" field.
func PassportNumberIn(vs ...string) predicate.Passenger {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Passenger(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPassportNumber), v...))
	})
}

// PassportNumberNotIn applies the NotIn predicate on the "passport_number" field.
func PassportNumberNotIn(vs ...string) predicate.Passenger {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Passenger(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPassportNumber), v...))
	})
}

// PassportNumberGT applies the GT predicate on the "passport_number" field.
func PassportNumberGT(v string) predicate.Passenger {
	return predicate.Passenger(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPassportNumber), v))
	})
}

// PassportNumberGTE applies the GTE predicate on the "passport_number" field.
func PassportNumberGTE(v string) predicate.Passenger {
	return predicate.Passenger(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPassportNumber), v))
	})
}

// PassportNumberLT applies the LT predicate on the "passport_number" field.
func PassportNumberLT(v string) predicate.Passenger {
	return predicate.Passenger(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPassportNumber), v))
	})
}

// PassportNumberLTE applies the LTE predicate on the "passport_number" field.
func PassportNumberLTE(v string) predicate.Passenger {
	return predicate.Passenger(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPassportNumber), v))
	})
}

// PassportNumberContains applies the Contains predicate on the "passport_number" field.
func PassportNumberContains(v string) predicate.Passenger {
	return predicate.Passenger(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldPassportNumber), v))
	})
}

// PassportNumberHasPrefix applies the HasPrefix predicate on the "passport_number" field.
func PassportNumberHasPrefix(v string) predicate.Passenger {
	return predicate.Passenger(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldPassportNumber), v))
	})
}

// PassportNumberHasSuffix applies the HasSuffix predicate on the "passport_number" field.
func PassportNumberHasSuffix(v string) predicate.Passenger {
	return predicate.Passenger(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldPassportNumber), v))
	})
}

// PassportNumberEqualFold applies the EqualFold predicate on the "passport_number" field.
func PassportNumberEqualFold(v string) predicate.Passenger {
	return predicate.Passenger(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldPassportNumber), v))
	})
}

// PassportNumberContainsFold applies the ContainsFold predicate on the "passport_number" field.
func PassportNumberContainsFold(v string) predicate.Passenger {
	return predicate.Passenger(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldPassportNumber), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Passenger {
	return predicate.Passenger(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Passenger {
	return predicate.Passenger(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Passenger {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Passenger(func(s *sql.Selector) {
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
func CreatedAtNotIn(vs ...time.Time) predicate.Passenger {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Passenger(func(s *sql.Selector) {
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
func CreatedAtGT(v time.Time) predicate.Passenger {
	return predicate.Passenger(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Passenger {
	return predicate.Passenger(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Passenger {
	return predicate.Passenger(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Passenger {
	return predicate.Passenger(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Passenger {
	return predicate.Passenger(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Passenger {
	return predicate.Passenger(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Passenger {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Passenger(func(s *sql.Selector) {
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
func UpdatedAtNotIn(vs ...time.Time) predicate.Passenger {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Passenger(func(s *sql.Selector) {
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
func UpdatedAtGT(v time.Time) predicate.Passenger {
	return predicate.Passenger(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Passenger {
	return predicate.Passenger(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Passenger {
	return predicate.Passenger(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Passenger {
	return predicate.Passenger(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// HasFlightReservation applies the HasEdge predicate on the "flight_reservation" edge.
func HasFlightReservation() predicate.Passenger {
	return predicate.Passenger(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(FlightReservationTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, FlightReservationTable, FlightReservationColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasFlightReservationWith applies the HasEdge predicate on the "flight_reservation" edge with a given conditions (other predicates).
func HasFlightReservationWith(preds ...predicate.FlightReservation) predicate.Passenger {
	return predicate.Passenger(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(FlightReservationInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, FlightReservationTable, FlightReservationColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasFlightSeat applies the HasEdge predicate on the "flight_seat" edge.
func HasFlightSeat() predicate.Passenger {
	return predicate.Passenger(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(FlightSeatTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, FlightSeatTable, FlightSeatColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasFlightSeatWith applies the HasEdge predicate on the "flight_seat" edge with a given conditions (other predicates).
func HasFlightSeatWith(preds ...predicate.FlightSeat) predicate.Passenger {
	return predicate.Passenger(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(FlightSeatInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, FlightSeatTable, FlightSeatColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Passenger) predicate.Passenger {
	return predicate.Passenger(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Passenger) predicate.Passenger {
	return predicate.Passenger(func(s *sql.Selector) {
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
func Not(p predicate.Passenger) predicate.Passenger {
	return predicate.Passenger(func(s *sql.Selector) {
		p(s.Not())
	})
}
