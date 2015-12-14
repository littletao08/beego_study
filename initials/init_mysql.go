package initials
import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

func InitMysql() {
	userName := beego.AppConfig.String("mysql_user_name")
	userPass := beego.AppConfig.String("mysql_user_pass")
	ip := beego.AppConfig.String("mysql_ip")
	port, err := beego.AppConfig.Int("mysql_port")
	dbName := beego.AppConfig.String("mysql_db_name")
	if nil != err {
		port = 3306
	}
	orm.Debug = true
	driver_url := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&loc=Local", userName, userPass, ip, port, dbName)
	beego.Info("driver_url:", driver_url)
	orm.RegisterDataBase("default", "mysql", driver_url)

}
