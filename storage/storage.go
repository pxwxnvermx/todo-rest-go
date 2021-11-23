package storage

import (
	"github.com/pxwxnvermx/todo-rest/models"
	"gorm.io/gorm"
)

type Storage interface {
	GetAllTodo() ([]models.Todo, error)
	GetTodoById(id string) (models.Todo, error)
	CreateTodo(todo models.Todo) error
	UpdateTodo(id string, todo models.Todo) error
	DeleteTodo(id string) error
}

type GormStorage struct {
	db *gorm.DB
}

func NewStorage(db *gorm.DB) *GormStorage {
	return &GormStorage{db}
}
