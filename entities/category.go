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


func NewCategories(userId int64,names []string) []Category{
	var categories []Category
	for _,name :=range names{
	  categories = append(categories,Category{UserId:userId,Name:name})
	}
	return categories
}
