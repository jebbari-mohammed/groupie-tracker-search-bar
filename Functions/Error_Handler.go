package functions

import (
	"html/template"
	"net/http"
)

var (
	Message string
	tmpl    = template.Must(template.ParseFiles("Template/Error.html"))
)

func ErrorHandler(w http.ResponseWriter, r *http.Request, code int) {
	switch code {
	case 400:
		Message = "Bad Request"
	case 403:
		Message = "Forbidden"
	case 404:
		Message = "NoT Found"
	case 405:
		Message = "Method Not Allowed "
	case 500:
		Message = "Internal Server Error"
	default:
		Message = "Error"
	}
	w.WriteHeader(code)

	data := struct {
		Message string
		Code    int
	}{
		Message: Message,
		Code:    code,
	}

	tmpl.Execute(w, data)
}
