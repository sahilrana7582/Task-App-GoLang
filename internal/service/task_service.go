package service

import (
	"github.com/sahilrana7582/Task-App-GoLang/internal/database"
	"github.com/sahilrana7582/Task-App-GoLang/internal/model"
)

type TaskService interface {
	CreateTask(task *model.Task) error
	GetAllTasks() ([]model.Task, error)
	GetTaskByID(id int) (*model.Task, error)
	UpdateTask(task *model.Task) error
	DeleteTask(id int) error
}

type taskService struct {
	repo database.TaskRepository
}

func NewTaskService(repo database.TaskRepository) TaskService {
	return &taskService{repo: repo}
}

func (s *taskService) CreateTask(task *model.Task) error {
	return s.repo.CreateTask(task)
}

func (s *taskService) GetAllTasks() ([]model.Task, error) {
	return s.repo.GetAllTasks()
}

func (s *taskService) GetTaskByID(id int) (*model.Task, error) {
	return s.repo.GetTaskByID(id)
}

func (s *taskService) UpdateTask(task *model.Task) error {
	return s.repo.UpdateTask(task)
}

func (s *taskService) DeleteTask(id int) error {
	return s.repo.DeleteTask(id)
}
