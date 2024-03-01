package database

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type (
	// Transactioner represents a database transaction.
	Transactioner interface {
		Commit(ctx context.Context) error
		Rollback(ctx context.Context) error
		QueryExecutor
	}

	// Transaction represents a database transaction.
	Transaction struct {
		tx pgx.Tx
	}

	// TxOptions are transaction modes within a transaction block.
	TxOptions struct {
		IsoLevel       TxIsoLevel
		AccessMode     TxAccessMode
		DeferrableMode TxDeferrableMode
	}

	// TxIsoLevel is the transaction isolation level (serializable, repeatable read, read committed or read uncommitted).
	TxIsoLevel string

	// TxAccessMode is the transaction access mode (read write or read only).
	TxAccessMode string

	// TxDeferrableMode is the transaction deferrable mode (deferrable or not deferrable).
	TxDeferrableMode string
)

var _ Transactioner = (*Transaction)(nil)

// Transaction isolation levels.
const (
	Serializable    = TxIsoLevel("serializable")
	RepeatableRead  = TxIsoLevel("repeatable read")
	ReadCommitted   = TxIsoLevel("read committed")
	ReadUncommitted = TxIsoLevel("read uncommitted")
)

// Transaction access modes.
const (
	ReadWrite = TxAccessMode("read write")
	ReadOnly  = TxAccessMode("read only")
)

// Transaction deferrable modes.
const (
	Deferrable    = TxDeferrableMode("deferrable")
	NotDeferrable = TxDeferrableMode("not deferrable")
)

// Commit commits the transaction.
func (t *Transaction) Commit(ctx context.Context) error {
	return t.tx.Commit(ctx)
}

// Rollback rolls back the transaction.
func (t *Transaction) Rollback(ctx context.Context) error {
	return t.tx.Rollback(ctx)
}

// CopyFrom copies data from a source into a table.
func (t *Transaction) CopyFrom(
	ctx context.Context,
	tableName pgx.Identifier,
	columnNames []string,
	rowSrc pgx.CopyFromSource) (int64, error) {
	return t.tx.CopyFrom(ctx, tableName, columnNames, rowSrc)
}

// Exec executes a SQL query.
func (t *Transaction) Exec(ctx context.Context, sql string, args ...any) (int64, error) {
	result, err := t.tx.Exec(ctx, sql, args...)

	return result.RowsAffected(), err
}

// Query executes a query on the database and returns rows.
func (t *Transaction) Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error) {
	return t.tx.Query(ctx, sql, args...)
}

// QueryRow executes a query on the database and returns a single row.
func (t *Transaction) QueryRow(ctx context.Context, sql string, args ...any) pgx.Row {
	return t.tx.QueryRow(ctx, sql, args...)
}
