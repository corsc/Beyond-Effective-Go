package _3_reduce_lock_coverage

import (
	"context"
	"database/sql"
	"sync"
)

var (
	cache      = map[int]Country{}
	cacheMutex = &sync.Mutex{}

	db *sql.DB
)

func UpdateCountryByIDV1(ctx context.Context, countryID int) error {
	cacheMutex.Lock()
	defer cacheMutex.Unlock()

	// load country from DB
	query := "SELECT name, population FROM countries WHERE ID = ?"
	result := db.QueryRowContext(ctx, query, countryID)

	country := Country{}
	err := result.Scan(&country.Name, &country.Population)
	if err != nil {
		return err
	}

	// update the cache
	cache[countryID] = country

	return nil
}

type Country struct {
	Name       string
	Population int
}
