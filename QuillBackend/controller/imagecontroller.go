package controller

import (
	"math/rand"
	"os"

	"github.com/gofiber/fiber/v2"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyz")

func randLetter(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func Upload(c *fiber.Ctx) error {
	// Ensure uploads directory exists
	if err := os.MkdirAll("./uploads", 0755); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Failed to create uploads directory",
			"error":   err.Error(),
		})
	}

	form, err := c.MultipartForm()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Failed to parse form",
			"error":   err.Error(),
		})
	}

	files := form.File["image"]
	if len(files) == 0 {
		return c.Status(400).JSON(fiber.Map{
			"message": "No image file provided",
		})
	}

	fileName := ""
	for _, file := range files {
		fileName = randLetter(5) + "-" + file.Filename
		if err := c.SaveFile(file, "./uploads/"+fileName); err != nil {
			return c.Status(500).JSON(fiber.Map{
				"message": "Failed to save file",
				"error":   err.Error(),
			})
		}
	}

	// Get backend URL from environment or construct it
	backendURL := os.Getenv("BACKEND_URL")
	if backendURL == "" {
		backendURL = "https://quill-backend-lgxs.onrender.com"
	}

	// Return the correct static path - app.Static serves /api/uploads
	return c.JSON(fiber.Map{
		"url": backendURL + "/api/uploads/" + fileName,
	})
}
