package inflow

import (
	"fmt"

	"github.com/Inflowenger/dev-backend/models"
	compiler "github.com/Inflowenger/inflow-fusion/compilers/vueFlow"
	inflowModels "github.com/Inflowenger/inflow-fusion/models"
)

/*
in front side we have a pallett and node types that in compile time those will map to inflow generics types
const items: PaletteItem[] = [

	// Extensions tab
	{ type: 'database', title: 'Database', icon: 'M4 7c0-1.657 3.582-3 8-3s8 1.343 8 3v10c0 1.657-3.582 3-8 3s-8-1.343-8-3V7zM4 12c0 1.657 3.582 3 8 3s8-1.343 8-3', tab: 'extensions' },

	// Generics tab
	{ type: 'startNode', title: 'Start', icon: 'M12 2C6.48 2 2 6.48 2 12s4.48 10 10 10 10-4.48 10-10S17.52 2 12 2z M12 16c-2.21 0-4-1.79-4-4s1.79-4 4-4 4 1.79 4 4-1.79 4-4 4z', tab: 'generics' },
	{ type: 'plugin', title: 'Plugin', icon: 'M20.84 4.61a5.5 5.5 0 0 0-7.78 0L12 5.67l-1.06-1.06a5.5 5.5 0 1 0-7.78 7.78l1.06 1.06L12 21.23l7.78-7.78 1.06-1.06a5.5 5.5 0 0 0 0-7.78z', tab: 'generics' },
	{ type: 'code', title: 'Code', icon: 'M16 18l6-6-6-6 M8 6l-6 6 6 6', tab: 'generics' },
	{ type: 'contract', title: 'Contract', icon: 'M14 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8z M14 2l6 6 M14 2v6h6', tab: 'generics' },
	{ type: 'event', title: 'Event', icon: 'M12 2C6.48 2 2 6.48 2 12s4.48 10 10 10 10-4.48 10-10S17.52 2 12 2zm-1 17.93c-3.95-.49-7-3.85-7-7.93 0-.62.08-1.21.21-1.79L9 15v1c0 1.1.9 2 2 2v1.93zm6.9-2.54c-.26-.81-1-1.39-1.9-1.39h-1v-3c0-.55-.45-1-1-1H8v-2h2c.55 0 1-.45 1-1V7h2c1.1 0 2-.9 2-2v-.41c2.93 1.19 5 4.06 5 7.41 0 2.08-.8 3.97-2.1 5.39z', tab: 'generics' },
	{ type: 'goto', title: 'Goto', icon: 'M7 17L17 7 M7 7h10v10', tab: 'generics' },
	{ type: 'void', title: 'Void', icon: 'M12 2C6.48 2 2 6.48 2 12s4.48 10 10 10 10-4.48 10-10S17.52 2 12 2z M12 16c-2.21 0-4-1.79-4-4s1.79-4 4-4 4 1.79 4 4-1.79 4-4 4z', tab: 'generics' },

]
*/
const (
	NODE_PLUGIN    = "plugin"
	NODE_START     = "startNode"
	NODE_VOID      = "void"
	NODE_CONTRACT  = "contract"
	NODE_CODE      = "code"
	NODE_EVENT_SVC = "event"
	NODE_GOTO      = "goto"
)

func FLowCompiler(f models.FlowRecord) (map[string]*inflowModels.Node, error) {
	startNodeId := ""
	for _, n := range f.ViewFlow.Nodes {
		if n.Type == NODE_START {
			startNodeId = n.ID
			break
		}
	}
	if startNodeId == "" {
		return nil, fmt.Errorf("start node is required")
	}
	cmpr := compiler.NewVueFlowCompiler(compiler.WithEachNodeFunc(NodeBuilder))
	if cmpr == nil {
		return nil, fmt.Errorf("error occurred in compile process")
	}
	l, errs := cmpr.Compile(startNodeId, f.ViewFlow)
	for _, e := range errs {
		return l, e
	}

	return l, nil
}

func NodeBuilder(vfn compiler.VueFlowNode) (*inflowModels.Node, error) {

	nodeData, ok := vfn.Data.(map[string]any)
	if !ok {
		return nil, fmt.Errorf("invalid node data ")
	}
	inflowNode := inflowModels.Node{
		ID:    vfn.ID,
		Title: nodeData["title"].(string),
	}
	if nodeData["key"] != nil {
		inflowNode.Key = nodeData["key"].(string)
	}
	if nodeData["scope"] != nil {
		inflowNode.Scope = nodeData["scope"].(string)

	}
	switch vfn.Type {
	case NODE_START:
		inflowNode.Type = inflowModels.VoidNodeType
	case NODE_CODE:
		inflowNode.Code = &inflowModels.CodeRule{
			Lang:      nodeData["lang"].(string),
			LogicRule: nodeData["logic_rule"].(string),
			OpaData:   map[string]any{},
			OpaResult: nodeData["opa_result"].(string),
		}
	case NODE_CONTRACT:
		inflowNode.Contract = &inflowModels.ContractRule{
			Lang:       nodeData["lang"].(string),
			LogicRule:  nodeData["logic_rule"].(string),
			Conditions: map[string]any{},
			OpaResult:  nodeData["opa_result"].(string),
		}
		if conds, ok := nodeData["conditions"].([]any); ok {
			for _, el := range conds {
				if field, ok := el.(map[string]any); ok {
					inflowNode.Contract.Conditions[field["key"].(string)] = field["value"]
				}
			}
		}
	case NODE_EVENT_SVC:
		inflowNode.Type = inflowModels.EventNodeType

	case NODE_GOTO:
		inflowNode.Type = inflowModels.GoToNodeType

	case NODE_PLUGIN:
		inflowNode.Type = inflowModels.PluginNodeType

	case NODE_VOID:
		inflowNode.Type = inflowModels.VoidNodeType

	}

	return &inflowNode, nil
}
