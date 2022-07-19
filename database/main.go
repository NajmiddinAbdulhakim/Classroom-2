package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Country struct {
	ID         string
	Name       string
	Currency   string
	Population int
}

type City struct {
	ID         string
	Name       string
	Population int
	CountryID  string
}
type Response struct {
	CityId            string
	CityName          string
	CityPopulation    int64
	CountryId         string
	CountryName       string
	CountryPopulation int64
	CountryCurrency   string
}

var (
	countries = []Country{
		{
			ID:         "586ea10e-4a0f-44de-aee1-9c6064e97a0f",
			Name:       "USA",
			Currency:   "USD",
			Population: 350_000_000,
		},
		{
			ID:         "309ccb5d-e558-4100-89cb-a1c7ca788df0",
			Name:       "Uzbekistan",
			Currency:   "UZS",
			Population: 35_000_000,
		},
		{
			ID:         "02b3a3e9-6841-4154-aa88-9c5309ae8fdc",
			Name:       "Russia",
			Currency:   "RUB",
			Population: 150_000_000,
		},
		{
			ID:         "fba3c521-799e-420f-a695-a1bbad60dc82",
			Name:       "China",
			Currency:   "YUAN",
			Population: 1_500_000_000,
		},
	}

	cities = []City{
		{
			ID:         "4d2ab5ff-3015-4517-ada7-c30f941b1b48",
			Name:       "New-York",
			Population: 8_000_000,
			CountryID:  "586ea10e-4a0f-44de-aee1-9c6064e97a0f",
		},
		{
			ID:         "c285e603-ae59-45bf-98ab-7b847cf9433c",
			Name:       "Moscow",
			Population: 4_000_000,
			CountryID:  "02b3a3e9-6841-4154-aa88-9c5309ae8fdc",
		},
		{
			ID:         "17726ad8-4ed2-44ba-aa61-e62c91958a78",
			Name:       "Tashkent",
			Population: 2_000_000,
			CountryID:  "309ccb5d-e558-4100-89cb-a1c7ca788df0",
		},
		{
			ID:         "365c36a1-597f-44fa-a752-a4a5f4a6e5f0",
			Name:       "San-Francisco",
			Population: 6_000_000,
			CountryID:  "586ea10e-4a0f-44de-aee1-9c6064e97a0f",
		},
		{
			ID:         "889aa2a4-59cb-4792-97b0-cc8c1c0db00a",
			Name:       "Beijing",
			Population: 15_000_000,
			CountryID:  "fba3c521-799e-420f-a695-a1bbad60dc82",
		},
	}
)

func connectDB() *sql.DB {
	db, err := sql.Open(
		"postgres",
		"host=localhost port=5432 user=najmiddin password=1234 dbname=test sslmode=disable",
	)
	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		log.Panicf("could not ping db: %v", err)
	}
	fmt.Println("connected!")

	return db
}

func main() {
	db := connectDB()
	defer db.Close()

	// for _, c := range countries {
	// 	_, err := db.Exec(
	// 		"INSERT INTO countries VALUES ($1, $2, $3, $4)",
	// 		c.ID, c.Name, c.Population, c.Currency,
	// 	)
	// 	if err != nil {
	// 		log.Panicf("failed to insert: %v", err)
	// 	}
	// }

	// for _, c := range cities {
	// 	_, err := db.Exec(
	// 		"INSERT INTO cities VALUES ($1, $2, $3, $4)",
	// 		c.ID, c.Name, c.Population, c.CountryID,
	// 	)
	// 	if err != nil {
	// 		log.Panicf("failed to insert: %v", err)
	// 	}
	// }
	rows, err := db.Query(`SELECT 
    ci.id AS city_id,
    ci.name AS city_name,
    ci.population AS city_population,
    co.id AS country_id,
    co.name AS country_name,
    co.population AS country_population,
    co.currency AS country_currency
	FROM countries AS co
	LEFT OUTER JOIN cities AS ci
	ON ci.country_id = co.id`)
	if err != nil {
		panic(err)
	}
	var cities []Response
	for rows.Next() {
		var res Response
		err = rows.Scan(
			&res.CityId,
			&res.CityName,
			&res.CityPopulation,
			&res.CountryId,
			&res.CountryName,
			&res.CountryPopulation,
			&res.CountryCurrency,
		)
		if err != nil {
			panic(err)
		}
		cities = append(cities, res)
	}
	for _, city := range cities {
		fmt.Println(city)
	}

}
