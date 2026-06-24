package models

type ProcessRequestInput struct {
	FlowId    string `json:"flowId" validate:"r_required"`
	ContextId string `json:"contextId" validate:"r_required"`
}
