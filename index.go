package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/sun053123/fiber-react-authen/database"
	"github.com/sun053123/fiber-react-authen/routes"
)

func main() {
	database.Connect() // connect กับ pgsql

	defer database.DB.Close() // ทำเพื่อความชัวร์ เมื่อปิด webapp ควร ปิด DB

	app := fiber.New() //ใช้ fiber

	app.Use(cors.New(cors.Config{
		AllowCredentials: true, //frontend can connect cookie จากหลังบ้านที่เราส่งไปได้
		AllowOrigins:     "http://localhost:3000",
		AllowMethods:     "*",
		AllowHeaders:     "*",
	}))

	routes.Setup(app)

	app.Listen(":8000")

}
