package main

import (
	"net/http"
	"school1/backend"
	"strings"
)

func main() {
	// Настраиваем обработку CSS файлов с явным указанием MIME-типа
	http.HandleFunc("/styles/", func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, ".css") {
			w.Header().Set("Content-Type", "text/css")
		}
		fs := http.StripPrefix("/styles/", http.FileServer(http.Dir("frontend/styles")))
		fs.ServeHTTP(w, r)
	})

	// Настраиваем обработку остальных статических файлов
	fs := http.FileServer(http.Dir("frontend"))
	http.Handle("/scripts/", http.StripPrefix("/scripts/", fs))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	// Обработчики страниц
	http.HandleFunc("/", backend.HomePage)
	http.HandleFunc("/predlogin/", backend.Predlogin)
	http.HandleFunc("/login/", backend.Login)

	http.ListenAndServe(":9090", nil)
}
