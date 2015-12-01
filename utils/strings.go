package utils
import (
	"strings"
)


var mobileUserAgentMap = map[string]bool{
	"nokia":true,
	"samsung":true,
	"midp-2":true,
	"cldc1.1":true,
	"symbianos":true,
	"maui":true,
	"untrusted/1.0":true,
	"windows ce":true,
	"iphone":true,
	"ipad":true,
	"android":true,
	"blackberry":true,
	"brew":true,
	"j2me":true,
	"yulong":true,
	"coolpad":true,
	"tianyu":true,
	"ty-":true,
	"k-touch":true,
	"haier":true,
	"dopod":true,
	"lenovo":true,
	"huaqin":true,
	"aigo-":true,
	"ctc/1.0":true,
	"ctc/2.0":true,
	"cmcc":true,
	"daxian":true,
	"mot-":true,
	"sonyericsson":true,
	"gionee":true,
	"htc":true,
	"zte":true,
	"huawei":true,
	"webos":true,
	"gobrowser":true,
	"iemobile":true,
	"wap2.0":true}
func IsMobileDevice(userAgent string) bool  {
	userAgent = strings.ToLower(userAgent)
	return mobileUserAgentMap[userAgent]
}
