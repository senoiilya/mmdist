package main

import (
	"github.com/gorilla/mux"
	"log"
	"mmdist/pkg"
	"net/http"
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

// статичный файл
//func mainHandler(w http.ResponseWriter, req *http.Request) {
//	http.ServeFile(w, req, "ui/html/layout.html")
//}

func main() {
	//data := ViewData{
	//	Title:   "New Page2",
	//	Message: "Hello World!",
	//}
	router := mux.NewRouter()

	// Использование шаблонов для создания динамических html страниц
	router.HandleFunc("/", home)

	router.HandleFunc("/test", test)
	router.HandleFunc("/new", newTest)
	//router.HandleFunc("/edit", func(w http.ResponseWriter, req *http.Request) {
	//	tmpl, _ := template.ParseFiles("./ui/html/template.html")
	//	tmpl.Execute(w, data)
	//})
	router.HandleFunc("/login", login)

	// Вывод классов
	//for _, typeName := range types {
	//	computer := pkg.New(typeName)
	//	if computer == nil {
	//		continue
	//	}
	//	computer.PrintDetails()
	//}

	log.Println("Сервер запущен на localhost:4000")
	err := http.ListenAndServe(":4000", router)
	log.Fatal(err)
}
