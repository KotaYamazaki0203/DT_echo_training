package usecase

import (
	"app/src/entities"
)

type Interactor struct {
	Repository Repository
}

// アプリケーション固有のビジネスルール
// このファイルでは取得したデータを組み合わせたりしてユースケースを実現する

func (i Interactor) GetAllArticle() (article []entities.Article, err error) {
	return i.Repository.GetAllArticle()
}

func (i Interactor) GetUndeletedTodos() (todo []entities.Todo, err error) {
	return i.Repository.GetUndeletedTodos()
}

func (i Interactor) GetTodo(todoId uint) (todo *entities.Todo, err error) {
	return i.Repository.GetTodo(todoId)
}

func (i Interactor) InsertNewTodo(title string, content string) error {
	return i.Repository.InsertNewTodo(title, content)
}

func (i Interactor) UpdateTodo(todoId uint, title string, content string) error {
	return i.Repository.UpdateTodo(todoId, title, content)
}

func (i Interactor) DeleteTodo(todoId uint) error {
	return i.Repository.DeleteTodo(todoId)
}
