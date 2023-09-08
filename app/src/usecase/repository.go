package usecase

import (
	"app/src/entities"
	"app/src/model"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

// このファイルではDBからのデータ取得やDBへのinsertなど、DB操作を記述する

func (r *Repository) GetAllArticle() (articles []entities.Article, err error) {
	// 以下は実際にはDBを使って記事の全データを取得したりする
	var article entities.Article
	article.ID = 1
	article.Title = "Deep Track"
	articles = append(articles, article)
	return articles, nil
}

func (r *Repository) GetUndeletedTodos() (convertedTodos []entities.Todo, err error) {
	todos := r.DB.Model(model.Todos{}).Find(&convertedTodos)

	return convertedTodos, todos.Error
}

func (r *Repository) GetTodo(todoId uint) (convertedTodo *entities.Todo, err error) {
	todo := r.DB.Model(model.Todos{}).First(&convertedTodo, todoId)

	return convertedTodo, todo.Error
}

func (r *Repository) InsertNewTodo(title string, content string) error {
	todo := model.Todos{
		TITLE:   title,
		CONTENT: content,
	}
	result := r.DB.Create(&todo)

	return result.Error
}

func (r *Repository) UpdateTodo(todoId uint, title string, content string) error {
	todo := model.Todos{
		ID:      todoId,
		TITLE:   title,
		CONTENT: content,
	}
	result := r.DB.Updates(&todo)

	return result.Error
}

func (r *Repository) DeleteTodo(todoId uint) error {
	todo := model.Todos{
		ID: todoId,
	}
	result := r.DB.Delete(&todo)

	return result.Error
}
