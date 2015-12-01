package utils_test
import (
"testing"
	"regexp"
"github.com/astaxie/beego"
)

func TestIsMobileDevice(t *testing.T)  {

	var userAgent="Mozilla/5.0 (iPhone; CPU iPhone OS 7_0_4 like Mac OS X) AppleWebKit/537.51.1 (KHTML, like Gecko) Version/7.0 Mobile/11B554a Safari/9537.53"
	//userAgent="Mozilla/5.0 (Linux; Android 4.3; zh-cn; SAMSUNG-GT-I9308_TD/1.0 Android/4.3 Release/11.15.2013 Browser/AppleWebKit534.30 Build/JSS15J) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30";
	patten := ";\\s?(\\S*?\\s?\\S*?)\\s?(Build)?/"
	reg ,_ := regexp.Compile(patten)
	var matchData = reg.Find([]byte(userAgent))
	userAgent = string(matchData[:len(matchData)])

	beego.Debug("**********user-agent:%s************",userAgent )
}
