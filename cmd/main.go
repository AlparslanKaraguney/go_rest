package main

import (
	"fmt"
	"log"
	"rest_service/db"
	"rest_service/user"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database, err := db.Connect()
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	data := user.NewData(database)
	err = data.Migrate()
	if err != nil {
		log.Fatalf("Error migrating database: %v", err)
	}

	fmt.Println("Successfully migrated database")

	service := user.NewService(data)
	handler := user.NewHandler(service)

	app := fiber.New()
	app.Get("/users", handler.GetAll)
	app.Get("/users/:id", handler.GetByID)
	app.Post("/users", handler.Create)

	fmt.Println("Server started on port 8000")
	app.Listen(":8000")

}
