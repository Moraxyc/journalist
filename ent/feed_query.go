// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/mrusme/journalist/ent/feed"
	"github.com/mrusme/journalist/ent/item"
	"github.com/mrusme/journalist/ent/predicate"
	"github.com/mrusme/journalist/ent/subscription"
	"github.com/mrusme/journalist/ent/user"
)

// FeedQuery is the builder for querying Feed entities.
type FeedQuery struct {
	config
	limit               *int
	offset              *int
	unique              *bool
	order               []OrderFunc
	fields              []string
	predicates          []predicate.Feed
	withItems           *ItemQuery
	withSubscribedUsers *UserQuery
	withSubscriptions   *SubscriptionQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the FeedQuery builder.
func (fq *FeedQuery) Where(ps ...predicate.Feed) *FeedQuery {
	fq.predicates = append(fq.predicates, ps...)
	return fq
}

// Limit adds a limit step to the query.
func (fq *FeedQuery) Limit(limit int) *FeedQuery {
	fq.limit = &limit
	return fq
}

// Offset adds an offset step to the query.
func (fq *FeedQuery) Offset(offset int) *FeedQuery {
	fq.offset = &offset
	return fq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (fq *FeedQuery) Unique(unique bool) *FeedQuery {
	fq.unique = &unique
	return fq
}

// Order adds an order step to the query.
func (fq *FeedQuery) Order(o ...OrderFunc) *FeedQuery {
	fq.order = append(fq.order, o...)
	return fq
}

// QueryItems chains the current query on the "items" edge.
func (fq *FeedQuery) QueryItems() *ItemQuery {
	query := &ItemQuery{config: fq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := fq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := fq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(feed.Table, feed.FieldID, selector),
			sqlgraph.To(item.Table, item.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, feed.ItemsTable, feed.ItemsColumn),
		)
		fromU = sqlgraph.SetNeighbors(fq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QuerySubscribedUsers chains the current query on the "subscribed_users" edge.
func (fq *FeedQuery) QuerySubscribedUsers() *UserQuery {
	query := &UserQuery{config: fq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := fq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := fq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(feed.Table, feed.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, feed.SubscribedUsersTable, feed.SubscribedUsersPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(fq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QuerySubscriptions chains the current query on the "subscriptions" edge.
func (fq *FeedQuery) QuerySubscriptions() *SubscriptionQuery {
	query := &SubscriptionQuery{config: fq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := fq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := fq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(feed.Table, feed.FieldID, selector),
			sqlgraph.To(subscription.Table, subscription.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, feed.SubscriptionsTable, feed.SubscriptionsColumn),
		)
		fromU = sqlgraph.SetNeighbors(fq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Feed entity from the query.
// Returns a *NotFoundError when no Feed was found.
func (fq *FeedQuery) First(ctx context.Context) (*Feed, error) {
	nodes, err := fq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{feed.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (fq *FeedQuery) FirstX(ctx context.Context) *Feed {
	node, err := fq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Feed ID from the query.
// Returns a *NotFoundError when no Feed ID was found.
func (fq *FeedQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = fq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{feed.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (fq *FeedQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := fq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Feed entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Feed entity is found.
// Returns a *NotFoundError when no Feed entities are found.
func (fq *FeedQuery) Only(ctx context.Context) (*Feed, error) {
	nodes, err := fq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{feed.Label}
	default:
		return nil, &NotSingularError{feed.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (fq *FeedQuery) OnlyX(ctx context.Context) *Feed {
	node, err := fq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Feed ID in the query.
// Returns a *NotSingularError when more than one Feed ID is found.
// Returns a *NotFoundError when no entities are found.
func (fq *FeedQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = fq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{feed.Label}
	default:
		err = &NotSingularError{feed.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (fq *FeedQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := fq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Feeds.
func (fq *FeedQuery) All(ctx context.Context) ([]*Feed, error) {
	if err := fq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return fq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (fq *FeedQuery) AllX(ctx context.Context) []*Feed {
	nodes, err := fq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Feed IDs.
func (fq *FeedQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := fq.Select(feed.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (fq *FeedQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := fq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (fq *FeedQuery) Count(ctx context.Context) (int, error) {
	if err := fq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return fq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (fq *FeedQuery) CountX(ctx context.Context) int {
	count, err := fq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (fq *FeedQuery) Exist(ctx context.Context) (bool, error) {
	if err := fq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return fq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (fq *FeedQuery) ExistX(ctx context.Context) bool {
	exist, err := fq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the FeedQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (fq *FeedQuery) Clone() *FeedQuery {
	if fq == nil {
		return nil
	}
	return &FeedQuery{
		config:              fq.config,
		limit:               fq.limit,
		offset:              fq.offset,
		order:               append([]OrderFunc{}, fq.order...),
		predicates:          append([]predicate.Feed{}, fq.predicates...),
		withItems:           fq.withItems.Clone(),
		withSubscribedUsers: fq.withSubscribedUsers.Clone(),
		withSubscriptions:   fq.withSubscriptions.Clone(),
		// clone intermediate query.
		sql:    fq.sql.Clone(),
		path:   fq.path,
		unique: fq.unique,
	}
}

// WithItems tells the query-builder to eager-load the nodes that are connected to
// the "items" edge. The optional arguments are used to configure the query builder of the edge.
func (fq *FeedQuery) WithItems(opts ...func(*ItemQuery)) *FeedQuery {
	query := &ItemQuery{config: fq.config}
	for _, opt := range opts {
		opt(query)
	}
	fq.withItems = query
	return fq
}

// WithSubscribedUsers tells the query-builder to eager-load the nodes that are connected to
// the "subscribed_users" edge. The optional arguments are used to configure the query builder of the edge.
func (fq *FeedQuery) WithSubscribedUsers(opts ...func(*UserQuery)) *FeedQuery {
	query := &UserQuery{config: fq.config}
	for _, opt := range opts {
		opt(query)
	}
	fq.withSubscribedUsers = query
	return fq
}

// WithSubscriptions tells the query-builder to eager-load the nodes that are connected to
// the "subscriptions" edge. The optional arguments are used to configure the query builder of the edge.
func (fq *FeedQuery) WithSubscriptions(opts ...func(*SubscriptionQuery)) *FeedQuery {
	query := &SubscriptionQuery{config: fq.config}
	for _, opt := range opts {
		opt(query)
	}
	fq.withSubscriptions = query
	return fq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		URL string `json:"url,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Feed.Query().
//		GroupBy(feed.FieldURL).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (fq *FeedQuery) GroupBy(field string, fields ...string) *FeedGroupBy {
	grbuild := &FeedGroupBy{config: fq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := fq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return fq.sqlQuery(ctx), nil
	}
	grbuild.label = feed.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		URL string `json:"url,omitempty"`
//	}
//
//	client.Feed.Query().
//		Select(feed.FieldURL).
//		Scan(ctx, &v)
//
func (fq *FeedQuery) Select(fields ...string) *FeedSelect {
	fq.fields = append(fq.fields, fields...)
	selbuild := &FeedSelect{FeedQuery: fq}
	selbuild.label = feed.Label
	selbuild.flds, selbuild.scan = &fq.fields, selbuild.Scan
	return selbuild
}

func (fq *FeedQuery) prepareQuery(ctx context.Context) error {
	for _, f := range fq.fields {
		if !feed.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if fq.path != nil {
		prev, err := fq.path(ctx)
		if err != nil {
			return err
		}
		fq.sql = prev
	}
	return nil
}

func (fq *FeedQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Feed, error) {
	var (
		nodes       = []*Feed{}
		_spec       = fq.querySpec()
		loadedTypes = [3]bool{
			fq.withItems != nil,
			fq.withSubscribedUsers != nil,
			fq.withSubscriptions != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*Feed).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &Feed{config: fq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, fq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := fq.withItems; query != nil {
		if err := fq.loadItems(ctx, query, nodes,
			func(n *Feed) { n.Edges.Items = []*Item{} },
			func(n *Feed, e *Item) { n.Edges.Items = append(n.Edges.Items, e) }); err != nil {
			return nil, err
		}
	}
	if query := fq.withSubscribedUsers; query != nil {
		if err := fq.loadSubscribedUsers(ctx, query, nodes,
			func(n *Feed) { n.Edges.SubscribedUsers = []*User{} },
			func(n *Feed, e *User) { n.Edges.SubscribedUsers = append(n.Edges.SubscribedUsers, e) }); err != nil {
			return nil, err
		}
	}
	if query := fq.withSubscriptions; query != nil {
		if err := fq.loadSubscriptions(ctx, query, nodes,
			func(n *Feed) { n.Edges.Subscriptions = []*Subscription{} },
			func(n *Feed, e *Subscription) { n.Edges.Subscriptions = append(n.Edges.Subscriptions, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (fq *FeedQuery) loadItems(ctx context.Context, query *ItemQuery, nodes []*Feed, init func(*Feed), assign func(*Feed, *Item)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*Feed)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Item(func(s *sql.Selector) {
		s.Where(sql.InValues(feed.ItemsColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.feed_items
		if fk == nil {
			return fmt.Errorf(`foreign-key "feed_items" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "feed_items" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (fq *FeedQuery) loadSubscribedUsers(ctx context.Context, query *UserQuery, nodes []*Feed, init func(*Feed), assign func(*Feed, *User)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[uuid.UUID]*Feed)
	nids := make(map[uuid.UUID]map[*Feed]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(feed.SubscribedUsersTable)
		s.Join(joinT).On(s.C(user.FieldID), joinT.C(feed.SubscribedUsersPrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(feed.SubscribedUsersPrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(feed.SubscribedUsersPrimaryKey[1]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	neighbors, err := query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
		assign := spec.Assign
		values := spec.ScanValues
		spec.ScanValues = func(columns []string) ([]interface{}, error) {
			values, err := values(columns[1:])
			if err != nil {
				return nil, err
			}
			return append([]interface{}{new(uuid.UUID)}, values...), nil
		}
		spec.Assign = func(columns []string, values []interface{}) error {
			outValue := *values[0].(*uuid.UUID)
			inValue := *values[1].(*uuid.UUID)
			if nids[inValue] == nil {
				nids[inValue] = map[*Feed]struct{}{byID[outValue]: struct{}{}}
				return assign(columns[1:], values[1:])
			}
			nids[inValue][byID[outValue]] = struct{}{}
			return nil
		}
	})
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "subscribed_users" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (fq *FeedQuery) loadSubscriptions(ctx context.Context, query *SubscriptionQuery, nodes []*Feed, init func(*Feed), assign func(*Feed, *Subscription)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*Feed)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.Where(predicate.Subscription(func(s *sql.Selector) {
		s.Where(sql.InValues(feed.SubscriptionsColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.FeedID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "feed_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (fq *FeedQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := fq.querySpec()
	_spec.Node.Columns = fq.fields
	if len(fq.fields) > 0 {
		_spec.Unique = fq.unique != nil && *fq.unique
	}
	return sqlgraph.CountNodes(ctx, fq.driver, _spec)
}

func (fq *FeedQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := fq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (fq *FeedQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   feed.Table,
			Columns: feed.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: feed.FieldID,
			},
		},
		From:   fq.sql,
		Unique: true,
	}
	if unique := fq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := fq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, feed.FieldID)
		for i := range fields {
			if fields[i] != feed.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := fq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := fq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := fq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := fq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (fq *FeedQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(fq.driver.Dialect())
	t1 := builder.Table(feed.Table)
	columns := fq.fields
	if len(columns) == 0 {
		columns = feed.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if fq.sql != nil {
		selector = fq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if fq.unique != nil && *fq.unique {
		selector.Distinct()
	}
	for _, p := range fq.predicates {
		p(selector)
	}
	for _, p := range fq.order {
		p(selector)
	}
	if offset := fq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := fq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// FeedGroupBy is the group-by builder for Feed entities.
type FeedGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (fgb *FeedGroupBy) Aggregate(fns ...AggregateFunc) *FeedGroupBy {
	fgb.fns = append(fgb.fns, fns...)
	return fgb
}

// Scan applies the group-by query and scans the result into the given value.
func (fgb *FeedGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := fgb.path(ctx)
	if err != nil {
		return err
	}
	fgb.sql = query
	return fgb.sqlScan(ctx, v)
}

func (fgb *FeedGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range fgb.fields {
		if !feed.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := fgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := fgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (fgb *FeedGroupBy) sqlQuery() *sql.Selector {
	selector := fgb.sql.Select()
	aggregation := make([]string, 0, len(fgb.fns))
	for _, fn := range fgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(fgb.fields)+len(fgb.fns))
		for _, f := range fgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(fgb.fields...)...)
}

// FeedSelect is the builder for selecting fields of Feed entities.
type FeedSelect struct {
	*FeedQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (fs *FeedSelect) Scan(ctx context.Context, v interface{}) error {
	if err := fs.prepareQuery(ctx); err != nil {
		return err
	}
	fs.sql = fs.FeedQuery.sqlQuery(ctx)
	return fs.sqlScan(ctx, v)
}

func (fs *FeedSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := fs.sql.Query()
	if err := fs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
