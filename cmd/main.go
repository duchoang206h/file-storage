package main

import (
	"log"

	"github.com/duchoang206h/file-storage/config"
	"github.com/duchoang206h/file-storage/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())
	config.LoadConfig()
	router.SetupRoute(app)
	log.Fatal(app.Listen(":4000"))
}
