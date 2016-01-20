package entities

import "time"

type User struct {
	Id        int64
	Name      string
	Nick      string
	Head      string
	Password  string
	Age       int
	Sex       int
	Cell      string
	Mail      string
	Qq        string
	Province  string
	City      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (c *User) Valid(name string, password string) bool {
	if nil == c {
		return false
	}

	if (len(c.Name) == 0 || len(c.Password) == 0) {
		return false
	}

	if (c.Name != name || c.Password != password) {
		return false
	}
	return true
}

func (u User) InitFromOpenUser(openUser *OpenUser) (User){
	if nil != openUser {
		u.Age = openUser.Age
		u.Nick = openUser.Nick
		u.Sex = openUser.Sex
		u.City = openUser.City
		u.Province = openUser.Province
		u.Head = openUser.Head
	}

	return u
}