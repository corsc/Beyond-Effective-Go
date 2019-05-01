package _3_reduce_lock_coverage

import (
	"context"
)

func UpdateCountryByIDV2(ctx context.Context, countryID int) error {
	// load country from DB
	query := "SELECT name, population FROM countries WHERE ID = ?"
	result := db.QueryRowContext(ctx, query, countryID)

	country := Country{}
	err := result.Scan(&country.Name, &country.Population)
	if err != nil {
		return err
	}

	// update the cache
	cacheMutex.Lock()
	cache[countryID] = country
	cacheMutex.Unlock()

	return nil
}
