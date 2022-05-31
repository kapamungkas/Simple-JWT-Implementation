package services

import (
	"errors"
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
	find_user, err := r.r.FindUsernamePassword(dataLogin.Username)
	if err != nil {
		return err
	}

	if dataLogin.Username == find_user.Username && dataLogin.Password == dataLogin.Password {
		return nil
	}
	return errors.New("Please check your username and password")
}
