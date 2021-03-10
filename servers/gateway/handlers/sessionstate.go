package handlers

import (
	"study-buddy/servers/gateway/models/students"
	"time"
)

// SessionState struct includes start time and user
type SessionState struct {
	StartTime time.Time        `json:"startTime"`
	Student   students.Student `json:"student"`
}
