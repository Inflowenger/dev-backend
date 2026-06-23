package apiHandlers

import (
	"fmt"

	contextControllers "github.com/Inflowenger/dev-backend/api/context"
	flowControllers "github.com/Inflowenger/dev-backend/api/flow"
	"github.com/Inflowenger/dev-backend/env"
	"github.com/Inflowenger/dev-backend/etc"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/proxy"
	recoverer "github.com/gofiber/fiber/v3/middleware/recover"
)

func RegisterAll(api fiber.Router) {
	api.Use(etc.HS256SecKeyHandler())
	api.Use(recoverer.New())
	api.All("/infra/*", infraProxyHandler)
	flowControllers.Register(api)
	contextControllers.Register(api)
}

func infraProxyHandler(c fiber.Ctx) error {
	url := fmt.Sprintf("%s/%s?%s", env.GetInfraApiUrl(), c.Params("*1"), c.Request().URI().QueryString())
	return proxy.Forward(url)(c)
}
