package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "bel:zebu@/golang")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("INSERT INTO users(id, name) VALUES (?, ?)")

	stmt.Exec(2000, "Bia")
	stmt.Exec(2001, "Carlos")

	_, err = stmt.Exec(1, "Thiago")

	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	tx.Commit()
}
