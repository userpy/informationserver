package app

import (
	"context"
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"informationserver/models"
	u "informationserver/utils"
	"net/http"
	"os"
	"strings"
)

//Проверка JWT -токена
var JwtAuthentication = func(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//Список эндпоинтов, для которых  требуется авторизация, с проверкой роли или без
		Auth := map[string]interface{}{
			"/api/contacts/new":       []string{"admin", "user"},
			"/api/user/\\d*/contacts": []string{"admin"},
			"/api/user/contacts":      []string{"admin"}}

		requestPath := r.URL.Path //текущий путь запроса
		//проверяем, не требует ли запрос аутентификации, обслуживаем запрос, если он  нужен
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
			//Записываем значение в контекст
			fmt.Printf(">*** %s\n", tk.UserId)
			r = r.WithContext(context.WithValue(r.Context(), "user", tk.UserId))
			r = r.WithContext(context.WithValue(r.Context(), "role", tk.Role))
			next.ServeHTTP(w, r) //передать управление следующему обработчику!

		} else {
			//fmt.Printf("-")
			next.ServeHTTP(w, r)
			return
		}
	})
}
