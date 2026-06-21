package apiHandlers

import (
	"github.com/Inflowenger/dev-backend/etc"
	"github.com/gofiber/fiber/v3"
)

func RegisterAll(api fiber.Router) {
	api.Use(etc.HS256SecKeyHandler())
}
