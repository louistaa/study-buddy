package classes

import (
	"database/sql"
)

// GetByType is an enumerate for GetBy* functions implemented
// by MySQLStore structs
type GetByType string

// These are the enumerates for GetByType
const (
	ID GetByType = "ID"
)

// MySQLStore is a user.Store backed by MySQL
type MySQLStore struct {
	Database *sql.DB
}

// NewMySQLStore constructs a new MySQLStore, and returns an error
// if there is a problem along the way.
func NewMySQLStore(dataSourceName string) (*MySQLStore, error) {
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return nil, err
	}

	return &MySQLStore{db}, nil
}

// getByProvidedType gets a specific user given the provided type.
// This requires the GetByType to be "unique" in the database.
func (ms *MySQLStore) getByProvidedType(t GetByType, arg interface{}) (*Class, error) {
	sel := string("select ID, Name, Department, ProfessorID from Courses where " + t + " = ?")

	rows, err := ms.Database.Query(sel, arg)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	class := &Class{}

	// Should never have more than one row, so only grab one
	rows.Next()
	if err := rows.Scan(
		&class.ID,
		&class.Name,
		&class.DepartmentName,
		&class.ProfessorName,
		&class.QuarterName); err != nil {
		return nil, err
	}
	return class, nil
}

//GetByID returns the Class with the given ID
func (ms *MySQLStore) GetByID(id int64) (*Class, error) {
	return ms.getByProvidedType(ID, id)
}

//Insert inserts the class into the database, and returns
//the newly-inserted Class, complete with the DBMS-assigned ID
func (ms *MySQLStore) Insert(class *Class) (*Class, error) {
	ins := "insert into Courses(Name, DepartmentName, ProfessorName, QuarterName) values (?,?,?,?)"
	res, err := ms.Database.Exec(ins, class.Name, class.DepartmentName, class.ProfessorName, class.QuarterName)
	if err != nil {
		return nil, err
	}

	lid, lidErr := res.LastInsertId()
	if lidErr != nil {
		return nil, lidErr
	}

	class.ID = lid
	return class, nil
}

//Update applies ClassUpdates to the given class ID
//and returns the newly-updated class
func (ms *MySQLStore) Update(id int64, updates *Updates) (*Class, error) {
	// Assumes updates ALWAYS includes FirstName and LastName
	upd := "update Courses set Name = ?, DepartmentName = ?, ProfessorName = ?, QuarterName = ? where ID = ?"
	res, err := ms.Database.Exec(upd, updates.Name, updates.DepartmentName, updates.ProfessorName, updates.QuarterName, id)
	if err != nil {
		return nil, err
	}

	rowsAffected, rowsAffectedErr := res.RowsAffected()
	if rowsAffectedErr != nil {
		return nil, rowsAffectedErr
	}

	if rowsAffected != 1 {
		return nil, ErrUserNotFound
	}

	// Get the class using GetByID
	return ms.GetByID(id)
}

//Delete deletes the class with the given ID
func (ms *MySQLStore) Delete(id int64) error {
	del := "delete from Courses where ID = ?"
	res, err := ms.Database.Exec(del, id)
	if err != nil {
		return err
	}

	rowsAffected, rowsAffectedErr := res.RowsAffected()
	if rowsAffectedErr != nil {
		return rowsAffectedErr
	}

	if rowsAffected != 1 {
		return ErrUserNotFound
	}

	return nil
}
