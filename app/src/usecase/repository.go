package usecase

import (
	"app/src/entities"
	"app/src/infrastructure/sqlhandler"
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

func (r *Repository) GetUndeletedTodos(sqlhandler *sqlhandler.SqlHandler) (convertedTodos []entities.Todo, err error) {
	todos := sqlhandler.DB.Model(model.Todos{}).Find(&convertedTodos)

	if todos.Error != nil {
		return nil, err
	}

	return convertedTodos, nil
}

func (r *Repository) GetTodo(sqlhandler *sqlhandler.SqlHandler, todoId uint) (convertedTodo *entities.Todo, err error) {
	todo := sqlhandler.DB.Model(model.Todos{}).First(&convertedTodo, todoId)

	if todo.Error != nil {
		return nil, err
	}

	return convertedTodo, nil
}

func (r *Repository) InsertNewTodo(sqlhandler *sqlhandler.SqlHandler, title string, content string) {
	todo := model.Todos{
		TITLE:   title,
		CONTENT: content,
	}
	sqlhandler.DB.Create(&todo)
}

func (r *Repository) UpdateTodo(sqlhandler *sqlhandler.SqlHandler, todoId uint, title string, content string) {
	todo := model.Todos{
		ID:      todoId,
		TITLE:   title,
		CONTENT: content,
	}
	sqlhandler.DB.Updates(&todo)
}

func (r *Repository) DeleteTodo(sqlhandler *sqlhandler.SqlHandler, todoId uint) {
	todo := model.Todos{
		ID: todoId,
	}
	sqlhandler.DB.Delete(&todo)
}
