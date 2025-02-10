package main

import (
	"fmt"
	"gomysql/db"
	"gomysql/models"
)

func main() {
	db.Connect()
	// db.Ping()

	// db.ExistsTable("users")
	// db.CreateTable(models.UserSchema, "users")
	// db.TruncateTable("users")
	// user := models.CreateUser("otro", "1234", "otro@gmail.com")
	// fmt.Println(user)
	user := models.GetUser(2)
	fmt.Println(user)
	// user.Username = "juan"
	// user.Password = "juan2451"
	// user.Email = "juan@gmail.com"
	// user.Save()
	// fmt.Println(models.ListUsers())
	user.Delete()

	db.Close()
}