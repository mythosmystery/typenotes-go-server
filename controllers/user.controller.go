package controllers

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go/v4"
	"github.com/gofiber/fiber/v2"
	"github.com/mythosmystery/typenotes-go-server/config"
	"github.com/mythosmystery/typenotes-go-server/models"
	"github.com/mythosmystery/typenotes-go-server/services"
	"golang.org/x/crypto/bcrypt"
)

type UserInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string
	User  models.User
}

func generatetoken(user models.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = user.ID
	claims["username"] = user.Username
	claims["email"] = user.Email
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	signedToken, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}
	return signedToken, nil
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

	token, err := generatetoken(user)
	if err != nil {
		return c.Status(500).SendString("Internal server error")
	}

	return c.JSON(LoginResponse{
		Token: token,
		User:  user,
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
		Hash:     hash,
		Email:    body.Email,
	}
	config.DB.Create(&user)
	token, err := generatetoken(user)
	if err != nil {
		return c.Status(500).SendString("Internal server error")
	}

	return c.JSON(LoginResponse{
		Token: token,
		User:  user,
	})
}

func GetUsers(c *fiber.Ctx) error {
	var users []models.User
	config.DB.Preload("Notes").Find(&users)
	return c.JSON(users)
}

func GetMe(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(401).SendString("Unauthorized")
	}
	token, err := jwt.Parse(authHeader, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		var user models.User
		config.DB.Find(&user, claims["id"])
		return c.JSON(user)
	}
	return err
}

func ForgotPassword(c *fiber.Ctx) error {
	var body struct{ Email string }
	if err := c.BodyParser(&body); err != nil {
		return err
	}
	var user models.User
	config.DB.Where("email = ?", body.Email).First(&user)
	if user.ID == 0 {
		return c.Status(401).SendString("Invalid email")
	}
	token, err := generatetoken(user)
	if err != nil {
		return c.Status(500).SendString("Internal server error")
	}
	if err := services.SendResetEmail(user.Email, token); err != nil {
		fmt.Println(err)
		return c.Status(500).SendString("Internal server error")
	}
	return c.Status(200).SendString("Email sent")
}

func ResetPassword(c *fiber.Ctx) error {
	authToken := c.Query("token")
	if authToken == "" {
		return c.Status(401).SendString("Unauthorized")
	}
	var body LoginInput
	if err := c.BodyParser(&body); err != nil {
		return c.Status(500).SendString("Internal server error")
	}
	token, err := jwt.Parse(authToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		var user models.User
		config.DB.Find(&user, claims["id"])
		if user.Email != body.Email {
			return c.Status(401).SendString("Invalid email")
		}
		user.Hash, _ = bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
		config.DB.Save(&user)
		return c.Status(200).SendString("Password changed")
	}
	return err
}
