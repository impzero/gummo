package main

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Email    string   `json:"email,omitempty" db:"email"`
	Password string   `json:"-" db:"password"`
	Name     string   `json:"name,omitempty" db:"name"`
	Avatar   string   `json:"avatar,omitempty" db:"avatar"`
	ID       string   `json:"id,omitempty" db:"id"`
	Location Location `json:"location,omitempty" db:"-"`
}

func NewUser(email, password, name, avatar string) (*User, error) {
	hashed, err := hashPassword(password)
	if err != nil {
		return nil, err
	}

	return &User{
		Email:    email,
		Password: hashed,
		Name:     name,
		Avatar:   avatar,
		ID:       uuid.New().String(),
	}, nil
}

func (u *User) CompareHashAndPassword(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)) == nil
}

func hashPassword(pwd string) (string, error) {
	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(pwd), 10)
	if err != nil {
		return "", err
	}

	return string(hashedPwd), nil
}
