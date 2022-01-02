// Code generated by entc, DO NOT EDIT.

package flightschedule

import (
	"airbound/internal/ent/customtypes"
	"airbound/internal/ent/enums"
	"airbound/internal/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.FlightSchedule {
	return predicate.FlightSchedule(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.FlightSchedule {
	return predicate.FlightSchedule(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.FlightSchedule {
	return predicate.FlightSchedule(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.FlightSchedule {
	return predicate.FlightSchedule(func(s *sql.Selector) {
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
func IDNotIn(ids ...uuid.UUID) predicate.FlightSchedule {
	return predicate.FlightSchedule(func(s *sql.Selector) {
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
func IDGT(id uuid.UUID) predicate.FlightSchedule {
	return predicate.FlightSchedule(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.FlightSchedule {
	return predicate.FlightSchedule(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.FlightSchedule {
	return predicate.FlightSchedule(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.FlightSchedule {
	return predicate.FlightSchedule(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CustomDate applies equality check predicate on the "custom_date" field. It's identical to CustomDateEQ.
func CustomDate(v customtypes.Date) predicate.FlightSchedule {
	return predicate.FlightSchedule(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCustomDate), v))
	})
}

// DepartsAt applies equality check predicate on the "departs_at" field. It's identical to DepartsAtEQ.
func DepartsAt(v customtypes.Time) predicate.FlightSchedule {
	return predicate.FlightSchedule(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDepartsAt), v))
	})
}

// ArrivesAt applies equality check predicate on the "arrives_at" field. It's identical to ArrivesAtEQ.
func ArrivesAt(v customtypes.Time) predicate.FlightSchedule {
	return predicate.FlightSchedule(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldArrivesAt), v))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.FlightSchedule {
	return predicate.FlightSchedule(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.FlightSchedule {
	return predicate.FlightSchedule(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// WeekdayEQ applies the EQ predicate on the "weekday" field.
func WeekdayEQ(v enums.WeekDay) predicate.FlightSchedule {
	return predicate.FlightSchedule(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldWeekday), v))
	})
}

// WeekdayNEQ applies the NEQ predicate on the "weekday" field.
func WeekdayNEQ(v enums.WeekDay) predicate.FlightSchedule {
	return predicate.FlightSchedule(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldWeekday), v))
	})
}

// WeekdayIn applies the In predicate on the "weekday" field.
func WeekdayIn(vs ...enums.WeekDay) predicate.FlightSchedule {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FlightSchedule(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldWeekday), v...))
	})
}

// WeekdayNotIn applies the NotIn predicate on the "weekday" field.
func WeekdayNotIn(vs ...enums.WeekDay) predicate.FlightSchedule {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FlightSchedule(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldWeekday), v...))
	})
}

// WeekdayIsNil applies the IsNil predicate on the "weekday" field.
func WeekdayIsNil() predicate.FlightSchedule {
	return predicate.FlightSchedule(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldWeekday)))
	})
}

// WeekdayNotNil applies the NotNil predicate on the "weekday" field.
func WeekdayNotNil() predicate.FlightSchedule {
	return predicate.FlightSchedule(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldWeekday)))
	})
}

// ScheduleTypeEQ applies the EQ predicate on the "schedule_type" field.
func ScheduleTypeEQ(v enums.FlightScheduleType) predicate.FlightSchedule {
	vc := v
	return predicate.FlightSchedule(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldScheduleType), vc))
	})
}

// ScheduleTypeNEQ applies the NEQ predicate on the "schedule_type" field.
func ScheduleTypeNEQ(v enums.FlightScheduleType) predicate.FlightSchedule {
	vc := v
	return predicate.FlightSchedule(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldScheduleType), vc))
	})
}

// ScheduleTypeIn applies the In predicate on the "schedule_type" field.
func ScheduleTypeIn(vs ...enums.FlightScheduleType) predicate.FlightSchedule {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FlightSchedule(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldScheduleType), v...))
	})
}

