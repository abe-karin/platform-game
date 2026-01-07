package model

import "errors"

type User struct {
	ID       int
	Username string
	Active bool
}

func (u User) Validate() error {
	if u.ID <= 0 {
		return errors.New("invalid user ID")
	}
	if u.Username == "" {
		return errors.New("username cannot be empty")
	}
	return nil
}
func (u User) IsValid() bool {
	if u.ID <= 0 {
		return false
	}
	if u.Username == "" {
		return false
	}
	return true
}