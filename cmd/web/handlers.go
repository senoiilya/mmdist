package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
)

// Controllers

func home(w http.ResponseWriter, req *http.Request) {
	data := ViewData{
		Title:   "Layout Page",
		Message: "Здесь находится рут страница",
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

func test(w http.ResponseWriter, req *http.Request) {
	dataArray := ViewData2{
		Title:   "List of users",
		Message: []string{"Bob", "Sam", "Tom"},
	}
	tmpl, err := template.ParseFiles("./ui/html/template.html")
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}
	err = tmpl.Execute(w, dataArray)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}

func login(w http.ResponseWriter, r *http.Request) {
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

	tmpl, err := template.ParseFiles("./ui/html/login.html")
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}

func newTest(w http.ResponseWriter, req *http.Request) {
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
	tmpl, err := template.ParseFiles("./ui/html/new-page.html")
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}
