package providers

import "go.uber.org/fx"

var Module = fx.Module(
	"providers",
	fx.Provide(NewCacheProvider),
	fx.Provide(NewDatabaseProvider),
	fx.Provide(NewLoggerProvider),
	fx.Provide(NewStorageProvider),
)
