package handlers

import (
	"time"

	"github.com/louistaa/study-buddy/servers/gateway/models/students"
)

// SessionState struct includes start time and user
type SessionState struct {
	StartTime time.Time        `json:"startTime"`
	Student   students.Student `json:"student"`
}
