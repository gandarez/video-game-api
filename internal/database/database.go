package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
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
		Close(ctx context.Context)
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
		conn             *pgx.Conn
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
	if c.conn != nil {
		if err := c.conn.Ping(ctx); err != nil {
			return fmt.Errorf("failed to reach database: %w", err)
		}

		_ = c.conn.Close(ctx)
	}

	conn, err := pgx.Connect(ctx, c.ConnectionString)
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	c.conn = conn

	return nil
}

// Close closes a database connection.
func (c *Client) Close(ctx context.Context) {
	_ = c.conn.Close(ctx)
}

// Check pings the database.
func (c *Client) Check(ctx context.Context) error {
	return c.conn.Ping(ctx)
}

// Begin begins a transaction.
func (c *Client) Begin(ctx context.Context) (*Transaction, error) {
	tx, err := c.conn.Begin(ctx)
	if err != nil {
		return nil, err
	}

	return &Transaction{
		tx: tx,
	}, nil
}

// BeginTx begins a transaction with options.
func (c *Client) BeginTx(ctx context.Context, txOptions TxOptions) (*Transaction, error) {
	tx, err := c.conn.BeginTx(ctx, pgx.TxOptions{
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
	return c.conn.CopyFrom(ctx, tableName, columnNames, rowSrc)
}

// Exec executes a SQL query.
func (c *Client) Exec(ctx context.Context, sql string, args ...any) (int64, error) {
	result, err := c.conn.Exec(ctx, sql, args...)

	return result.RowsAffected(), err
}

// Query executes a query on the database and returns rows.
func (c *Client) Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error) {
	return c.conn.Query(ctx, sql, args...)
}

// QueryRow executes a query on the database and returns a single row.
func (c *Client) QueryRow(ctx context.Context, sql string, args ...any) pgx.Row {
	return c.conn.QueryRow(ctx, sql, args...)
}
