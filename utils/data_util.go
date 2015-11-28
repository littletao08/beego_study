package utils
import (
"reflect"
"fmt"
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
