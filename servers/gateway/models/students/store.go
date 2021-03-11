package students

import (
	"errors"
	"time"
)

//ErrUserNotFound is returned when the student can't be found
var ErrUserNotFound = errors.New("student not found")

//Store represents a store for Users
type Store interface {
	//GetByID returns the User with the given ID
	GetByID(id int64) (*Student, error)

	//GetByEmail returns the User with the given email
	GetByEmail(email string) (*Student, error)

	//GetByUserName returns the User with the given Username
	GetByUserName(username string) (*Student, error)

	//Insert inserts the student into the database, and returns
	//the newly-inserted User, complete with the DBMS-assigned ID
	Insert(student *Student) (*Student, error)

	//Update applies UserUpdates to the given student ID
	//and returns the newly-updated student
	Update(id int64, updates *Updates) (*Student, error)

	//Delete deletes the student with the given ID
	Delete(id int64) error

	//Log successful student sign in
	LogSignIn(student *Student, time time.Time, clientIP string) error
}
