package service

import (
	"context"
	"errors"
	"math"
	"strings"

	"github.com/Pasdf0/task-management-app/backend/internal/model"
	"github.com/Pasdf0/task-management-app/backend/internal/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ErrTaskAlreadyExists : Tarea con el mismo título ya existe
var ErrTaskAlreadyExists = errors.New("ya existe una tarea con ese título")

// TaskService : Servicio de tareas
type TaskService struct {
	repo *repository.TaskRepository
}

// NewTaskService : Constructor de TaskService
func NewTaskService(repo *repository.TaskRepository) *TaskService {
	return &TaskService{
		repo: repo,
	}
}

// CreateTask : Crea una nueva tarea
func (s *TaskService) CreateTask(ctx context.Context, task *model.Task) error {
	task.Title = strings.TrimSpace(task.Title)
	task.Description = strings.TrimSpace(task.Description)

	// Descripçión opcional, si está vacía se asigna un valor por defecto
	if task.Description == "" {
		task.Description = "Sin descripción"
	}

	// Limpiar y normalizar tags
	if len(task.Tags) > 0 {
		var cleanedTags []string

		for _, tag := range task.Tags {
			cleanTag := strings.ToLower(strings.TrimSpace(tag))

			if cleanTag != "" {
				cleanedTags = append(cleanedTags, cleanTag)
			}
		}
		task.Tags = cleanedTags
	}

	// Evitar Titulos duplicados
	_, err := s.repo.FindByTitle(ctx, task.Title)
	if err == nil {
		return ErrTaskAlreadyExists
	}
	if !errors.Is(err, repository.ErrTaskNotFound) {
		return err
	}

	// Guardar en repositorio
	err = s.repo.Create(ctx, task)
	if err != nil {
		return err
	}

	return nil
}

// GetAllTasks : Obtiene todas las tareas
func (s *TaskService) GetAllTasks(ctx context.Context) ([]model.Task, error) {
	tasks, err := s.repo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

// GetSomeTasks : Obtiene un número limitado de tareas
func (s *TaskService) GetSomeTasks(ctx context.Context, page int, limit int) (*model.PaginatedTasks, error) {
	// Validar parámetros
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	// Calcular offset
	offset := (page - 1) * limit

	tasks, totalItems, err := s.repo.FindSomeTasks(ctx, offset, limit)
	if err != nil {
		return nil, err
	}

	// Calcular total de páginas
	totalPages := int(math.Ceil(float64(totalItems) / float64(limit)))

	// Crear respuesta paginada
	response := &model.PaginatedTasks{
		Data: tasks,
		Meta: model.Pagination{
			TotalItems:  totalItems,
			TotalPages:  totalPages,
			CurrentPage: page,
			PageSize:    limit,
		},
	}

	return response, nil
}

// GetTaskByID : Obtiene una tarea por ID
func (s *TaskService) GetTaskByID(ctx context.Context, idString string) (*model.Task, error) {
	// Convertir string a ObjectID
	id, err := primitive.ObjectIDFromHex(idString)
	if err != nil {
		return nil, errors.New("ID inválido")
	}

	// Buscar en repositorio
	task, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return task, nil
}

// CompleteTask : Marca una tarea como completada
func (s *TaskService) CompleteTask(ctx context.Context, idString string) error {
	// Convertir string a ObjectID
	id, err := primitive.ObjectIDFromHex(idString)
	if err != nil {
		return errors.New("ID inválido")
	}

	// Marcar tarea como completada
	err = s.repo.CompleteTask(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

// DeleteTask : Elimina una tarea por ID
func (s *TaskService) DeleteTask(ctx context.Context, idString string) error {
	// Convertir string a ObjectID
	id, err := primitive.ObjectIDFromHex(idString)
	if err != nil {
		return errors.New("ID inválido")
	}

	// Eliminar tarea
	err = s.repo.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

// AddTagsToTask : Agrega una o más etiquetas a una tarea
func (s *TaskService) AddTagsToTask(ctx context.Context, idString string, tag []string) error {
	// Convertir string a ObjectID
	id, err := primitive.ObjectIDFromHex(idString)
	if err != nil {
		return errors.New("ID inválido")
	}

	// Limpiar y normalizar tags
	var cleanedTags []string
	for _, t := range tag {
		cleanTag := strings.ToLower(strings.TrimSpace(t))
		if cleanTag != "" {
			cleanedTags = append(cleanedTags, cleanTag)
		}
	}

	if len(cleanedTags) == 0 {
		return errors.New("no se proporcionaron tags válidos")
	}

	// Agregar tags a la tarea
	err = s.repo.AddTagsToTask(ctx, id, cleanedTags)
	if err != nil {
		return err
	}

	return nil
}
