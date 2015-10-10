package utils
import (
	"github.com/gogather/com/log"
)

func Regain(err interface{}) {
	if r := recover(); r != nil {
		log.Redf("%v", err)
	}
}
