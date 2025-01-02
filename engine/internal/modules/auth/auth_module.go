package auth

import (
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/modules/auth/auth_accounts"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"auth-module",
	fx.Provide(
		auth_accounts.NewAuthAccount,
		NewAuthProvider,
		NewAuthFootstep,
		NewAuthService,
	),
)
