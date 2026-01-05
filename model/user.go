package model
type User struct {
	ID       int
	Username string
	Active bool
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