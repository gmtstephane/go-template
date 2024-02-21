// package models holds repository models declarations
package models

type User struct {
	ID       int
	Username string
	Picture  string
}

func (u User) IsValid() bool {
	return u.ID != 0
}
