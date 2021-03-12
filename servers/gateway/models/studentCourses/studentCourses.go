package studentCourses

//Student represents a student account in the database
type StudentCourses struct {
	ID        int64 `json:"id"`
	StudentID int64 `json:"studentID"`
	CourseID  int64 `json:"courseID"`
}
