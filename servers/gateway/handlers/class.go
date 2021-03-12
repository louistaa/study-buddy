package handlers

import "net/http"

// ClassHandler handles requests for "classes" resource
func ClassHandler(w http.ResponseWriter, r *http.Request) {
	// GET /students/{id}/class gets all classes student with given id is enrolled in (would this go in the students handler?)
	
	// POST /class - create a new course with is class ID and professor ID
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

		w.Header().Set("Content-Type", "application/json")
		studentJSON, _ := json.Marshal(student)
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(studentJSON))
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// SpecificClassHandler handles requests for a specific class resource
func SpecificClassHandler(w http.ResponseWriter, r *http.Request) {
	// GET /class/{id} - gets class details
	// GET /class/{id}/people - gets list of people in a class
	// DELETE /class/{id} - delete course
}
