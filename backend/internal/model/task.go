package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Task : Tarea struct
type Task struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Title       string             `json:"title" binding:"required,min=3,max=100" bson:"title"`
	Description string             `json:"description" binding:"max=250" bson:"description"`
	Completed   bool               `json:"completed" bson:"completed"`
	Tags        []string           `json:"tags" binding:"omitempty,dive,min=1" bson:"tags,omitempty"`
	CreatedAt   time.Time          `json:"createdAt" bson:"createdAt"`
}
