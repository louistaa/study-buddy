package handlers

import (
	"time"

	"github.com/louistaa/study-buddy/servers/gateway/models/users"
)

// SessionState struct includes start time and user
type SessionState struct {
	StartTime time.Time  `json:"startTime"`
	User      users.User `json:"user"`
}
