package routing

import (
	"app/src/infrastructure/sqlhandler"
	"app/src/interfaces/controllers"
	"github.com/labstack/echo/v4"
	"net/http"
)

// このファイルにはリクエストのルーティング処理を実装する

func SetRouting(e *echo.Echo) {
	controller := controllers.NewController(sqlhandler.NewSqlHandler())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Echo World!!")
	})

	e.GET("/allArticles", func(c echo.Context) error {
		return controller.Index(c)
	})

	e.GET("/detail", func(c echo.Context) error {
		return controller.Detail(c)
	})

	e.GET("/all_todos", func(c echo.Context) error {
		return controller.AllTodos(c)
	})

	e.GET("/new_todo", func(c echo.Context) error {
		return controller.NewTodo(c)
	})

	e.POST("/new_todo/submit", func(c echo.Context) error {
		return controller.NewTodoSubmit(c)
	})

	e.GET("/edit", func(c echo.Context) error {
		return controller.EditTodo(c)
	})

	e.POST("/edit/submit", func(c echo.Context) error {
		return controller.EditTodoSubmit(c)
	})

	e.GET("/delete_todo", func(c echo.Context) error {
		return controller.DeleteTodo(c)
	})
}
