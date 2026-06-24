package flowControllers

import (
	"fmt"

	"github.com/Inflowenger/dev-backend/etc"
	"github.com/Inflowenger/dev-backend/inflow"
	"github.com/Inflowenger/dev-backend/models"
	"github.com/Inflowenger/dev-backend/repository"
	inflowfuse "github.com/Inflowenger/inflow-fusion/inflow"
	"github.com/gofiber/fiber/v3"
)

func newProcess(c fiber.Ctx) error {
	input := models.ProcessRequestInput{}
	if err := c.Bind().Body(&input); err != nil {
		return etc.Send(c, fiber.StatusBadRequest, nil, models.ErrorResponse{Code: fiber.ErrBadRequest.Code, Message: fiber.ErrBadRequest.Message})
	}
	rec, err := repository.GetFlowById(input.FlowId)
	if err != nil {
		return etc.Send(c, fiber.StatusNotFound, nil, models.ErrorResponse{Code: fiber.ErrBadRequest.Code, Message: "given flow id not found or error  occured " + err.Error()})
	}
	startNodeId, err := inflow.GetStartNodeId(*rec)
	if err != nil {
		return etc.Send(c, fiber.StatusBadRequest, nil, models.ErrorResponse{Message: err.Error()})
	}

	proc, err := inflowfuse.NewProcess(startNodeId,
		inflowfuse.WithContextDocument(input.ContextId),
		inflowfuse.WithFlowId(input.FlowId),
		inflowfuse.WithInflowInstanceUrl("http://mate-Predator-PHN16-73:9001"),
		inflowfuse.WithMeta(map[string]string{"account": "dev"}),
	)
	if err != nil {
		fmt.Println(err.Error())
	}
	resp, err := proc.Exec(c.Context())
	return etc.Send(c, fiber.StatusAccepted, map[string]any{"pid": resp.Data.PID, "selected_resource": proc.GetResource()}, err)
}
func compile(c fiber.Ctx) error {
	input := models.ProcessRequestInput{}
	if err := c.Bind().Body(&input); err != nil {
		return etc.Send(c, fiber.StatusBadRequest, nil, models.ErrorResponse{Code: fiber.ErrBadRequest.Code, Message: fiber.ErrBadRequest.Message})
	}
	rec, err := repository.GetFlowById(input.FlowId)
	if err != nil {
		return etc.Send(c, fiber.StatusNotFound, nil, models.ErrorResponse{Code: fiber.ErrBadRequest.Code, Message: "given flow id not found or error  occured " + err.Error()})
	}
	startNodeId, cmp, err := inflow.FLowCompiler(*rec)
	if err != nil {
		return etc.Send(c, fiber.StatusBadRequest, nil, models.ErrorResponse{Message: err.Error()})
	}

	proc, err := inflowfuse.NewProcess(startNodeId,
		inflowfuse.WithContextDocument(input.ContextId),
		inflowfuse.WithFlowId(input.FlowId),
		inflowfuse.WithInflowInstanceUrl("http://mate-Predator-PHN16-73:9001"),
		inflowfuse.WithMeta(map[string]string{"account": "dev"}),
	)
	if err != nil {
		fmt.Println(err.Error())
	}
	return etc.Send(c, fiber.StatusAccepted, map[string]any{"selected_resource": proc.GetResource(), "process_req": proc.GetRequest(), "compiled": cmp}, err)
}
func stopByPid(c fiber.Ctx) error {

	return nil
}
