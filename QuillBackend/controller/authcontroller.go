package controller

import (
	"QuillBackend/database"
	"QuillBackend/models"
	"QuillBackend/utils"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"

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

// Login function to handle user login
func Login(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		fmt.Println("Unable to parse body:", err)
	}
	// Check if the email exists in the database
	// basically telling that if user is trying to login without prior signup,Telling them to signup first before login
	var user models.User
	database.DB.Where("email=?", data["email"]).First(&user)
	if user.Id == 0 {
		c.Status(404)
		return c.JSON(fiber.Map{"message": "User not found"})
	}
	// Check if the password is correct and matches the hashed password
	if err := user.CheckPassword(data["password"]); err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{"message": "Incorrect password"})
	}
	// Generate JWT token upon successful login
	token, err := utils.GenerateJWT(strconv.Itoa(int(user.Id)))
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return nil
	}
	cookie := fiber.Cookie{
		Name:     "jwt",                          // cookie name used to store the token
		Value:    token,                          // token as value used for authentication
		Expires:  time.Now().Add(time.Hour * 24), // 1 day
		HTTPOnly: true,                           // to prevent client side js access
	}
	c.Cookie(&cookie)
	return c.JSON(fiber.Map{"message": "Login successful"})

}
