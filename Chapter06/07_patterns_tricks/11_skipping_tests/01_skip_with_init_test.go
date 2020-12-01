package _1_skipping_tests

import (
	"context"
	"database/sql"
	"testing"
	"time"
)

func TestSkipWithInit(t *testing.T) {
	if !mysqlAvailable {
		t.Skip("test skipped as MySQL is not available")
	}

	// rest of the test
}

var mysqlAvailable bool

func init() {
	db, err := sql.Open("mysql", "root@0.0.0.0/test")
	if err != nil {
		mysqlAvailable = false
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		mysqlAvailable = false
		return
	}

	mysqlAvailable = true
}
