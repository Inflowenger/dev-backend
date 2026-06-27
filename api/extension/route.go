package extensionControllers




import "github.com/gofiber/fiber/v3"

func Register(api fiber.Router) {
	extGroup := api.Group("extension")
	extGroup.Post("", addNewExt)
	extGroup.Get("",list)
	extGroup.Get("/id/:extId", getExtensionById)
	extGroup.Delete("/id/:extId", deleteExtById)

}
