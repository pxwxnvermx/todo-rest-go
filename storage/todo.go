package storage

import (
	"fmt"

	"github.com/pxwxnvermx/todo-rest/models"
)

func (gs *GormStorage) GetAllTodo() ([]models.Todo, error) {
	var todos []models.Todo

	err := gs.db.Find(&todos).Error

	if err != nil {
		fmt.Println(err.Error())
		return todos, err
	}

	return todos, nil
}

func (gs *GormStorage) GetTodoById(id string) (models.Todo, error) {
	var todo models.Todo

	err := gs.db.First(&todo, id).Error

	if err != nil {
		fmt.Println(err.Error())
		return todo, err
	}

	return todo, nil
}

func (gs *GormStorage) CreateTodo(todo models.Todo) error {
	err := gs.db.Create(&todo).Error

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	return nil
}

func (gs *GormStorage) UpdateTodo(id string, todo models.Todo) error {
	err := gs.db.Model(&models.Todo{}).Where("id = ?", id).Updates(
		map[string]interface{}{
			"description":  todo.Description,
			"is_completed": todo.IsCompleted,
		}).Error

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	return nil
}

func (gs *GormStorage) DeleteTodo(id string) error {
	err := gs.db.Delete(&models.Todo{}, id).Error

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	return nil
}
