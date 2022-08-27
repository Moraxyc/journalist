// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/mrusme/journalist/ent/item"
	"github.com/mrusme/journalist/ent/read"
	"github.com/mrusme/journalist/ent/user"
)

// ReadCreate is the builder for creating a Read entity.
type ReadCreate struct {
	config
	mutation *ReadMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetUserID sets the "user_id" field.
func (rc *ReadCreate) SetUserID(u uuid.UUID) *ReadCreate {
	rc.mutation.SetUserID(u)
	return rc
}

// SetItemID sets the "item_id" field.
func (rc *ReadCreate) SetItemID(u uuid.UUID) *ReadCreate {
	rc.mutation.SetItemID(u)
	return rc
}

// SetCreatedAt sets the "created_at" field.
func (rc *ReadCreate) SetCreatedAt(t time.Time) *ReadCreate {
	rc.mutation.SetCreatedAt(t)
	return rc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (rc *ReadCreate) SetNillableCreatedAt(t *time.Time) *ReadCreate {
	if t != nil {
		rc.SetCreatedAt(*t)
	}
	return rc
}

// SetID sets the "id" field.
func (rc *ReadCreate) SetID(u uuid.UUID) *ReadCreate {
	rc.mutation.SetID(u)
	return rc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (rc *ReadCreate) SetNillableID(u *uuid.UUID) *ReadCreate {
	if u != nil {
		rc.SetID(*u)
	}
	return rc
}

// SetUser sets the "user" edge to the User entity.
func (rc *ReadCreate) SetUser(u *User) *ReadCreate {
	return rc.SetUserID(u.ID)
}

// SetItem sets the "item" edge to the Item entity.
func (rc *ReadCreate) SetItem(i *Item) *ReadCreate {
	return rc.SetItemID(i.ID)
}

// Mutation returns the ReadMutation object of the builder.
func (rc *ReadCreate) Mutation() *ReadMutation {
	return rc.mutation
}

// Save creates the Read in the database.
func (rc *ReadCreate) Save(ctx context.Context) (*Read, error) {
	var (
		err  error
		node *Read
	)
	rc.defaults()
	if len(rc.hooks) == 0 {
		if err = rc.check(); err != nil {
			return nil, err
		}
		node, err = rc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ReadMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = rc.check(); err != nil {
				return nil, err
			}
			rc.mutation = mutation
			if node, err = rc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(rc.hooks) - 1; i >= 0; i-- {
			if rc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = rc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, rc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Read)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from ReadMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (rc *ReadCreate) SaveX(ctx context.Context) *Read {
	v, err := rc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rc *ReadCreate) Exec(ctx context.Context) error {
	_, err := rc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rc *ReadCreate) ExecX(ctx context.Context) {
	if err := rc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (rc *ReadCreate) defaults() {
	if _, ok := rc.mutation.CreatedAt(); !ok {
		v := read.DefaultCreatedAt()
		rc.mutation.SetCreatedAt(v)
	}
	if _, ok := rc.mutation.ID(); !ok {
		v := read.DefaultID()
		rc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rc *ReadCreate) check() error {
	if _, ok := rc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "Read.user_id"`)}
	}
	if _, ok := rc.mutation.ItemID(); !ok {
		return &ValidationError{Name: "item_id", err: errors.New(`ent: missing required field "Read.item_id"`)}
	}
	if _, ok := rc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Read.created_at"`)}
	}
	if _, ok := rc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user", err: errors.New(`ent: missing required edge "Read.user"`)}
	}
	if _, ok := rc.mutation.ItemID(); !ok {
		return &ValidationError{Name: "item", err: errors.New(`ent: missing required edge "Read.item"`)}
	}
	return nil
}

func (rc *ReadCreate) sqlSave(ctx context.Context) (*Read, error) {
	_node, _spec := rc.createSpec()
	if err := sqlgraph.CreateNode(ctx, rc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	return _node, nil
}

func (rc *ReadCreate) createSpec() (*Read, *sqlgraph.CreateSpec) {
	var (
		_node = &Read{config: rc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: read.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: read.FieldID,
			},
		}
	)
	_spec.OnConflict = rc.conflict
	if id, ok := rc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := rc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: read.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if nodes := rc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   read.UserTable,
			Columns: []string{read.UserColumn},
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
		_node.UserID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := rc.mutation.ItemIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   read.ItemTable,
			Columns: []string{read.ItemColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: item.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.ItemID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Read.Create().
//		SetUserID(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ReadUpsert) {
//			SetUserID(v+v).
//		}).
//		Exec(ctx)
//
func (rc *ReadCreate) OnConflict(opts ...sql.ConflictOption) *ReadUpsertOne {
	rc.conflict = opts
	return &ReadUpsertOne{
		create: rc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Read.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (rc *ReadCreate) OnConflictColumns(columns ...string) *ReadUpsertOne {
	rc.conflict = append(rc.conflict, sql.ConflictColumns(columns...))
	return &ReadUpsertOne{
		create: rc,
	}
}

type (
	// ReadUpsertOne is the builder for "upsert"-ing
	//  one Read node.
	ReadUpsertOne struct {
		create *ReadCreate
	}

	// ReadUpsert is the "OnConflict" setter.
	ReadUpsert struct {
		*sql.UpdateSet
	}
)

// SetUserID sets the "user_id" field.
func (u *ReadUpsert) SetUserID(v uuid.UUID) *ReadUpsert {
	u.Set(read.FieldUserID, v)
	return u
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *ReadUpsert) UpdateUserID() *ReadUpsert {
	u.SetExcluded(read.FieldUserID)
	return u
}

// SetItemID sets the "item_id" field.
func (u *ReadUpsert) SetItemID(v uuid.UUID) *ReadUpsert {
	u.Set(read.FieldItemID, v)
	return u
}

// UpdateItemID sets the "item_id" field to the value that was provided on create.
func (u *ReadUpsert) UpdateItemID() *ReadUpsert {
	u.SetExcluded(read.FieldItemID)
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *ReadUpsert) SetCreatedAt(v time.Time) *ReadUpsert {
	u.Set(read.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *ReadUpsert) UpdateCreatedAt() *ReadUpsert {
	u.SetExcluded(read.FieldCreatedAt)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Read.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(read.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *ReadUpsertOne) UpdateNewValues() *ReadUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(read.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.Read.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *ReadUpsertOne) Ignore() *ReadUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ReadUpsertOne) DoNothing() *ReadUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ReadCreate.OnConflict
// documentation for more info.
func (u *ReadUpsertOne) Update(set func(*ReadUpsert)) *ReadUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ReadUpsert{UpdateSet: update})
	}))
	return u
}

// SetUserID sets the "user_id" field.
func (u *ReadUpsertOne) SetUserID(v uuid.UUID) *ReadUpsertOne {
	return u.Update(func(s *ReadUpsert) {
		s.SetUserID(v)
	})
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *ReadUpsertOne) UpdateUserID() *ReadUpsertOne {
	return u.Update(func(s *ReadUpsert) {
		s.UpdateUserID()
	})
}

// SetItemID sets the "item_id" field.
func (u *ReadUpsertOne) SetItemID(v uuid.UUID) *ReadUpsertOne {
	return u.Update(func(s *ReadUpsert) {
		s.SetItemID(v)
	})
}

// UpdateItemID sets the "item_id" field to the value that was provided on create.
func (u *ReadUpsertOne) UpdateItemID() *ReadUpsertOne {
	return u.Update(func(s *ReadUpsert) {
		s.UpdateItemID()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *ReadUpsertOne) SetCreatedAt(v time.Time) *ReadUpsertOne {
	return u.Update(func(s *ReadUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *ReadUpsertOne) UpdateCreatedAt() *ReadUpsertOne {
	return u.Update(func(s *ReadUpsert) {
		s.UpdateCreatedAt()
	})
}

// Exec executes the query.
func (u *ReadUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ReadCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ReadUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *ReadUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: ReadUpsertOne.ID is not supported by MySQL driver. Use ReadUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *ReadUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// ReadCreateBulk is the builder for creating many Read entities in bulk.
type ReadCreateBulk struct {
	config
	builders []*ReadCreate
	conflict []sql.ConflictOption
}

// Save creates the Read entities in the database.
func (rcb *ReadCreateBulk) Save(ctx context.Context) ([]*Read, error) {
	specs := make([]*sqlgraph.CreateSpec, len(rcb.builders))
	nodes := make([]*Read, len(rcb.builders))
	mutators := make([]Mutator, len(rcb.builders))
	for i := range rcb.builders {
		func(i int, root context.Context) {
			builder := rcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ReadMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, rcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = rcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, rcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, rcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (rcb *ReadCreateBulk) SaveX(ctx context.Context) []*Read {
	v, err := rcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rcb *ReadCreateBulk) Exec(ctx context.Context) error {
	_, err := rcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rcb *ReadCreateBulk) ExecX(ctx context.Context) {
	if err := rcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Read.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ReadUpsert) {
//			SetUserID(v+v).
//		}).
//		Exec(ctx)
//
func (rcb *ReadCreateBulk) OnConflict(opts ...sql.ConflictOption) *ReadUpsertBulk {
	rcb.conflict = opts
	return &ReadUpsertBulk{
		create: rcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Read.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (rcb *ReadCreateBulk) OnConflictColumns(columns ...string) *ReadUpsertBulk {
	rcb.conflict = append(rcb.conflict, sql.ConflictColumns(columns...))
	return &ReadUpsertBulk{
		create: rcb,
	}
}

// ReadUpsertBulk is the builder for "upsert"-ing
// a bulk of Read nodes.
type ReadUpsertBulk struct {
	create *ReadCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Read.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(read.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *ReadUpsertBulk) UpdateNewValues() *ReadUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(read.FieldID)
				return
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Read.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *ReadUpsertBulk) Ignore() *ReadUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ReadUpsertBulk) DoNothing() *ReadUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ReadCreateBulk.OnConflict
// documentation for more info.
func (u *ReadUpsertBulk) Update(set func(*ReadUpsert)) *ReadUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ReadUpsert{UpdateSet: update})
	}))
	return u
}

// SetUserID sets the "user_id" field.
func (u *ReadUpsertBulk) SetUserID(v uuid.UUID) *ReadUpsertBulk {
	return u.Update(func(s *ReadUpsert) {
		s.SetUserID(v)
	})
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *ReadUpsertBulk) UpdateUserID() *ReadUpsertBulk {
	return u.Update(func(s *ReadUpsert) {
		s.UpdateUserID()
	})
}

// SetItemID sets the "item_id" field.
func (u *ReadUpsertBulk) SetItemID(v uuid.UUID) *ReadUpsertBulk {
	return u.Update(func(s *ReadUpsert) {
		s.SetItemID(v)
	})
}

// UpdateItemID sets the "item_id" field to the value that was provided on create.
func (u *ReadUpsertBulk) UpdateItemID() *ReadUpsertBulk {
	return u.Update(func(s *ReadUpsert) {
		s.UpdateItemID()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *ReadUpsertBulk) SetCreatedAt(v time.Time) *ReadUpsertBulk {
	return u.Update(func(s *ReadUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *ReadUpsertBulk) UpdateCreatedAt() *ReadUpsertBulk {
	return u.Update(func(s *ReadUpsert) {
		s.UpdateCreatedAt()
	})
}

// Exec executes the query.
func (u *ReadUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the ReadCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ReadCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ReadUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
