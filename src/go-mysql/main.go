package main

import (
	"go-mysql/database"
	"go-mysql/handlers"
	"go-mysql/models"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Establecer conexión a la base de datos
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	// Crear una instancia de Contact
	newContact := models.Contact{
		Name: "Nuevo Usuario",
		Email: "nuevo@example.com",
		Phone: "091 235 634",
	}

	handlers.CreateContact(db, newContact)

	handlers.ListContacts(db)

	// handlers.GetContactByID(db, 2)
}
