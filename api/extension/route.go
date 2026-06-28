package extensionControllers

import (
	"github.com/Inflowenger/dev-backend/models"
	"github.com/gofiber/fiber/v3"
	validation "github.com/mehdi-shokohi/fiberValidation"
)

func Register(api fiber.Router) {
	extGroup := api.Group("extension")
	extGroup.Post("", validation.ValidateBodyAs[models.ExtensionRecord](),addNewExt)
	extGroup.Get("",list)
	extGroup.Get("/id/:extId", getExtensionById)
	extGroup.Delete("/id/:extId", deleteExtById)
	extGroup.Get("/extrinsics",listOfExtHandlers)

}
