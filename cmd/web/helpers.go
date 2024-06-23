package main

import (
	"fmt"
	"net/http"
	"runtime/debug"
)

// Помощник serverError записывает сообщение об ошибке в errorLog и
// затем отправляет пользователю ответ 500 "Внутренняя ошибка сервера".
func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errorLog.Println(trace)

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

// debug.Stack() позволяет увидеть трассировку стека для текущей горутины и добавить её в логгер

// Помощник clientError отправляет определённый код состояния и соответствущее описание
// пользователю, когда есть проблема с пользовательским запросом. Например, 400 "Bad Request"
func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
	// http.StatusText(status) возвращает понятный ответ, например http.StatusText(400) вернет строку "Bad Request"
}

// Помощник notFound это удобная оболочка вокруг clientError,
// которая отправляет пользователю ответ "404 Страница не найдена".
func (app *application) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}
