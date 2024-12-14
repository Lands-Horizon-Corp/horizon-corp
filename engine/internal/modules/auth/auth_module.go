package auth

import "go.uber.org/fx"

var Module = fx.Module(
	"auth-module",
	fx.Provide(
		NewAuthProvider,
		NewAuthService,
	),
)
