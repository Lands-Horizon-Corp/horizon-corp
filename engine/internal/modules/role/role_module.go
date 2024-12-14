package role

import "go.uber.org/fx"

var Module = fx.Module(
	"role-module",
	fx.Provide(
		NewRoleService,
	),
)
