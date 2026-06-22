package flowControllers

import "github.com/gofiber/fiber/v3"

func Register(api fiber.Router) {
	flowGroup := api.Group("flow")
	flowGroup.Post("", addNewFlow)
	flowGroup.Get("",list)
	flowGroup.Get("/id/:flowId", getFlowById)
	flowGroup.Delete("/id/:flowId", deleteFlowById)

}
