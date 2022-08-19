// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/mrusme/journalist/ent/feed"
	"github.com/mrusme/journalist/ent/predicate"
	"github.com/mrusme/journalist/ent/subscription"
	"github.com/mrusme/journalist/ent/user"
)

// SubscriptionUpdate is the builder for updating Subscription entities.
type SubscriptionUpdate struct {
	config
	hooks    []Hook
	mutation *SubscriptionMutation
}

// Where appends a list predicates to the SubscriptionUpdate builder.
func (su *SubscriptionUpdate) Where(ps ...predicate.Subscription) *SubscriptionUpdate {
	su.mutation.Where(ps...)
	return su
}

// SetUserID sets the "user_id" field.
func (su *SubscriptionUpdate) SetUserID(u uuid.UUID) *SubscriptionUpdate {
	su.mutation.SetUserID(u)
	return su
}

// SetFeedID sets the "feed_id" field.
func (su *SubscriptionUpdate) SetFeedID(u uuid.UUID) *SubscriptionUpdate {
	su.mutation.SetFeedID(u)
	return su
}

// SetGroup sets the "group" field.
func (su *SubscriptionUpdate) SetGroup(s string) *SubscriptionUpdate {
	su.mutation.SetGroup(s)
	return su
}

// SetCreatedAt sets the "created_at" field.
func (su *SubscriptionUpdate) SetCreatedAt(t time.Time) *SubscriptionUpdate {
	su.mutation.SetCreatedAt(t)
	return su
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (su *SubscriptionUpdate) SetNillableCreatedAt(t *time.Time) *SubscriptionUpdate {
	if t != nil {
		su.SetCreatedAt(*t)
	}
	return su
}

// SetUser sets the "user" edge to the User entity.
func (su *SubscriptionUpdate) SetUser(u *User) *SubscriptionUpdate {
	return su.SetUserID(u.ID)
}

// SetFeed sets the "feed" edge to the Feed entity.
func (su *SubscriptionUpdate) SetFeed(f *Feed) *SubscriptionUpdate {
	return su.SetFeedID(f.ID)
}

// Mutation returns the SubscriptionMutation object of the builder.
func (su *SubscriptionUpdate) Mutation() *SubscriptionMutation {
	return su.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (su *SubscriptionUpdate) ClearUser() *SubscriptionUpdate {
	su.mutation.ClearUser()
	return su
}

// ClearFeed clears the "feed" edge to the Feed entity.
func (su *SubscriptionUpdate) ClearFeed() *SubscriptionUpdate {
	su.mutation.ClearFeed()
	return su
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (su *SubscriptionUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(su.hooks) == 0 {
		if err = su.check(); err != nil {
			return 0, err
		}
		affected, err = su.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SubscriptionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = su.check(); err != nil {
				return 0, err
			}
			su.mutation = mutation
			affected, err = su.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(su.hooks) - 1; i >= 0; i-- {
			if su.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = su.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, su.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (su *SubscriptionUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *SubscriptionUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *SubscriptionUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (su *SubscriptionUpdate) check() error {
	if v, ok := su.mutation.Group(); ok {
		if err := subscription.GroupValidator(v); err != nil {
			return &ValidationError{Name: "group", err: fmt.Errorf(`ent: validator failed for field "Subscription.group": %w`, err)}
		}
	}
	if _, ok := su.mutation.UserID(); su.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Subscription.user"`)
	}
	if _, ok := su.mutation.FeedID(); su.mutation.FeedCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Subscription.feed"`)
	}
	return nil
}

func (su *SubscriptionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   subscription.Table,
			Columns: subscription.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: subscription.FieldID,
			},
		},
	}
	if ps := su.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := su.mutation.Group(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: subscription.FieldGroup,
		})
	}
	if value, ok := su.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: subscription.FieldCreatedAt,
		})
	}
	if su.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   subscription.UserTable,
			Columns: []string{subscription.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   subscription.UserTable,
			Columns: []string{subscription.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if su.mutation.FeedCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   subscription.FeedTable,
			Columns: []string{subscription.FeedColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: feed.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.FeedIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   subscription.FeedTable,
			Columns: []string{subscription.FeedColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: feed.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{subscription.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// SubscriptionUpdateOne is the builder for updating a single Subscription entity.
type SubscriptionUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *SubscriptionMutation
}

// SetUserID sets the "user_id" field.
func (suo *SubscriptionUpdateOne) SetUserID(u uuid.UUID) *SubscriptionUpdateOne {
	suo.mutation.SetUserID(u)
	return suo
}

// SetFeedID sets the "feed_id" field.
func (suo *SubscriptionUpdateOne) SetFeedID(u uuid.UUID) *SubscriptionUpdateOne {
	suo.mutation.SetFeedID(u)
	return suo
}

// SetGroup sets the "group" field.
func (suo *SubscriptionUpdateOne) SetGroup(s string) *SubscriptionUpdateOne {
	suo.mutation.SetGroup(s)
	return suo
}

// SetCreatedAt sets the "created_at" field.
func (suo *SubscriptionUpdateOne) SetCreatedAt(t time.Time) *SubscriptionUpdateOne {
	suo.mutation.SetCreatedAt(t)
	return suo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (suo *SubscriptionUpdateOne) SetNillableCreatedAt(t *time.Time) *SubscriptionUpdateOne {
	if t != nil {
		suo.SetCreatedAt(*t)
	}
	return suo
}

// SetUser sets the "user" edge to the User entity.
func (suo *SubscriptionUpdateOne) SetUser(u *User) *SubscriptionUpdateOne {
	return suo.SetUserID(u.ID)
}

// SetFeed sets the "feed" edge to the Feed entity.
func (suo *SubscriptionUpdateOne) SetFeed(f *Feed) *SubscriptionUpdateOne {
	return suo.SetFeedID(f.ID)
}

// Mutation returns the SubscriptionMutation object of the builder.
func (suo *SubscriptionUpdateOne) Mutation() *SubscriptionMutation {
	return suo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (suo *SubscriptionUpdateOne) ClearUser() *SubscriptionUpdateOne {
	suo.mutation.ClearUser()
	return suo
}

// ClearFeed clears the "feed" edge to the Feed entity.
func (suo *SubscriptionUpdateOne) ClearFeed() *SubscriptionUpdateOne {
	suo.mutation.ClearFeed()
	return suo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (suo *SubscriptionUpdateOne) Select(field string, fields ...string) *SubscriptionUpdateOne {
	suo.fields = append([]string{field}, fields...)
	return suo
}

// Save executes the query and returns the updated Subscription entity.
func (suo *SubscriptionUpdateOne) Save(ctx context.Context) (*Subscription, error) {
	var (
		err  error
		node *Subscription
	)
	if len(suo.hooks) == 0 {
		if err = suo.check(); err != nil {
			return nil, err
		}
		node, err = suo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SubscriptionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = suo.check(); err != nil {
				return nil, err
			}
			suo.mutation = mutation
			node, err = suo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(suo.hooks) - 1; i >= 0; i-- {
			if suo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = suo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, suo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Subscription)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from SubscriptionMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (suo *SubscriptionUpdateOne) SaveX(ctx context.Context) *Subscription {
	node, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (suo *SubscriptionUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *SubscriptionUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (suo *SubscriptionUpdateOne) check() error {
	if v, ok := suo.mutation.Group(); ok {
		if err := subscription.GroupValidator(v); err != nil {
			return &ValidationError{Name: "group", err: fmt.Errorf(`ent: validator failed for field "Subscription.group": %w`, err)}
		}
	}
	if _, ok := suo.mutation.UserID(); suo.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Subscription.user"`)
	}
	if _, ok := suo.mutation.FeedID(); suo.mutation.FeedCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Subscription.feed"`)
	}
	return nil
}

func (suo *SubscriptionUpdateOne) sqlSave(ctx context.Context) (_node *Subscription, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   subscription.Table,
			Columns: subscription.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: subscription.FieldID,
			},
		},
	}
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Subscription.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := suo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, subscription.FieldID)
		for _, f := range fields {
			if !subscription.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != subscription.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := suo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := suo.mutation.Group(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: subscription.FieldGroup,
		})
	}
	if value, ok := suo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: subscription.FieldCreatedAt,
		})
	}
	if suo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   subscription.UserTable,
			Columns: []string{subscription.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   subscription.UserTable,
			Columns: []string{subscription.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if suo.mutation.FeedCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   subscription.FeedTable,
			Columns: []string{subscription.FeedColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: feed.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.FeedIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   subscription.FeedTable,
			Columns: []string{subscription.FeedColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: feed.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Subscription{config: suo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{subscription.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
