package entities
import "time"

type Tag struct {
	Id        int32
	Name      string
	ArticleCount int32
	CreatedAt time.Time
	UpdatedAt time.Time
}
