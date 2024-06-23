package main

import (
	"github.com/gorilla/mux"
	"github.com/senoiilya/mmdist/pkg"
	"log"
	"net/http"
	"time"
)

var types = []string{pkg.PersonalComputerType, pkg.NotebookType, pkg.ServerType, "mono-block"}

type ViewData struct {
	Title   string
	Message string
}

type ViewData2 struct {
	Title   string
	Message []string
}

type ViewData3 struct {
	Title   string
	Message string
	Id      int
}

type ViewLayout struct {
	Title    string
	Message  string
	Computer string
}

// статичный файл
//func mainHandler(w http.ResponseWriter, req *http.Request) {
//	http.ServeFile(w, req, "ui/html/layout.html")
//}

func main() {
	router := mux.NewRouter()

	// Использование шаблонов для создания динамических html страниц
	router.HandleFunc("/", home)
	router.HandleFunc("/login", login)
	router.HandleFunc("/products", products)
	router.HandleFunc("/registration", registration)
	router.HandleFunc("/userpage", userPage)

	// Вывод классов
	//for _, typeName := range types {
	//	computer := pkg.New(typeName)
	//	if computer == nil {
	//		continue
	//	}
	//	computer.PrintDetails()
	//}

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fileServer))

	srv := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:4000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("Сервер запущен на localhost:4000")
	err := srv.ListenAndServe()
	log.Fatal(err)
}
