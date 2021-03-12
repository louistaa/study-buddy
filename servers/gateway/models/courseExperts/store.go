package courseExpert

import (
	"errors"
)

//ErrUserNotFound is returned when the student can't be found
var ErrUserNotFound = errors.New("courseExpert not found")

//Store represents a store for Users
type Store interface {
	//GetByID returns the User with the given ID
	GetByID(id int64) (*CourseExpert, error)

	//GetByStudentID returns the student and correlating classes with the given student ID
	GetByExpertID(expertID int64) (*[]int64, error)

	//GetByClassID returns the class and correlating students with the given class ID
	GetByClassID(classID int64) (*[]int64, error)

	//Insert inserts the student into the database, and returns
	//the newly-inserted User, complete with the DBMS-assigned ID
	Insert(courseExpert *CourseExpert) (*CourseExpert, error)

	//Delete deletes the student with the given ID
	Delete(id int64) error
}
