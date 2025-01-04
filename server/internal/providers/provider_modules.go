package providers

import "go.uber.org/fx"

var Module = fx.Module(
	"providers",
	fx.Provide(
		NewLoggerProvider,
		NewCacheProvider,
		NewDatabaseProvider,
		NewEmailProvider,
		NewStorageProvider,
		NewSMSProvider,
		NewEngineProvider,
		NewOTPProvider,
		NewTokenProvider,
		NewWebSocketProvider,
	),
)
