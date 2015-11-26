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
}
