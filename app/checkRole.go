package app

import (
	"fmt"
	set "github.com/deckarep/golang-set"
	"reflect"
	"regexp"
)

//Вхождение маршрута  проверку регистрации, добавление роли
func CheckUrlPath(val string, array map[string]interface{}, role *[]interface{}) (exists bool) {
	exists = false
	for pattern, role_from_pattern := range array {
		matched, _ := regexp.Match(pattern, []byte(val))
		if matched {
			v := reflect.ValueOf(role_from_pattern)
			switch v.Kind() {
			case reflect.Slice:
				for i := 0; i < v.Len(); i++ {
					//fmt.Printf("+++ %s %s\n", i, v.Index(i).String())
					*role = append(*role, v.Index(i).String())
				}
			case reflect.Bool:
				*role = append(*role, role_from_pattern)
			}
			exists = true
			return exists
		}
	}
	return exists
}

//Перевод массива интерфейсов в массив строк
func ArrInterfaceToArrayString(interface_arr []interface{}) (string_arr []interface{}) {
	for _, v := range interface_arr {
		string_arr = append(string_arr, fmt.Sprintf("%s", v))
	}
	return
}

// Проверка роли пользователя
func CheckRole(role_token []interface{}, role_check_in_server []interface{}) (exist bool) {
	if role_check_in_server[0] == false {
		exist = false
		return
	} else {
		if len(set.NewSetFromSlice(ArrInterfaceToArrayString(role_token)).Intersect(set.NewSetFromSlice(ArrInterfaceToArrayString(role_check_in_server))).ToSlice()) != 0 {
			//fmt.Printf("Пересекаются\n")
			exist = false
			return
		}
		exist = true
		return
	}
}
