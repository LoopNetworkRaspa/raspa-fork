package flowcontext

import "github.com/LoopNetworkRaspa/Raspa-Fork/infrastructure/config"

// Config returns an instance of *config.Config associated to the flow context.
func (f *FlowContext) Config() *config.Config {
	return f.cfg
}
