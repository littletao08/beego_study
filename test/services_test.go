package test
import (
	"fmt"
	"beego_study/models"
	"testing"
	_ "beego_study/initials"
)

func TestGetUsers(t *testing.T) {

	fmt.Println("*************************")
	var users = models.GetUser(1);
	fmt.Println("users", users[1])
	for _, user := range users {

		fmt.Println(user)
	}
}
