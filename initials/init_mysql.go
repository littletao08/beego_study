package initials
import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/gogather/com/log"
)

func InitMysql() {
	userName := beego.AppConfig.String("mysql_user_name")
	userPass := beego.AppConfig.String("mysql_user_pass")
	ip := beego.AppConfig.String("mysql_ip")
	port, err := beego.AppConfig.Int("mysql_port")
	dbName := beego.AppConfig.String("mysql_db_name")

	log.Greenf("userName:%v,userPass:%v,ip:%v,port:%v,dbName:%v \n", userName, userPass, ip, port, dbName)
	if nil != err {
		port = 3306
	}
	orm.Debug = true
	driver_url := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", userName, userPass, ip, port, dbName)
	log.Greenf("driver_url:%v \n", driver_url)
	orm.RegisterDataBase("default", "mysql", driver_url)
	//orm.RegisterDataBase("default", "mysql", "root:root@/orm_test?charset=utf8")

}