// ScheduleTypeNotIn applies the NotIn predicate on the "schedule_type" field.
func ScheduleTypeNotIn(vs ...enums.FlightScheduleType) predicate.FlightSchedule {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FlightSchedule(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldScheduleType), v...))
	})
}

// CustomDateEQ applies the EQ predicate on the "custom_date" field.
func CustomDateEQ(v customtypes.Date) predicate.FlightSchedule {
	return predicate.FlightSchedule(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCustomDate), v))
	})
}

// CustomDateNEQ applies the NEQ predicate on the "custom_date" field.
func CustomDateNEQ(v customtypes.Date) predicate.FlightSchedule {
	return predicate.FlightSchedule(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCustomDate), v))
	})
}

// CustomDateIn applies the In predicate on the "custom_date" field.
func CustomDateIn(vs ...customtypes.Date) predicate.FlightSchedule {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FlightSchedule(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCustomDate), v...))
	})
}

// CustomDateNotIn applies the NotIn predicate on the "custom_date" field.
func CustomDateNotIn(vs ...customtypes.Date) predicate.FlightSchedule {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FlightSchedule(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCustomDate), v...))
	})
}

// CustomDateGT applies the GT predicate on the "custom_date" field.
func CustomDateGT(v customtypes.Date) predicate.FlightSchedule {
	return predicate.FlightSchedule(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCustomDate), v))
	})
}

// CustomDateGTE applies the GTE predicate on the "custom_date" field.
func CustomDateGTE(v customtypes.Date) predicate.FlightSchedule {
	return predicate.FlightSchedule(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCustomDate), v))
	})
}

// CustomDateLT applies the LT predicate on the "custom_date" field.
func CustomDateLT(v customtypes.Date) predicate.FlightSchedule {
	return predicate.FlightSchedule(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCustomDate), v))
	})
}

// CustomDateLTE applies the LTE predicate on the "custom_date" field.
func CustomDateLTE(v customtypes.Date) predicate.FlightSchedule {
	return predicate.FlightSchedule(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCustomDate), v))
	})
}

// CustomDateContains applies the Contains predicate on the "custom_date" field.
func CustomDateContains(v customtypes.Date) predicate.FlightSchedule {
	vc := v.String()
	return predicate.FlightSchedule(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldCustomDate), vc))
	})
}

// CustomDateHasPrefix applies the HasPrefix predicate on the "custom_date" field.
func CustomDateHasPrefix(v customtypes.Date) predicate.FlightSchedule {
	vc := v.String()
	return predicate.FlightSchedule(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldCustomDate), vc))
	})
}

// CustomDateHasSuffix applies the HasSuffix predicate on the "custom_date" field.
func CustomDateHasSuffix(v customtypes.Date) predicate.FlightSchedule {
	vc := v.String()
	return predicate.FlightSchedule(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldCustomDate), vc))
	})
}

// CustomDateIsNil applies the IsNil predicate on the "custom_date" field.
func CustomDateIsNil() predicate.FlightSchedule {
	return predicate.FlightSchedule(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldCustomDate)))
	})
}

// CustomDateNotNil applies the NotNil predicate on the "custom_date" field.
func CustomDateNotNil() predicate.FlightSchedule {
	return predicate.FlightSchedule(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldCustomDate)))
	})
}

// CustomDateEqualFold applies the EqualFold predicate on the "custom_date" field.
func CustomDateEqualFold(v customtypes.Date) predicate.FlightSchedule {
	vc := v.String()
	return predicate.FlightSchedule(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldCustomDate), vc))
	})
}

// CustomDateContainsFold applies the ContainsFold predicate on the "custom_date" field.
func CustomDateContainsFold(v customtypes.Date) predicate.FlightSchedule {
	vc := v.String()
	return predicate.FlightSchedule(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldCustomDate), vc))
	})
}

// DepartsAtEQ applies the EQ predicate on the "departs_at" field.
func DepartsAtEQ(v customtypes.Time) predicate.FlightSchedule {
	return predicate.FlightSchedule(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDepartsAt), v))
	})
}

