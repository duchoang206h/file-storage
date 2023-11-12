package handler

import (
	"fmt"
	"path/filepath"

	"github.com/duchoang206h/file-storage/util"
	"github.com/gofiber/fiber/v2"
)

func UploadImageHandler(c *fiber.Ctx) error {
	image, err := c.FormFile("image")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Bad request"})
	}
	ext := filepath.Ext(image.Filename)
	fileName := util.TimestampFileName() + ext
	c.SaveFile(image, fmt.Sprintf("./static/%s", fileName))
	fileUrl := util.FormatFileUrl(fileName)
	return c.Status(200).JSON(fiber.Map{"result": fileUrl})
}
