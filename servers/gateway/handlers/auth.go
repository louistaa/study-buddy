package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/louistaa/study-buddy/servers/gateway/models/classes"
	"github.com/louistaa/study-buddy/servers/gateway/models/students"
	"github.com/louistaa/study-buddy/servers/gateway/sessions"

	"golang.org/x/crypto/bcrypt"
)

// StudentsHandler handles requests for the "users" resource
func (hc *HandlerContext) StudentsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		if !strings.HasPrefix(r.Header.Get("Content-Type"), "application/json") {
			http.Error(w, "request body must be in JSON", http.StatusUnsupportedMediaType)
			return
		}
		newStudent := &students.NewStudent{}
		body, _ := ioutil.ReadAll(r.Body)
		err := json.Unmarshal([]byte(body), newStudent)
		if err != nil {
			http.Error(w, "JSON error: "+err.Error(), http.StatusBadRequest)
			return
		}
		err = newStudent.Validate()
		if err != nil {
			http.Error(w, "Invalid User: "+err.Error(), http.StatusBadRequest)
			return
		}
		student, err := newStudent.ToStudent()
		if err != nil {
			http.Error(w, "Error creating user: "+err.Error(), http.StatusInternalServerError)
			return
		}
		insertStudent, err := hc.StudentStore.Insert(student)
		if err != nil {
			http.Error(w, "Error inserting to StudentStore: "+err.Error(), http.StatusInternalServerError)
			return
		}

		sessionState := &SessionState{StartTime: time.Now(), Student: *insertStudent}
		_, err = sessions.BeginSession(hc.SigningKey, hc.SessionStore, sessionState, w)
		w.Header().Set("Content-Type", "application/json")
		studentJSON, _ := json.Marshal(student)
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(studentJSON))
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// SpecificStudentHandler handles requests for a specific student
func (hc *HandlerContext) SpecificStudentHandler(w http.ResponseWriter, r *http.Request) {
	auth := r.Header.Get("Authorization")
	sessionID := sessions.SessionID(strings.TrimPrefix(auth, "Bearer "))
	sessionState := &SessionState{}
	err := hc.SessionStore.Get(sessionID, sessionState)
	if err != nil {
		http.Error(w, "Authorization required", http.StatusUnauthorized)
		return
	}

	if r.Method == "GET" {
		var userID int64

		base, endpoint := path.Split(path.Clean(r.URL.Path))
		if endpoint == "me" {
			userID = sessionState.Student.ID
		} else if endpoint == "classes" {
			userID, _ = strconv.ParseInt(path.Base(base), 10, 64)
			classIDs, err := hc.StudentCourses.GetByStudentID(userID)
			if err != nil { // if no class is found with that ID
				http.Error(w, "No student is found with that ID", http.StatusNotFound)
				return
			}

			classes := []*classes.Class{}
			for _, classID := range classIDs {
				class, err := hc.ClassStore.GetByID(classID)
				if err != nil {
					http.Error(w, "Error marshaling class", http.StatusNotFound)
				}
				classes = append(classes, class)
			}

			w.Header().Add("Content-Type", "application/json")
			classesJSON, _ := json.Marshal(classes)
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(classesJSON))
			return
		} else {
			userID, _ = strconv.ParseInt(path.Base(r.URL.Path), 10, 64)
		}
		student, err := hc.StudentStore.GetByID(userID)
		if err != nil {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		studentJSON, _ := json.Marshal(student)
		w.Write([]byte(studentJSON))
	} else if r.Method == "PATCH" {
		if path.Base(r.URL.Path) != "me" {
			userID, _ := strconv.ParseInt(path.Base(r.URL.Path), 10, 64)
			if userID != sessionState.Student.ID {
				http.Error(w, "User not authenticated", http.StatusForbidden)
				return
			}
		}
		if !strings.HasPrefix(r.Header.Get("Content-Type"), "application/json") {
			http.Error(w, "request body must be in JSON", http.StatusUnsupportedMediaType)
			return
		}
		updates := &students.Updates{}
		body, _ := ioutil.ReadAll(r.Body)
		err = json.Unmarshal([]byte(body), updates)
		if err != nil {
			http.Error(w, "Unable to create update from JSON: "+err.Error(), http.StatusBadRequest)
			return
		}
		user, err := hc.StudentStore.Update(sessionState.Student.ID, updates)
		if err != nil {
			http.Error(w, "Couldn't update user: "+err.Error(), http.StatusInternalServerError)
			return
		}
		sessionState.Student = *user
		err = hc.SessionStore.Save(sessionID, sessionState)
		if err != nil {
			http.Error(w, "Couldn't update session: "+err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		studentJSON, _ := json.Marshal(user)
		w.Write([]byte(studentJSON))
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

// SessionsHandler handles requests for the "sessions" resource
func (hc *HandlerContext) SessionsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		if r.Header.Get("Content-Type") != "application/json" {
			http.Error(w, "Body must be in json", http.StatusUnsupportedMediaType)
			return
		}
		credentials := &students.Credentials{}
		body, _ := ioutil.ReadAll(r.Body)
		err := json.Unmarshal(body, credentials)
		if err != nil {
			http.Error(w, "Unable to create credentials from JSON: "+err.Error(), http.StatusBadRequest)
			return
		}
		user, err := hc.StudentStore.GetByEmail(credentials.Email)
		if err != nil {
			bcrypt.GenerateFromPassword([]byte(credentials.Password), bcrypt.DefaultCost)
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
			return
		}
		err = bcrypt.CompareHashAndPassword(user.PassHash, []byte(credentials.Password))
		if err != nil {
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
			return
		}
		sessionState := &SessionState{StartTime: time.Now(), Student: *user}
		_, err = sessions.BeginSession(hc.SigningKey, hc.SessionStore, sessionState, w)
		if err != nil {
			http.Error(w, "Unable to begin session: "+err.Error(), http.StatusInternalServerError)
			return
		}
		clientIP := r.RemoteAddr
		if clientIP == "" {
			clientIP = r.Header.Get("X-Forwarded-For")
		}

		err = hc.StudentStore.LogSignIn(user, time.Now(), clientIP)
		if err != nil {
			http.Error(w, "Failed to log user sign in: "+err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")
		userJSON, _ := json.Marshal(user)
		w.Write([]byte(userJSON))
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

// SpecificSessionHanlder handles requests for a specific session
func (hc *HandlerContext) SpecificSessionHanlder(w http.ResponseWriter, r *http.Request) {
	if r.Method == "DELETE" {
		if path.Base(r.URL.Path) != "mine" {
			http.Error(w, "Can't delete session for other users", http.StatusForbidden)
			return
		}
		_, err := sessions.EndSession(r, hc.SigningKey, hc.SessionStore)
		if err != nil {
			http.Error(w, "Unable to end session: "+err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write([]byte("signed out"))
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
}
