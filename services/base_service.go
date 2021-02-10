package services

import "go.mongodb.org/mongo-driver/mongo"

type BaseService struct {
	db *mongo.Database
}

func NewBaseService(db *mongo.Database) *BaseService {
	return &BaseService{db: db}
}
