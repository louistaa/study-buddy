package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	courseExpert "github.com/louistaa/study-buddy/servers/gateway/models/courseExperts"
	"github.com/louistaa/study-buddy/servers/gateway/models/studentCourses"
	"github.com/louistaa/study-buddy/servers/gateway/sessions"
)

const contentType = "Content-Type"
const applicationJSON = "application/json"

// Register handles requests for a class registration
func (hc *HandlerContext) RegisterClass(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Status method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if !strings.HasPrefix(r.Header.Get("Content-Type"), "application/json") {
		http.Error(w, "request body must be in JSON", http.StatusUnsupportedMediaType)
		return
	}

	auth := r.Header.Get("Authorization")
	sessionID := sessions.SessionID(strings.TrimPrefix(auth, "Bearer "))
	sessionState := &SessionState{}
	err := hc.SessionStore.Get(sessionID, sessionState)
	if err != nil {
		http.Error(w, "Authorization required", http.StatusUnauthorized)
		return
	}

	registration := &studentCourses.StudentCourses{}
	body, _ := ioutil.ReadAll(r.Body)
	err = json.Unmarshal([]byte(body), registration)
	if err != nil {
		http.Error(w, "JSON error: "+err.Error(), http.StatusBadRequest)
		return
	}

	if registration.StudentID != sessionState.Student.ID {
		http.Error(w, "Student not logged in! You can't register other students for classes"+err.Error(), http.StatusBadRequest)
		return
	}

	_, err = hc.ClassStore.GetByID(registration.CourseID)
	if err != nil {
		http.Error(w, "Invalid class"+err.Error(), http.StatusBadRequest)
		return
	}

	registration, err = hc.StudentCourses.Insert(registration)
	if err != nil {
		http.Error(w, "Couldn't register student: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	registerJSON, _ := json.Marshal(registration)
	w.Write([]byte(registerJSON))
}

// RegisterExpert handles requests for a expert registration
func (hc *HandlerContext) RegisterExpert(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Status method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if !strings.HasPrefix(r.Header.Get("Content-Type"), "application/json") {
		http.Error(w, "request body must be in JSON", http.StatusUnsupportedMediaType)
		return
	}

	auth := r.Header.Get("Authorization")
	sessionID := sessions.SessionID(strings.TrimPrefix(auth, "Bearer "))
	sessionState := &SessionState{}
	err := hc.SessionStore.Get(sessionID, sessionState)
	if err != nil {
		http.Error(w, "Authorization required", http.StatusUnauthorized)
		return
	}

	registration := &courseExpert.CourseExpert{}
	body, _ := ioutil.ReadAll(r.Body)
	err = json.Unmarshal([]byte(body), registration)
	if err != nil {
		http.Error(w, "JSON error: "+err.Error(), http.StatusBadRequest)
		return
	}

	if registration.ExpertID != sessionState.Student.ID {
		http.Error(w, "Student not logged in! You can't register other students as an expert"+err.Error(), http.StatusBadRequest)
		return
	}

	_, err = hc.ClassStore.GetByID(registration.CourseID)
	if err != nil {
		http.Error(w, "Invalid class"+err.Error(), http.StatusBadRequest)
		return
	}

	registration, err = hc.CourseExpert.Insert(registration)
	if err != nil {
		http.Error(w, "Couldn't register student: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	registerJSON, _ := json.Marshal(registration)
	w.Write([]byte(registerJSON))
}

func jsonWriting(w http.ResponseWriter, r *http.Request, code int, structure interface{}) {
	w.Header().Add(contentType, applicationJSON)
	w.WriteHeader(code)
	enc := json.NewEncoder(w)
	if encodingErr := enc.Encode(structure); encodingErr != nil {
		http.Error(w, encodingErr.Error(), code)
		return
	}
}
