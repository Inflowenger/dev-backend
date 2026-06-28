package models


type ExtensionType string

const (
	ExtEventBaseType  ExtensionType = "extrinsic"
	ExtPluginBaseType ExtensionType = "plugin"
)

type ExtensionRecord struct {
	ID          string         `json:"id"`
	Type        ExtensionType  `json:"type" validate:"r_required,oneof=plugin extrinsic"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Icon        Icon           `json:"icon"`
	Parameters  FormParameters `json:"params"`
	BindTo      Bind           `json:"bindTo"`
	CreatedAt   int64          `json:"createdAt"`
	UpdatedAt   int64          `json:"updatedAt"`
}
type Bind struct {
	TopicKey string            `json:"topic_key"`
	Values   map[string]string `json:"values"`
}
type Icon struct {
	Class string         `json:"class"`
	Name  string         `json:"name"`
	Meta  map[string]any `json:"meta"`
}

type FormParameters struct {
	Schema map[string]any `json:"schema"`
	UI     map[string]any `json:"ui"`
}
