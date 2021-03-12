package students

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"net/mail"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

//bcryptCost is the default bcrypt cost to use when hashing passwords
var bcryptCost = 13

//Class represents a class in the database
type Class struct {
	ID        	int64  `json:"id"`
	Name     	string `json:"name"` 
	Department  string `json:"department"`
	ProfessorID string `json:"professorID"`
}

//NewClass represents a new student signing up for an account
type NewClass struct {
	Name     	string `json:"name"` 
	Department  string `json:"department"`
}

//Updates represents allowed updates to a class
type Updates struct {
	Name     	string `json:"name"` 
	Department  string `json:"department"`
}

//Validate validates the new class and returns an error if
//any of the validation rules fail, or nil if its valid
func (nc *NewClass) Validate() error {
	if len(nc.Name) == 0 || strings.Contains(nc.Name, " ") {
		return fmt.Errorf("Name must be non-zero length and may not contain spaces")
	}

	if len(nc.Department) == 0 || strings.Contains(nc.Department, " ") {
		return fmt.Errorf("Department must be non-zero length and may not contain spaces")
	}

	return nil
}

//ToClass converts the NewClass to a Class
func (nc *NewClass) ToClass() (*Class, error) {
	validationErr := ns.Validate()
	if validationErr != nil {
		return nil, validationErr
	}

	newClass := &Class{
		Email:     nc.Name,
		UserName:  nc.Department
	}
	return newClass, nil
}

//Name returns the name of the class
func (c *Class) Name() string {
	return c.Name
}

//ApplyUpdates applies the updates to the class. An error
//is returned if the updates are invalid
func (c *Class) ApplyUpdates(updates *Updates) error {
	if updates.Name != "" {
		c.Name = updates.Name
	}

	if updates.Department != "" {
		c.Department = updates.Department
	}

	return nil
}
