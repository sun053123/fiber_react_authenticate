package routes

import (
	"os"

	"github.com/gofiber/fiber/v2"
	// jwtware "github.com/gofiber/jwt/v2"
	"github.com/sun053123/fiber-react-authen/controllers"
)

var jwtSecret = os.Getenv("JWTSECRET")

func Setup(app *fiber.App) {

	// app.Use("/api/post", jwtware.New(jwtware.Config{
	// 	SigningMethod: "HS256",
	// 	SigningKey:    []byte(jwtSecret),
	// 	SuccessHandler: func(c *fiber.Ctx) error {
	// 		return c.Next() // เมื่อ Token ผ่าน จะ Next ไปทำงาน Handler เลย
	// 	},
	// 	ErrorHandler: func(c *fiber.Ctx, e error) error {
	// 		return fiber.ErrUnauthorized
	// 	},
	// }))

	app.Post("/api/register", controllers.HandlerRegister)
	app.Post("/api/login", controllers.HandlerLogin)
	app.Get("/api/logout", controllers.HandlerLogout)
	app.Get("/api/user", controllers.User)

	app.Get("/api/home", controllers.HandlerHome)
	app.Post("/api/home", controllers.HandlerNewPost)

	app.Get("/api/post/:postid", controllers.HandlerSinglePost)
	app.Delete("/api/post/:postid", controllers.HandlerDeletePost)
	app.Post("/api/post/:postid", controllers.HandlerCreateComment)
	app.Delete("/api/post/:postid/:comment", controllers.HandlerDeleteComment)

}
