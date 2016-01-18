package entities

const (
	OPEN_USER_TYPE_QQ=1
	OPEN_USER_TYPE_SINA=2
)
type OpenUser struct {
	Id int64
	OpenId string
	UserId int64
	Type int
	Nick string
	Head string
	Sex  int
	Age int
	Province string
    City string
	AccessToken string `orm:"-"`

}

func (c OpenUser) HasBindUser() bool {
	return c.UserId > 0
}
