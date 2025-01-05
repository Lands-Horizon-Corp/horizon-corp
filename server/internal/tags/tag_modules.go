package tags

import "go.uber.org/fx"

var Module = fx.Module(
	"tags",
	fx.Provide(
		NewRecordTag,
		NewRequestTag,
		NewRecordTag,
	),
)
