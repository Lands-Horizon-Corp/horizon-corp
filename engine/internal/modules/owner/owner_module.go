package owner

import "go.uber.org/fx"

var Module = fx.Module(
	"owner-module",
	fx.Provide(
		NewOwnerService,
	),
)
