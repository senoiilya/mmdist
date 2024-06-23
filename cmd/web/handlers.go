package main

import (
	"github.com/senoiilya/mmdist/pkg"
	"html/template"
	"net/http"
	"strconv"
)

// Controllers

func (app *application) home(w http.ResponseWriter, req *http.Request) {
	data := ViewLayout{
		Title:    "Домашняя страница",
		Message:  "Здесь находится рут страница",
		Computer: pkg.New(pkg.NotebookType).String(),
	}

	files := []string{"./ui/html/layout.html", "./ui/html/viewLayout.html"}

	// если страница не главная, то откроет страницу Page Not Found
	if req.URL.Path != "/" {
		app.notFound(w)
		return
	}

	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
		return
	}
	err = tmpl.ExecuteTemplate(w, "layout", data)
	if err != nil {
		app.serverError(w, err)
		return
	}
}

func (app *application) login(w http.ResponseWriter, req *http.Request) {
	dataArray := ViewData2{
		Title:   "List of users",
		Message: []string{"Bob", "Sam", "Tom"},
	}
	/// тут чекнуть
	if req.Method != "GET" {
		w.Header().Set("Allow", http.MethodGet)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	files := []string{"./ui/html/layout.html", "./ui/html/login.html"}

	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
		return
	}

	err = tmpl.ExecuteTemplate(w, "layout", dataArray)
	if err != nil {
		app.serverError(w, err)
		return
	}
}

func (app *application) registration(w http.ResponseWriter, r *http.Request) {
	data := ViewData{
		Title:   "Login",
		Message: "Войти",
	}
	files := []string{"./ui/html/layout.html", "../ui/html/login.html"}

	// упраление картой заголовков напрямую
	//w.Header()["Date"] = nil

	// проверка на метод POST
	if r.Method != http.MethodPost {
		// устаналвиваем какой HTTP - заголовок нам разрешён
		w.Header().Set("Allow", http.MethodPost)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
		return
	}
	err = tmpl.ExecuteTemplate(w, "layout", data)
	if err != nil {
		app.serverError(w, err)
		return
	}
}

func (app *application) userPage(w http.ResponseWriter, req *http.Request) {
	// Извлекаем значение параметра id из URL и попытаемся
	// конвертировать строку в integer используя функцию strconv.Atoi(). Если его нельзя
	// конвертировать в integer, или значение меньше 1, возвращаем ответ
	// 404 - страница не найдена!
	id, err := strconv.Atoi(req.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
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
		app.serverError(w, err)
		return
	}

	err = tmpl.ExecuteTemplate(w, "layout", data)
	if err != nil {
		app.serverError(w, err)
		return
	}
}

func (app *application) products(w http.ResponseWriter, req *http.Request) {
	files := []string{"./ui/html/layout.html", "./ui/html/products.html"}
	data := ViewLayout{Title: "Products", Message: "fsdfsd"}

	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
		return
	}

	err = tmpl.ExecuteTemplate(w, "layout", data)
	if err != nil {
		app.serverError(w, err)
		return
	}
}
