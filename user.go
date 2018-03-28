package main

//TODO : make models

//User contains information about snippets stored
type User struct {
	Name     string
	Password string
}

//Save saves the User
func (s *User) Save() error {
	err := update(s.Name, "user", s)
	if err != nil {
		return err
	}
	return nil
}
