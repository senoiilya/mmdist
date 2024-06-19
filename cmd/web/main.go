package main

import (
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"mmdist/pkg"
	"net/http"
	"strconv"
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

	// упраление картой заголовков напрямую
	//w.Header()["Date"] = nil

	// проверка на метод POST
	if r.Method != http.MethodPost {
		// Не продвинутый подход к отправке ответа
		//w.WriteHeader(405) // Записываем заголовок
		//w.Write([]byte("GET-метод запрещён!"))
		//return
		http.Error(w, "Метод запрещён!", 405)
		return
	}

	tmpl, _ := template.ParseFiles("./ui/html/login.html")
	tmpl.Execute(w, data)
}

func newTestHandler(w http.ResponseWriter, req *http.Request) {
	// Извлекаем значение параметра id из URL и попытаемся
	// конвертировать строку в integer используя функцию strconv.Atoi(). Если его нельзя
	// конвертировать в integer, или значение меньше 1, возвращаем ответ
	// 404 - страница не найдена!
	id, err := strconv.Atoi(req.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, req)
		return
	}
	data := ViewData3{
		Title:   "New Page2",
		Message: "Hello World!",
		Id:      id,
	}
	tmpl, _ := template.ParseFiles("./ui/html/new-page.html")
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
	router.HandleFunc("/new", newTestHandler)
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
