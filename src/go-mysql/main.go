package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	dns := "daniel:rootroot@tcp(localhost:3306)/db_contacts"

	// Abrir una conexión a la base de datos
	db, err := sql.Open("mysql", dns)
	if err != nil {
		log.Fatal(err)
	}

	// Verificar la conexión
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	log.Println("Conexión a la base de datos MySQL exitosa")
}
