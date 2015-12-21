package entities
import (
	"time"
	"strings"
	"github.com/gogather/com"
	"html"
	"github.com/PuerkitoBio/goquery"
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

func (a Article) ShortContent(subLength int) string {
	var content = html.UnescapeString(a.Content)
	if len(content) == 0 {
		return ""
	}
	reader := strings.NewReader(content)
	doc, _ := goquery.NewDocumentFromReader(reader)
	text := doc.Text()
	text = strings.Replace(text, " ", "", -1)
	text = strings.Replace(text, "\n", "", -1)
	text = strings.Replace(text, "#", "", -1)
	text = strings.Replace(text, "*", "", -1)
	subText := com.SubString(text, 0, 160)
	return subText
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


func (a Article) SetCategories(categories string) {
	if len(categories) > 0 {
		categories = strings.ToLower(categories)
		a.Categories = categories
	}
}


func (a Article) SetTags(tags string) {
	if len(tags) > 0 {
		tags = strings.ToLower(tags)
		a.Tags = tags
	}
}



