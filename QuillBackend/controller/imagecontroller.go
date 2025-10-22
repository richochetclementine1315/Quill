package controller

import (
	"math/rand"

	"github.com/gofiber/fiber/v2"
)

// letters is a list of characters used for generating random strings because
// file names should be unique to avoid overwriting existing files.
var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// randLetter generates a random string of length n using
// the letters slice it is necessary to create unique file names for uploaded images.
func randLetter(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// Upload handles image upload requests.
func Upload(c *fiber.Ctx) error {
	// Parse the multipart form. A multipart form is used for file uploads.
	form, err := c.MultipartForm()
	if err != nil {
		return err
	}
	// Get the files from the "image" field in the form.
	files := form.File["image"]
	fileName := ""
	// Loop through each file and save it to the "uploads" directory with a unique name.
	for _, file := range files {
		// Generate a unique file name by prepending a random string to the original file name.
		fileName = randLetter(5) + "_" + file.Filename

		if err := c.SaveFile(file, "./uploads/"+fileName); err != nil {
			return nil
		}
	}
	// Return the URL of the uploaded image as a JSON response.
	return c.JSON(fiber.Map{
		"url": "http://localhost:8080/uploads/" + fileName,
	})
}
