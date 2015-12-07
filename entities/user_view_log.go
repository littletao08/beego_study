package entities
import "time"

type UserViewLog struct {
	Id int64
	UserId int64
	ArticleId int64
	Ip string
	CreatedAt time.Time
	UpdatedAt time.Time
}
