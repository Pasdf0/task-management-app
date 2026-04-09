package server

import (
	"log"

	"github.com/Pasdf0/task-management-app/backend/internal/handler"
	"github.com/Pasdf0/task-management-app/backend/internal/middleware"
	"github.com/gin-gonic/gin"
)

type Server struct {
	engine *gin.Engine
}

// NewServer : Constructor de Server
func NewServer(taskHandler *handler.TaskHandler) *Server {
	r := gin.Default()

	// Middleware
	r.Use(middleware.Cors())

	api := r.Group("/api")
	{
		api.POST("/tasks", taskHandler.CreateTask)
		api.GET("/tasks", taskHandler.GetAllTasks)
		api.GET("/tasks/filter", taskHandler.GetSomeTasks)

		api.GET("/tasks/:id", taskHandler.GetTask)
		api.PUT("/tasks/:id/complete", taskHandler.CompleteTask)
		api.DELETE("/tasks/:id", taskHandler.DeleteTask)

		api.POST("/tasks/:id/tags", taskHandler.AddTagsToTask)
	}

	return &Server{
		engine: r,
	}
}

// Run : Inicia el servidor
func (s *Server) Run(port string) {
	if err := s.engine.Run(port); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}
