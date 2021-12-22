package controllers

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/sun053123/fiber-react-authen/database"
	"github.com/sun053123/fiber-react-authen/models"
	"golang.org/x/crypto/bcrypt"
)

var jwtSecret = os.Getenv("JWTSECRET")

func HandlerRegister(c *fiber.Ctx) error {

	var user_exist models.User
	request := RegisterRequest{}
	err := c.BodyParser(&request)

	if err != nil {
		return err
	}

	if request.Email == "" || request.Password == "" || request.Username == "" {
		return fiber.ErrUnprocessableEntity
	}

	//check if email already exist
	database.DB.Where("email =?", request.Email).First(&user_exist)
	if user_exist.ID != 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "This email already exist ...",
		})
	}

	password, err := bcrypt.GenerateFromPassword([]byte(request.Password), 12)
	if err != nil {
		return fiber.NewError(fiber.StatusUnprocessableEntity, err.Error())
	}

	user := models.User{
		Username: request.Username,
		Email:    request.Email,
		Password: string(password),
	}

	createdUser := database.DB.Create(&user)
	err = createdUser.Error

	if err != nil {
		return fiber.NewError(fiber.StatusUnprocessableEntity, err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON("Success")
}

func HandlerLogin(c *fiber.Ctx) error {
	fmt.Println("BRUH")

	request := LoginRequest{}
	err := c.BodyParser(&request)
	if err != nil {
		return err
	}

	if request.Email == "" || request.Password == "" {
		return fiber.ErrUnprocessableEntity
	}

	var user models.User

	database.DB.Where("email =?", request.Email).First(&user)

	if user.ID == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "User not found ...",
		})
	}

	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "Incorrect username or password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)) //bcrypt compare hash ที่ได้จาก db กับ password ที่ user ใส่เข้ามา จะไม่ไป compare ใน database
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "Incorrect username or password")
	}

	claims := jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(user.ID)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := jwtToken.SignedString([]byte(jwtSecret))
	if err != nil {
		return fiber.ErrInternalServerError
	}
	// สร้าง cookie แล้้วนำ cliams มาใส่เพื่อตรวจสอบ ส่งให้ front ทาง header , front ไม่สามารถดึงได้โดยตรง
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true, //frontend cannot access this cookie
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "success",
	})
}

func User(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthorized",
		})
	}

	// cliams ไม่สามารถ return issue ได้ จึงต้องแปลงเป็น jwt standard cliams ก่อน
	claims := token.Claims.(*jwt.StandardClaims)

	var user models.User // มันจะส่ง user struct ทุกตัว ต้องเข้าไป tag กำกับว่าให้ส่งอะไรบ้าง

	database.DB.Where("id = ?", claims.Issuer).First(&user)

	return c.JSON(user)

}

func HandlerLogout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"messgae": "success",
	})

}
