package registry

import (
	"github.com/google/wire"
)

// ProviderSet is registry providers.
var ProviderSet = wire.NewSet(NewRegistrarServer)
