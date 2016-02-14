package models
import (
	_"beego_study/initials"
	"testing"
	"beego_study/entities"
	"beego_study/models"
	"beego_study/db"
)

func TestSaveCategories(t *testing.T) {
	categories := entities.NewCategories(2, []string{"go1", "go2"})
	models.BatchSaveOrUpdateCategory(db.NewDB(),categories)
}
