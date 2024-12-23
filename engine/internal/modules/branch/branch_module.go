package branch

import "go.uber.org/fx"

var Module = fx.Module(
	"branch-module",
	fx.Provide(
		NewBranchService,
	),
)
