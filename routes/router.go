package routes

import (
	"github.com/gin-gonic/gin"
	db "github.com/pxwxnvermx/todo-rest-go/db/sqlc"
	"github.com/pxwxnvermx/todo-rest-go/handler"
	"github.com/pxwxnvermx/todo-rest-go/middleware"
)

func NewRouter(todoStorage *db.Store) *gin.Engine {
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
