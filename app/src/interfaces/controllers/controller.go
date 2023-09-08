package controllers

import (
	"app/src/infrastructure/sqlhandler"
	"app/src/usecase"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"strconv"
)

type Controller struct {
	Interactor usecase.Interactor
}

/*
このファイルには外部からのリクエストで受け取ったデータをusecaseで使えるように変形したり、
内部からのデータを外部機能に向けて便利な形式に変換したりする
例)　外部からのデータをArticleエンティティに変換
*/

func NewController(sqlhandler *sqlhandler.SqlHandler) *Controller {
	return &Controller{
		Interactor: usecase.Interactor{
			Repository: usecase.Repository{
				DB: sqlhandler.DB,
			},
		},
	}
}

func (c Controller) Index(ctx echo.Context) error {
	articles, err := c.Interactor.GetAllArticle()
	if err != nil {
		log.Print(err)
		return ctx.Render(500, "article_list.html", nil)
	}
	return ctx.Render(http.StatusOK, "article_list.html", articles)
}

func (c Controller) AllTodos(ctx echo.Context, sqlhandler *sqlhandler.SqlHandler) error {
	todos, err := c.Interactor.GetUndeletedTodos(sqlhandler)
	if err != nil {
		log.Print(err)
		return ctx.Render(http.StatusInternalServerError, "500.html", nil)
	}

	return ctx.Render(http.StatusOK, "all_todos.html", todos)
}

func (c Controller) Detail(ctx echo.Context, sqlhandler *sqlhandler.SqlHandler) error {
	todoId, strconvErr := c.convertTodoIdToUint(ctx.QueryParam("todo_id"))
	if strconvErr != nil {
		log.Print(strconvErr)
		return ctx.Render(http.StatusInternalServerError, "500.html", nil)
	}

	todo, err := c.Interactor.GetTodo(sqlhandler, todoId)
	if err != nil {
		log.Print(err)
		return ctx.Render(http.StatusInternalServerError, "500.html", nil)
	}

	return ctx.Render(http.StatusOK, "detail.html", todo)
}

func (c Controller) NewTodo(ctx echo.Context) error {
	return ctx.Render(http.StatusOK, "new_todo.html", nil)
}

func (c Controller) NewTodoSubmit(ctx echo.Context, sqlhandler *sqlhandler.SqlHandler) error {
	title := ctx.FormValue("title")
	content := ctx.FormValue("content")
	c.Interactor.InsertNewTodo(sqlhandler, title, content)
	return ctx.Redirect(http.StatusFound, "/all_todos")
}

func (c Controller) EditTodo(ctx echo.Context, sqlhandler *sqlhandler.SqlHandler) error {
	todoId, strconvErr := c.convertTodoIdToUint(ctx.QueryParam("todo_id"))
	if strconvErr != nil {
		log.Print(strconvErr)
		return ctx.Render(http.StatusInternalServerError, "500.html", nil)
	}

	todo, err := c.Interactor.GetTodo(sqlhandler, todoId)
	if err != nil {
		log.Print(err)
		return ctx.Render(http.StatusInternalServerError, "500.html", nil)
	}

	return ctx.Render(http.StatusFound, "edit.html", todo)
}

func (c Controller) EditTodoSubmit(ctx echo.Context, sqlhandler *sqlhandler.SqlHandler) error {
	todoId, strconvErr := c.convertTodoIdToUint(ctx.QueryParam("todo_id"))

	if strconvErr != nil {
		log.Print(strconvErr)
		return ctx.Render(http.StatusInternalServerError, "500.html", nil)
	}

	title := ctx.FormValue("title")
	content := ctx.FormValue("content")
	c.Interactor.UpdateTodo(sqlhandler, todoId, title, content)
	return ctx.Redirect(http.StatusFound, "/all_todos")
}

func (c Controller) DeleteTodo(ctx echo.Context, sqlhandler *sqlhandler.SqlHandler) error {
	todoId, strconvErr := c.convertTodoIdToUint(ctx.QueryParam("todo_id"))

	if strconvErr != nil {
		log.Print(strconvErr)
		return ctx.Render(http.StatusInternalServerError, "500.html", nil)
	}

	c.Interactor.DeleteTodo(sqlhandler, todoId)
	return ctx.Redirect(http.StatusFound, "/all_todos")
}

func (c Controller) convertTodoIdToUint(todoId string) (uint, error) {
	id, strconvErr := strconv.ParseUint(todoId, 10, 64)
	return uint(id), strconvErr
}
