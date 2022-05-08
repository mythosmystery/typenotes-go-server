package middleware

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

func LogOp(c *fiber.Ctx) error {
	fmt.Println("-----------------------------------------------------------------------------------")
	fmt.Printf("Time: %s -- IP: %s\n", time.Now(), c.IP())
	fmt.Printf("Path: %s -- Method: %s\n\n", c.Path(), c.Method())
	var body map[string]interface{}
	c.BodyParser(&body)

	fmt.Println("Body:")
	for key, value := range body {
		if key != "password" {
			fmt.Printf("  %s: %v\n", key, value)
		}
	}

	fmt.Println("Request Headers:")
	fmt.Println("  Authorization:", c.Get("Authorization"))

	fmt.Println("Params:")
	for key, value := range c.AllParams() {
		fmt.Printf("  %s: %v\n", key, value)
	}

	fmt.Println("-----------------------------------------------------------------------------------")
	return c.Next()
}