// DepartsAtNEQ applies the NEQ predicate on the "departs_at" field.
func DepartsAtNEQ(v customtypes.Time) predicate.FlightSchedule {
	return predicate.FlightSchedule(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDepartsAt), v))
	})
}

// DepartsAtIn applies the In predicate on the "departs_at" field.
func DepartsAtIn(vs ...customtypes.Time) predicate.FlightSchedule {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FlightSchedule(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDepartsAt), v...))
	})
}

// DepartsAtNotIn applies the NotIn predicate on the "departs_at" field.
func DepartsAtNotIn(vs ...customtypes.Time) predicate.FlightSchedule {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FlightSchedule(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDepartsAt), v...))
	})
}

// DepartsAtGT applies the GT predicate on the "departs_at" field.
func DepartsAtGT(v customtypes.Time) predicate.FlightSchedule {
	return predicate.FlightSchedule(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDepartsAt), v))
	})
}

// DepartsAtGTE applies the GTE predicate on the "departs_at" field.
func DepartsAtGTE(v customtypes.Time) predicate.FlightSchedule {
	return predicate.FlightSchedule(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDepartsAt), v))
	})
}

// DepartsAtLT applies the LT predicate on the "departs_at" field.
func DepartsAtLT(v customtypes.Time) predicate.FlightSchedule {
	return predicate.FlightSchedule(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDepartsAt), v))
	})
}

// DepartsAtLTE applies the LTE predicate on the "departs_at" field.
func DepartsAtLTE(v customtypes.Time) predicate.FlightSchedule {
	return predicate.FlightSchedule(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDepartsAt), v))
	})
}

// DepartsAtContains applies the Contains predicate on the "departs_at" field.
func DepartsAtContains(v customtypes.Time) predicate.FlightSchedule {
	vc := v.String()
	return predicate.FlightSchedule(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldDepartsAt), vc))
	})
}

// DepartsAtHasPrefix applies the HasPrefix predicate on the "departs_at" field.
func DepartsAtHasPrefix(v customtypes.Time) predicate.FlightSchedule {
	vc := v.String()
	return predicate.FlightSchedule(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldDepartsAt), vc))
	})
}

// DepartsAtHasSuffix applies the HasSuffix predicate on the "departs_at" field.
func DepartsAtHasSuffix(v customtypes.Time) predicate.FlightSchedule {
	vc := v.String()
	return predicate.FlightSchedule(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldDepartsAt), vc))
	})
}

// DepartsAtEqualFold applies the EqualFold predicate on the "departs_at" field.
func DepartsAtEqualFold(v customtypes.Time) predicate.FlightSchedule {
	vc := v.String()
	return predicate.FlightSchedule(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldDepartsAt), vc))
	})
}

// DepartsAtContainsFold applies the ContainsFold predicate on the "departs_at" field.
func DepartsAtContainsFold(v customtypes.Time) predicate.FlightSchedule {
	vc := v.String()
	return predicate.FlightSchedule(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldDepartsAt), vc))
	})
}

// ArrivesAtEQ applies the EQ predicate on the "arrives_at" field.
func ArrivesAtEQ(v customtypes.Time) predicate.FlightSchedule {
	return predicate.FlightSchedule(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldArrivesAt), v))
	})
}

// ArrivesAtNEQ applies the NEQ predicate on the "arrives_at" field.
func ArrivesAtNEQ(v customtypes.Time) predicate.FlightSchedule {
	return predicate.FlightSchedule(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldArrivesAt), v))
	})
}

// ArrivesAtIn applies the In predicate on the "arrives_at" field.
func ArrivesAtIn(vs ...customtypes.Time) predicate.FlightSchedule {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FlightSchedule(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldArrivesAt), v...))
	})
}

// ArrivesAtNotIn applies the NotIn predicate on the "arrives_at" field.
func ArrivesAtNotIn(vs ...customtypes.Time) predicate.FlightSchedule {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FlightSchedule(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldArrivesAt), v...))
	})
}

