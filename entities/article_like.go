package entities
import "time"

type ArticleLike struct {
	Id int64
	UserId int64
	ArticleId int64
	Ip string
	Valid int
	CreatedAt time.Time
	UpdatedAt time.Time
}
