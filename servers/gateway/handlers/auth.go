package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/louistaa/study-buddy/servers/gateway/models/users"
	"github.com/louistaa/study-buddy/servers/gateway/sessions"
	"golang.org/x/crypto/bcrypt"
)

// UsersHandler handles requests for the "users" resource
func (hc *HandlerContext) UsersHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		if !strings.HasPrefix(r.Header.Get("Content-Type"), "application/json") {
			http.Error(w, "request body must be in JSON", http.StatusUnsupportedMediaType)
			return
		}
		newUser := &users.NewUser{}
		body, _ := ioutil.ReadAll(r.Body)
		err := json.Unmarshal([]byte(body), newUser)
		if err != nil {
			http.Error(w, "JSON error: "+err.Error(), http.StatusBadRequest)
			return
		}
		err = newUser.Validate()
		if err != nil {
			http.Error(w, "Invalid User: "+err.Error(), http.StatusBadRequest)
			return
		}
		user, err := newUser.ToUser()
		if err != nil {
			http.Error(w, "Error creating user: "+err.Error(), http.StatusInternalServerError)
			return
		}
		insertUser, err := hc.UserStore.Insert(user)
		if err != nil {
			http.Error(w, "Error inserting to UserStore: "+err.Error(), http.StatusInternalServerError)
			return
		}

		sessionState := &SessionState{StartTime: time.Now(), User: *insertUser}
		_, err = sessions.BeginSession(hc.SigningKey, hc.SessionStore, sessionState, w)
		w.Header().Set("Content-Type", "application/json")
		userJSON, _ := json.Marshal(user)
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(userJSON))
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// SpecificUserHandler handles requests for a specific users
func (hc *HandlerContext) SpecificUserHandler(w http.ResponseWriter, r *http.Request) {
	auth := r.Header.Get("Authorization")
	sessionID := sessions.SessionID(strings.TrimPrefix(auth, "Bearer "))
	sessionState := &SessionState{}
	err := hc.SessionStore.Get(sessionID, sessionState)
	if err != nil {
		http.Error(w, "Authorization required", http.StatusUnauthorized)
		return
	}

	if r.Method == "GET" {
		userID := int64(-1)

		if path.Base(r.URL.Path) == "me" {
			userID = sessionState.User.ID
		} else {
			userID, _ = strconv.ParseInt(path.Base(r.URL.Path), 10, 64)
		}
		user, err := hc.UserStore.GetByID(userID)
		if err != nil {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		userJSON, _ := json.Marshal(user)
		w.Write([]byte(userJSON))
	} else if r.Method == "PATCH" {
		if path.Base(r.URL.Path) != "me" {
			userID, _ := strconv.ParseInt(path.Base(r.URL.Path), 10, 64)
			if userID != sessionState.User.ID {
				http.Error(w, "User not authenticated", http.StatusForbidden)
				return
			}
		}
		if !strings.HasPrefix(r.Header.Get("Content-Type"), "application/json") {
			http.Error(w, "request body must be in JSON", http.StatusUnsupportedMediaType)
			return
		}
		updates := &users.Updates{}
		body, _ := ioutil.ReadAll(r.Body)
		err = json.Unmarshal([]byte(body), updates)
		if err != nil {
			http.Error(w, "Unable to create update from JSON: "+err.Error(), http.StatusBadRequest)
			return
		}
		user, err := hc.UserStore.Update(sessionState.User.ID, updates)
		if err != nil {
			http.Error(w, "Couldn't update user: "+err.Error(), http.StatusInternalServerError)
			return
		}
		sessionState.User = *user
		err = hc.SessionStore.Save(sessionID, sessionState)
		if err != nil {
			http.Error(w, "Couldn't update session: "+err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		userJSON, _ := json.Marshal(user)
		w.Write([]byte(userJSON))
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
		credentials := &users.Credentials{}
		body, _ := ioutil.ReadAll(r.Body)
		err := json.Unmarshal(body, credentials)
		if err != nil {
			http.Error(w, "Unable to create credentials from JSON: "+err.Error(), http.StatusBadRequest)
			return
		}
		user, err := hc.UserStore.GetByEmail(credentials.Email)
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
		sessionState := &SessionState{StartTime: time.Now(), User: *user}
		_, err = sessions.BeginSession(hc.SigningKey, hc.SessionStore, sessionState, w)
		if err != nil {
			http.Error(w, "Unable to begin session: "+err.Error(), http.StatusInternalServerError)
			return
		}
		clientIP := r.RemoteAddr
		if clientIP == "" {
			clientIP = r.Header.Get("X-Forwarded-For")
		}

		err = hc.UserStore.LogSignIn(user, time.Now(), clientIP)
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