package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/Pasdf0/task-management-app/backend/internal/model"
	"github.com/Pasdf0/task-management-app/backend/internal/service"
	"github.com/gin-gonic/gin"
)

// TaskHandler : Handler de tareas
type TaskHandler struct {
	service *service.TaskService
}

// NewTaskHandler : Constructor de TaskHandler
func NewTaskHandler(service *service.TaskService) *TaskHandler {
	return &TaskHandler{service: service}
}

// CreateTask : Handler para crear una nueva tarea
func (h *TaskHandler) CreateTask(c *gin.Context) {
	var newTask model.Task

	// Validar datos de entrada y bindear a struct
	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Datos no válidos",
			"details": err.Error(),
		})
		return
	}

	// Validar lógica de negocio
	err := h.service.CreateTask(c.Request.Context(), &newTask)
	if err != nil {
		if errors.Is(err, service.ErrTaskAlreadyExists) {
			c.JSON(http.StatusConflict, gin.H{
				"error":   "Error al crear la tarea",
				"details": err.Error(),
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Error al crear la tarea",
			"details": err.Error(),
		})
		return
	}

	// Tarea creada exitosamente
	c.JSON(http.StatusCreated, newTask)
}

// GetAllTasks : Handler para obtener todas las tareas
func (h *TaskHandler) GetAllTasks(c *gin.Context) {
	tasks, err := h.service.GetAllTasks(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener tareas"})
		return
	}

	c.JSON(http.StatusOK, tasks)
}

// GetSomeTasks : Handler para obtener tareas según parámetros
func (h *TaskHandler) GetSomeTasks(c *gin.Context) {
	// Leer parámetros
	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "10")

	// Convertir a enteros
	page, err1 := strconv.Atoi(pageStr)
	limit, err2 := strconv.Atoi(limitStr)

	if err1 != nil || err2 != nil || page < 1 || limit < 1 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Parámetros inválidos",
		})
		return
	}

	// Obtener tareas del servicio
	paginatedResult, err := h.service.GetSomeTasks(c.Request.Context(), page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener tareas"})
		return
	}

	c.JSON(http.StatusOK, paginatedResult)
}

// GetTask : Handler para obtener una tarea por ID
func (h *TaskHandler) GetTask(c *gin.Context) {
	idString := c.Param("id")

	// Obtener Tarea por ID
	task, err := h.service.GetTaskByID(c.Request.Context(), idString)

	if err != nil {
		if err.Error() == "tarea no encontrada" {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "No se encontró la tarea",
			})
			return
		}

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, task)
}

// CompleteTask : Handler para marcar una tarea como completada
func (h *TaskHandler) CompleteTask(c *gin.Context) {
	idString := c.Param("id")

	err := h.service.CompleteTask(c.Request.Context(), idString)
	if err != nil {
		if err.Error() == "tarea no encontrada" {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "No se encontró la tarea",
			})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Tarea marcada como completada"})
}

// DeleteTask : Handler para eliminar una tarea
func (h *TaskHandler) DeleteTask(c *gin.Context) {
	idString := c.Param("id")

	err := h.service.DeleteTask(c.Request.Context(), idString)
	if err != nil {
		if err.Error() == "tarea no encontrada" {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "No se encontró la tarea",
			})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Tarea eliminada"})
}

// AddTagsToTask : Handler para agregar tags a una tarea
func (h *TaskHandler) AddTagsToTask(c *gin.Context) {
	idString := c.Param("id")

	var req struct {
		Tags []string `json:"tags" binding:"required,dive,min=1"`
	}

	// Validar datos de entrada y bindear a struct
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Datos no válidos",
			"details": err.Error(),
		})
		return
	}

	// Agregar tags a la tarea
	err := h.service.AddTagsToTask(c.Request.Context(), idString, req.Tags)
	if err != nil {
		if err.Error() == "tarea no encontrada" {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "No se encontró la tarea",
			})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Tags agregados a la tarea"})
}
