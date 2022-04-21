package controllers

import (
	"time"

	"github.com/dgrijalva/jwt-go/v4"
	"github.com/gofiber/fiber/v2"
	"github.com/mythosmystery/typenotes-go-server/config"
	"github.com/mythosmystery/typenotes-go-server/models"
	"golang.org/x/crypto/bcrypt"
)

type UserInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email 	 string `json:"email"`
}

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string
	User  models.User
}

func LoginUser(c *fiber.Ctx) error {
	var body LoginInput
	if err := c.BodyParser(&body); err != nil {
		return err
	}
	var user models.User
	config.DB.Where("email = ?", body.Email).First(&user)
	if user.ID == 0 {
		return c.Status(401).SendString("Invalid email or password")
	}
	if err := bcrypt.CompareHashAndPassword(user.Hash, []byte(body.Password)); err != nil {
		return c.Status(401).SendString("Invalid password or password")
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": user.ID,
		"username": user.Username,
		"email": user.Email,
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	})
	token, err := claims.SignedString([]byte(config.JWTSecret))
	if err != nil {
		return c.Status(500).SendString("Internal server error")
	}

	return c.JSON(LoginResponse{
		Token: token,
		User: user,
	})
}

func RegisterUser(c *fiber.Ctx) error {
	var body UserInput
	if err := c.BodyParser(&body); err != nil {
		return err
	}
	hash, _ := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	user := models.User{
		Username: body.Username,
		Hash: hash,
		Email: body.Email,
	}
	config.DB.Create(&user)
	return c.JSON(user)
}

func GetUsers(c *fiber.Ctx) error {
	var users []models.User
	config.DB.Find(&users)
	return c.JSON(users)
}