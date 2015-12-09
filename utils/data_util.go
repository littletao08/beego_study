package utils
import (
	"reflect"
	"fmt"
	"errors"
	"strconv"
	"strings"
)

func ToSlice(container interface{}) []interface{} {
	val := reflect.ValueOf(container)
	sInd := reflect.Indirect(val)
	if sInd.Kind() != reflect.Slice {
		panic(fmt.Errorf("container must be slice "))
	}
	l := sInd.Len()
	ret := make([]interface{}, l)
	for i := 0; i < l; i++ {
		ret[i] = sInd.Index(i).Interface()
	}
	return ret
}

func ExtractFieldValues(source interface{}, field string) ([]interface{}, error) {

	val := reflect.ValueOf(source)
	sInd := reflect.Indirect(val)
	if (sInd.Kind() != reflect.Slice) {
		return make([]interface{}, 0), errors.New("source must be slice type")
	}

	defer Regain("extract field values failure !")

	var slice = ToSlice(source)
	var filedValues = make([]interface{}, 0)
	for _, obj := range slice {
		var objVal = reflect.ValueOf(obj)
		var val = objVal.FieldByName(field).Interface()
		filedValues = append(filedValues, val)
	}

	return filedValues, nil

}


func SliceToString(a interface{}, sep string) string {
	var strSlice [] string
	var slice = ToSlice(a)
	for _, v := range slice {
		switch v.(type) {
		case string:
			strSlice = append(strSlice, v.(string))
			continue
		case int,int8,int16,int32,int64:
			strV := strconv.FormatInt(reflect.ValueOf(v).Int(),10)
			strSlice = append(strSlice, strV)
			continue
		default:
			panic("params type not supported")
		}
	}
	return strings.Join(strSlice, ",")
}
