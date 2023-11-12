package router

import (
	"github.com/duchoang206h/file-storage/handler"
	"github.com/gofiber/fiber/v2"
)

func SetupRoute(app *fiber.App) {
	images := app.Group("/images")
	images.Static("/", "./static")
	images.Post("/", handler.UploadImageHandler)
}
