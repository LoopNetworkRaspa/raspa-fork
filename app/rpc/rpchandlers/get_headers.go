package rpchandlers

import (
	"github.com/LoopNetworkRaspa/Raspa-Fork/app/appmessage"
	"github.com/LoopNetworkRaspa/Raspa-Fork/app/rpc/rpccontext"
	"github.com/LoopNetworkRaspa/Raspa-Fork/infrastructure/network/netadapter/router"
)

// HandleGetHeaders handles the respectively named RPC command
func HandleGetHeaders(context *rpccontext.Context, _ *router.Router, request appmessage.Message) (appmessage.Message, error) {
	response := &appmessage.GetHeadersResponseMessage{}
	response.Error = appmessage.RPCErrorf("not implemented")
	return response, nil
}
