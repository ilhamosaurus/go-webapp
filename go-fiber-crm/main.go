package main

import (
	"fmt"

	"go-fiber-crm/database"
	"go-fiber-crm/lead"

	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
)

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/lead", lead.GetLeads)
	app.Get("/api/v1/lead/:id", lead.GetLead)
	app.Post("/api/v1/lead", lead.CreateLead)
	app.Delete("/api/v1/lead/:id", lead.DeleteLead)
}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("mysql", "ilhamo:qwerty@tcp(127.0.0.1:3306)/ilhamodb?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	fmt.Println("Connection Opened to Database")
	database.DBConn.AutoMigrate(&lead.Lead{})
	fmt.Println("Database Migrated")
}

func main() {
	app := fiber.New()
	initDatabase()
	setupRoutes(app)
	app.Listen(3000)
	defer database.DBConn.Close()
}
