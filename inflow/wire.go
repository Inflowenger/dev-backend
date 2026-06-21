package inflow

import (
	"fmt"

	"github.com/nats-io/nats.go"
)

type InflowWire struct{}

func (isvc *InflowWire) RetrieveContext(msg *nats.Msg) {
	fmt.Println(msg.Header)
	// get context data from db
	msg.Respond([]byte(`{"header":{},"data":"{\"node1\":{\"b\":2,\"sum\":3,\"a\":1},\"node2\":{\"a\":1,\"b\":2,\"sum\":3}}"}`))
}

func (isvc *InflowWire) UpdateContext(msg *nats.Msg) {
	fmt.Println(string(msg.Data)) // save to db
	msg.Respond([]byte(`accepted`))
}

func (isvc *InflowWire) RetrieveFlow(msg *nats.Msg) {

	msg.Respond([]byte(`{}`))

}
