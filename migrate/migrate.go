package main

import (
	"fmt"
	"go-rest-api/db"
	"go-rest-api/models"
)

func main() {
	dbConnection := db.NewDB()

	defer fmt.Println("Successfully Migrated!")
	defer db.CloseDB(dbConnection)
	dbConnection.AutoMigrate(&models.User{}, &models.Task{})
}
