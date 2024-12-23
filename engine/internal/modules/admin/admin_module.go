package admin

import "go.uber.org/fx"

var Module = fx.Module(
	"admin-module",
	fx.Provide(
		NewAdminRepository,
		NewAdminService,
		NewAdminProvider,
	),
)
