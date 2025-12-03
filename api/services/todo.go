package services

import (
	"GoGin/api/dao"
	"GoGin/internal/model"
	"time"
)

type TodoService struct {
	TodoRepo dao.TodoRepository
}

func NewTodoService(todoRepo dao.TodoRepository) *TodoService {
	return &TodoService{TodoRepo: todoRepo}
}

func (s *TodoService) CreateTodoTask(req model.CreateTodoRequest, userID int) (model.TodoTask, error) {
	// 封装
	var todoTask = model.TodoTask{
		UserID:      userID,
		Title:       req.Title,
		Description: req.Description,
		Completed:   false,
		CreatedTime: time.Now(),
	}

	// 调用数据层
	err := s.TodoRepo.CreateTodoTask(&todoTask)
	if err != nil {
		return model.TodoTask{}, err
	}

	return todoTask, nil
}

func (s *TodoService) FinishTodoTask(taskID int) error {
	//调用数据层
	err := s.TodoRepo.FinishTodoTask(taskID)
	if err != nil {
		return err
	}

	return nil
}

func (s *TodoService) DeleteTodoTask(taskID int) error {
	//调用数据层
	err := s.TodoRepo.DeleteTodoTask(taskID)
	if err != nil {
		return err
	}

	return nil
}

func (s *TodoService) GetInfo(userID int) ([]model.TodoTask, []model.TodoTask, error) {
	//调用数据层
	todos, dones, err := s.TodoRepo.CheckTodoTask(userID)
	if err != nil {
		return []model.TodoTask{}, []model.TodoTask{}, err
	}

	return todos, dones, nil
}
