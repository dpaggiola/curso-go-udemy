package handlers

import (
	"database/sql"
	"fmt"
	"go-mysql/models"
	"log"
)

func ListContacts(db *sql.DB) {
	// Consulta SQL para seleccionar todos los contactos
	query := "SELECT * FROM contact"

	// Ejecutar la consulta
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	// Iterar sobre los resultados y mostrarlos
	fmt.Println("\nLISTA DE CONTACTOS:")
	fmt.Println("--------------------------------------------------------------- ")
	for rows.Next(){
		// Instancia de modelo Contact
		contact := models.Contact{}

		var valueEmail sql.NullString
		err := rows.Scan(&contact.Id, &contact.Name, &valueEmail, &contact.Phone)
		if err != nil {
			log.Fatal(err)
		}

		if valueEmail.Valid {
			contact.Email = valueEmail.String
		} else {
			contact.Email = "Sin correo electrónico"
		}
		
		fmt.Printf("ID: %d, Nombre: %s, Email: %s, Telefono: %s\n",
			contact.Id, contact.Name, contact.Email, contact.Phone)
		
		fmt.Println("--------------------------------------------------------------- ")
	}
}

func GetContactByID(db *sql.DB, contactID int) {
	query := "SELECT * FROM contact WHERE id = ?"

	row := db.QueryRow(query, contactID)

	// Instancia de modelo Contact
	contact := models.Contact{}
	var valueEmail sql.NullString // Para manejar valor null

	// Escanear el resultado en el modelo contact
	err := row.Scan(&contact.Id, &contact.Name, &valueEmail, &contact.Phone)
	if err != nil {
		if err == sql.ErrNoRows{
			log.Fatalf("No se encontró ningún contacto con el ID %d", contactID)
		}
	}

	if valueEmail.Valid {
		contact.Email = valueEmail.String
	} else {
		contact.Email = "Sin correo electrónico"
	}

	fmt.Println("\nLISTA DE UN CONTACTO:")
	fmt.Println("--------------------------------------------------------------- ")
	fmt.Printf("ID: %d, Nombre: %s, Email: %s, Telefono: %s\n",
		contact.Id, contact.Name, contact.Email, contact.Phone)
	fmt.Println("--------------------------------------------------------------- ")
}

func CreateContact(db *sql.DB, contact models.Contact){
	// Sentencia SQL para insertar un nuevo contacto
	query := "INSERT INTO contact (name, email, phone) VALUES (?, ?, ?)"

	// Ejecutar la sentencia SQL
	_, err := db.Exec(query, contact.Name, contact.Email, contact.Phone)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Nuevo contacto registrado con éxito")
}