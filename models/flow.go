package models


type FlowRecord struct{
	ID string `json:"id"`
	Title string `json:"title"`
	CreatedAt int64 `json:"createdAt"`
	ViewFlow ViewFlow `json:"view_flow"`
}



type ViewFlow struct {
	Nodes []VueFlowNode `json:"nodes"`
	Edges []Edges       `json:"edges"`
}
type Dimensions struct {
	Width  float32 `json:"width"`
	Height float32 `json:"height"`
}
type ComputedPosition struct {
	X float32 `json:"x"`
	Y float32 `json:"y"`
	Z float32 `json:"z"`
}
type Position struct {
	X float32 `json:"x"`
	Y float32 `json:"y"`
}

type Source struct {
	ID       string `json:"id"`
	Type     string `json:"type"`
	NodeID   string `json:"nodeId"`
	Position string `json:"position"`
	X        float32    `json:"x"`
	Y        float32    `json:"y"`
	Width    float32    `json:"width"`
	Height   float32    `json:"height"`
}
type Target struct {
	ID       string `json:"id"`
	Type     string `json:"type"`
	NodeID   string `json:"nodeId"`
	Position string `json:"position"`
	X        float32    `json:"x"`
	Y        float32    `json:"y"`
	Width    float32    `json:"width"`
	Height   float32    `json:"height"`
}
type HandleBounds struct {
	Source []Source `json:"source"`
	Target []Target `json:"target"`
}
type VueFlowNode struct {
	ID               string           `json:"id"`
	Type             string           `json:"type"`
	Dimensions       Dimensions       `json:"dimensions"`
	ComputedPosition ComputedPosition `json:"computedPosition"`
	Selected         bool             `json:"selected"`
	Dragging         bool             `json:"dragging"`
	Resizing         bool             `json:"resizing"`
	Initialized      bool             `json:"initialized"`
	IsParent         bool             `json:"isParent"`
	Position         Position         `json:"position"`
	Data             any              `json:"data"`
	Events           any              `json:"events"`
	HandleBounds     HandleBounds     `json:"handleBounds,omitempty"`
}
type EdgePayload struct {
	Tags     []string `json:"tags"`
	EdgeType string   `json:"edgeType"`
}

type Edges struct {
	ID           string      `json:"id"`
	Type         string      `json:"type"`
	Source       string      `json:"source"`
	Target       string      `json:"target"`
	SourceHandle string      `json:"sourceHandle"`
	TargetHandle string      `json:"targetHandle"`
	Data         EdgePayload `json:"data"`
	Events       any         `json:"events"`
	Label        string      `json:"label"`
	MarkerEnd    string      `json:"markerEnd"`
	Animated     bool        `json:"animated"`
	SourceX      float32         `json:"sourceX"`
	SourceY      float32         `json:"sourceY"`
	TargetX      float32         `json:"targetX"`
	TargetY      float32         `json:"targetY"`
	SourceNode   VueFlowNode `json:"sourceNode"`
	TargetNode   VueFlowNode `json:"targetNode"`
}