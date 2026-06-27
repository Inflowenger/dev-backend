package models

import "github.com/Inflowenger/inflow-fusion/models"

type ExtensionType string

const (
	ExtEventBaseType  ExtensionType = "event"
	ExtPluginBaseType ExtensionType = "plugin"
)

type ExtensionRecord struct {
	ID          string         `json:"id"`
	Type        ExtensionType  `json:"type" validate:"oneof=plugin event"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Icon        Icon           `json:"icon"`
	Parameters  map[string]any `json:"params"`
	BindTo      models.Node    `json:"bindTo"`
	CreatedAt   int64          `json:"createdAt"`
	UpdatedAt   int64          `json:"updatedAt"`
}

type Icon struct {
	Class string         `json:"class"`
	Name  string         `json:"name"`
	Meta  map[string]any `json:"meta"`
}
