package app

import (
	"context"
	"fmt"
	set "github.com/deckarep/golang-set"
	jwt "github.com/dgrijalva/jwt-go"
	"informationserver/models"
	u "informationserver/utils"
	"net/http"
	"os"
	"reflect"
	"regexp"
	"strings"
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
					fmt.Printf("+++ %s %s\n", i, v.Index(i).String())
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
			fmt.Printf("Пересекаются\n")
			exist = false
			return
		}
		exist = true
		return
	}
}

var JwtAuthentication = func(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//Список эндпоинтов, для которых  требуется авторизация, с проверкой роли или без
		Auth := map[string]interface{}{
			"/api/contacts/new":       []string{"admin", "user"},
			"/api/user/\\d*/contacts": []string{"admin"}}

		requestPath := r.URL.Path //текущий путь запроса
		//проверяем, не требует ли запрос аутентификации, обслуживаем запрос, если он не нужен
		var role []interface{}
		if CheckUrlPath(requestPath, Auth, &role) {
			response := make(map[string]interface{})
			tokenHeader := r.Header.Get("Authorization") //Получение токена
			if tokenHeader == "" {                       //Токен отсутствует, возвращаем  403 http-код Unauthorized
				response = u.Message(false, "Missing auth token")
				w.WriteHeader(http.StatusForbidden)
				w.Header().Add("Content-Type", "application/json")
				u.Respond(w, response)
				return
			}
			splitted := strings.Split(tokenHeader, " ")
			//Токен обычно поставляется в формате `Bearer {token-body}`,
			// мы проверяем, соответствует ли полученный токен этому требованию
			if len(splitted) != 2 {
				response = u.Message(false, "Invalid/Malformed auth token")
				w.WriteHeader(http.StatusForbidden)
				w.Header().Add("Content-Type", "application/json")
				u.Respond(w, response)
				return
			}

			tokenPart := splitted[1] //Получаем вторую часть токена
			tk := &models.Token{}
			//fmt.Printf("- %s\n", tokenPart)
			token, err := jwt.ParseWithClaims(tokenPart, tk, func(token *jwt.Token) (interface{}, error) {
				return []byte(os.Getenv("token_password")), nil
			})
			//Неправильный токен, как правило, возвращает 403 http-код
			if err != nil {
				response = u.Message(false, "Malformed authentication token")
				w.WriteHeader(http.StatusForbidden)
				w.Header().Add("Content-Type", "application/json")
				u.Respond(w, response)
				return
			}
			//токен недействителен, возможно, не подписан на этом сервере
			if !token.Valid {
				response = u.Message(false, "Token is not valid.")
				w.WriteHeader(http.StatusForbidden)
				w.Header().Add("Content-Type", "application/json")
				u.Respond(w, response)
				return
			}
			//Всё прошло хорошо, продолжаем выполнение запроса
			//Полезно для мониторинга
			if CheckRole(tk.Role, role) {
				response = u.Message(false, "Role is not valid.")
				w.WriteHeader(http.StatusForbidden)
				w.Header().Add("Content-Type", "application/json")
				u.Respond(w, response)
				return
			}
			ctx := context.WithValue(r.Context(), "user", tk.UserId)
			r = r.WithContext(ctx)
			//fmt.Printf(">>>>>>\n")
			next.ServeHTTP(w, r) //передать управление следующему обработчику!

		} else {
			fmt.Printf("-")
			next.ServeHTTP(w, r)
			return
		}
	})
}
