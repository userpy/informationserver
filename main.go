package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"informationserver/app"
	"informationserver/controllers"
	"net/http"
	"os"
)

func main() {
	//https://tproger.ru/translations/deploy-a-secure-golang-rest-api/
	router := mux.NewRouter()
	router.Use(app.JwtAuthentication) // добавляем middleware проверки JWT-токена
	router.HandleFunc("/api/user/new", controllers.CreateAccount).Methods("POST")
	router.HandleFunc("/api/user/login", controllers.Authenticate).Methods("POST")
	router.HandleFunc("/api/contacts/new", controllers.CreateContact).Methods("GET")
	router.HandleFunc("/api/user/{id}/contacts", controllers.GetContactsFor).Methods("GET")
	port := os.Getenv("PORT") //Получить порт из файла .env; мы не указали порт, поэтому при локальном тестировании должна возвращаться пустая строка
	if port == "" {
		port = "8000" //localhost
	}

	fmt.Println(port)

	err := http.ListenAndServe(":"+port, router) //Запустите приложение, посетите localhost:8000/api

	if err != nil {
		fmt.Print(err)
	}
}
