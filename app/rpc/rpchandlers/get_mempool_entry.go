package rpchandlers

import (
	"github.com/LoopNetworkRaspa/Raspa-Fork/app/appmessage"
	"github.com/LoopNetworkRaspa/Raspa-Fork/app/rpc/rpccontext"
	"github.com/LoopNetworkRaspa/Raspa-Fork/domain/consensus/model/externalapi"
	"github.com/LoopNetworkRaspa/Raspa-Fork/domain/consensus/utils/transactionid"
	"github.com/LoopNetworkRaspa/Raspa-Fork/infrastructure/network/netadapter/router"
)

// HandleGetMempoolEntry handles the respectively named RPC command
func HandleGetMempoolEntry(context *rpccontext.Context, _ *router.Router, request appmessage.Message) (appmessage.Message, error) {

	transaction := &externalapi.DomainTransaction{}
	var found bool
	var isOrphan bool

	getMempoolEntryRequest := request.(*appmessage.GetMempoolEntryRequestMessage)

	transactionID, err := transactionid.FromString(getMempoolEntryRequest.TxID)
	if err != nil {
		errorMessage := &appmessage.GetMempoolEntryResponseMessage{}
		errorMessage.Error = appmessage.RPCErrorf("Transaction ID could not be parsed: %s", err)
		return errorMessage, nil
	}

	mempoolTransaction, isOrphan, found := context.Domain.MiningManager().GetTransaction(transactionID, !getMempoolEntryRequest.FilterTransactionPool, getMempoolEntryRequest.IncludeOrphanPool)

	if !found {
		errorMessage := &appmessage.GetMempoolEntryResponseMessage{}
		errorMessage.Error = appmessage.RPCErrorf("Transaction %s was not found", transactionID)
		return errorMessage, nil
	}

	rpcTransaction := appmessage.DomainTransactionToRPCTransaction(mempoolTransaction)
	err = context.PopulateTransactionWithVerboseData(rpcTransaction, nil)
	if err != nil {
		return nil, err
	}
	return appmessage.NewGetMempoolEntryResponseMessage(transaction.Fee, rpcTransaction, isOrphan), nil
}
