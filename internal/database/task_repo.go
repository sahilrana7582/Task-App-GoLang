package database

import (
	"database/sql"
	"fmt"

	"github.com/sahilrana7582/Task-App-GoLang/internal/model"
)

type TaskRepository interface {
	CreateTask(task *model.Task) error
	GetAllTasks() ([]model.Task, error)
	GetTaskByID(id int) (*model.Task, error)
	UpdateTask(task *model.Task) error
	DeleteTask(id int) error
}

type taskRepository struct {
	db *sql.DB
}

func NewTaskRepository(db *sql.DB) *taskRepository {
	return &taskRepository{db: db}
}

func (repo *taskRepository) CreateTask(task *model.Task) error {
	query := `INSERT INTO tasks (title, description, status) VALUES ($1, $2, $3) RETURNING id`
	err := repo.db.QueryRow(query, task.Title, task.Description, task.Status).Scan(&task.ID)
	if err != nil {
		fmt.Println("Error creating task:", err)
		return err
	}

	return nil
}

func (repo *taskRepository) GetAllTasks() ([]model.Task, error) {
	query := `SELECT id, title, description, status, created_at, updated_at FROM tasks`
	rows, err := repo.db.Query(query)
	if err != nil {
		fmt.Println("Error fetching tasks:", err)
		return nil, err
	}
	defer rows.Close()

	var tasks []model.Task
	for rows.Next() {
		var task model.Task
		if err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Status, &task.CreatedAt, &task.UpdatedAt); err != nil {
			fmt.Println("Error scanning task:", err)
			return nil, err
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (repo *taskRepository) GetTaskByID(id int) (*model.Task, error) {
	query := `SELECT id, title, description, status, created_at, updated_at FROM tasks WHERE id = $1`
	row := repo.db.QueryRow(query, id)

	var task model.Task
	if err := row.Scan(&task.ID, &task.Title, &task.Description, &task.Status, &task.CreatedAt, &task.UpdatedAt); err != nil {
		fmt.Println("Error fetching task by ID:", err)
		return nil, err
	}

	return &task, nil
}

func (repo *taskRepository) UpdateTask(task *model.Task) error {
	query := `UPDATE tasks SET title = $1, description = $2, status = $3, updated_at = CURRENT_TIMESTAMP WHERE id = $4`
	_, err := repo.db.Exec(query, task.Title, task.Description, task.Status, task.ID)
	if err != nil {
		fmt.Println("Error updating task:", err)
		return err
	}

	return nil
}

func (repo *taskRepository) DeleteTask(id int) error {
	query := `DELETE FROM tasks WHERE id = $1`
	_, err := repo.db.Exec(query, id)
	if err != nil {
		fmt.Println("Error deleting task:", err)
		return err
	}

	return nil
}
