package classes

import (
	"errors"
)

//ErrUserNotFound is returned when the student can't be found
var ErrUserNotFound = errors.New("class not found")

//Store represents a store for Users
type Store interface {
	//GetByID returns the Class with the given ID
	GetByID(id int64) (*Class, error)

	//Insert inserts the class into the database, and returns
	//the newly-inserted Class, complete with the DBMS-assigned ID
	Insert(class *Class) (*Class, error)

	//Update applies ClassUpdates to the given class ID
	//and returns the newly-updated class
	Update(id int64, updates *Updates) (*Class, error)

	//Delete deletes the class with the given ID
	Delete(id int64) error
}
