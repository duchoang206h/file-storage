package main

import (
	"log"

	"github.com/duchoang206h/file-storage/config"
	"github.com/duchoang206h/file-storage/router"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	config.LoadConfig()
	router.SetupRoute(app)
	log.Fatal(app.Listen(":4000"))
}
