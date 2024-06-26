package main

import (
	"net/http"
	"path/filepath"

	"github.com/gorilla/mux"
	"github.com/senoiilya/mmdist/middlewares"
)

// Настраиваемая файловая система, не позволяет пользователю открывать папки в static на сайте

type neuteredFileSystem struct {
	fs http.FileSystem
}

func (nfs neuteredFileSystem) Open(path string) (http.File, error) {
	f, err := nfs.fs.Open(path)
	if err != nil {
		return nil, err
	}

	s, _ := f.Stat()
	if s.IsDir() {
		index := filepath.Join(path, "index.html")
		if _, err := nfs.fs.Open(index); err != nil {
			closeErr := f.Close()
			if closeErr != nil {
				return nil, closeErr
			}

			return nil, err
		}
	}

	return f, nil
}

func (app *application) routes() *mux.Router {
	// маршрутизатор HTTP запросов
	router := mux.NewRouter()

	// Использование шаблонов для создания динамических html страниц
	router.HandleFunc("/", app.home)
	// router.HandleFunc("/login", app.login).Methods("POST")
	router.HandleFunc("/products", app.products)
	router.HandleFunc("/registration", app.registration)
	router.HandleFunc("/profile", app.userProfile)
	router.HandleFunc("/cart", app.cart)
	router.HandleFunc("/postLogin", app.postLogin)

	router.HandleFunc("/login", app.Login).Methods("POST")
	router.HandleFunc("/signup", app.Signup).Methods("POST")
	router.Handle("/home", middlewares.IsAuthorized(http.HandlerFunc(app.Home))).Methods("GET")
	router.Handle("/premium", middlewares.IsAuthorized(http.HandlerFunc(app.Premium))).Methods("GET")
	router.HandleFunc("/logout", app.Logout).Methods("GET")

	fileServer := http.FileServer(neuteredFileSystem{http.Dir("./ui/static/")})
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fileServer))

	return router
}
