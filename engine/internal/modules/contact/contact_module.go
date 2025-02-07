package contact

import "go.uber.org/fx"

var Module = fx.Module(
	"contact-module",
	fx.Provide(
		NewContactService,
	),
)
