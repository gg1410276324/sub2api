package service

import (
	"testing"
	"time"
)

func TestAffiliateSalesPeriodStarts(t *testing.T) {
	now := time.Date(2026, time.July, 26, 20, 0, 0, 0, time.FixedZone("UTC+8", 8*60*60))
	week, month := affiliateSalesPeriodStarts(now, "Asia/Shanghai")

	wantWeek := time.Date(2026, time.July, 20, 0, 0, 0, 0, time.FixedZone("UTC+8", 8*60*60)).UTC()
	wantMonth := time.Date(2026, time.July, 1, 0, 0, 0, 0, time.FixedZone("UTC+8", 8*60*60)).UTC()
	if !week.Equal(wantWeek) || !month.Equal(wantMonth) {
		t.Fatalf("unexpected periods: week=%s month=%s", week, month)
	}
}
