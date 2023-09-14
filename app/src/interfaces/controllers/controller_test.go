package controllers

import (
	"app/src/entities"
	"app/src/infrastructure/sqlhandler"
	"app/src/test/testutil"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestController_NewTodoSubmit(t *testing.T) {
	e := echo.New()

	// Connect to db for testing
	testSqlHandler := testutil.NewSqlHandler()
	controller := NewController((*sqlhandler.SqlHandler)(testSqlHandler))

	tests := []struct {
		title   string
		content string
	}{
		{"test title", "test content"},
		{"", "test content"},
		{"test title", ""},
		{"", ""},
	}

	for _, tt := range tests {
		values := url.Values{}
		values.Set("title", tt.title)
		values.Add("content", tt.content)
		req := httptest.NewRequest(http.MethodPost, "/new_todos/submit", strings.NewReader(values.Encode()))
		req.Header.Set(echo.HeaderContentType, "application/x-www-form-urlencoded")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		err := controller.NewTodoSubmit(c)
		if err != nil {
			t.Errorf("unexpected error")
		}

		foundTodo := []entities.Todo{}
		result := testSqlHandler.DB.Find(&foundTodo)

		if tt.title != "" && tt.content != "" {
			t.Run("valid title and valid content are submitted", func(t *testing.T) {
				if result.RowsAffected == 0 {
					t.Errorf("Expect: todo is inserted. But todo is not inserted")
				}

				for _, todo := range foundTodo {
					if todo.Title != "test title" {
						t.Errorf("Expect: inputted title is inserted but unexpected title is inserted")
					}

					if todo.Content != "test content" {
						t.Errorf("Expect: inputted content is inserted but unexpected content is inserted")
					}

					if todo.CreatedAt.IsZero() != false {
						t.Errorf("Expect: created_at is inserted but it is NULL")
					}

					if todo.UpdatedAt.IsZero() != false {
						t.Errorf("Expect: updated_at is inserted but it is NULL")
					}

					if todo.DeletedAt.IsZero() != true {
						t.Errorf("Expect: deleted_at is NULL but it is inserted")
					}
				}

				assert.Equal(t, http.StatusFound, rec.Code)
				assert.Equal(t, "/all_todos", rec.HeaderMap.Get("Location"))
			})
		}

		if tt.title == "" || tt.content == "" {
			t.Run("invalid title or invalid content is submitted", func(t *testing.T) {
				if result.RowsAffected == 1 {
					t.Errorf("Expect: todo is not inserted. But todo is inserted")
				}

				assert.Equal(t, http.StatusFound, rec.Code)
				assert.Equal(t, "/new_todo", rec.HeaderMap.Get("Location"))
			})
		}

		if tt.title == "" && tt.content == "" {
			t.Run("invalid title and invalid content is submitted", func(t *testing.T) {
				if result.RowsAffected == 1 {
					t.Errorf("Expect: todo is not inserted. But todo is inserted")
				}

				assert.Equal(t, http.StatusFound, rec.Code)
				assert.Equal(t, "/new_todo", rec.HeaderMap.Get("Location"))
			})
		}

		testutil.TruncateTodoTable(*testSqlHandler)
	}
}
