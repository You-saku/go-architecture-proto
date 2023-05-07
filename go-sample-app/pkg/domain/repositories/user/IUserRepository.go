package repositories

import (
	"go-sample-app/pkg/domain/models"
)

type IUserRepository interface {
	FindUserById(id string) models.User
	GetAllUsers() []models.User
	NewUser()
	UpdateUser(id string)
	DeleteUser(id string)
}
