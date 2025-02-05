package main

import (
	"gomysql/db"
	"gomysql/models"
)

func main() {
	db.Connect()
	// db.Ping()

	db.ExistsTable("users")
	db.CreateTable(models.UserSchema, "users")

	db.Close()
}