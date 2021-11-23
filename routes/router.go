package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/pxwxnvermx/todo-rest/handler"
	"github.com/pxwxnvermx/todo-rest/middleware"
	"github.com/pxwxnvermx/todo-rest/storage"
	"github.com/sirupsen/logrus"
)

func NewRouter(logger *logrus.Logger, todoStorage storage.Storage) *gin.Engine {
	r := gin.Default()

	r.Use(middleware.Cors())

	v1 := r.Group("/api/v1")
	{
		v1.GET("/todo", handler.GetAllTodo(todoStorage))
		v1.GET("/todo/:id", handler.GetTodoById(todoStorage))
		v1.POST("/todo", handler.CreateTodo(todoStorage))
		v1.PUT("/todo/:id", handler.UpdateTodo(todoStorage))
		v1.DELETE("/todo/:id", handler.DeleteTodo(todoStorage))
	}

	return r
}
