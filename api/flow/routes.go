package flowControllers

import (
	"github.com/Inflowenger/dev-backend/models"
	validation "github.com/mehdi-shokohi/fiberValidation"

	"github.com/gofiber/fiber/v3"
)
	

func Register(api fiber.Router) {
	flowGroup := api.Group("flow")
	flowGroup.Post("", addNewFlow)
	flowGroup.Get("",list)
	flowGroup.Get("/id/:flowId", getFlowById)
	flowGroup.Delete("/id/:flowId", deleteFlowById)

	// process requests
	procGroup:=api.Group("ps")
	procGroup.Post("",validation.ValidateBodyAs(models.ProcessRequestInput{}),newProcess)
	procGroup.Post("/compile",validation.ValidateBodyAs(models.ProcessRequestInput{}),compile)

	procGroup.Post("/stop/:pid",stopByPid)
}
