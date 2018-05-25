package models

//User contains information about snippets stored
type User struct {
	Name     string
	Password string
	Endpoint string
}

//NewUser creates new user
func NewUser(name, password string) *User {
	return &User{
		Name:     name,
		Password: password,
	}
}

func (u *User) BucketName() string {
	return "user"
}

func (u *User) ID() string {
	return u.Name
}

func (u *User) Value() interface{} {
	return u
}
