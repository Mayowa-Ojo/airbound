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

// Level applies equality check predicate on the "level" field. It's identical to LevelEQ.
func Level(v int) predicate.Admin {
	return predicate.Admin(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLevel), v))
	})
}

// SecurityQuestion applies equality check predicate on the "security_question" field. It's identical to SecurityQuestionEQ.
func SecurityQuestion(v string) predicate.Admin {
	return predicate.Admin(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSecurityQuestion), v))
	})
}

// SecurityAnswer applies equality check predicate on the "security_answer" field. It's identical to SecurityAnswerEQ.
func SecurityAnswer(v string) predicate.Admin {
	return predicate.Admin(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSecurityAnswer), v))
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

// LevelEQ applies the EQ predicate on the "level" field.
func LevelEQ(v int) predicate.Admin {
	return predicate.Admin(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLevel), v))
	})
}

// LevelNEQ applies the NEQ predicate on the "level" field.
func LevelNEQ(v int) predicate.Admin {
	return predicate.Admin(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLevel), v))
	})
}

// LevelIn applies the In predicate on the "level" field.
func LevelIn(vs ...int) predicate.Admin {
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
		s.Where(sql.In(s.C(FieldLevel), v...))
	})
}

// LevelNotIn applies the NotIn predicate on the "level" field.
func LevelNotIn(vs ...int) predicate.Admin {
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
		s.Where(sql.NotIn(s.C(FieldLevel), v...))
	})
}

// LevelGT applies the GT predicate on the "level" field.
func LevelGT(v int) predicate.Admin {
	return predicate.Admin(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldLevel), v))
	})
}

// LevelGTE applies the GTE predicate on the "level" field.
func LevelGTE(v int) predicate.Admin {
	return predicate.Admin(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldLevel), v))
	})
}

// LevelLT applies the LT predicate on the "level" field.
func LevelLT(v int) predicate.Admin {
	return predicate.Admin(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldLevel), v))
	})
}

// LevelLTE applies the LTE predicate on the "level" field.
func LevelLTE(v int) predicate.Admin {
	return predicate.Admin(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldLevel), v))
	})
}

// SecurityQuestionEQ applies the EQ predicate on the "security_question" field.
func SecurityQuestionEQ(v string) predicate.Admin {
	return predicate.Admin(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSecurityQuestion), v))
	})
}

// SecurityQuestionNEQ applies the NEQ predicate on the "security_question" field.
func SecurityQuestionNEQ(v string) predicate.Admin {
	return predicate.Admin(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSecurityQuestion), v))
	})
}

// SecurityQuestionIn applies the In predicate on the "security_question" field.
func SecurityQuestionIn(vs ...string) predicate.Admin {
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
		s.Where(sql.In(s.C(FieldSecurityQuestion), v...))
	})
}

// SecurityQuestionNotIn applies the NotIn predicate on the "security_question" field.
func SecurityQuestionNotIn(vs ...string) predicate.Admin {
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
		s.Where(sql.NotIn(s.C(FieldSecurityQuestion), v...))
	})
}

// SecurityQuestionGT applies the GT predicate on the "security_question" field.
func SecurityQuestionGT(v string) predicate.Admin {
	return predicate.Admin(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSecurityQuestion), v))
	})
}

// SecurityQuestionGTE applies the GTE predicate on the "security_question" field.
func SecurityQuestionGTE(v string) predicate.Admin {
	return predicate.Admin(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSecurityQuestion), v))
	})
}

// SecurityQuestionLT applies the LT predicate on the "security_question" field.
func SecurityQuestionLT(v string) predicate.Admin {
	return predicate.Admin(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSecurityQuestion), v))
	})
}

// SecurityQuestionLTE applies the LTE predicate on the "security_question" field.
func SecurityQuestionLTE(v string) predicate.Admin {
	return predicate.Admin(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSecurityQuestion), v))
	})
}

// SecurityQuestionContains applies the Contains predicate on the "security_question" field.
func SecurityQuestionContains(v string) predicate.Admin {
	return predicate.Admin(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldSecurityQuestion), v))
	})
}

// SecurityQuestionHasPrefix applies the HasPrefix predicate on the "security_question" field.
func SecurityQuestionHasPrefix(v string) predicate.Admin {
	return predicate.Admin(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldSecurityQuestion), v))
	})
}

