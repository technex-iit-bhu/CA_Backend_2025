package main

import (
	"CA_Backend/database"
	"CA_Backend/router"
	_ "fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"
	"os"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())
	router.Route(app)
	if err := database.Init(); err != nil {
		log.Fatal("unable to connect to database")
	}
	defer database.Disconnect()
	port := os.Getenv("PORT")
	if port == "" {
		port = "6969"
	}
	app.Listen(":" + port)
	log.Println("Server started on port " + port)
}
