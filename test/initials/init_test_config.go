package initials

import
(
	"github.com/astaxie/beego"
)

func init() {
	beego.AppConfigPath = "/Users/zhanglida/go_path/src/beego_study/conf/app.conf"
	beego.ParseConfig()
	beego.Error("apppath", beego.AppConfigPath)
}



