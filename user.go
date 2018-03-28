package main

//TODO : make models

//User contains information about snippets stored
type User struct {
	Name     string
	Password string
}

//NewUser creates new user
func NewUser(name, password string) *User {
	return &User{
		Name:     name,
		Password: password,
	}
}

//Save saves the User
func (s *User) Save() error {
	err := update(s.Name, "user", s)
	if err != nil {
		return err
	}
	return nil
}
