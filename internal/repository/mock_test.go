package repository_test

import (
	"context"
	"testing"

	"github.com/gandarez/video-game-api/internal/database"

	"github.com/jackc/pgx/v5"
	"github.com/pashagolub/pgxmock/v3"
	"github.com/stretchr/testify/require"
)

type mockDb struct {
	CopyFromFn func(
		ctx context.Context,
		tableName pgx.Identifier,
		columnNames []string,
		rowSrc pgx.CopyFromSource) (int64, error)
	ExecFn     func(ctx context.Context, sql string, args ...any) (int64, error)
	QueryFn    func(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRowFn func(ctx context.Context, sql string, args ...any) pgx.Row
}

func (mock mockDb) CopyFrom(
	ctx context.Context,
	tableName pgx.Identifier,
	columnNames []string,
	rowSrc pgx.CopyFromSource) (int64, error) {
	return mock.CopyFromFn(ctx, tableName, columnNames, rowSrc)
}

func (mock mockDb) Exec(ctx context.Context, sql string, args ...any) (int64, error) {
	return mock.ExecFn(ctx, sql, args...)
}

func (mock mockDb) Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error) {
	return mock.QueryFn(ctx, sql, args...)
}

func (mock mockDb) QueryRow(ctx context.Context, sql string, args ...any) pgx.Row {
	return mock.QueryRowFn(ctx, sql, args...)
}

type mockTx struct {
	CommitFn   func(ctx context.Context) error
	RollbackFn func(ctx context.Context) error
	CopyFromFn func(
		ctx context.Context,
		tableName pgx.Identifier,
		columnNames []string,
		rowSrc pgx.CopyFromSource) (int64, error)
	ExecFn     func(ctx context.Context, sql string, args ...any) (int64, error)
	QueryFn    func(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRowFn func(ctx context.Context, sql string, args ...any) pgx.Row
}

func (mock mockTx) Commit(ctx context.Context) error {
	return mock.CommitFn(ctx)
}

func (mock mockTx) Rollback(ctx context.Context) error {
	return mock.RollbackFn(ctx)
}

func (mock mockTx) CopyFrom(
	ctx context.Context,
	tableName pgx.Identifier,
	columnNames []string,
	rowSrc pgx.CopyFromSource) (int64, error) {
	return mock.CopyFromFn(ctx, tableName, columnNames, rowSrc)
}

func (mock mockTx) Exec(ctx context.Context, sql string, args ...any) (int64, error) {
	return mock.ExecFn(ctx, sql, args...)
}

func (mock mockTx) Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error) {
	return mock.QueryFn(ctx, sql, args...)
}

func (mock mockTx) QueryRow(ctx context.Context, sql string, args ...any) pgx.Row {
	return mock.QueryRowFn(ctx, sql, args...)
}

// nolint:revive
func setupTestDb(t *testing.T) (database.QueryExecutor, database.Transactioner, pgxmock.PgxConnIface, func()) {
	conn, err := pgxmock.NewConn()
	require.NoError(t, err)

	conn.ExpectBegin()

	tx, err := conn.Begin(context.Background())
	require.NoError(t, err)

	mockdb := mockDb{
		CopyFromFn: func(ctx context.Context, tableName pgx.Identifier, columnNames []string, rowSrc pgx.CopyFromSource) (int64, error) {
			return conn.CopyFrom(ctx, tableName, columnNames, rowSrc)
		},
		ExecFn: func(ctx context.Context, sql string, args ...any) (int64, error) {
			result, err := conn.Exec(ctx, sql, args...)
			return result.RowsAffected(), err
		},
		QueryFn: func(ctx context.Context, sql string, args ...any) (pgx.Rows, error) {
			return conn.Query(ctx, sql, args...)
		},
		QueryRowFn: func(ctx context.Context, sql string, args ...any) pgx.Row {
			return conn.QueryRow(ctx, sql, args...)
		},
	}

	mocktx := mockTx{
		CommitFn: func(ctx context.Context) error {
			return tx.Commit(ctx)
		},
		RollbackFn: func(ctx context.Context) error {
			return tx.Rollback(ctx)
		},
		CopyFromFn: func(ctx context.Context, tableName pgx.Identifier, columnNames []string, rowSrc pgx.CopyFromSource) (int64, error) {
			return tx.CopyFrom(ctx, tableName, columnNames, rowSrc)
		},
		ExecFn: func(ctx context.Context, sql string, args ...any) (int64, error) {
			result, err := tx.Exec(ctx, sql, args...)
			return result.RowsAffected(), err
		},
		QueryFn: func(ctx context.Context, sql string, args ...any) (pgx.Rows, error) {
			return tx.Query(ctx, sql, args...)
		},
		QueryRowFn: func(ctx context.Context, sql string, args ...any) pgx.Row {
			return tx.QueryRow(ctx, sql, args...)
		},
	}

	return mockdb, mocktx, conn, func() {
		err = conn.ExpectationsWereMet()
		require.NoError(t, err)
	}
}
