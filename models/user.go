package models

//User contains information about snippets stored
type User struct {
	Name     string
	Password string
	Endpoint string
	Slack    *Slack
	Email    *Email
}

type Slack struct {
	Token string
}

type Email struct {
	Server      string
	Port        string
	Address     string
	Password    string
	SenderName  string
	SenderEmail string
	Enabled     bool
}

//NewUser creates new user
func NewUser(name, password string) *User {
	return &User{
		Name:     name,
		Password: password,
		Slack:    &Slack{},
		Email:    &Email{},
	}
}

// NewEmailSettings creates new email
func NewEmailSettings(server, port, address, password, sendername, senderemail string) *Email {
	return &Email{
		Server:      server,
		Port:        port,
		Address:     address,
		Password:    password,
		SenderName:  sendername,
		SenderEmail: senderemail,
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
