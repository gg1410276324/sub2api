package repository

import (
	"context"
	"testing"
	"time"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"github.com/DATA-DOG/go-sqlmock"
	dbent "github.com/Wei-Shaw/sub2api/ent"
	"github.com/stretchr/testify/require"
)

func TestAffiliateSalesCountsCompletedBalanceOrders(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	t.Cleanup(func() { _ = db.Close() })

	client := dbent.NewClient(dbent.Driver(sql.OpenDB(dialect.Postgres, db)))
	repo := NewAffiliateRepository(client, db)
	weekStart := time.Date(2026, time.July, 20, 0, 0, 0, 0, time.UTC)
	monthStart := time.Date(2026, time.July, 1, 0, 0, 0, 0, time.UTC)

	mock.ExpectQuery(`(?s)\$2::timestamptz.*\$3::timestamptz.*FROM user_affiliates ua.*po\.order_type = 'balance'.*po\.status = 'COMPLETED'.*LEAST\(\$2::timestamptz, \$3::timestamptz\)`).
		WithArgs(int64(9), weekStart, monthStart).
		WillReturnRows(sqlmock.NewRows([]string{"invitee_count", "week_sales", "month_sales"}).AddRow(3, 12.5, 40.25))

	sales, err := repo.GetAffiliateSales(context.Background(), 9, weekStart, monthStart)

	require.NoError(t, err)
	require.Equal(t, 3, sales.InviteeCount)
	require.InDelta(t, 12.5, sales.WeekSales, 0.000001)
	require.InDelta(t, 40.25, sales.MonthSales, 0.000001)
	require.NoError(t, mock.ExpectationsWereMet())
}
