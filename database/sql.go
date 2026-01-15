package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)
//run go get github.com/lib/pq
//also tellme what postgrse stuff to run on my terinal to see my table

type Product struct {
	Name string
	Price float64
	Available bool
}

func main() {
	connStr := "postgres://postgres:postgres@localhost:5432/testdb?sslmode=disable"

	dp, err := sql.Open("postgres", connStr)

	defer dp.Close()

	if err != nil {
		log.Fatal(err)
	}

	if err = dp.Ping(); err != nil {
		log.Fatal(err)
	}

	createProductTable(dp)

	product := Product{Name: "Book", Price: 15.55, Available: true}

	pk := insertProduct(dp, product)

	fmt.Printf("ID = %d\n", pk)
}

func createProductTable(db *sql.DB) {
/*
Product Table
- ID
- Name
- Price
- Available
- Data Created
*/

query := `CREATE TABLE IF NOT EXISTS product (
	id SERIAL PRIMARY KEY,
	name VARCHAR(100) NOT NULL,
	price NUMERIC(10, 2) NOT NULL,
	available BOOLEAN DEFAULT true,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	)`

	_, err := db.Exec(query)

	if err != nil {
		log.Fatal(err)
	}
}

func insertProduct(dp *sql.DB, product Product) int {
	query := `INSERT INTO product (name, price, available)
		VALUES ($1, $2, $3) RETURNING id`
	
		var pk int
		err := dp.QueryRow(query, product.Name, product.Price, product.Available).Scan(&pk)

		if err != nil {
			log.Fatal(err)
		}

		return pk
}