// SecurityQuestionHasSuffix applies the HasSuffix predicate on the "security_question" field.
func SecurityQuestionHasSuffix(v string) predicate.Admin {
	return predicate.Admin(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldSecurityQuestion), v))
	})
}

// SecurityQuestionIsNil applies the IsNil predicate on the "security_question" field.
func SecurityQuestionIsNil() predicate.Admin {
	return predicate.Admin(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldSecurityQuestion)))
	})
}

// SecurityQuestionNotNil applies the NotNil predicate on the "security_question" field.
func SecurityQuestionNotNil() predicate.Admin {
	return predicate.Admin(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldSecurityQuestion)))
	})
}

// SecurityQuestionEqualFold applies the EqualFold predicate on the "security_question" field.
func SecurityQuestionEqualFold(v string) predicate.Admin {
	return predicate.Admin(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldSecurityQuestion), v))
	})
}

// SecurityQuestionContainsFold applies the ContainsFold predicate on the "security_question" field.
func SecurityQuestionContainsFold(v string) predicate.Admin {
	return predicate.Admin(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldSecurityQuestion), v))
	})
}

// SecurityAnswerEQ applies the EQ predicate on the "security_answer" field.
func SecurityAnswerEQ(v string) predicate.Admin {
	return predicate.Admin(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSecurityAnswer), v))
	})
}

// SecurityAnswerNEQ applies the NEQ predicate on the "security_answer" field.
func SecurityAnswerNEQ(v string) predicate.Admin {
	return predicate.Admin(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSecurityAnswer), v))
	})
}

// SecurityAnswerIn applies the In predicate on the "security_answer" field.
func SecurityAnswerIn(vs ...string) predicate.Admin {
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
		s.Where(sql.In(s.C(FieldSecurityAnswer), v...))
	})
}

// SecurityAnswerNotIn applies the NotIn predicate on the "security_answer" field.
func SecurityAnswerNotIn(vs ...string) predicate.Admin {
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
		s.Where(sql.NotIn(s.C(FieldSecurityAnswer), v...))
	})
}

// SecurityAnswerGT applies the GT predicate on the "security_answer" field.
func SecurityAnswerGT(v string) predicate.Admin {
	return predicate.Admin(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSecurityAnswer), v))
	})
}

// SecurityAnswerGTE applies the GTE predicate on the "security_answer" field.
func SecurityAnswerGTE(v string) predicate.Admin {
	return predicate.Admin(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSecurityAnswer), v))
	})
}

// SecurityAnswerLT applies the LT predicate on the "security_answer" field.
func SecurityAnswerLT(v string) predicate.Admin {
	return predicate.Admin(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSecurityAnswer), v))
	})
}

// SecurityAnswerLTE applies the LTE predicate on the "security_answer" field.
func SecurityAnswerLTE(v string) predicate.Admin {
	return predicate.Admin(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSecurityAnswer), v))
	})
}

// SecurityAnswerContains applies the Contains predicate on the "security_answer" field.
func SecurityAnswerContains(v string) predicate.Admin {
	return predicate.Admin(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldSecurityAnswer), v))
	})
}

// SecurityAnswerHasPrefix applies the HasPrefix predicate on the "security_answer" field.
func SecurityAnswerHasPrefix(v string) predicate.Admin {
	return predicate.Admin(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldSecurityAnswer), v))
	})
}

// SecurityAnswerHasSuffix applies the HasSuffix predicate on the "security_answer" field.
func SecurityAnswerHasSuffix(v string) predicate.Admin {
	return predicate.Admin(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldSecurityAnswer), v))
	})
}

// SecurityAnswerIsNil applies the IsNil predicate on the "security_answer" field.
func SecurityAnswerIsNil() predicate.Admin {
	return predicate.Admin(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldSecurityAnswer)))
	})
}

// SecurityAnswerNotNil applies the NotNil predicate on the "security_answer" field.
func SecurityAnswerNotNil() predicate.Admin {
	return predicate.Admin(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldSecurityAnswer)))
	})
}

// SecurityAnswerEqualFold applies the EqualFold predicate on the "security_answer" field.
func SecurityAnswerEqualFold(v string) predicate.Admin {
	return predicate.Admin(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldSecurityAnswer), v))
	})
}

// SecurityAnswerContainsFold applies the ContainsFold predicate on the "security_answer" field.
func SecurityAnswerContainsFold(v string) predicate.Admin {
	return predicate.Admin(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldSecurityAnswer), v))
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
