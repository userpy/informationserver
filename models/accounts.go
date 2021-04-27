package models

import (
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	u "informationserver/utils"
	"os"
	"strings"
)

/*
Структура прав доступа JWT
*/
type Token struct {
	UserId uint
	Email  string
	User   string
	Role   []interface{}
	jwt.StandardClaims
}

//структура для учётной записи пользователя
type Account struct {
	gorm.Model
	Email    string `json:"email"`
	Password string `json:"password"`
	Token    string `json:"token";sql:"-"`
	User     string `json:"user"`
}

type Role struct {
	UserId   uint
	RoleName string
}

//Получить роли
func (account *Account) GetRoles() []interface{} {
	responseRole := []Role{}
	err := GetDB().Table("roles").Where("user_id = ?", account.ID).Find(&responseRole).Error
	if err != nil {
		return nil
	}
	var responseRoleArr []interface{}
	for _, val := range responseRole {
		responseRoleArr = append(responseRoleArr, val.RoleName)
	}
	return responseRoleArr
}

//Проверить входящие данные пользователя ...
func (account *Account) Validate() (map[string]interface{}, bool) {

	if !strings.Contains(account.Email, "@") {
		return u.Message(false, "Email address is required"), false
	}

	if len(account.Password) < 6 {
		return u.Message(false, "Password is required"), false
	}
	if len(account.User) != 0 {
		return u.Message(false, "Password is required"), false
	}
	//Email должен быть уникальным
	temp := &Account{}

	//проверка на наличие ошибок и дубликатов электронных писем
	err := GetDB().Table("accounts").Where("email = ?", account.Email).First(temp).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return u.Message(false, "Connection error. Please retry"), false
	}
	//Проверка e-mail
	if temp.Email != "" {
		return u.Message(false, "Email address already in use by another user."), false
	}
	err = GetDB().Table("accounts").Where("user = ?", account.User).First(temp).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return u.Message(false, "Connection error. Please retry"), false
	}
	//Проверка пользователя
	if temp.User != "" {
		return u.Message(false, "User address already in use by another user."), false
	}
	return u.Message(false, "Requirement passed"), true
}

//Создание пользователя
func (account *Account) Create() map[string]interface{} {
	if resp, ok := account.Validate(); !ok {
		return resp
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(account.Password), bcrypt.DefaultCost)
	account.Password = string(hashedPassword)

	GetDB().Create(account)
	GetDB().Create(&Role{UserId: account.ID, RoleName: "user"})
	if account.ID <= 0 {
		return u.Message(false, "Failed to create account, connection error.")
	}
	//Создать новый токен JWT для новой зарегистрированной учётной записи
	tk := &Token{UserId: account.ID, Email: account.Email, User: account.User, Role: account.GetRoles()}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	account.Token = tokenString

	account.Password = "" //удалить пароль

	response := u.Message(true, "Account has been created")
	response["account"] = account
	return response
}

//Вход пользователя
func Login(email, password string) map[string]interface{} {

	account := &Account{}
	err := GetDB().Table("accounts").Where("email = ?", email).First(account).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return u.Message(false, "Email address not found")
		}
		return u.Message(false, "Connection error. Please retry")
	}

	err = bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword { //Пароль не совпадает!!
		return u.Message(false, "Invalid login credentials. Please try again")
	}
	//Работает! Войти в систему
	account.Password = ""
	//Создать новый токен JWT для новой зарегистрированной учётной записи
	tk := &Token{UserId: account.ID, Email: account.Email, User: account.User, Role: account.GetRoles()}
	//Создать токен JWT
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	account.Token = tokenString // Сохраните токен в ответе

	resp := u.Message(true, "Logged In")
	resp["account"] = account
	return resp
}

func GetUser(u uint) *Account {

	acc := &Account{}
	GetDB().Table("accounts").Where("id = ?", u).First(acc)
	if acc.Email == "" { //Пользователь не найден!
		return nil
	}

	acc.Password = ""
	return acc
}
