package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("db", "root:senha@/cursogo")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	tx, _ := db.Begin() // Inicio da Transaction
	stmt, _ := tx.Prepare("insert into usuarios(id, nome) values(?,?)")

	stmt.Exec(2000, "Bia")
	stmt.Exec(2001, "Carlos")
	result, err := stmt.Exec(1, "Tiago")

	if err != nil {
		// E NESCESSARIO CHAMAR O ROLL BACK
		tx.Rollback() // desfaz tudo dentro da transaction
		log.Fatal(err)
	}
	log.Println(result.LastInsertId())

	tx.Commit()
}
