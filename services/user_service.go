package services

import (
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
)

type UserService struct {
	*BaseService
}

func NewUserService(db *mongo.Database) *UserService {
	service := NewBaseService(db)
	return &UserService{BaseService: service}
}

func (s UserService) Register() {
	fmt.Println("service")
}
