package apiHandlers

import (
	contextControllers "github.com/Inflowenger/dev-backend/api/context"
	flowControllers "github.com/Inflowenger/dev-backend/api/flow"
	"github.com/Inflowenger/dev-backend/etc"
	"github.com/gofiber/fiber/v3"
)

func RegisterAll(api fiber.Router) {
	api.Use(etc.HS256SecKeyHandler())
	flowControllers.Register(api)
	contextControllers.Register(api)
}
