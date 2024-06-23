package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func (app *application) routes() *mux.Router {
	// маршрутизатор HTTP запросов
	router := mux.NewRouter()

	// Использование шаблонов для создания динамических html страниц
	router.HandleFunc("/", app.home)
	router.HandleFunc("/login", app.login)
	router.HandleFunc("/products", app.products)
	router.HandleFunc("/registration", app.registration)
	router.HandleFunc("/user_page", app.userPage)

	fileServer := http.FileServer(neuteredFileSystem{http.Dir("./ui/static/")})
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fileServer))

	return router
}
