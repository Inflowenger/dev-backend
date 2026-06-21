package main

import (
	apiHandlers "github.com/Inflowenger/dev-backend/api"
	"github.com/Inflowenger/dev-backend/env"
	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/logger"
)

func main() {
	app := fiber.New(fiber.Config{
		JSONEncoder: sonic.Marshal,
		JSONDecoder: sonic.Unmarshal,
		AppName:     "inflow-developer",
	})

	app.Use(cors.New())

	app.Use(logger.New())
	apiHandlers.RegisterAll(app)
	if err := app.Listen(env.GetApiPort()); err != nil {
		panic(err)
	}
}
