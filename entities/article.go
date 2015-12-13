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
	Tags         string
	Categories   string
	Content      string
	ViewCount    int
	LikeCount    int
	CommentCount int
	HasLike      bool `orm:"-"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}


func (a Article) SliceTags() []string {
	if len(a.Tags) == 0 {
		return []string{}
	}
	tag := strings.TrimSpace(a.Tags);
	tags := strings.Split(tag, ",")

	return tags
}

func (a Article) SliceCategories() []string {
	if len(a.Categories) == 0 {
		return []string{}
	}
	category := strings.TrimSpace(a.Categories);
	categories := strings.Split(category, ",")

	return categories
}

