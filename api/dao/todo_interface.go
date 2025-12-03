package dao

import "GoGin/internal/model"

type TodoRepository interface {
	CreateTodoTask(task *model.TodoTask) error
	DeleteTodoTask(taskID int) error
	FinishTodoTask(taskID int) error
	CheckTodoTask(userID int) ([]model.TodoTask, []model.TodoTask, error)
}
