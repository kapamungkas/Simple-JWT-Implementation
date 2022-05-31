package repositories

import (
	"simple-jwt-golang/entities"
)

type AuthRepository interface {
	FindUsernamePassword(username string) (entities.AuthEntity, error)
}

type authRepository struct{}

func NewAuthRepository() *authRepository {
	return &authRepository{}
}

func (r authRepository) FindUsernamePassword(username string) (entities.AuthEntity, error) {
	return entities.AuthEntity{}, nil
}
