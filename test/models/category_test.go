package services
import (
	_"beego_study/initials"
	"testing"
	"beego_study/entities"
	"beego_study/services"
	"beego_study/db"
)

func TestSaveCategories(t *testing.T) {
	categories := entities.NewCategories(2, []string{"go1", "go2"})
	services.BatchSaveOrUpdateCategory(db.NewDB(),categories)
}
