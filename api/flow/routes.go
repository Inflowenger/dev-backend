package flowControllers

import "github.com/gofiber/fiber/v3"

func Register(api fiber.Router) {
	flowGroup := api.Group("flow")
	flowGroup.Post("", addNewFlow)
	flowGroup.Get("",list)
	flowGroup.Get("/id/:flowId", getFlowById)
	flowGroup.Delete("/id/:flowId", deleteFlowById)

	// process requests
	procGroup:=api.Group("ps")
	procGroup.Post("/flow/:flowId",newProcess)
	procGroup.Post("/stop/:pid",stopByPid)
}
