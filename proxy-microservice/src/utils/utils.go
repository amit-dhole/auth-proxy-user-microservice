package utils

import (
	"net/http"
)

//GenerateResponse : Generate APIResponse Instance
func GenerateResponse(data []byte, code int, w http.ResponseWriter) {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(code)
	_, errJSON := w.Write(data)
	if errJSON != nil {
		return
	}
	return
}

//ToString converts an interface type to string
func ToString(v interface{}) string {
	t, ok := v.(string)
	if ok {
		return t
	}
	return ""
}
