package initials
import (
	"github.com/astaxie/beego/orm"
	"beego_study/entities"
)

func InitEntity()  {
	orm.RegisterModel(new(entities.Article))
	orm.RegisterModel(new(entities.Category))
	orm.RegisterModel(new(entities.Parameter))
	orm.RegisterModel(new(entities.User))
	orm.RegisterModel(new(entities.ArticleLike))
	orm.RegisterModel(new(entities.ArticleView))
	orm.RegisterModel(new(entities.Tag))
	orm.RegisterModel(new(entities.OpenUser))
}
