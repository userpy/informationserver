package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
)

func Message(status bool, message string) map[string]interface{} {
	return map[string]interface{}{"status": status, "message": message}
}

func Respond(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func GetContextValueUint(r *http.Request, val string) (response uint64, err error) {
	defer func() {
		if r := recover(); r != nil {
			response = 0
			err = fmt.Errorf("Context valuet not found")
			return
		}
	}()
	return reflect.ValueOf(r.Context().Value(val)).Uint(), nil
}

func GetContextValueInt(r *http.Request, val string) (response int64, err error) {
	defer func() {
		if r := recover(); r != nil {
			response = 0
			err = fmt.Errorf("Context valuet not found")
			return
		}
	}()
	return reflect.ValueOf(r.Context().Value(val)).Int(), nil
}
