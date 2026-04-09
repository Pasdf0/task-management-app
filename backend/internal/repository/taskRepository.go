package repository

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/Pasdf0/task-management-app/backend/internal/model"
)

// ErrTaskNotFound : Tarea no encontrada
var ErrTaskNotFound = errors.New("tarea no encontrada")

// TaskRepository : Repositorio de tareas
type TaskRepository struct {
	collection *mongo.Collection
}

// NewTaskRepository : Constructor de TaskRepository
func NewTaskRepository(db *mongo.Database) *TaskRepository {
	return &TaskRepository{
		collection: db.Collection("tasks"),
	}
}

// Create : Crea una nueva tarea en la base de datos
func (r *TaskRepository) Create(ctx context.Context, task *model.Task) error {
	task.CreatedAt = time.Now().Truncate(time.Millisecond)

	// Guardar en base de datos
	result, err := r.collection.InsertOne(ctx, task)
	if err != nil {
		return err
	}

	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		task.ID = oid
	}

	return nil
}

// FindAll : Busca todas las tareas
func (r *TaskRepository) FindAll(ctx context.Context) ([]model.Task, error) {
	var tasks []model.Task

	// Ejecutar consulta para obtener todas las tareas
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	// Iterar sobre resultados
	if err = cursor.All(ctx, &tasks); err != nil {
		return nil, err
	}

	return tasks, nil
}

// FindSomeTasks : Busca tareas con paginación
func (r *TaskRepository) FindSomeTasks(ctx context.Context, offset int, limit int) ([]model.Task, int64, error) {
	var tasks []model.Task

	// Contar Items
	totalItems, err := r.collection.CountDocuments(ctx, bson.M{})
	if err != nil {
		return nil, 0, err
	}

	// Configurar opciones de paginación
	findOptions := options.Find()
	findOptions.SetSkip(int64(offset))
	findOptions.SetLimit(int64(limit))
	findOptions.SetSort(bson.D{{Key: "createdAt", Value: -1}})

	// Ejecutar consulta
	cursor, err := r.collection.Find(ctx, bson.M{}, findOptions)
	if err != nil {
		return nil, 0, err
	}
	defer cursor.Close(ctx)

	if err = cursor.All(ctx, &tasks); err != nil {
		return nil, 0, err
	}

	// Retornar tareas y total de items
	return tasks, totalItems, nil
}

// FindByID : Busca una tarea por ID
func (r *TaskRepository) FindByID(ctx context.Context, id primitive.ObjectID) (*model.Task, error) {
	var task model.Task

	// Ejecutar consulta por ID
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&task)
	if err != nil {
		return nil, ErrTaskNotFound
	}

	return &task, nil
}

// CompleteTask : Marca una tarea como completada
func (r *TaskRepository) CompleteTask(ctx context.Context, id primitive.ObjectID) error {
	// Ejecutar actualización para marcar como completada
	update := bson.M{"$set": bson.M{"completed": true}}
	result, err := r.collection.UpdateOne(ctx, bson.M{"_id": id}, update)
	if err != nil {
		return err
	}

	// Verificar si se encontró la tarea
	if result.MatchedCount == 0 {
		return ErrTaskNotFound
	}

	return nil
}

// Delete : Elimina una tarea por ID
func (r *TaskRepository) Delete(ctx context.Context, id primitive.ObjectID) error {
	// Ejecutar eliminación por ID
	result, err := r.collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return err
	}

	// Verificar si se eliminó alguna tarea
	if result.DeletedCount == 0 {
		return ErrTaskNotFound
	}

	return nil
}

// AddTagsToTask : Agrega etiquetas a una tarea
func (r *TaskRepository) AddTagsToTask(ctx context.Context, id primitive.ObjectID, tags []string) error {
	// Ejecutar actualización para agregar tags sin duplicados
	update := bson.M{"$addToSet": bson.M{"tags": bson.M{"$each": tags}}}
	result, err := r.collection.UpdateOne(ctx, bson.M{"_id": id}, update)
	if err != nil {
		return err
	}

	// Verificar si se encontró la tarea
	if result.MatchedCount == 0 {
		return ErrTaskNotFound
	}

	return nil
}

// FindByTitle : Busca una tarea por título
func (r *TaskRepository) FindByTitle(ctx context.Context, title string) (*model.Task, error) {
	var task model.Task

	// Ejecutar consulta por título
	err := r.collection.FindOne(ctx, bson.M{"title": title}).Decode(&task)
	if err != nil {
		return nil, ErrTaskNotFound
	}

	return &task, nil
}
