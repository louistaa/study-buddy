package studentCourses

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
func (ms *MySQLStore) getByProvidedType(t GetByType, arg interface{}) (*StudentCourses, error) {
	sel := string("select ID, StudentID, CourseID from StudentCourse where " + t + " = ?")

	rows, err := ms.Database.Query(sel, arg)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	studentCourse := &StudentCourses{}

	// Should never have more than one row, so only grab one
	rows.Next()
	if err := rows.Scan(
		&studentCourse.ID,
		&studentCourse.StudentID,
		&studentCourse.CourseID); err != nil {
		return nil, err
	}
	return studentCourse, nil
}

//GetByID returns the User with the given ID
func (ms *MySQLStore) GetByID(id int64) (*StudentCourses, error) {
	return ms.getByProvidedType(ID, id)
}

//Insert inserts the user into the database, and returns
//the newly-inserted User, complete with the DBMS-assigned ID
func (ms *MySQLStore) Insert(studentCourse *StudentCourses) (*StudentCourses, error) {
	ins := "insert into StudentCourse(StudentID, CourseID) values (?,?)"
	res, err := ms.Database.Exec(ins, studentCourse.StudentID, studentCourse.CourseID)
	if err != nil {
		return nil, err
	}

	lid, lidErr := res.LastInsertId()
	if lidErr != nil {
		return nil, lidErr
	}

	studentCourse.ID = lid

	return studentCourse, nil
}

//Delete deletes the user with the given ID
func (ms *MySQLStore) Delete(id int64) error {
	del := "delete from StudentCourse where ID = ?"
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

//GetByStudentID returns the student and correlating classes with the given student ID
func (ms *MySQLStore) GetByStudentID(studentID int64) ([]int64, error) {
	sel := string("select CourseID from StudentCourse where StudentID = ?")

	rows, err := ms.Database.Query(sel, studentID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var courses []int64

	for rows.Next() {
		var courseID int64
		err = rows.Scan(&courseID)
		if err != nil {
			return nil, err
		}
		courses = append(courses, courseID)
	}

	return courses, nil
}

//GetByClassID returns the class and correlating students with the given class ID
func (ms *MySQLStore) GetByClassID(classID int64) ([]int64, error) {
	sel := string("select StudentID from StudentCourse where CourseID = ?")

	rows, err := ms.Database.Query(sel, classID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var students []int64

	for rows.Next() {
		var studentID int64
		err = rows.Scan(&studentID)
		if err != nil {
			return nil, err
		}
		students = append(students, studentID)
	}

	return students, nil
}
