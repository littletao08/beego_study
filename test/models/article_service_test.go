package models_test
import (
	"testing"
	"time"
	"beego_study/models"
	"beego_study/entities"
	"fmt"
	_"beego_study/initials"
)



func TestSave(t *testing.T) {
	article := entities.Article{UserId:1, Title:"title", Content:"content", CreatedAt:time.Now()}
	err := models.Save(&article)
	fmt.Println("err:", err)
}
