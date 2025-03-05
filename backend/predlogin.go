package backend

import (
	"html/template"
	"net/http"
)

func Predlogin(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("frontend/predlogin.html")
	if err != nil {
		http.Error(w, "Ошибка при загрузке шаблона.", http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
