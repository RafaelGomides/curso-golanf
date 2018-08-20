package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "bel:zebu@/golang")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	stmt, _ := db.Prepare("UPDATE users SET name = ? WHERE id = ?")

	stmt.Exec("SAPORRA É NOSSA", 1)
	stmt.Exec("BRASIL!", 2)

	stmt2, _ := db.Prepare("DELETE FROM users WHERE id = ?")

	stmt2.Exec(3)
}
