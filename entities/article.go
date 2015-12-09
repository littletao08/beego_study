package entities
import (
	"time"
	"strings"
)
/**
 *文章
 */

type Article struct {
	Id           int64
	UserId       int64
	Title        string
	Tag          string
	Categories   []*Category `orm:"rel(m2m)"`
	Content      string
	ViewCount    int
	LikeCount    int
	CommentCount int
	HasLike      bool
	CreatedAt    time.Time
	UpdatedAt    time.Time
}


func (a Article) Tags() []string {
	if len(a.Tag) == 0 {
		return []string{}
	}
	tag := strings.TrimSpace(a.Tag);
	tags := strings.Split(tag, ",")

	return tags
}

