package main

import (
	"github.com/senoiilya/mmdist/pkg"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

// Controllers

func home(w http.ResponseWriter, req *http.Request) {
	data := ViewLayout{
		Title:    "Домашняя страница",
		Message:  "Здесь находится рут страница",
		Computer: pkg.New(pkg.NotebookType).String(),
	}

	files := []string{"./ui/html/layout.html", "./ui/html/viewLayout.html"}

	// если страница не главная, то откроет страницу Page Not Found
	if req.URL.Path != "/" {
		http.NotFound(w, req)
		return
	}

	//tmpl, err := template.ParseFiles("./ui/html/layout.html")
	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}
	err = tmpl.ExecuteTemplate(w, "layout", data)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}
}

func login(w http.ResponseWriter, req *http.Request) {
	dataArray := ViewData2{
		Title:   "List of users",
		Message: []string{"Bob", "Sam", "Tom"},
	}

	if req.Method != "GET" {
		http.Error(w, "Method Not Allowed", 405)
		return
	}

	files := []string{"./ui/html/layout.html", "./ui/html/login.html"}

	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	err = tmpl.ExecuteTemplate(w, "layout", dataArray)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}

func registration(w http.ResponseWriter, r *http.Request) {
	data := ViewData{
		Title:   "Login",
		Message: "Войти",
	}
	files := []string{"./ui/html/layout.html", "../ui/html/login.html"}
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

	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}
	err = tmpl.ExecuteTemplate(w, "layout", data)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}

func userPage(w http.ResponseWriter, req *http.Request) {
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
	files := []string{"./ui/html/layout.html", "./ui/html/registration.html"}

	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	err = tmpl.ExecuteTemplate(w, "layout", data)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}

func products(w http.ResponseWriter, req *http.Request) {
	files := []string{"./ui/html/layout.html", "./ui/html/products.html"}

	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	data := ViewLayout{Title: "Products", Message: "fsdfsd"}

	err = tmpl.ExecuteTemplate(w, "layout", data)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}
