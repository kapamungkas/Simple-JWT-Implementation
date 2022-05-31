package services

import (
	"simple-jwt-golang/entities"
	"simple-jwt-golang/repositories"
)

type AuthService interface {
	Login(dataLogin entities.AuthEntity) error
}

type authSerAuthService struct {
	r repositories.AuthRepository
}

func NewAuthService(r repositories.AuthRepository) *authSerAuthService {
	return &authSerAuthService{r}
}

func (r authSerAuthService) Login(dataLogin entities.AuthEntity) error {
	return nil
}
