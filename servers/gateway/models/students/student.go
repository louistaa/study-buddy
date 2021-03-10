package students

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"net/mail"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

//gravatarBasePhotoURL is the base URL for Gravatar image requests.
//See https://id.gravatar.com/site/implement/images/ for details
const gravatarBasePhotoURL = "https://www.gravatar.com/avatar/"

//bcryptCost is the default bcrypt cost to use when hashing passwords
var bcryptCost = 13

//Student represents a student account in the database
type Student struct {
	ID        int64  `json:"id"`
	Email     string `json:"-"` //never JSON encoded/decoded
	PassHash  []byte `json:"-"` //never JSON encoded/decoded
	UserName  string `json:"userName"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	PhotoURL  string `json:"photoURL"`
}

//Credentials represents student sign-in credentials
type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

//NewStudent represents a new student signing up for an account
type NewStudent struct {
	Email        string `json:"email"`
	Password     string `json:"password"`
	PasswordConf string `json:"passwordConf"`
	UserName     string `json:"userName"`
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
}

//Updates represents allowed updates to a student profile
type Updates struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

//Validate validates the new student and returns an error if
//any of the validation rules fail, or nil if its valid
func (ns *NewStudent) Validate() error {
	if _, emailErr := mail.ParseAddress(ns.Email); emailErr != nil {
		return emailErr
	}

	if len(ns.Password) < 6 {
		return fmt.Errorf("Password must be at least 6 characters")
	}

	if ns.Password != ns.PasswordConf {
		return fmt.Errorf("Password and confirmation must match")
	}

	if len(ns.UserName) == 0 || strings.Contains(ns.UserName, " ") {
		return fmt.Errorf("UserName must be non-zero length and may not contain spaces")
	}

	return nil
}

//ToStudent converts the NewStudent to a Student, setting the
//PhotoURL and PassHash fields appropriately
func (ns *NewStudent) ToStudent() (*Student, error) {
	validationErr := ns.Validate()
	if validationErr != nil {
		return nil, validationErr
	}

	newStudent := &Student{
		Email:     ns.Email,
		UserName:  ns.UserName,
		FirstName: ns.FirstName,
		LastName:  ns.LastName,
	}

	passwordHashErr := newStudent.SetPassword(ns.Password)
	if passwordHashErr != nil {
		return nil, passwordHashErr
	}

	GetGravitar(newStudent, ns.Email)
	return newStudent, nil
}

// GetGravitar calculates the gravitar hash based on the string given and
// stores it for the student
func GetGravitar(student *Student, str string) {
	photoURLHash := md5.Sum([]byte(strings.ToLower(strings.TrimSpace(str))))
	photoURLHashString := hex.EncodeToString(photoURLHash[:])
	student.PhotoURL = gravatarBasePhotoURL + photoURLHashString
}

//FullName returns the student's full name, in the form:
// "<FirstName> <LastName>"
//If either first or last name is an empty string, no
//space is put between the names. If both are missing,
//this returns an empty string
func (s *Student) FullName() string {
	totalStringArr := []string{}
	if s.FirstName != "" {
		totalStringArr = append(totalStringArr, s.FirstName)
	}

	if s.LastName != "" {
		totalStringArr = append(totalStringArr, s.LastName)
	}

	return strings.Join(totalStringArr, " ")
}

//SetPassword hashes the password and stores it in the PassHash field
func (s *Student) SetPassword(password string) error {
	passHash, err := bcrypt.GenerateFromPassword([]byte(password), bcryptCost)
	if err != nil {
		return err
	}
	s.PassHash = passHash
	return nil
}

//Authenticate compares the plaintext password against the stored hash
//and returns an error if they don't match, or nil if they do
func (s *Student) Authenticate(password string) error {
	err := bcrypt.CompareHashAndPassword(s.PassHash, []byte(password))
	if err != nil {
		return err
	}
	return nil
}

//ApplyUpdates applies the updates to the user. An error
//is returned if the updates are invalid
func (s *Student) ApplyUpdates(updates *Updates) error {
	// Sure hope there isn't a catch to this function. I don't think
	// it said to modify Updates in any way, and if it doesn't change then
	// this is valid.
	if updates.FirstName != "" {
		s.FirstName = updates.FirstName
	}

	if updates.LastName != "" {
		s.LastName = updates.LastName
	}

	return nil
}
