package handlers

import (
	"github.com/louistaa/study-buddy/servers/gateway/models/students"
	"github.com/louistaa/study-buddy/servers/gateway/sessions"
)

// HandlerContext struct provides access to
// globals, such as the key used for signing
// and verifying SessionIDs, the session store
// and the user store
type HandlerContext struct {
	SigningKey   string
	SessionStore sessions.Store
	StudentStore students.Store
}
