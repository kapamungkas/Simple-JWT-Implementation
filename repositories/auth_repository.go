package repositories

import (
	"errors"
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
	if username == "admin" {
		return entities.AuthEntity{
			Username: "admin",
			Password: "admin",
		}, nil
	}
	return entities.AuthEntity{}, errors.New("user not found")
}
