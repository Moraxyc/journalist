// Code generated by ent, DO NOT EDIT.

package subscription

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
	"github.com/mrusme/journalist/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Subscription {
	return predicate.Subscription(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Subscription {
	return predicate.Subscription(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Subscription {
	return predicate.Subscription(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Subscription {
	return predicate.Subscription(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Subscription {
	return predicate.Subscription(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Subscription {
	return predicate.Subscription(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Subscription {
	return predicate.Subscription(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Subscription {
	return predicate.Subscription(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Subscription {
	return predicate.Subscription(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v uuid.UUID) predicate.Subscription {
	return predicate.Subscription(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserID), v))
	})
}

// FeedID applies equality check predicate on the "feed_id" field. It's identical to FeedIDEQ.
func FeedID(v uuid.UUID) predicate.Subscription {
	return predicate.Subscription(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFeedID), v))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Subscription {
	return predicate.Subscription(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// Group applies equality check predicate on the "group" field. It's identical to GroupEQ.
func Group(v string) predicate.Subscription {
	return predicate.Subscription(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldGroup), v))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Subscription {
	return predicate.Subscription(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v uuid.UUID) predicate.Subscription {
	return predicate.Subscription(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserID), v))
	})
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v uuid.UUID) predicate.Subscription {
	return predicate.Subscription(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUserID), v))
	})
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...uuid.UUID) predicate.Subscription {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Subscription(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUserID), v...))
	})
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...uuid.UUID) predicate.Subscription {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Subscription(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUserID), v...))
	})
}

// FeedIDEQ applies the EQ predicate on the "feed_id" field.
func FeedIDEQ(v uuid.UUID) predicate.Subscription {
	return predicate.Subscription(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFeedID), v))
	})
}

// FeedIDNEQ applies the NEQ predicate on the "feed_id" field.
func FeedIDNEQ(v uuid.UUID) predicate.Subscription {
	return predicate.Subscription(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldFeedID), v))
	})
}

// FeedIDIn applies the In predicate on the "feed_id" field.
func FeedIDIn(vs ...uuid.UUID) predicate.Subscription {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Subscription(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldFeedID), v...))
	})
}

// FeedIDNotIn applies the NotIn predicate on the "feed_id" field.
func FeedIDNotIn(vs ...uuid.UUID) predicate.Subscription {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Subscription(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldFeedID), v...))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Subscription {
	return predicate.Subscription(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Subscription {
	return predicate.Subscription(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Subscription {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Subscription(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldName), v...))
	})
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Subscription {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Subscription(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldName), v...))
	})
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Subscription {
	return predicate.Subscription(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Subscription {
	return predicate.Subscription(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Subscription {
	return predicate.Subscription(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Subscription {
	return predicate.Subscription(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Subscription {
	return predicate.Subscription(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Subscription {
	return predicate.Subscription(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Subscription {
	return predicate.Subscription(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Subscription {
	return predicate.Subscription(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Subscription {
	return predicate.Subscription(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// GroupEQ applies the EQ predicate on the "group" field.
func GroupEQ(v string) predicate.Subscription {
	return predicate.Subscription(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldGroup), v))
	})
}

// GroupNEQ applies the NEQ predicate on the "group" field.
func GroupNEQ(v string) predicate.Subscription {
	return predicate.Subscription(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldGroup), v))
	})
}

// GroupIn applies the In predicate on the "group" field.
func GroupIn(vs ...string) predicate.Subscription {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Subscription(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldGroup), v...))
	})
}

// GroupNotIn applies the NotIn predicate on the "group" field.
func GroupNotIn(vs ...string) predicate.Subscription {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Subscription(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldGroup), v...))
	})
}

// GroupGT applies the GT predicate on the "group" field.
func GroupGT(v string) predicate.Subscription {
	return predicate.Subscription(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldGroup), v))
	})
}

// GroupGTE applies the GTE predicate on the "group" field.
func GroupGTE(v string) predicate.Subscription {
	return predicate.Subscription(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldGroup), v))
	})
}

// GroupLT applies the LT predicate on the "group" field.
func GroupLT(v string) predicate.Subscription {
	return predicate.Subscription(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldGroup), v))
	})
}

// GroupLTE applies the LTE predicate on the "group" field.
func GroupLTE(v string) predicate.Subscription {
	return predicate.Subscription(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldGroup), v))
	})
}

// GroupContains applies the Contains predicate on the "group" field.
func GroupContains(v string) predicate.Subscription {
	return predicate.Subscription(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldGroup), v))
	})
}

// GroupHasPrefix applies the HasPrefix predicate on the "group" field.
func GroupHasPrefix(v string) predicate.Subscription {
	return predicate.Subscription(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldGroup), v))
	})
}

// GroupHasSuffix applies the HasSuffix predicate on the "group" field.
func GroupHasSuffix(v string) predicate.Subscription {
	return predicate.Subscription(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldGroup), v))
	})
}

// GroupEqualFold applies the EqualFold predicate on the "group" field.
func GroupEqualFold(v string) predicate.Subscription {
	return predicate.Subscription(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldGroup), v))
	})
}

// GroupContainsFold applies the ContainsFold predicate on the "group" field.
func GroupContainsFold(v string) predicate.Subscription {
	return predicate.Subscription(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldGroup), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Subscription {
	return predicate.Subscription(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Subscription {
	return predicate.Subscription(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Subscription {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Subscription(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Subscription {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Subscription(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Subscription {
	return predicate.Subscription(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Subscription {
	return predicate.Subscription(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Subscription {
	return predicate.Subscription(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Subscription {
	return predicate.Subscription(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.Subscription {
	return predicate.Subscription(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.Subscription {
	return predicate.Subscription(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, UserTable, UserColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasFeed applies the HasEdge predicate on the "feed" edge.
func HasFeed() predicate.Subscription {
	return predicate.Subscription(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(FeedTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, FeedTable, FeedColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasFeedWith applies the HasEdge predicate on the "feed" edge with a given conditions (other predicates).
func HasFeedWith(preds ...predicate.Feed) predicate.Subscription {
	return predicate.Subscription(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(FeedInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, FeedTable, FeedColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Subscription) predicate.Subscription {
	return predicate.Subscription(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Subscription) predicate.Subscription {
	return predicate.Subscription(func(s *sql.Selector) {
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
func Not(p predicate.Subscription) predicate.Subscription {
	return predicate.Subscription(func(s *sql.Selector) {
		p(s.Not())
	})
}
