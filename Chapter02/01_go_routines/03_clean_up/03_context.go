package _3_clean_up

import (
	"context"
	"database/sql"
	"time"
)

var db *sql.DB

func Example3() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	go func() {
		// call to database
		query := "SELECT name, population FROM countries"
		results, err := db.QueryContext(ctx, query)
		if err != nil {
			// kill the goroutine and free resources
			return
		}

		useResults(results)
	}()

}

func useResults(rows *sql.Rows) {
	// not implemented
}
