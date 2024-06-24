package main

import (
	"github.com/senoiilya/mmdist/pkg"
	"html/template"
	"net/http"
	"strconv"
)

type ViewData struct {
	Title   string
	Message string
	IsAuth  bool
}

type ViewData2 struct {
	Title   string
	Message []string
	IsAuth  bool
}

type ViewData3 struct {
	Title   string
	Message string
	Id      int
	IsAuth  bool
}

type ViewLayout struct {
	Title    string
	Message  string
	IsAuth   bool
	Username string
	Computer string
}

// Controllers

func (app *application) home(w http.ResponseWriter, req *http.Request) {
	data := ViewLayout{
		Title:    "Домашняя страница",
		Message:  "Здесь находится рут страница",
		IsAuth:   false,
		Username: "Илья",
		Computer: pkg.New(pkg.NotebookType).String(),
	}

	files := []string{"./ui/html/layout.html", "./ui/html/viewLayout.html", "./ui/html/loginpartial.html"}

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
		Title:   "Войти",
		IsAuth:  false,
		Message: []string{"Bob", "Sam", "Tom"},
	}
	/// тут чекнуть
	if req.Method != "GET" {
		w.Header().Set("Allow", http.MethodGet)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	files := []string{"./ui/html/layout.html", "./ui/html/login.html", "./ui/html/loginpartial.html"}

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
		Title:   "Регистрация",
		Message: "Пожалуйста, заполните форму регистрации",
		IsAuth:  false,
	}
	files := []string{"./ui/html/layout.html", "./ui/html/registration.html", "./ui/html/loginpartial.html"}

	// упраление картой заголовков напрямую
	//w.Header()["Date"] = nil

	// проверка на метод POST
	//if r.Method != http.MethodPost {
	//	// устаналвиваем какой HTTP - заголовок нам разрешён
	//	w.Header().Set("Allow", http.MethodPost)
	//	app.clientError(w, http.StatusMethodNotAllowed)
	//	return
	//}

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

func (app *application) userProfile(w http.ResponseWriter, req *http.Request) {
	// Извлекаем значение параметра id из URL и попытаемся
	// конвертировать строку в integer используя функцию strconv.Atoi(). Если его нельзя
	// конвертировать в integer, или значение меньше 1, возвращаем ответ
	// 404 - страница не найдена!
	id, err := strconv.Atoi(req.URL.Query().Get("id"))
	//if err != nil || id < 1 {
	//	app.notFound(w)
	//	return
	//}
	// позже вернуть
	data := ViewData3{
		Title:   "New Page2",
		Message: "Hello World!",
		Id:      id,
		IsAuth:  false,
	}
	files := []string{"./ui/html/layout.html", "./ui/html/userpage.html", "./ui/html/loginpartial.html"}

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
	files := []string{"./ui/html/layout.html", "./ui/html/products.html", "./ui/html/loginpartial.html"}
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

func (app *application) cart(w http.ResponseWriter, r *http.Request) {
	files := []string{"./ui/html/layout.html", "./ui/html/cart.html", "./ui/html/loginpartial.html"}

	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
		return
	}
	err = tmpl.ExecuteTemplate(w, "layout", nil)
	if err != nil {
		app.serverError(w, err)
	}
}

func (app *application) postLogin(w http.ResponseWriter, r *http.Request) {

}
