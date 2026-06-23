package flowControllers

import (
	"github.com/Inflowenger/dev-backend/etc"
	"github.com/Inflowenger/dev-backend/inflow"
	"github.com/Inflowenger/dev-backend/models"
	"github.com/Inflowenger/dev-backend/repository"
	"github.com/gofiber/fiber/v3"
)

func newProcess(c fiber.Ctx)error{
	flowId:=c.Params("flowId")
	rec,err:=repository.GetFlowById(flowId)
	if  err!=nil{
		return etc.Send(c,fiber.StatusNotFound,nil,models.ErrorResponse{Code: fiber.ErrBadRequest.Code,Message: "given flow id not found or error  occured "+err.Error()})
	}

	cmp,err:=inflow.FLowCompiler(*rec)
	if err!=nil{
		return etc.Send(c,fiber.StatusBadRequest,nil,models.ErrorResponse{Message:err.Error()})
	}
	return etc.Send(c,fiber.StatusAccepted,cmp,nil)
}


func stopByPid(c fiber.Ctx)error{

	return nil
}