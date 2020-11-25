package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	// pq is the libary that allows us to connect
	// to postgres with databases/sql.
	_ "github.com/lib/pq"
)
func main() {
	pgURL := os.Getenv("PGURL")
	if pgURL == "" {
		log.Fatal("PGURL empty")
	}

	// Open a database value.  Specify the postgres driver
	// for databases/sql.
	db, err := sql.Open("postgres", pgURL)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query(`
			SELECT
					sepal.length as sLength,
					sepal.width as sWidth,
					petal.length as pLength,
					petal.width as pWidth
			FROM iris 
			WHERE variety = $1`, "setosa")
    if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var (
			sLength float64
			sWidth float64
			pLength float64
			pWidth float64
		)

		if err := rows.Scan(&sLength, &sWidth, &pLength, &pWidth); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%.2f, %.2f, %.2f, %.2f\n", sLength, sWidth, pLength, pWidth)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}


}
