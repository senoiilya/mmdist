package main

import (
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

//var types = []string{pkg.PersonalComputerType, pkg.NotebookType, pkg.ServerType, "mono-block"}

type FlagsConfig struct {
	Addr      string
	StaticDir string
}

// Используем патерн Dependency Injection
type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

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

type ViewLayout struct {
	Title    string
	Message  string
	Computer string
}

// Настраиваемая файловая система, не позволяет пользователю открывать папки в static на сайте

type neuteredFileSystem struct {
	fs http.FileSystem
}

func (nfs neuteredFileSystem) Open(path string) (http.File, error) {
	f, err := nfs.fs.Open(path)
	if err != nil {
		return nil, err
	}

	s, _ := f.Stat()
	if s.IsDir() {
		index := filepath.Join(path, "index.html")
		if _, err := nfs.fs.Open(index); err != nil {
			closeErr := f.Close()
			if closeErr != nil {
				return nil, closeErr
			}

			return nil, err
		}
	}

	return f, nil
}

func main() {
	flagCfg := new(FlagsConfig)
	// флаги при билде приложения через консоль
	flag.StringVar(&flagCfg.Addr, "addr", ":4000", "Сетевой адресс HTTP")
	flag.StringVar(&flagCfg.StaticDir, "static-dir", "./ui/static/", "Каталог, в котором будут храниться статичные файлы.\nПо умолчанию используется:")
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

	// маршрутизатор HTTP запросов
	router := mux.NewRouter()

	// Использование шаблонов для создания динамических html страниц
	router.HandleFunc("/", app.home)
	router.HandleFunc("/login", app.login)
	router.HandleFunc("/products", app.products)
	router.HandleFunc("/registration", app.registration)
	router.HandleFunc("/user_page", app.userPage)

	// Вывод классов
	//for _, typeName := range types {
	//	computer := pkg.New(typeName)
	//	if computer == nil {
	//		continue
	//	}
	//	computer.PrintDetails()
	//}

	fileServer := http.FileServer(neuteredFileSystem{http.Dir(flagCfg.StaticDir)})
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fileServer))

	srv := &http.Server{
		Handler:      router,
		Addr:         fmt.Sprintf("127.0.0.1%s", flagCfg.Addr),
		ErrorLog:     errorLog,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	infoLog.Printf("Сервер запущен на localhost%s", flagCfg.Addr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}
