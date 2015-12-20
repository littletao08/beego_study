package utils
import (
	"time"
	"github.com/astaxie/beego"
	"strconv"
)

func ZhTime(t time.Time) string {
	nowTimeSecond := time.Now().Second()
	createTimeSecond := t.Second()
	beego.Error("nowTimeSecond:",nowTimeSecond,"createTimeSecond:",createTimeSecond)
	diff := nowTimeSecond - createTimeSecond;
	zhTime := "刚刚";
	if diff > 0 {
		if diff < 60 {
			zhTime = strconv.Itoa(diff) + "秒前";
		}else if diff >= 60 && diff < 60 * 60 {
			zhTime = strconv.Itoa(diff / 60) + "分钟前";
		}else if diff >= 60 * 60 && diff < 60 * 60 * 24 {
			zhTime = strconv.Itoa((diff / (60 * 60))) + "小时前";
		}else if diff > 60 * 60 * 24 && diff < 60 * 60 * 24 * 365 {
			zhTime = strconv.Itoa((diff / (60 * 60 * 24))) + "天前";
		}else {
			zhTime = "1年前";
		}
	}
	return zhTime;
}
