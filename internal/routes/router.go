package router

import (
	"github.com/gorilla/mux"
	"github.com/sahilrana7582/Task-App-GoLang/internal/database"
	"github.com/sahilrana7582/Task-App-GoLang/internal/handler"
	"github.com/sahilrana7582/Task-App-GoLang/internal/service"
)

func RegisterRoutes(r *mux.Router) {
	taskRepo := database.NewTaskRepository(database.DB)
	taskService := service.NewTaskService(taskRepo)
	taskHandler := handler.NewTaskHandler(taskService)

	r.HandleFunc("/api/v1/tasks", taskHandler.CreateTask).Methods("POST")
	r.HandleFunc("/api/v1/tasks", taskHandler.GetAllTasks).Methods("GET")
	r.HandleFunc("/api/v1/tasks/{id}", taskHandler.GetTaskByID).Methods("GET")
	r.HandleFunc("/api/v1/tasks/{id}", taskHandler.UpdateTask).Methods("PUT")
	r.HandleFunc("/api/v1/tasks/{id}", taskHandler.DeleteTask).Methods("DELETE")
}
