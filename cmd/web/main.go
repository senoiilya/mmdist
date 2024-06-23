package main

import (
	"flag"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"time"
)

//var types = []string{pkg.PersonalComputerType, pkg.NotebookType, pkg.ServerType, "mono-block"}

// Структура для флагов
type FlagsConfig struct {
	Addr string
}

// Используем патерн Dependency Injection
type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

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

type User struct {
	Name     string
	Age      uint
	Birthday time.Time
}

func main() {
	dsn := "host=localhost port=5432 dbname=test_db user=postgres password="
	db, er := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if er != nil {
		log.Fatal(er)
	}
	user := User{Name: "Ilia", Age: 18, Birthday: time.Now()}
	result := db.Create(&user)
	fmt.Println(result.RowsAffected)
	flagCfg := new(FlagsConfig)
	// флаги при билде приложения через консоль
	flag.StringVar(&flagCfg.Addr, "addr", ":4000", "Сетевой адресс HTTP")
	// парсим, написанные в консоли флаги
	flag.Parse()

	// Логгер для информационных логов
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	// Логгер для логов об ошибках
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	// Вывод классов
	//for _, typeName := range types {
	//	computer := pkg.New(typeName)
	//	if computer == nil {
	//		continue
	//	}
	//	computer.PrintDetails()
	//}

	srv := &http.Server{
		Handler:      app.routes(),
		Addr:         fmt.Sprintf("127.0.0.1%s", flagCfg.Addr),
		ErrorLog:     errorLog,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	infoLog.Printf("Сервер запущен на localhost%s", flagCfg.Addr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}
