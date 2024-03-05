package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type (
	// Configuration contains database configuration.
	Configuration struct {
		DbName   string
		Host     string
		User     string
		Password string
		Port     int
	}

	// Connector wraps the basic database operations.
	Connector interface {
		Open(ctx context.Context) error
		Close()
		DBChecker
		TransactionOpener
		QueryExecutor
	}

	// DBChecker checks if the database is healthy.
	DBChecker interface {
		Check(ctx context.Context) error
	}

	// QueryExecutor executes queries on the database.
	QueryExecutor interface {
		CopyFrom(
			ctx context.Context,
			tableName pgx.Identifier,
			columnNames []string,
			rowSrc pgx.CopyFromSource) (int64, error)
		Exec(ctx context.Context, sql string, args ...any) (int64, error)
		Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
		QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
	}

	// TransactionOpener opens transactions on the database.
	TransactionOpener interface {
		Begin(ctx context.Context) (*Transaction, error)
		BeginTx(ctx context.Context, txOptions TxOptions) (*Transaction, error)
	}

	// Client connects to the database.
	Client struct {
		ConnectionString string
		pool             *pgxpool.Pool
	}
)

var _ Connector = (*Client)(nil)

// NewClient creates a database client.
func NewClient(c Configuration) *Client {
	connString := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		c.User, c.Password, c.Host, c.Port, c.DbName,
	)

	return &Client{
		ConnectionString: connString,
	}
}

// Open opens a database connection.
func (c *Client) Open(ctx context.Context) error {
	if c.pool != nil {
		if err := c.pool.Ping(ctx); err != nil {
			return fmt.Errorf("failed to reach database: %w", err)
		}

		c.pool.Close()
	}

	pool, err := pgxpool.New(ctx, c.ConnectionString)
	if err != nil {
		return fmt.Errorf("failed to connect to database: %s", err)
	}

	c.pool = pool

	return nil
}

// Close closes a database connection.
func (c *Client) Close() {
	c.pool.Close()
}

// Check pings the database.
func (c *Client) Check(ctx context.Context) error {
	return c.pool.Ping(ctx)
}

// Begin begins a transaction.
func (c *Client) Begin(ctx context.Context) (*Transaction, error) {
	tx, err := c.pool.Begin(ctx)
	if err != nil {
		return nil, err
	}

	return &Transaction{
		tx: tx,
	}, nil
}

// BeginTx begins a transaction with options.
func (c *Client) BeginTx(ctx context.Context, txOptions TxOptions) (*Transaction, error) {
	tx, err := c.pool.BeginTx(ctx, pgx.TxOptions{
		IsoLevel:       pgx.TxIsoLevel(txOptions.IsoLevel),
		AccessMode:     pgx.TxAccessMode(txOptions.AccessMode),
		DeferrableMode: pgx.TxDeferrableMode(txOptions.DeferrableMode),
	})
	if err != nil {
		return nil, err
	}

	return &Transaction{
		tx: tx,
	}, nil
}

// CopyFrom copies data from a source into a table.
func (c *Client) CopyFrom(
	ctx context.Context,
	tableName pgx.Identifier,
	columnNames []string,
	rowSrc pgx.CopyFromSource) (int64, error) {
	return c.pool.CopyFrom(ctx, tableName, columnNames, rowSrc)
}

// Exec executes a SQL query.
func (c *Client) Exec(ctx context.Context, sql string, args ...any) (int64, error) {
	result, err := c.pool.Exec(ctx, sql, args...)

	return result.RowsAffected(), err
}

// Query executes a query on the database and returns rows.
func (c *Client) Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error) {
	return c.pool.Query(ctx, sql, args...)
}

// QueryRow executes a query on the database and returns a single row.
func (c *Client) QueryRow(ctx context.Context, sql string, args ...any) pgx.Row {
	return c.pool.QueryRow(ctx, sql, args...)
}
