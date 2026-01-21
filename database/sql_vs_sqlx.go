package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

/*
=====================================
SCHEMA
=====================================
*/

var schema = `
CREATE TABLE IF NOT EXISTS authors (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT NOT NULL,
	email TEXT NOT NULL
);
`

/*
=====================================
STRUCTS
=====================================
*/

type Author struct {
	ID    int    `db:"id"`
	Name  string `db:"name"`
	Email string `db:"email"`
}

/*
=====================================
MAIN
=====================================
*/

func main() {
	fmt.Println("========== SQL vs SQLX DEMO ==========")

	// -------------------------------------------------
	// database/sql
	// -------------------------------------------------
	fmt.Println("\n--- USING database/sql ---")

	sqlDB, err := sql.Open("sqlite3", "demo.db")
	if err != nil {
		log.Fatal(err)
	}
	defer sqlDB.Close()

	_, err = sqlDB.Exec(schema)
	if err != nil {
		log.Fatal(err)
	}

	_, err = sqlDB.Exec(
		"INSERT INTO authors (name, email) VALUES (?, ?)",
		"J.K. Rowling",
		"jk@hogwarts.com",
	)
	if err != nil {
		log.Fatal(err)
	}

	// Query ONE row
	var id int
	var name, email string

	row := sqlDB.QueryRow(
		"SELECT id, name, email FROM authors WHERE id = ?",
		1,
	)
	err = row.Scan(&id, &name, &email)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("database/sql → Author: id=%d name=%s email=%s\n", id, name, email)

	// Query MANY rows
	rows, err := sqlDB.Query("SELECT id, name, email FROM authors")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&id, &name, &email)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("database/sql → Row: %d %s %s\n", id, name, email)
	}

	// -------------------------------------------------
	// sqlx
	// -------------------------------------------------
	fmt.Println("\n--- USING sqlx ---")

	dbx, err := sqlx.Connect("sqlite3", "demo.db")
	if err != nil {
		log.Fatal(err)
	}
	defer dbx.Close()

	// Query ONE row → struct
	var author Author
	err = dbx.Get(&author, "SELECT * FROM authors WHERE id = ?", 1)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("sqlx → Author struct: %+v\n", author)

	// Query MANY rows → slice
	var authors []Author
	err = dbx.Select(&authors, "SELECT * FROM authors")
	if err != nil {
		log.Fatal(err)
	}

	for _, a := range authors {
		fmt.Printf("sqlx → Row struct: %+v\n", a)
	}

	// Named Query
	rowsX, err := dbx.NamedQuery(
		"SELECT * FROM authors WHERE name = :name",
		map[string]interface{}{"name": "J.K. Rowling"},
	)
	if err != nil {
		log.Fatal(err)
	}
	defer rowsX.Close()

	for rowsX.Next() {
		var a Author
		rowsX.StructScan(&a)
		fmt.Printf("sqlx → NamedQuery result: %+v\n", a)
	}

	// IN clause
	ids := []int{1}

	query, args, err := sqlx.In(
		"SELECT * FROM authors WHERE id IN (?)",
		ids,
	)
	if err != nil {
		log.Fatal(err)
	}
	query = dbx.Rebind(query)

	var inAuthors []Author
	err = dbx.Select(&inAuthors, query, args...)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("sqlx → IN clause result:", inAuthors)

	fmt.Println("\n========== DONE ==========")
}