// ArrivesAtGT applies the GT predicate on the "arrives_at" field.
func ArrivesAtGT(v customtypes.Time) predicate.FlightSchedule {
	return predicate.FlightSchedule(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldArrivesAt), v))
	})
}

// ArrivesAtGTE applies the GTE predicate on the "arrives_at" field.
func ArrivesAtGTE(v customtypes.Time) predicate.FlightSchedule {
	return predicate.FlightSchedule(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldArrivesAt), v))
	})
}

// ArrivesAtLT applies the LT predicate on the "arrives_at" field.
func ArrivesAtLT(v customtypes.Time) predicate.FlightSchedule {
	return predicate.FlightSchedule(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldArrivesAt), v))
	})
}

// ArrivesAtLTE applies the LTE predicate on the "arrives_at" field.
func ArrivesAtLTE(v customtypes.Time) predicate.FlightSchedule {
	return predicate.FlightSchedule(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldArrivesAt), v))
	})
}

// ArrivesAtContains applies the Contains predicate on the "arrives_at" field.
func ArrivesAtContains(v customtypes.Time) predicate.FlightSchedule {
	vc := v.String()
	return predicate.FlightSchedule(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldArrivesAt), vc))
	})
}

// ArrivesAtHasPrefix applies the HasPrefix predicate on the "arrives_at" field.
func ArrivesAtHasPrefix(v customtypes.Time) predicate.FlightSchedule {
	vc := v.String()
	return predicate.FlightSchedule(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldArrivesAt), vc))
	})
}

// ArrivesAtHasSuffix applies the HasSuffix predicate on the "arrives_at" field.
func ArrivesAtHasSuffix(v customtypes.Time) predicate.FlightSchedule {
	vc := v.String()
	return predicate.FlightSchedule(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldArrivesAt), vc))
	})
}

// ArrivesAtEqualFold applies the EqualFold predicate on the "arrives_at" field.
func ArrivesAtEqualFold(v customtypes.Time) predicate.FlightSchedule {
	vc := v.String()
	return predicate.FlightSchedule(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldArrivesAt), vc))
	})
}

// ArrivesAtContainsFold applies the ContainsFold predicate on the "arrives_at" field.
func ArrivesAtContainsFold(v customtypes.Time) predicate.FlightSchedule {
	vc := v.String()
	return predicate.FlightSchedule(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldArrivesAt), vc))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.FlightSchedule {
	return predicate.FlightSchedule(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.FlightSchedule {
	return predicate.FlightSchedule(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.FlightSchedule {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FlightSchedule(func(s *sql.Selector) {
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
func CreatedAtNotIn(vs ...time.Time) predicate.FlightSchedule {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FlightSchedule(func(s *sql.Selector) {
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
func CreatedAtGT(v time.Time) predicate.FlightSchedule {
	return predicate.FlightSchedule(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.FlightSchedule {
	return predicate.FlightSchedule(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.FlightSchedule {
	return predicate.FlightSchedule(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.FlightSchedule {
	return predicate.FlightSchedule(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.FlightSchedule {
	return predicate.FlightSchedule(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.FlightSchedule {
	return predicate.FlightSchedule(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.FlightSchedule {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FlightSchedule(func(s *sql.Selector) {
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
func UpdatedAtNotIn(vs ...time.Time) predicate.FlightSchedule {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FlightSchedule(func(s *sql.Selector) {
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
func UpdatedAtGT(v time.Time) predicate.FlightSchedule {
	return predicate.FlightSchedule(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.FlightSchedule {
	return predicate.FlightSchedule(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.FlightSchedule {
	return predicate.FlightSchedule(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.FlightSchedule {
	return predicate.FlightSchedule(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.FlightSchedule) predicate.FlightSchedule {
	return predicate.FlightSchedule(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.FlightSchedule) predicate.FlightSchedule {
	return predicate.FlightSchedule(func(s *sql.Selector) {
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
func Not(p predicate.FlightSchedule) predicate.FlightSchedule {
	return predicate.FlightSchedule(func(s *sql.Selector) {
		p(s.Not())
	})
}
