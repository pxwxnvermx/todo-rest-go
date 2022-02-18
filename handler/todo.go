package handler

import (
	"database/sql"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	db "github.com/pxwxnvermx/todo-rest-go/db/sqlc"
	logger "github.com/sirupsen/logrus"
)

type TodoResponse struct {
	Id          int64     `json:"id"`
	Description string    `json:"description"`
	IsCompleted bool      `json:"is_completed"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func serializeTodo(todoDb db.Todo) TodoResponse {
	return TodoResponse{
		Id:          todoDb.ID,
		Description: todoDb.Description,
		IsCompleted: todoDb.IsCompleted,
		CreatedAt:   todoDb.CreatedAt.Time,
		UpdatedAt:   todoDb.UpdatedAt.Time,
	}
}

func serializeGetTodoResponse(todoResDB []db.Todo) []TodoResponse {
	var todoResponse []TodoResponse

	for _, todo := range todoResDB {
		tempTodoRes := serializeTodo(todo)
		todoResponse = append(todoResponse, tempTodoRes)
	}

	return todoResponse
}

func GetAllTodo(s *db.Store) gin.HandlerFunc {
	return func(c *gin.Context) {
		code := http.StatusOK

		response, err := s.Queries.GetTodo(c, db.GetTodoParams{Limit: 20, Offset: 0})

		if err != nil {
			code = http.StatusBadRequest
			logger.Errorln(err.Error())
		}

		c.JSON(code, gin.H{"data": serializeGetTodoResponse(response)})
	}
}

func GetTodoById(s *db.Store) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 0)
		code := http.StatusOK

		if err != nil {
			code = http.StatusBadRequest
			logger.Errorln(err.Error())
		}

		response, err := s.Queries.GetTodoById(c, id)

		if err != nil {
			code = http.StatusBadRequest
			logger.Errorln(err.Error())
		}

		c.JSON(code, gin.H{"data": serializeTodo(response)})
	}
}

type createTodoRequestParams struct {
	Description string `json:"description"`
	IsCompleted bool   `json:"is_completed"`
}

func CreateTodo(s *db.Store) gin.HandlerFunc {
	return func(c *gin.Context) {
		var todo createTodoRequestParams

		err := c.ShouldBindJSON(&todo)

		code := http.StatusOK

		if err != nil {
			code = http.StatusBadRequest
			logger.Errorln(err.Error())
		}

		response, err := s.Queries.CreateTodo(c, db.CreateTodoParams{
			Description: todo.Description,
			IsCompleted: todo.IsCompleted,
			CreatedAt:   sql.NullTime{Time: time.Now(), Valid: true},
			UpdatedAt:   sql.NullTime{Time: time.Now(), Valid: true},
		})

		if err != nil {
			code = http.StatusBadRequest
			logger.Errorln(err.Error())
		}

		c.JSON(code, gin.H{"data": serializeTodo(response)})
	}
}

func UpdateTodo(s *db.Store) gin.HandlerFunc {
	return func(c *gin.Context) {
		code := http.StatusOK

		id, err := strconv.ParseInt(c.Param("id"), 10, 0)
		if err != nil {
			code = http.StatusBadRequest
			logger.Errorln(err.Error())
		}

		todoBody, err := s.Queries.GetTodoById(c, id)
		if err != nil {
			code = http.StatusBadRequest
			logger.Errorln(err.Error())
		}

		err = c.ShouldBindJSON(&todoBody)
		if err != nil {
			code = http.StatusBadRequest
			logger.Errorln(err.Error())
		}

		todo, err := s.Queries.UpdateTodo(c, db.UpdateTodoParams{
			Description: todoBody.Description,
			IsCompleted: todoBody.IsCompleted,
			ID:          id,
		})
		if err != nil {
			code = http.StatusBadRequest
			logger.Errorln(err.Error())
		}

		c.JSON(code, gin.H{"data": serializeTodo(todo)})
	}
}

func DeleteTodo(s *db.Store) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 0)
		code := http.StatusOK

		if err != nil {
			code = http.StatusBadRequest
		}

		err = s.Queries.DeleteTodo(c, id)

		if err != nil {
			code = http.StatusBadRequest
		}

		c.Status(code)
	}
}
