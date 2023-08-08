package models

import (
	"html"
	"strings"

	"github.com/ayowilfred95/Golang-RESTful-API/database"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// Notice that the JSON binding for the Password field is -.
// This ensures that the user’s password is not returned in the JSON response.

type User struct {
	gorm.Model
	Username string `gorm:"size:255;not null;unique" json:"username"`
	Password string  `gorm:"size:255;not null;" json:"-"`
	Loggings []Logs
}


func(user *User) SaveUser() (*User, error) {
	if err := user.hashPassword(); err != nil {
        return &User{}, err
    }
	
	err:= database.Database.Create(&user).Error
	if err != nil {
		return &User{}, err
	}
	return user, nil

}

func(user *User) hashPassword() error{
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password),bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(passwordHash)
	user.Username = html.EscapeString(strings.TrimSpace(user.Username))
	return nil
}


// login
func(user *User) ValidatePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
}

func FindUserByUsername(username string) (User, error) {
	var user User
	err:= database.Database.Where("username=?", username).Find(&user).Error
	if err != nil{
		return User{}, err
	}
	return user, nil
}
