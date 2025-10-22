package controller

import (
	"QuillBackend/database"
	"QuillBackend/models"
	"QuillBackend/utils"
	"errors"
	"fmt"
	"math"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CreatePost(c *fiber.Ctx) error {
	// Initialize a Blog model which means this will hold the data for a new blog post
	var blogpost models.Blog
	if err := c.BodyParser(&blogpost); err != nil {
		fmt.Println("Error parsing body:", err)
		c.Status(400)
		return c.JSON(fiber.Map{"message": "Invalid request body"})
	}

	// Extract user ID from JWT cookie
	cookie := c.Cookies("jwt")
	id, err := utils.ParseJWT(cookie)
	if err != nil {
		c.Status(401)
		return c.JSON(fiber.Map{"message": "Unauthenticated"})
	}

	// Convert string ID to uint and set the UserID for the blog post
	userID, _ := strconv.Atoi(id)
	blogpost.UserID = uint(userID)

	if err := database.DB.Create(&blogpost).Error; err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{"message": "Invalid Payload!"})
	}
	return c.JSON(fiber.Map{"message": "Blog Post Created Successfully!", "post": blogpost})
}

// AllPosts gets all the posts with pagination
func AllPosts(c *fiber.Ctx) error {
	// Get the current page from the query parameters because we are doing pagination
	// pagination means we are dividing the posts into multiple pages.
	// It is beneficial because it allows us to load a small number of posts at a time,
	// improving performance and user experience.
	page, _ := strconv.Atoi(c.Query("page", "1"))
	// Limit to 5 posts per page
	limit := 5
	// Calculate the offset. Offset is the number of posts to skip
	offset := (page - 1) * limit
	var total int64
	var getblog []models.Blog
	// preload user data associated with each blog post. Preload is a GORM feature that allows
	// us to load related data in a single query, reducing the number of database calls and improving performance.
	database.DB.Preload("User").Offset(offset).Limit(limit).Find(&getblog)
	// Get the total count of blog posts in the database
	database.DB.Model(&models.Blog{}).Count(&total)
	return c.JSON(fiber.Map{
		"data": getblog,
		"meta": fiber.Map{
			"page":      page,
			"total":     total,
			"last_page": math.Ceil(float64(int(total) / limit)),
		},
	})

}

// DetailPost gets the detail of a single post by ID
func DetailPost(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	var blogpost models.Blog
	database.DB.Where("id=?", id).Preload("User").First(&blogpost)
	return c.JSON(fiber.Map{
		"data": blogpost,
	})
}

// UpdatePost updates a blog post by ID basically
//
//	it allows us to modify the content of an existing blog post.
func UpdatePost(c *fiber.Ctx) error {
	// Get the ID from the URL parameters
	// means we are extracting the post ID from the request
	id, _ := strconv.Atoi(c.Params("id"))
	blog := models.Blog{
		ID: uint(id),
	}
	// we are (parsing) extracting the updated post data from the request body
	// and populating the blog struct with it
	if err := c.BodyParser(&blog); err != nil {
		fmt.Println("Unable to parse body:", err)
	}
	database.DB.Model(&blog).Updates(blog)
	return c.JSON(blog)

}
func UniquePost(c *fiber.Ctx) error {
	// Get the JWT cookie from the request because
	//  it contains the user's authentication token and we need it to identify the user
	cookie := c.Cookies("jwt")
	id, _ := utils.ParseJWT(cookie)
	var blog []models.Blog
	database.DB.Model(&blog).Where("user_id=?", id).Preload("User").Find(&blog)
	return c.JSON(blog)
}
func DeletePost(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	blog := models.Blog{
		ID: uint(id),
	}
	deleteQuery := database.DB.Delete(&blog)
	if errors.Is(deleteQuery.Error, gorm.ErrRecordNotFound) {
		c.Status(400)
		return c.JSON(fiber.Map{"message": "Oops! Record not found"})
	}
	return c.JSON(fiber.Map{"message": "Post deleted successfully"})
}
