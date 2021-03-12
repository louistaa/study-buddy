package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"path"
	"strconv"
	"strings"

	"github.com/louistaa/study-buddy/servers/gateway/models/classes"
	"github.com/louistaa/study-buddy/servers/gateway/models/students"
	"github.com/louistaa/study-buddy/servers/gateway/sessions"
)

// ClassHandler handles requests for "classes" resource
func (hc *HandlerContext) ClassHandler(w http.ResponseWriter, r *http.Request) {
	auth := r.Header.Get("Authorization")
	sessionID := sessions.SessionID(strings.TrimPrefix(auth, "Bearer "))
	sessionState := &SessionState{}
	err := hc.SessionStore.Get(sessionID, sessionState)
	if err != nil {
		http.Error(w, "Authorization required", http.StatusUnauthorized)
		return
	}

	// POST /class - create a new course with is class ID and professor ID
	if r.Method == "POST" {
		if !strings.HasPrefix(r.Header.Get("Content-Type"), "application/json") {
			http.Error(w, "request body must be in JSON", http.StatusUnsupportedMediaType)
			return
		}
		newClass := &classes.NewClass{}
		body, _ := ioutil.ReadAll(r.Body)
		err := json.Unmarshal([]byte(body), newClass)
		if err != nil {
			http.Error(w, "JSON error: "+err.Error(), http.StatusBadRequest)
			return
		}
		err = newClass.Validate()
		if err != nil {
			http.Error(w, "Invalid Class: "+err.Error(), http.StatusBadRequest)
			return
		}
		class, err := newClass.ToClass()
		if err != nil {
			http.Error(w, "Error creating class: "+err.Error(), http.StatusInternalServerError)
			return
		}
		_, err = hc.ClassStore.Insert(class)
		if err != nil {
			http.Error(w, "Error inserting to ClassStore: "+err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		classJSON, _ := json.Marshal(class)
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(classJSON))
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// SpecificClassHandler handles requests for a specific class resource
func (hc *HandlerContext) SpecificClassHandler(w http.ResponseWriter, r *http.Request) {
	auth := r.Header.Get("Authorization")
	sessionID := sessions.SessionID(strings.TrimPrefix(auth, "Bearer "))
	sessionState := &SessionState{}
	err := hc.SessionStore.Get(sessionID, sessionState)
	if err != nil {
		http.Error(w, "Authorization required", http.StatusUnauthorized)
		return
	}

	// get the user id from the base path
	urlPath := path.Clean(r.URL.Path)
	base, endpoint := path.Split(urlPath)

	var classID int64
	getPeople := false
	getExperts := false

	if endpoint == "people" {
		getPeople = true
		classID, err = strconv.ParseInt(path.Base(base), 10, 64)
		// get people in class
	} else if endpoint == "experts" {
		getExperts = true
	} else { //parse base path for id int
		classID, err = strconv.ParseInt(endpoint, 10, 64)
		// get class info
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}
	// GET /class/{id} - gets class details
	// GET /class/{id}/people - gets list of people in a class
	if r.Method == "GET" {
		if getPeople {
			studentIDs, err := hc.StudentCourses.GetByClassID(classID)
			if err != nil { // if no class is found with that ID
				http.Error(w, "No class is found with that ID", http.StatusNotFound)
				return
			}

			students := []*students.Student{}
			for _, studentID := range studentIDs {
				student, err := hc.StudentStore.GetByID(studentID)
				if err != nil {
					http.Error(w, "Error marshaling student in class", http.StatusNotFound)
				}
				students = append(students, student)
			}

			w.Header().Add("Content-Type", "application/json")
			studentsJSON, _ := json.Marshal(students)
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(studentsJSON))
			return
		} else if getExperts {
			expertIDs, err := hc.CourseExpert.GetByClassID(classID)
			if err != nil { // if no class is found with that ID
				http.Error(w, "No class is found with that ID", http.StatusNotFound)
				return
			}

			experts := []*students.Student{}
			for _, expertID := range expertIDs {
				expert, err := hc.StudentStore.GetByID(expertID)
				if err != nil {
					http.Error(w, "Error marshaling student in class", http.StatusNotFound)
				}
				experts = append(experts, expert)
			}

			w.Header().Add("Content-Type", "application/json")
			expertsJSON, _ := json.Marshal(experts)
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(expertsJSON))
			return
		} else {
			// get class associated with requested class id from store
			class, err := hc.ClassStore.GetByID(classID)
			if err != nil { // if no class is found with that ID
				http.Error(w, "No class is found with that ID", http.StatusNotFound)
				return
			}
			w.Header().Add("Content-Type", "application/json")
			classJSON, _ := json.Marshal(class)
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(classJSON))
			return
		}
	} else if r.Method == "DELETE" { // DELETE /class/{id} - delete course
		_, err := hc.ClassStore.GetByID(classID)
		if err != nil { // if no class is found with that ID
			http.Error(w, "No class is found with that ID", http.StatusNotFound)
			return
		}
		err = hc.ClassStore.Delete(classID)
		if err != nil {
			http.Error(w, "Error inserting to ClassStore: "+err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write([]byte("Class deleted"))
		return
	} else if r.Method == "PATCH" {
		// TODO check if user is part of the staff
		// if classID != sess.Class.ID {
		// 	http.Error(w, "Class ID does not match up with the currently-authenticated user", http.StatusForbidden)
		// 	return
		// }
		if !strings.HasPrefix(r.Header.Get("Content-Type"), "application/json") {
			http.Error(w, "The request body must be in JSON", http.StatusUnsupportedMediaType)
			return
		}
		// update class with JSON request body
		u := &classes.Updates{}
		updates := json.NewDecoder(r.Body).Decode(u)
		if updates != nil {
			http.Error(w, updates.Error(), http.StatusBadRequest)
			return
		}
		class, err := hc.ClassStore.Update(classID, u)
		if err != nil {
			http.Error(w, "Error adding updates to Class", http.StatusInternalServerError)
			return
		}

		w.Header().Add("Content-Type", "application/json")
		classJSON, _ := json.Marshal(class)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(classJSON))
		return
	} else {
		http.Error(w, "Method is not GET, PATCH, or DELETE", http.StatusMethodNotAllowed)
		return
	}
}
