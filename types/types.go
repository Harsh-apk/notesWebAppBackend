package types

import (
	"regexp"

	"github.com/Harsh-apk/notesWebApp/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	bcryptCost      = 12
	MinFirstNameLen = 2
	MinLastNameLen  = 2
	MinPasswordLen  = 7
)

// USER ->
type IncomingUser struct {
	FirstName string `json:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty"`
	Email     string `json:"email,omitempty"`
	Password  string `json:"password,omitempty"`
}

type User struct {
	ID                primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	FirstName         string             `bson:"firstName,omitempty" json:"firstName,omitempty"`
	LastName          string             `bson:"lastName,omitempty" json:"lastName,omitempty"`
	Email             string             `bson:"email,omitempty" json:"email,omitempty"`
	EncryptedPassword string             `bson:"encryptedPassword,omitempty" json:"encryptedPassword,omitempty"`
}
type IncomingLoginUser struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

func (user *IncomingUser) ValidateUser() *map[string]string {
	errors := map[string]string{}
	if len(user.FirstName) < MinFirstNameLen {
		errors["firstNameError"] = "first name is too short"
	}
	if len(user.LastName) < MinLastNameLen {
		errors["lastNameError"] = "last name is too short"
	}
	if len(user.Password) < MinPasswordLen {
		errors["passwordError"] = "inadequate password"
	}
	if !isValidEmail(&user.Email) {
		errors["invalidEmail"] = "invalid email address"
	}
	return &errors

}

func isValidEmail(email *string) bool {
	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	return re.MatchString(*email)

}

func UserFromIncomingUser(data *IncomingUser) (*User, error) {
	encpw, err := utils.EncryptPassword(&data.Password)
	if err != nil {
		return nil, err
	}
	return &User{
		FirstName:         data.FirstName,
		LastName:          data.LastName,
		Email:             data.Email,
		EncryptedPassword: *encpw,
	}, nil

}

//NOTES ->

type IncomingNote struct {
	UserID   string `json:"userID,omitempty"`
	Title    string `json:"title,omitempty"`
	Content  string `json:"content,omitempty"`
	Category string `json:"category,omitempty"`
}
type Note struct {
	NoteID   primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	UserID   primitive.ObjectID `bson:"userId,omitempty" json:"userId,omitempty"`
	Title    string             `bson:"title,omitempty" json:"title,omitempty"`
	Content  string             `bson:"content,omitempty" json:"content,omitempty"`
	Category string             `bson:"category,omitempty" json:"category,omitempty"`
}

func NoteFromIncomingNote(data *IncomingNote) (*Note, error) {
	oid, err := primitive.ObjectIDFromHex(data.UserID)
	if err != nil {
		return nil, err
	}
	return &Note{
		UserID:   oid,
		Title:    data.Title,
		Content:  data.Content,
		Category: data.Category,
	}, nil
}
