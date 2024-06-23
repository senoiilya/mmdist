package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"path/filepath"
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
	router.HandleFunc("/login", app.login)
	router.HandleFunc("/products", app.products)
	router.HandleFunc("/registration", app.registration)
	router.HandleFunc("/user_page", app.userPage)

	fileServer := http.FileServer(neuteredFileSystem{http.Dir("./ui/static/")})
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fileServer))

	return router
}
