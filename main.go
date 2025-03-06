package main

import (
	"net/http"
	"school1/backend"
)

func main() {
	// Обработка JavaScript файлов
	http.HandleFunc("/scripts/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/javascript")
		http.StripPrefix("/scripts/", http.FileServer(http.Dir("frontend/scripts"))).ServeHTTP(w, r)
	})

	// Обработка CSS файлов
	http.HandleFunc("/styles/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/css")
		http.StripPrefix("/styles/", http.FileServer(http.Dir("frontend/styles"))).ServeHTTP(w, r)
	})
	// Обработка png файлов
	http.HandleFunc("/photo/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/png")
		http.StripPrefix("/photo/", http.FileServer(http.Dir("assets/photo"))).ServeHTTP(w, r)
	})
	// Обработка svg файлов
	http.HandleFunc("/icon/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/svg+xml")
		http.StripPrefix("/icon/", http.FileServer(http.Dir("assets/icon"))).ServeHTTP(w, r)
	})

	// Обработчики страниц
	http.HandleFunc("/", backend.HomePage)
	http.HandleFunc("/predlogin/", backend.Predlogin)
	http.HandleFunc("/login/", backend.Login)

	http.ListenAndServe(":9090", nil)
}
