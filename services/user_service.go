package services

import (
	"fmt"
)

type UserService struct {
	*BaseService
}

func NewUserService() *UserService {
	service := NewBaseService()
	return &UserService{BaseService: service}
}

func (s UserService) Register() {
	fmt.Println("service")
}
