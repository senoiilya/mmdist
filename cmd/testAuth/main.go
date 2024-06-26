package main

import (
	"encoding/json"
	"html/template"
	"net/http"
	"sync"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var (
	jwtKey            = []byte("my_secret_key")
	loginTemplate     *template.Template
	protectedTemplate *template.Template
	registerTemplate  *template.Template
	users             = make(map[string]string) // In-memory storage for user credentials
	usersMutex        sync.Mutex                // Mutex to protect access to the users map
)

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func main() {
	// Загрузка шаблонов
	loginTemplate = template.Must(template.ParseFiles("./templates/login.html"))
	protectedTemplate = template.Must(template.ParseFiles("./templates/protected.html"))
	registerTemplate = template.Must(template.ParseFiles("./templates/register.html"))

	http.HandleFunc("/login", ShowLoginPage)
	http.HandleFunc("/login-action", Login)
	http.HandleFunc("/register", ShowRegisterPage)
	http.HandleFunc("/register-action", Register)
	http.HandleFunc("/protected", Authenticate(ShowProtectedPage))
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./ui/static/"))))

	http.ListenAndServe(":8080", nil)
}

func ShowLoginPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		loginTemplate.Execute(w, nil)
	}
}

func ShowRegisterPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		registerTemplate.Execute(w, nil)
	}
}

func Register(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var creds Credentials
		err := json.NewDecoder(r.Body).Decode(&creds)
		if err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		usersMutex.Lock()
		if _, exists := users[creds.Username]; exists {
			usersMutex.Unlock()
			http.Error(w, "User already exists", http.StatusConflict)
			return
		}

		users[creds.Username] = creds.Password
		usersMutex.Unlock()

		w.WriteHeader(http.StatusCreated)
	}
}

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var creds Credentials
		err := json.NewDecoder(r.Body).Decode(&creds)
		if err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		usersMutex.Lock()
		storedPassword, exists := users[creds.Username]
		usersMutex.Unlock()
		if !exists || storedPassword != creds.Password {
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
			return
		}

		expirationTime := time.Now().Add(5 * time.Minute)
		claims := &Claims{
			Username: creds.Username,
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: &jwt.NumericDate{expirationTime},
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err := token.SignedString(jwtKey)
		if err != nil {
			http.Error(w, "Could not create token", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
	}
}

func Authenticate(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}

		claims := &Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}

		r.Header.Set("username", claims.Username)
		next(w, r)
	}
}

func ShowProtectedPage(w http.ResponseWriter, r *http.Request) {
	username := r.Header.Get("username")
	protectedTemplate.Execute(w, map[string]string{"username": username})
}
