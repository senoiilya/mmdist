package main

import (
	"encoding/json"
	"html/template"
	"net/http"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/senoiilya/mmdist/pkg"
	"github.com/senoiilya/mmdist/pkg/models"
	"github.com/senoiilya/mmdist/utils"
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

var jwtKey = []byte("secret_key_for_auth")

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

func (app *application) Login(w http.ResponseWriter, r *http.Request) {
	var user models.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var existingUser models.User
	models.DB.Where("email = ?", user.Email).First(&existingUser)

	if existingUser.ID == 0 {
		http.Error(w, "user does not exist", http.StatusBadRequest)
		return
	}

	errHash := utils.CompareHashPassword(user.Password, existingUser.Password)
	if !errHash {
		http.Error(w, "invalid password", http.StatusBadRequest)
		return
	}

	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &models.Claims{
		Role: existingUser.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   existingUser.Email,
			ExpiresAt: &jwt.NumericDate{expirationTime},
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		http.Error(w, "could not generate token", http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    tokenString,
		Expires:  expirationTime,
		Path:     "/",
		HttpOnly: true,
	})

	json.NewEncoder(w).Encode(map[string]string{"success": "user logged in"})
}

func (app *application) Signup(w http.ResponseWriter, r *http.Request) {
	var user models.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var existingUser models.User
	models.DB.Where("email = ?", user.Email).First(&existingUser) // ORM

	// проверка на существование такого пользователя с таким email
	if existingUser.ID != 0 {
		http.Error(w, "user already exists", http.StatusBadRequest)
		return
	}

	var errHash error
	user.Password, errHash = utils.GenerateHashPassword(user.Password)
	if errHash != nil {
		http.Error(w, "could not generate password hash", http.StatusInternalServerError)
		return
	}

	models.DB.Create(&user)
	json.NewEncoder(w).Encode(map[string]string{"success": "user created"})
}

func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("token")
	if err != nil {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	claims, err := utils.ParseToken(cookie.Value)
	if err != nil || (claims.Role != "user" && claims.Role != "admin") {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{"success": "home page", "role": claims.Role})
}

func (app *application) Premium(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("token")
	if err != nil {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	claims, err := utils.ParseToken(cookie.Value)
	if err != nil || claims.Role != "admin" {
		http.Error(w, "Not enough permissions", http.StatusForbidden)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{"success": "premium page", "role": claims.Role})
}

func (app *application) Logout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    "",
		MaxAge:   -1,
		Path:     "/",
		HttpOnly: true,
	})

	json.NewEncoder(w).Encode(map[string]string{"success": "user logged out"})
}

// func ResetPassword(w http.ResponseWriter, r *http.Request) {
// 	var user models.User

// 	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	var existingUser models.User
// 	models.DB.Where("email = ?", user.Email).First(&existingUser)

// 	if existingUser.ID == 0 {
// 		http.Error(w, "user does not exist", http.StatusBadRequest)
// 		return
// 	}

// 	var errHash error
// 	user.Password, errHash = utils.GenerateHashPassword(user.Password)
// 	if errHash != nil {
// 		http.Error(w, "could not generate password hash", http.StatusInternalServerError)
// 		return
// 	}

// 	models.DB.Model(&existingUser).Update("password", user.Password)
// 	json.NewEncoder(w).Encode(map[string]string{"success": "password updated"})
// }
