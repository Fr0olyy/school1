package backend

import (
	"html/template"
	"net/http"
)

func Predregister(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("frontend/predregister.html")
	if err != nil {
		http.Error(w, "Ошибка при загрузке шаблона.", http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
