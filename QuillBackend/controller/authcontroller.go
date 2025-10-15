package controller

import (
	"QuillBackend/database"
	"QuillBackend/models"
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// Function to validate email format
func validateEmail(email string) bool {
	// A simple regex for email validation
	// This regex checks for the general structure of an email address
	Reg := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	// Check if the email matches the regex
	return Reg.MatchString(email)
}

// function to handle user registration
// So, c *fiber.Ctx gives the handler full access to the HTTP request and response.
func Register(c *fiber.Ctx) error {
	/* Parse the request body into a map
	   we are using a map here because we have multiple fields to parse
	   and we are using an interface{} to allow for any type of value*/
	var data map[string]interface{}
	var userData models.User
	// Parse the request body into the data map by using a BodyParser method
	if err := c.BodyParser(&data); err != nil {
		// if err found:
		fmt.Println("Error parsing body:", err)
	}
	// if password greater lesser than 6 characters
	// converting the []byte to string
	if len(data["password"].(string)) <= 6 {
		c.Status(400) //sending a 400bad request status code
		return c.JSON(fiber.Map{"message": "Password must be at least 6 characters long"})
	}
	// Validate email format
	if !validateEmail(strings.TrimSpace(data["email"].(string))) {
		c.Status(400)
		return c.JSON(fiber.Map{"message": "Invalid email format"})
	}
	// Check if the email already exists in the database
	database.DB.Where("email=?", strings.TrimSpace(data["email"].(string))).First(&userData)
	if userData.Id != 0 {
		c.Status(400)
		return c.JSON(fiber.Map{"message": "Email already registered"})
	}
	// Create a new user instance
	// Populate the user fields with the parsed data
	user := models.User{
		// frontend key names should match with these keys
		FirstName: data["first_name"].(string),
		LastName:  data["last_name"].(string),
		Phone:     data["phone"].(string),
		Email:     strings.TrimSpace(data["email"].(string)),
	}
	// Hash the password before storing it in the database
	user.SetPassword(data["password"].(string))
	// Save the user to the database
	err := database.DB.Create(&user)
	if err != nil {
		log.Println("Error creating user:", err)
	}
	c.Status(200)
	return c.JSON(fiber.Map{
		"user": user,
		// we are not sending the password back in the response
		"message": "Registration successful",
	})

}
