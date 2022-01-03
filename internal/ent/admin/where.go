// Code generated by entc, DO NOT EDIT.

package admin

import (
	"airbound/internal/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Admin {
	return predicate.Admin(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Admin {
	return predicate.Admin(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Admin {
	return predicate.Admin(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Admin {
	return predicate.Admin(func(s *sql.Selector) {
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
func IDNotIn(ids ...uuid.UUID) predicate.Admin {
	return predicate.Admin(func(s *sql.Selector) {
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
func IDGT(id uuid.UUID) predicate.Admin {
	return predicate.Admin(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Admin {
	return predicate.Admin(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Admin {
	return predicate.Admin(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Admin {
	return predicate.Admin(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// TwoFaSecret applies equality check predicate on the "two_fa_secret" field. It's identical to TwoFaSecretEQ.
func TwoFaSecret(v string) predicate.Admin {
	return predicate.Admin(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTwoFaSecret), v))
	})
}

// TwoFaCompleted applies equality check predicate on the "two_fa_completed" field. It's identical to TwoFaCompletedEQ.
func TwoFaCompleted(v bool) predicate.Admin {
	return predicate.Admin(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTwoFaCompleted), v))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Admin {
	return predicate.Admin(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Admin {
	return predicate.Admin(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// TwoFaSecretEQ applies the EQ predicate on the "two_fa_secret" field.
func TwoFaSecretEQ(v string) predicate.Admin {
	return predicate.Admin(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTwoFaSecret), v))
	})
}

// TwoFaSecretNEQ applies the NEQ predicate on the "two_fa_secret" field.
func TwoFaSecretNEQ(v string) predicate.Admin {
	return predicate.Admin(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTwoFaSecret), v))
	})
}

// TwoFaSecretIn applies the In predicate on the "two_fa_secret" field.
func TwoFaSecretIn(vs ...string) predicate.Admin {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Admin(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldTwoFaSecret), v...))
	})
}

// TwoFaSecretNotIn applies the NotIn predicate on the "two_fa_secret" field.
func TwoFaSecretNotIn(vs ...string) predicate.Admin {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Admin(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldTwoFaSecret), v...))
	})
}

// TwoFaSecretGT applies the GT predicate on the "two_fa_secret" field.
func TwoFaSecretGT(v string) predicate.Admin {
	return predicate.Admin(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTwoFaSecret), v))
	})
}

// TwoFaSecretGTE applies the GTE predicate on the "two_fa_secret" field.
func TwoFaSecretGTE(v string) predicate.Admin {
	return predicate.Admin(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTwoFaSecret), v))
	})
}

// TwoFaSecretLT applies the LT predicate on the "two_fa_secret" field.
func TwoFaSecretLT(v string) predicate.Admin {
	return predicate.Admin(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTwoFaSecret), v))
	})
}

// TwoFaSecretLTE applies the LTE predicate on the "two_fa_secret" field.
func TwoFaSecretLTE(v string) predicate.Admin {
	return predicate.Admin(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTwoFaSecret), v))
	})
}

// TwoFaSecretContains applies the Contains predicate on the "two_fa_secret" field.
func TwoFaSecretContains(v string) predicate.Admin {
	return predicate.Admin(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTwoFaSecret), v))
	})
}

// TwoFaSecretHasPrefix applies the HasPrefix predicate on the "two_fa_secret" field.
func TwoFaSecretHasPrefix(v string) predicate.Admin {
	return predicate.Admin(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTwoFaSecret), v))
	})
}

// TwoFaSecretHasSuffix applies the HasSuffix predicate on the "two_fa_secret" field.
func TwoFaSecretHasSuffix(v string) predicate.Admin {
	return predicate.Admin(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTwoFaSecret), v))
	})
}

// TwoFaSecretIsNil applies the IsNil predicate on the "two_fa_secret" field.
func TwoFaSecretIsNil() predicate.Admin {
	return predicate.Admin(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldTwoFaSecret)))
	})
}

// TwoFaSecretNotNil applies the NotNil predicate on the "two_fa_secret" field.
func TwoFaSecretNotNil() predicate.Admin {
	return predicate.Admin(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldTwoFaSecret)))
	})
}

// TwoFaSecretEqualFold applies the EqualFold predicate on the "two_fa_secret" field.
func TwoFaSecretEqualFold(v string) predicate.Admin {
	return predicate.Admin(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTwoFaSecret), v))
	})
}

// TwoFaSecretContainsFold applies the ContainsFold predicate on the "two_fa_secret" field.
func TwoFaSecretContainsFold(v string) predicate.Admin {
	return predicate.Admin(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTwoFaSecret), v))
	})
}

// TwoFaCompletedEQ applies the EQ predicate on the "two_fa_completed" field.
func TwoFaCompletedEQ(v bool) predicate.Admin {
	return predicate.Admin(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTwoFaCompleted), v))
	})
}

// TwoFaCompletedNEQ applies the NEQ predicate on the "two_fa_completed" field.
func TwoFaCompletedNEQ(v bool) predicate.Admin {
	return predicate.Admin(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTwoFaCompleted), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Admin {
	return predicate.Admin(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Admin {
	return predicate.Admin(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Admin {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Admin(func(s *sql.Selector) {
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
func CreatedAtNotIn(vs ...time.Time) predicate.Admin {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Admin(func(s *sql.Selector) {
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
func CreatedAtGT(v time.Time) predicate.Admin {
	return predicate.Admin(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Admin {
	return predicate.Admin(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Admin {
	return predicate.Admin(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Admin {
	return predicate.Admin(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Admin {
	return predicate.Admin(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Admin {
	return predicate.Admin(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Admin {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Admin(func(s *sql.Selector) {
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
func UpdatedAtNotIn(vs ...time.Time) predicate.Admin {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Admin(func(s *sql.Selector) {
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
func UpdatedAtGT(v time.Time) predicate.Admin {
	return predicate.Admin(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Admin {
	return predicate.Admin(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Admin {
	return predicate.Admin(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Admin {
	return predicate.Admin(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.Admin {
	return predicate.Admin(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.Admin {
	return predicate.Admin(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Admin) predicate.Admin {
	return predicate.Admin(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Admin) predicate.Admin {
	return predicate.Admin(func(s *sql.Selector) {
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
func Not(p predicate.Admin) predicate.Admin {
	return predicate.Admin(func(s *sql.Selector) {
		p(s.Not())
	})
}
