package api

import (
	"pinking-go/server/store"

	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tools/hook"
)

func RequireLockoutMiddleware() *hook.Handler[*core.RequestEvent] {
	return &hook.Handler[*core.RequestEvent]{
		Id:       "LockoutMiddleware",
		Func:     lockoutMiddleware,
		Priority: 0,
	}
}

func lockoutMiddleware(e *core.RequestEvent) error {
	if store.IsLockoutEnabled(e.Auth) {
		return e.ForbiddenError("error_lockout_enabled", nil)
	}

	return e.Next()
}
