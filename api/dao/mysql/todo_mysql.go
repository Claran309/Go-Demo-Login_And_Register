package mysql

import (
	"GoGin/api/dao"
	"GoGin/internal/model"
	"errors"
	"log"

	"gorm.io/gorm"
)

type mysqlTodoRepo struct {
	db *gorm.DB
}

func NewMysqlTodoRepo(db *gorm.DB) dao.TodoRepository {
	err := db.AutoMigrate(&model.TodoTask{})
	if err != nil {
		log.Fatal("Failed to migrate student & course table:", err)
	}

	return &mysqlTodoRepo{db: db}
}

func (repo *mysqlTodoRepo) CreateTodoTask(task *model.TodoTask) error {
	if err := repo.db.Create(task).Error; err != nil {
		return errors.New("failed to create task")
	}
	return nil
}

func (repo *mysqlTodoRepo) DeleteTodoTask(taskID int) error {
	if err := repo.db.Delete(&model.TodoTask{}, taskID).Error; err != nil {
		return errors.New("failed to delete task")
	}
	return nil
}

func (repo *mysqlTodoRepo) FinishTodoTask(taskID int) error {
	if err := repo.db.Where("task_id = ?", taskID).Update("completed", true).Error; err != nil {
		return errors.New("failed to finish task")
	}
	return nil
}

func (repo *mysqlTodoRepo) CheckTodoTask(userID int) ([]model.TodoTask, []model.TodoTask, error) {
	var todos []model.TodoTask
	var dones []model.TodoTask
	if err := repo.db.Where("user_id = ? AND complete = ?", userID, false).Find(&todos).Error; err != nil {
		return nil, nil, errors.New("failed to check task")
	}
	if err := repo.db.Where("user_id = ? AND complete = ?", userID, true).Where("complete = ?", true).Find(&dones).Error; err != nil {
		return nil, nil, errors.New("failed to check task")
	}
	return todos, dones, nil
}
