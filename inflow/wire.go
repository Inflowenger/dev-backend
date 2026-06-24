package inflow

import (
	"fmt"
	"strings"

	"github.com/Inflowenger/dev-backend/repository"
	inflowModels "github.com/Inflowenger/inflow-fusion/models"
	"github.com/bytedance/sonic"
	"github.com/nats-io/nats.go"
)

type InflowWire struct{}

func (isvc *InflowWire) RetrieveContext(msg *nats.Msg) {
	fmt.Println(msg.Header)
	// get context data from db
	parts := strings.Split(msg.Subject, ".")
	ctxId := parts[len(parts)-1]
	rec:=repository.GetContextById(ctxId)
	ctxData,_:=sonic.MarshalString(rec.Context)
	wantedContext:=inflowModels.ContextDoc{Header : rec.Header,Data:ctxData}
	b, _ := sonic.Marshal(wantedContext)
	msg.Respond(b)
	// msg.Respond([]byte(`{"header":{},"data":"{\"node1\":{\"b\":2,\"sum\":3,\"a\":1},\"node2\":{\"a\":1,\"b\":2,\"sum\":3}}"}`))
}

func (isvc *InflowWire) UpdateContext(msg *nats.Msg) {
	fmt.Println(string(msg.Data)) // save to db
	msg.Respond([]byte(`accepted`))
}

func (isvc *InflowWire) RetrieveFlow(msg *nats.Msg) {
	//pattern of get_flow is inflow.req.flow.get.{flowId}
	// get flowId
	parts := strings.Split(msg.Subject, ".")
	flowId := parts[len(parts)-1]
	rec, err := repository.GetFlowById(flowId)
	if err != nil {
		msg.Respond([]byte(`{}`))
		fmt.Println("given flow id not found or exception error occurred")
	}
	_, cmp, err := FLowCompiler(*rec)
	wantedFlow:=inflowModels.Flow{
		UUID: flowId,
		Nodes: []inflowModels.Node{},
	}
	for _, el := range cmp {
		wantedFlow.Nodes = append(wantedFlow.Nodes, *el)
	}
	b, _ := sonic.Marshal(wantedFlow)

	msg.Respond(b)

}
