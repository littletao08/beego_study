package entities
import (
	"time"
	"strings"
	"github.com/astaxie/beego"
)
/**
 *文章
 */

type Article struct {
	Id         int64
	UserId     int64
	Title      string
	Tag        string
	CategoryId int32
	Content    string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}


func (a Article) Tags() []string {
	if len(a.Tag) == 0 {
		return []string{}
	}
	tag := strings.TrimSpace(a.Tag);
	tags := strings.Split(tag, ",")

	beego.Error("******tag*************", len(tags))
	return tags
}