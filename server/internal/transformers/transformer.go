package transformers

import (
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/providers"
	"go.uber.org/fx"
)

type Transformer struct {
	storage *providers.StorageProvider
}

func NewModelResource(
	storage *providers.StorageProvider,
) *Transformer {
	return &Transformer{
		storage: storage,
	}
}

var Module = fx.Module(
	"transformers",
	fx.Provide(
		NewModelResource,
	),
)
