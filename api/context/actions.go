package contextControllers

import (
	"strings"
	"time"

	"github.com/Inflowenger/dev-backend/models"
	"github.com/Inflowenger/dev-backend/repository"
	"github.com/gofiber/fiber/v3"
)

func addNewContext(ctx fiber.Ctx) error {
	input := models.ContextRecord{}
	if err := ctx.Bind().Body(&input); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(models.Response{Data: nil, Error: models.ErrorResponse{Code: fiber.ErrBadRequest.Code, Message: fiber.ErrBadRequest.Message}})
	}
	if input.ID == "" {
		input.CreatedAt = time.Now().Unix()
	}
	if strings.TrimSpace(input.Title) == "" {
		input.Title = "untitled context"
	}
	err := repository.UpsertContext(&input)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(models.Response{Data: nil, Error: models.ErrorResponse{Code: fiber.ErrInternalServerError.Code, Message: fiber.ErrInternalServerError.Message}})

	}
	return ctx.JSON(input)
}

func getContextById(c fiber.Ctx) error {
	contextId := c.Params("contextId")
	if !strings.HasPrefix(contextId, repository.CONTEXT_INDEX_PREFIX) {
		contextId = repository.ContextIndexByString(contextId)
	}
	context := repository.GetContextById(contextId)
	if context == nil {
		return c.Status(fiber.StatusNotFound).JSON(models.Response{Data: nil, Error: models.ErrorResponse{Code: fiber.ErrNotFound.Code, Message: "given context id not found"}})

	}
	return c.JSON(models.Response{Data: context})
}
func deleteContextById(c fiber.Ctx) error {
	contextId := c.Params("contextId")
	if !strings.HasPrefix(contextId, repository.CONTEXT_INDEX_PREFIX) {
		contextId = repository.ContextIndexByString(contextId)
	}
	err := repository.Delete(contextId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{Data: nil, Error: models.ErrorResponse{Code: fiber.ErrInternalServerError.Code, Message: fiber.ErrInternalServerError.Message}})
	}
	return c.Status(fiber.StatusAccepted).JSON(models.Response{Data: map[string]any{"contextId": contextId}})

}
func list(c fiber.Ctx) error {
	q := models.PaginationParams{}
	if err := c.Bind().Query(&q); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{Data: nil, Error: models.ErrorResponse{Code: fiber.ErrBadRequest.Code, Message: fiber.ErrBadRequest.Message}})
	}
	l, cursor, err := repository.GetContextList(q.Cursor, int(q.PerPage))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{Data: nil, Error: models.ErrorResponse{Code: fiber.ErrInternalServerError.Code, Message: fiber.ErrInternalServerError.Message}})
	}
	return c.JSON(map[string]any{"list": l, "next": cursor, "error": nil})
}
