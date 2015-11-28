package db_test
import (
	"testing"
	_"beego_study/entities"
	_"beego_study/routers"
	"beego_study/db"
	"beego_study/entities"
	"fmt"
	"reflect"
	"beego_study/utils"
)

func TestPagination(t *testing.T) {
	fmt.Println("1111")
	var container []entities.User
	pagination, _ := db.NewDB().From("user").Select("id", "name").Pagination(&container, 1, 10)
	for _, value := range pagination.Data {
		fmt.Println("id", value.(entities.User).Id, "name", value.(entities.User).Name)
	}
}

func testMakeSliceByType(t *testing.T) {
	//fmt.Println(types.Struct{})
}

/*func printType(t Type) {
	fmt.Println(reflect.TypeOf(t))
}*/


func testDataUtil(t *testing.T) {
	var container []entities.User
	var user1 = entities.User{Id:1, Name:"name1"}
	var user2 = entities.User{Id:2, Name:"name2"}
	container = append(container, user1)
	container = append(container, user2)

	newContainer := utils.ToSlice(&container)
	for _, rowData := range newContainer {
		fmt.Println(rowData.(entities.User).Id)
	}
}


func testReflect(t *testing.T) {
	var container []entities.User
	var user1 = entities.User{Id:1, Name:"name1"}
	var user2 = entities.User{Id:2, Name:"name2"}
	container = append(container, user1)
	container = append(container, user2)

	parse(&container)
}

func parse(container interface{}) {

	val := reflect.ValueOf(container)
	sInd := reflect.Indirect(val)
	if val.Kind() != reflect.Ptr || sInd.Kind() != reflect.Slice {
		panic(fmt.Errorf("<RawSeter.QueryRows> all args must be use ptr slice"))
	}

	etyp := sInd.Type().Elem()
	typ := etyp
	if typ.Kind() == reflect.Ptr {
		fmt.Println("1111111111")
		typ = typ.Elem()
	}
	//"val", val, "sInd", sInd, "etyp", etyp, "typ", typ
	fmt.Println("2222222222222")
	fmt.Println("val", val)
	fmt.Println("sInd", sInd)
	fmt.Println("etyp", etyp)
	fmt.Println("typ", typ)
	fmt.Println("------------------------")
	if typ.Kind() == reflect.Struct {
		fmt.Println("fiels", "\\\\\\\\\\")
	}

	fmt.Println("------------------------")
	/*nInd := val.Elem()
	fmt.Println("item", nInd.Slice(0,1))
	fmt.Println("len", len(nInd))*/

	val1 := reflect.New(typ)
	ind := val1.Elem()

	fmt.Println("val1", val1)
	fmt.Println("ind", ind)

	fmt.Println("t type", sInd.Kind())
	fmt.Println("------------------------")


	switch sInd.Kind() {
	case reflect.Slice:
		for i := 0; i < sInd.Len(); i++ {
			fmt.Println(sInd.Index(i))
		}
	}

	//reflect.ValueOf(refs[cur]).Elem().Interface()


	/*val := reflect.New(typ)
	ind := val.Elem()
*/


}
