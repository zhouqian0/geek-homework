// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"dipper/app/server/service/internal/data/ent/migrate"

	"dipper/app/server/service/internal/data/ent/blacklist"
	"dipper/app/server/service/internal/data/ent/host"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Blacklist is the client for interacting with the Blacklist builders.
	Blacklist *BlacklistClient
	// Host is the client for interacting with the Host builders.
	Host *HostClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Blacklist = NewBlacklistClient(c.config)
	c.Host = NewHostClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:       ctx,
		config:    cfg,
		Blacklist: NewBlacklistClient(cfg),
		Host:      NewHostClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:       ctx,
		config:    cfg,
		Blacklist: NewBlacklistClient(cfg),
		Host:      NewHostClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Blacklist.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Blacklist.Use(hooks...)
	c.Host.Use(hooks...)
}

// BlacklistClient is a client for the Blacklist schema.
type BlacklistClient struct {
	config
}

// NewBlacklistClient returns a client for the Blacklist from the given config.
func NewBlacklistClient(c config) *BlacklistClient {
	return &BlacklistClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `blacklist.Hooks(f(g(h())))`.
func (c *BlacklistClient) Use(hooks ...Hook) {
	c.hooks.Blacklist = append(c.hooks.Blacklist, hooks...)
}

// Create returns a create builder for Blacklist.
func (c *BlacklistClient) Create() *BlacklistCreate {
	mutation := newBlacklistMutation(c.config, OpCreate)
	return &BlacklistCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Blacklist entities.
func (c *BlacklistClient) CreateBulk(builders ...*BlacklistCreate) *BlacklistCreateBulk {
	return &BlacklistCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Blacklist.
func (c *BlacklistClient) Update() *BlacklistUpdate {
	mutation := newBlacklistMutation(c.config, OpUpdate)
	return &BlacklistUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *BlacklistClient) UpdateOne(b *Blacklist) *BlacklistUpdateOne {
	mutation := newBlacklistMutation(c.config, OpUpdateOne, withBlacklist(b))
	return &BlacklistUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *BlacklistClient) UpdateOneID(id int64) *BlacklistUpdateOne {
	mutation := newBlacklistMutation(c.config, OpUpdateOne, withBlacklistID(id))
	return &BlacklistUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Blacklist.
func (c *BlacklistClient) Delete() *BlacklistDelete {
	mutation := newBlacklistMutation(c.config, OpDelete)
	return &BlacklistDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *BlacklistClient) DeleteOne(b *Blacklist) *BlacklistDeleteOne {
	return c.DeleteOneID(b.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *BlacklistClient) DeleteOneID(id int64) *BlacklistDeleteOne {
	builder := c.Delete().Where(blacklist.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &BlacklistDeleteOne{builder}
}

// Query returns a query builder for Blacklist.
func (c *BlacklistClient) Query() *BlacklistQuery {
	return &BlacklistQuery{
		config: c.config,
	}
}

// Get returns a Blacklist entity by its id.
func (c *BlacklistClient) Get(ctx context.Context, id int64) (*Blacklist, error) {
	return c.Query().Where(blacklist.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *BlacklistClient) GetX(ctx context.Context, id int64) *Blacklist {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *BlacklistClient) Hooks() []Hook {
	return c.hooks.Blacklist
}

// HostClient is a client for the Host schema.
type HostClient struct {
	config
}

// NewHostClient returns a client for the Host from the given config.
func NewHostClient(c config) *HostClient {
	return &HostClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `host.Hooks(f(g(h())))`.
func (c *HostClient) Use(hooks ...Hook) {
	c.hooks.Host = append(c.hooks.Host, hooks...)
}

// Create returns a create builder for Host.
func (c *HostClient) Create() *HostCreate {
	mutation := newHostMutation(c.config, OpCreate)
	return &HostCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Host entities.
func (c *HostClient) CreateBulk(builders ...*HostCreate) *HostCreateBulk {
	return &HostCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Host.
func (c *HostClient) Update() *HostUpdate {
	mutation := newHostMutation(c.config, OpUpdate)
	return &HostUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *HostClient) UpdateOne(h *Host) *HostUpdateOne {
	mutation := newHostMutation(c.config, OpUpdateOne, withHost(h))
	return &HostUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *HostClient) UpdateOneID(id int64) *HostUpdateOne {
	mutation := newHostMutation(c.config, OpUpdateOne, withHostID(id))
	return &HostUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Host.
func (c *HostClient) Delete() *HostDelete {
	mutation := newHostMutation(c.config, OpDelete)
	return &HostDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *HostClient) DeleteOne(h *Host) *HostDeleteOne {
	return c.DeleteOneID(h.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *HostClient) DeleteOneID(id int64) *HostDeleteOne {
	builder := c.Delete().Where(host.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &HostDeleteOne{builder}
}

// Query returns a query builder for Host.
func (c *HostClient) Query() *HostQuery {
	return &HostQuery{
		config: c.config,
	}
}

// Get returns a Host entity by its id.
func (c *HostClient) Get(ctx context.Context, id int64) (*Host, error) {
	return c.Query().Where(host.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *HostClient) GetX(ctx context.Context, id int64) *Host {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *HostClient) Hooks() []Hook {
	return c.hooks.Host
}
