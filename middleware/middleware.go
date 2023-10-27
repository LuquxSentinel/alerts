package middleware

import (
	"context"
	"github.com/gofiber/websocket/v2"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

func Authorization(c *fiber.Ctx) error {
	_, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	jwtToken := c.Get("authorization")
	if jwtToken == "" {
		return c.Status(fiber.StatusUnauthorized).JSON("missing authorization header")
	}

	// TODO: verify authorization header
	// TODO: add uid to request context

	return c.Next()
}


func WebSocketUpgrade(c *fiber.Ctx) error {
	log.Println("websocket upgrade")
	if websocket.IsWebSocketUpgrade(c) {
		c.Locals("allowed", true)
		return c.Next()
	}

	return fiber.ErrUpgradeRequired
}