package entities
type User struct {
	Id int64
	QcOpenId string
	Name string
	Nick string
	Password string
	Age int32
	Sex int32
	Cell string
	Mail string
	Qq string
	CreatedAt string
	UpdatedAt string
}

func (c *User) Valid(name string,password string) bool  {
	if nil == c{
		return false
	}

	if (len(c.Name) == 0 || len(c.Password) == 0){
		return false
	}

	if (c.Name != name || c.Password != password){
		return false
	}
	return true
}