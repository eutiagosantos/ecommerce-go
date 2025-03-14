package domain

import (
	"errors"
	"regexp"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `gorm:"type:varchar(100)" json:"name"`
	Password string `gorm:"type:varchar(255)" json:"password"`
	Email    string `gorm:"type:varchar(100);uniqueIndex" json:"email"`
}

func BuildUser(name string, password string, email string) (*User, error) {
	u := &User{
		Name:     name,
		Password: password,
		Email:    email,
	}
	if err := u.validate(); err != nil {
		return nil, err
	}
	return u, nil
}

func (u *User) validate() error {
	if u.Name == "" || u.Password == "" || u.Email == "" {
		return errors.New("name, password and email are required")
	}
	if len(u.Name) < 10 {
		return errors.New("name must be at least 10 characters")
	}
	if len(u.Password) < 10 {
		return errors.New("password must be at least 10 characters")
	}
	if !regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`).MatchString(u.Email) {
		return errors.New("invalid email")
	}
	return nil
}
