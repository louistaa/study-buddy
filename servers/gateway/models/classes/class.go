package classes

import (
	"fmt"
	"strings"
)

//Class represents a class in the database
type Class struct {
	ID             int64  `json:"id"`
	Name           string `json:"name"`
	DepartmentName string `json:"departmentName"`
	ProfessorName  string `json:"professorName"`
	QuarterName    string `json:"quarterName"`
}

//NewClass represents a new student signing up for an account
type NewClass struct {
	Name           string `json:"name"`
	DepartmentName string `json:"departmentName"`
	ProfessorName  string `json:"professorName"`
	QuarterName    string `json:"quarterName"`
}

//Updates represents allowed updates to a class
type Updates struct {
	Name           string `json:"name"`
	DepartmentName string `json:"departmentName"`
	ProfessorName  string `json:"professorName"`
	QuarterName    string `json:"quarterName"`
}

//Validate validates the new class and returns an error if
//any of the validation rules fail, or nil if its valid
func (nc *NewClass) Validate() error {
	if len(nc.Name) == 0 || strings.Contains(nc.Name, " ") {
		return fmt.Errorf("Name must be non-zero length and may not contain spaces")
	}

	if len(nc.DepartmentName) == 0 {
		return fmt.Errorf("Department must be non-zero length and may not contain spaces")
	}

	if len(nc.ProfessorName) == 0 {
		return fmt.Errorf("Department must be non-zero length and may not contain spaces")
	}

	return nil
}

//ToClass converts the NewClass to a Class
func (nc *NewClass) ToClass() (*Class, error) {
	validationErr := nc.Validate()
	if validationErr != nil {
		return nil, validationErr
	}

	newClass := &Class{
		Name:           nc.Name,
		DepartmentName: nc.DepartmentName,
		ProfessorName:  nc.ProfessorName,
		QuarterName:    nc.QuarterName,
	}

	return newClass, nil
}

//ApplyUpdates applies the updates to the class. An error
//is returned if the updates are invalid
func (c *Class) ApplyUpdates(updates *Updates) error {
	if updates.Name != "" {
		c.Name = updates.Name
	}

	if updates.DepartmentName != "" {
		c.DepartmentName = updates.DepartmentName
	}

	if updates.ProfessorName != "" {
		c.ProfessorName = updates.ProfessorName
	}

	if updates.QuarterName != "" {
		c.QuarterName = updates.QuarterName
	}

	return nil
}
