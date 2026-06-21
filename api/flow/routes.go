package flowControllers

import "github.com/gofiber/fiber/v3"


func Register(api fiber.Router) {
	flowGroup:=api.Group("flow")
	flowGroup.Post("",addNewFlow)
	flowGroup.Get("/id/:flowId",getFlowById)
}
