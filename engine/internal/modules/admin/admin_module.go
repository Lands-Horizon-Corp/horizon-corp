package admin

import "go.uber.org/fx"

var Module = fx.Module(
	"helpers",
	fx.Provide(NewAdminRepository),
	fx.Provide(NewAdminService),
	fx.Provide(NewAdminProvider),
)
