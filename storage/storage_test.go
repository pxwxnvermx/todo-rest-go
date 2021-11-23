package storage

import (
	"fmt"
	"testing"

	"github.com/pxwxnvermx/todo-rest/models"
	"github.com/pxwxnvermx/todo-rest/utils"
)

var db Storage

func init() {
	db = setup()
}

func setup() Storage {
	db, err := utils.InitDB()

	if err != nil {
		fmt.Print(err)
	}

	return NewStorage(db)
}

func TestCreateTodo(t *testing.T) {
	todo := models.Todo{
		Description: "Hello Testing",
	}

	err := db.CreateTodo(todo)

	if err != nil {
		t.Fail()
	}
}

func TestGetTodo(t *testing.T) {
	todos, err := db.GetAllTodo()

	if err != nil && len(todos) >= 1 {
		t.Fail()
	}
}
