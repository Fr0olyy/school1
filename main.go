package main

import (
	"net/http"
	"school1/backend"
)

func main() {
	// Настраиваем обработку статических файлов
	fs := http.FileServer(http.Dir("frontend"))
	http.Handle("/styles/", fs)
	http.Handle("/scripts/", fs)
	http.Handle("/assets/", fs)

	// Обработчик главной страницы
	http.HandleFunc("/", backend.HomePage)

	http.ListenAndServe(":9090", nil)
	//Соси Борис
}
