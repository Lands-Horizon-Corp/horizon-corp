package helpers

import "go.uber.org/fx"

var Module = fx.Module(
	"helpers",
	fx.Provide(NewHelperFunction),
)
