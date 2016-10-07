package models

import (
	"beego_study/utils"
	"time"
	"reflect"
)

func AsJSON(models []IModel) ([]map[string]interface{}) {
	result := []map[string]interface{}{}
	if utils.IsEmpty(models) {
		return result
	}

	for _, val := range models {
		result = append(result, val.AsJSON())
	}

	return result

}

func Rewrite(a interface{}) interface{} {
	if utils.IsMap(a) {
		return rewriteMap(a)
	}else if utils.IsSlice(a) {
		return rewriteSlice(a)
	} else if val, ok := a.(IModel); ok {
		return val.AsJSON()
	} else if val, ok := a.(time.Time); ok {
		return rewriteDate(val)
	}

	return a
}

func rewriteMap(a interface{}) (interface{}) {
	var result = make(map[interface{}]interface{})
	val := reflect.ValueOf(a)
	sInd := reflect.Indirect(val)
	for _, k := range sInd.MapKeys() {
		v := sInd.MapIndex(k)
		result[k.Interface()] = Rewrite(v.Interface())
	}
	return result
}

func rewriteSlice(a interface{}) interface{} {
	result := []interface{}{}
	values := utils.ToSlice(a)

	for _, v := range values {
		result = append(result, Rewrite(v))
	}
	return nil
}

func rewriteDate(a interface{}) interface{} {
	v, ok := a.(time.Time)

	if ok {
		t := utils.DateTimeFormat(v)
		return t
	}

	return a
}

func Slice(i interface{}, properties ... string)  (map[string]interface{}) {
	result := make(map[string]interface{})
	if utils.IsEmpty(i) {
		return result
	}
	val := reflect.ValueOf(i)
	sInd := reflect.Indirect(val)
	for _, k := range properties {
		field := sInd.FieldByName(k)
		result[k] = field.Interface()
	}

	return result
}

/*public static Map<String, Object> slice(Object obj, String... properties) {
Map<String, Object> result = Maps.newHashMap();
if (obj == null) {
return result;
}

BeanWrapper wrapper = PropertyAccessorFactory
.forBeanPropertyAccess(obj);
for (String property : properties) {
Object value = wrapper.getPropertyValue(property);
result.put(property, value);
}
return result;
}*/
