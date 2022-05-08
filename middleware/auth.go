package middleware

import (
	"fmt"
	"os"

	"github.com/dgrijalva/jwt-go/v4"
	"github.com/gofiber/fiber/v2"
)

type AuthUser struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func Authorized(c *fiber.Ctx) error {
	// Get token from header
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(401).SendString("Unauthorized")
	}
	token, err := jwt.Parse(authHeader, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		user := AuthUser{
			ID:       uint(claims["id"].(float64)),
			Username: claims["username"].(string),
			Email:    claims["email"].(string),
		}
		c.Locals("user", user)
	} else {
		fmt.Println(err)
		return c.Status(401).SendString("Unauthorized")
	}
	return c.Next()
}
