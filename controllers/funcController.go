package controllers

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/sun053123/fiber-react-authen/database"
	"github.com/sun053123/fiber-react-authen/models"
)

func HandlerHome(c *fiber.Ctx) error {

	return c.SendString("Home")
}

func HandlerNewPost(c *fiber.Ctx) error {
	request := CreatePostRequest{}
	err := c.BodyParser(&request)

	if err != nil {
		log.Fatal(err)
		return err
	}

	// validator
	if request.Body == "" {
		return fiber.ErrUnprocessableEntity
	}

	post := models.Post{
		Body: request.Body,
		URI:  request.URI,
	}

	// fmt.Println(&post)

	createdPost := database.DB.Create(&post)
	err = createdPost.Error

	return c.Status(fiber.StatusCreated).JSON("Success")
}

func HandlerSinglePost(c *fiber.Ctx) error {
	id, err := c.ParamsInt("postid")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(id)

	var post models.Post

	database.DB.Where("id= ?", id).First(&post)
	if post.ID != 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "This post doesn't exist ...",
		})
	}

	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "Incorrect username or password")
	}

	return c.JSON(&post)
}

func HandlerDeletePost(c *fiber.Ctx) error {
	return c.SendString("Home")
}

func HandlerCreateComment(c *fiber.Ctx) error {
	return c.SendString("Home")
}

func HandlerDeleteComment(c *fiber.Ctx) error {
	return c.SendString("Home")
}
