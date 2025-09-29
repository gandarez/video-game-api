package repository_test

import (
	"context"
	"testing"

	"github.com/gandarez/video-game-api/internal/repository"

	"github.com/jackc/pgx/v5"
	"github.com/pashagolub/pgxmock/v3"
	"github.com/stretchr/testify/require"
)

type mockDb struct {
	ExecFn     func(ctx context.Context, sql string, args ...any) (int64, error)
	QueryFn    func(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRowFn func(ctx context.Context, sql string, args ...any) pgx.Row
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

// nolint:revive
func setupTestDb(t *testing.T) (repository.DatabaseQueryExecutor, pgxmock.PgxConnIface, func()) {
	conn, err := pgxmock.NewConn()
	require.NoError(t, err)

	mockdb := mockDb{
		ExecFn: func(ctx context.Context, sql string, args ...any) (int64, error) {
			result, err := conn.Exec(ctx, sql, args...)
			return result.RowsAffected(), err
		},
		QueryRowFn: func(ctx context.Context, sql string, args ...any) pgx.Row {
			return conn.QueryRow(ctx, sql, args...)
		},
	}

	return mockdb, conn, func() {
		err = conn.ExpectationsWereMet()
		require.NoError(t, err)
	}
}
