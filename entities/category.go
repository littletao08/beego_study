package entities
import "time"
type Category struct {
	Id        int32
	UserId    int64
	Name      string
	ArticleCount int32
	Order     int32
	CreatedAt time.Time
	UpdatedAt time.Time
}
