package contextControllers

import "github.com/gofiber/fiber/v3"

func Register(api fiber.Router) {
	contextGroup := api.Group("context")
	contextGroup.Post("", addNewContext)
	contextGroup.Get("",list)
	contextGroup.Get("/id/:contextId", getContextById)
	contextGroup.Delete("/id/:contextId", deleteContextById)

}
