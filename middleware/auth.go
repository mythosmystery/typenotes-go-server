package middleware

import (
	"github.com/dgrijalva/jwt-go/v4"
	"github.com/gofiber/fiber/v2"
	"github.com/mythosmystery/typenotes-go-server/config"
)

type AuthUser struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`	
	Email    string `json:"email"`
}

func Authorized(c *fiber.Ctx) error {
	token := c.Get("Authorization")
	if token == "" {
		return c.Status(401).SendString("Unauthorized")
	}
	mapClaims := jwt.MapClaims{}
	parsed, err := jwt.ParseWithClaims(token, mapClaims, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.JWTSecret), nil
	})
	claims := parsed.Claims.(jwt.MapClaims)
	user := AuthUser{
		ID:       uint(claims["id"].(float64)),
		Username: claims["username"].(string),
		Email:    claims["email"].(string),
	}
	c.Locals("user", user)
	if err != nil {
		return c.Status(401).SendString("Unauthorized")
	}
	return c.Next()
}