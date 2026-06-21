package flowControllers

import (
	"strings"
	"time"

	"github.com/Inflowenger/dev-backend/models"
	"github.com/Inflowenger/dev-backend/repository"
	"github.com/gofiber/fiber/v3"
)

func addNewFlow(ctx fiber.Ctx) error {
	input := models.FlowRecord{}
	if err := ctx.Bind().Body(&input); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(models.Response{Data: nil, Error: models.ErrorResponse{Code: fiber.ErrBadRequest.Code, Message: fiber.ErrBadRequest.Message}})
	}
	if input.ID == "" {
		input.CreatedAt = time.Now().Unix()
	}
	if strings.TrimSpace(input.Title) == "" {
		input.Title = "untitled workflow"
	}
	err := repository.UpsertFlow(&input)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(models.Response{Data: nil, Error: models.ErrorResponse{Code: fiber.ErrInternalServerError.Code, Message: fiber.ErrInternalServerError.Message}})

	}
	return ctx.JSON(input)
}

func getFlowById(c fiber.Ctx) error {
	flowId := c.Params("flowId")
	if !strings.HasPrefix(flowId, repository.FLOW_INDEX_PREFIX) {
		flowId = repository.FlowIndexByString(flowId)
	}
	flow := repository.GetFlowById(flowId)
	if flow == nil {
		return c.Status(fiber.StatusNotFound).JSON(models.Response{Data: nil, Error: models.ErrorResponse{Code: fiber.ErrNotFound.Code, Message: "given flow id not found , or internal error occurred"}})

	}
	return c.JSON(models.Response{Data: flow})
}
