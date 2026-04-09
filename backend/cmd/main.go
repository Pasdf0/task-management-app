package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/Pasdf0/task-management-app/backend/internal/handler"
	"github.com/Pasdf0/task-management-app/backend/internal/repository"
	"github.com/Pasdf0/task-management-app/backend/internal/server"
	"github.com/Pasdf0/task-management-app/backend/internal/service"
)

func main() {
	// Cargar variables de entorno
	if err := godotenv.Load(); err != nil {
		log.Println("No se pudo cargar el archivo .env, usando variables de entorno del sistema")
	}

	// Configurar conexión a MongoDB
	mongoDBNAME := os.Getenv("MONGO_DB_NAME")
	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		log.Fatal("MONGO_URI no está definido en las variables de entorno")
	}

	// Conectar a MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatalf("Error al conectar a MongoDB: %v", err)
	}

	// Apagar el cliente al finalizar
	defer func() {
		if err = client.Disconnect(context.Background()); err != nil {
			log.Fatal("Error al desconectar: ", err)
		}
	}()

	db := client.Database(mongoDBNAME)

	// Inicializar repositorio, servicio y handler
	taskRepo := repository.NewTaskRepository(db)
	taskService := service.NewTaskService(taskRepo)
	taskHandler := handler.NewTaskHandler(taskService)

	// Configurar router y middleware
	app := server.NewServer(taskHandler)

	// Iniciar servidor
	app.Run(":8080")
}
