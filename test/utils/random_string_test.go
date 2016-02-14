package utils

import (
	"testing"
	"fmt"
	"beego_study/utils"
)

func TestCaptchaCreate(t *testing.T) {
	for i := 0; i < 10; i ++ {
		fmt.Println(utils.RandomIntCaptcha(6))
	}
}

