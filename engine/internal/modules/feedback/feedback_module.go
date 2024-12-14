package feedback

import "go.uber.org/fx"

var Module = fx.Module(
	"feedback-module",
	fx.Provide(
		NewFeedbackService,
	),
)
