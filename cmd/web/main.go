package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/senoiilya/mmdist/pkg/models"
	"gorm.io/gorm"
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

// type User struct {
// 	Name     string
// 	Age      uint
// 	Birthday time.Time
// }

func main() {
	// data base ORM
	//dsn := "host=localhost port=5432 dbname=test_db user=postgres password="
	//db, er := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	//if er != nil {
	//	log.Fatal(er)
	//}
	//user := User{Name: "Ilia", Age: 18, Birthday: time.Now()}
	//result := db.Create(&user)
	//fmt.Println(result.RowsAffected)

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

	// Загрузка .env файла и создание подключения к базе данных
	errEnv := godotenv.Load()
	if errEnv != nil {
		errorLog.Fatal("Error loading .env file")
	}

	DBconfig := models.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	}

	// Отладочный вывод
	// fmt.Println(DBconfig)

	// Инициализация базы данных
	models.InitDB(DBconfig)

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

	// SQL and ORM

	models.DB.Transaction(func(tx *gorm.DB) error {
		// return любой ошибки приведёт к откату(Rollback`у)
		if err := tx.Create(&models.Vendor{Name: "Asus"}).Error; err != nil {
			return err
		}
		if err := tx.Create(&models.Vendor{Name: "Acer"}).Error; err != nil {
			return err
		}
		if err := tx.Create(&models.Vendor{Name: "MSI"}).Error; err != nil {
			return err
		}

		// nil, если ошибок нет и транзакция успешно отработала
		return nil
	})

	models.DB.Transaction(func(tx *gorm.DB) error {
		// return любой ошибки приведёт к откату(Rollback`у)
		if err := tx.Delete(&models.Vendor{Name: "Asus"}).Error; err != nil {
			return err
		}
		if err := tx.Delete(&models.Vendor{Name: "Acer"}).Error; err != nil {
			return err
		}
		if err := tx.Delete(&models.Vendor{Name: "MSI"}).Error; err != nil {
			return err
		}

		// nil, если ошибок нет и транзакция успешно отработала
		return nil
	})

	models.DB.Transaction(func(tx *gorm.DB) error {
		// return любой ошибки приведёт к откату(Rollback`у)
		if err := tx.Delete(&models.Category{}, "3").Error; err != nil {
			return err
		}
		if err := tx.Delete(&models.Category{}, "2").Error; err != nil {
			return err
		}
		if err := tx.Delete(&models.Category{}, "1").Error; err != nil {
			return err
		}

		// nil, если ошибок нет и транзакция успешно отработала
		return nil
	})

	models.DB.Transaction(func(tx *gorm.DB) error {
		// return любой ошибки приведёт к откату(Rollback`у)
		if err := tx.Model(models.Category{}).Where("name = ?", "Laptop").Updates(models.Category{Name: "New Laptop"}).Error; err != nil {
			return err
		}
		// UPDATE categories SET name='New Laptop' WHERE name = 'Laptop';

		if err := tx.Model(models.User{}).Where("role = ?", "admin").Updates(models.User{Name: "Ilia"}).Error; err != nil {
			return err
		}
		// UPDATE users SET name='Ilia' WHERE role = 'admin';

		if err := tx.Save(&models.Category{Name: "New PC"}).Error; err != nil { // Update
			return err
		}
		var user models.User
		tx.First(&user) // Read

		user.Name = "Ilia"
		user.Role = "admin"

		if err := tx.Save(&user).Error; err != nil { // Update
			return err
		}

		// nil, если ошибок нет и транзакция успешно отработала
		return nil
	})

	models.DB.Transaction(func(tx *gorm.DB) error {
		// return любой ошибки приведёт к откату(Rollback`у)

		var (
			user     models.User
			product  models.Product
			category models.Category
		)
		tx.First(&product) // Read
		// SELECT * FROM products ORDER BY id LIMIT 1;

		tx.First(&user, "10") // Read
		// SELECT * FROM users WHERE id = 10;

		tx.Where("name = ?", "New PC").First(&category)
		// SELECT * FROM users WHERE name = 'Ilia' ORDER BY id LIMIT 1;

		// nil, если ошибок нет и транзакция успешно отработала
		return nil
	})

	models.DB.Transaction(func(tx *gorm.DB) error {
		// return любой ошибки приведёт к откату(Rollback`у)
		// READ
		var product models.Product
		tx.Raw("SELECT id, name FROM products WHERE name = ?", "Ноутбук ASUS TUF Gaming A15").Scan(&product)

		// UPDATE
		var users []models.User
		tx.Raw("UPDATE users SET name = ? RETURNING id, name", "Ilia").Scan(&users)

		// DELETE
		tx.Raw("DELETE FROM users WHERE id = 2").Scan(&users)

		// CREATE
		var categories []models.Category
		tx.Raw("INSERT INTO categories ('name') VALUES ('Smartphone')").Scan(&categories)

		// nil, если ошибок нет и транзакция успешно отработала
		return nil
	})

	infoLog.Printf("Сервер запущен на localhost%s", flagCfg.Addr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}
