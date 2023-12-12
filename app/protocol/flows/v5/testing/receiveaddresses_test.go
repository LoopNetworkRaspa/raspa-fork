package testing

import (
	"github.com/LoopNetworkRaspa/Raspa-Fork/app/protocol/flows/v5/addressexchange"
	"testing"
	"time"

	"github.com/LoopNetworkRaspa/Raspa-Fork/app/appmessage"
	peerpkg "github.com/LoopNetworkRaspa/Raspa-Fork/app/protocol/peer"
	"github.com/LoopNetworkRaspa/Raspa-Fork/domain/consensus"
	"github.com/LoopNetworkRaspa/Raspa-Fork/domain/consensus/utils/testutils"
	"github.com/LoopNetworkRaspa/Raspa-Fork/infrastructure/network/addressmanager"
	"github.com/LoopNetworkRaspa/Raspa-Fork/infrastructure/network/netadapter/router"
)

type fakeReceiveAddressesContext struct{}

func (f fakeReceiveAddressesContext) AddressManager() *addressmanager.AddressManager {
	return nil
}

func TestReceiveAddressesErrors(t *testing.T) {
	testutils.ForAllNets(t, true, func(t *testing.T, consensusConfig *consensus.Config) {
		incomingRoute := router.NewRoute("incoming")
		outgoingRoute := router.NewRoute("outgoing")
		peer := peerpkg.New(nil)
		errChan := make(chan error)
		go func() {
			errChan <- addressexchange.ReceiveAddresses(fakeReceiveAddressesContext{}, incomingRoute, outgoingRoute, peer)
		}()

		_, err := outgoingRoute.DequeueWithTimeout(time.Second)
		if err != nil {
			t.Fatalf("DequeueWithTimeout: %+v", err)
		}

		// Sending addressmanager.GetAddressesMax+1 addresses should trigger a ban
		err = incomingRoute.Enqueue(appmessage.NewMsgAddresses(make([]*appmessage.NetAddress,
			addressmanager.GetAddressesMax+1)))
		if err != nil {
			t.Fatalf("Enqueue: %+v", err)
		}

		select {
		case err := <-errChan:
			checkFlowError(t, err, true, true, "address count exceeded")
		case <-time.After(time.Second):
			t.Fatalf("timed out after %s", time.Second)
		}
	})
}
