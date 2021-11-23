package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pxwxnvermx/todo-rest/models"
	"github.com/pxwxnvermx/todo-rest/storage"
)

func GetAllTodo(s storage.Storage) gin.HandlerFunc {
	return func(c *gin.Context) {
		code := http.StatusOK

		response, err := s.GetAllTodo()

		if err != nil {
			code = http.StatusBadRequest
		}

		c.JSON(code, gin.H{"data": response})
	}
}

func GetTodoById(s storage.Storage) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		code := http.StatusOK

		response, err := s.GetTodoById(id)

		if err != nil {
			code = http.StatusBadRequest
		}

		c.JSON(code, gin.H{"data": response})
	}
}

func CreateTodo(s storage.Storage) gin.HandlerFunc {
	return func(c *gin.Context) {
		var todo models.Todo

		err := c.ShouldBindJSON(&todo)

		code := http.StatusOK

		if err != nil {
			code = http.StatusBadRequest
		}

		err = s.CreateTodo(models.Todo{
			Description: todo.Description,
		})

		if err != nil {
			code = http.StatusBadRequest
		}

		c.Status(code)
	}
}

func UpdateTodo(s storage.Storage) gin.HandlerFunc {
	return func(c *gin.Context) {
		var todo models.Todo

		id := c.Param("id")

		code := http.StatusOK

		err := c.ShouldBindJSON(&todo)

		if err != nil {
			code = http.StatusBadRequest
			fmt.Print(todo)
		}

		err = s.UpdateTodo(id, models.Todo{
			Description: todo.Description,
			IsCompleted: todo.IsCompleted,
		})

		if err != nil {
			code = http.StatusBadRequest
		}

		c.Status(code)
	}
}

func DeleteTodo(s storage.Storage) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		code := http.StatusOK

		err := s.DeleteTodo(id)

		if err != nil {
			code = http.StatusBadRequest
		}

		c.Status(code)
	}
}
