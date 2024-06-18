package main

import (
	"github.com/gorilla/mux"
	"html/template"
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

// статичный файл
func mainHandler(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "ui/html/layout.html")
}

// Controllers

func layoutHandler(w http.ResponseWriter, req *http.Request) {
	data := ViewData{
		Title:   "Layout Page",
		Message: "Здесь находится рут страница",
	}
	// если страница не главная, то откроет страницу Page Not Found
	if req.URL.Path != "/" {
		http.NotFound(w, req)
		return
	}
	tmpl, _ := template.ParseFiles("ui/html/layout.html")
	tmpl.Execute(w, data)
}

func testHandler(w http.ResponseWriter, req *http.Request) {
	dataArray := ViewData2{
		Title:   "List of users",
		Message: []string{"Bob", "Sam", "Tom"},
	}
	tmpl, _ := template.ParseFiles("./ui/html/template.html")
	tmpl.Execute(w, dataArray)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	data := ViewData{
		Title:   "Login",
		Message: "Войти",
	}
	// устаналвиваем какой HTTP - заголовок нам разрешён
	w.Header().Set("Allow", http.MethodPost)

	// проверка на метод POST
	if r.Method != http.MethodPost {
		w.WriteHeader(405)
		w.Write([]byte("GET-метод запрещён!"))
		return
	}

	tmpl, _ := template.ParseFiles("./ui/html/login.html")
	tmpl.Execute(w, data)
}

func main() {
	data := ViewData{
		Title:   "New Page2",
		Message: "Hello World!",
	}
	//router := http.NewServeMux()
	router := mux.NewRouter()
	// Использование шаблонов для создания динамических html страниц
	router.HandleFunc("/", layoutHandler)
	router.HandleFunc("/test", testHandler)
	router.HandleFunc("/new", func(w http.ResponseWriter, req *http.Request) {
		tmpl, _ := template.ParseFiles("./ui/html/new-page.html")
		tmpl.Execute(w, data)
	})
	router.HandleFunc("/edit", func(w http.ResponseWriter, req *http.Request) {
		tmpl, _ := template.ParseFiles("./ui/html/template.html")
		tmpl.Execute(w, data)
	})
	router.HandleFunc("/login", loginHandler)

	// Вывод классов
	//for _, typeName := range types {
	//	computer := pkg.New(typeName)
	//	if computer == nil {
	//		continue
	//	}
	//	computer.PrintDetails()
	//}

	log.Println("Сервер запущен на localhost:8080")
	err := http.ListenAndServe(":4000", router)
	log.Fatal(err)
}
