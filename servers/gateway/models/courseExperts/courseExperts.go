package courseExpert

//Student represents a student account in the database
type CourseExpert struct {
	ID       int64 `json:"id"`
	CourseID int64 `json:"courseID"`
	ExpertID int64 `json:"expertID"`
}
