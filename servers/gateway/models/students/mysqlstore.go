package students

import (
	"database/sql"
	"time"
)

// GetByType is an enumerate for GetBy* functions implemented
// by MySQLStore structs
type GetByType string

// These are the enumerates for GetByType
const (
	ID       GetByType = "ID"
	Email    GetByType = "Email"
	UserName GetByType = "UserName"
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
func (ms *MySQLStore) getByProvidedType(t GetByType, arg interface{}) (*Student, error) {
	sel := string("select ID, Email, PassHash, UserName, FirstName, LastName, PhotoURL, Major, PhoneNumber from Students where " + t + " = ?")

	rows, err := ms.Database.Query(sel, arg)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	student := &Student{}

	// Should never have more than one row, so only grab one
	rows.Next()
	if err := rows.Scan(
		&student.ID,
		&student.Email,
		&student.PassHash,
		&student.UserName,
		&student.FirstName,
		&student.LastName,
		&student.PhotoURL,
		&student.Major,
		&student.PhoneNumber); err != nil {
		return nil, err
	}
	return student, nil
}

//GetByID returns the User with the given ID
func (ms *MySQLStore) GetByID(id int64) (*Student, error) {
	return ms.getByProvidedType(ID, id)
}

//GetByEmail returns the User with the given email
func (ms *MySQLStore) GetByEmail(email string) (*Student, error) {
	return ms.getByProvidedType(Email, email)
}

//GetByUserName returns the User with the given Username
func (ms *MySQLStore) GetByUserName(username string) (*Student, error) {
	return ms.getByProvidedType(UserName, username)
}

//Insert inserts the user into the database, and returns
//the newly-inserted User, complete with the DBMS-assigned ID
func (ms *MySQLStore) Insert(student *Student) (*Student, error) {
	ins := "insert into Students(Email, PassHash, UserName, FirstName, LastName, PhotoURL, Major, PhoneNumber) values (?,?,?,?,?,?,?,?)"
	res, err := ms.Database.Exec(ins, student.Email, student.PassHash, student.UserName, student.FirstName, student.LastName, student.PhotoURL, student.Major, student.PhoneNumber)
	if err != nil {
		return nil, err
	}

	lid, lidErr := res.LastInsertId()
	if lidErr != nil {
		return nil, lidErr
	}

	student.ID = lid
	return student, nil
}

//Update applies UserUpdates to the given user ID
//and returns the newly-updated user
func (ms *MySQLStore) Update(id int64, updates *Updates) (*Student, error) {
	// Assumes updates ALWAYS includes FirstName and LastName
	upd := "update Students set FirstName = ?, LastName = ? where ID = ?"
	res, err := ms.Database.Exec(upd, updates.FirstName, updates.LastName, id)
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

	// Get the user using GetByID
	return ms.GetByID(id)
}

//Delete deletes the user with the given ID
func (ms *MySQLStore) Delete(id int64) error {
	del := "delete from Students where ID = ?"
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

//LogSignIn logs successful sign ins
func (ms *MySQLStore) LogSignIn(student *Student, time time.Time, clientIP string) error {
	ins := "insert into SignInLog(StudentID, Time, clientIP) values(?, ?, ?)"
	_, err := ms.Database.Exec(ins, student.ID, time, clientIP)
	return err
}
