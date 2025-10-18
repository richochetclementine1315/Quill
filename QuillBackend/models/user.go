package models

import "golang.org/x/crypto/bcrypt"

type User struct {
	Id        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	// Bcrypt hashed password
	// We removed the json tag to prevent it from being sent in responses
	Password []byte `json:"-"`
	Phone    string `json:"phone"`
}

func (user *User) SetPassword(password string) {
	// Hash the password using bcrypt
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	user.Password = hashedPassword

}

// CheckPassword method to verify the password during login
func (user *User) CheckPassword(password string) error {
	return bcrypt.CompareHashAndPassword(user.Password, []byte(password))
}

// Now we will migrate it to the DB
