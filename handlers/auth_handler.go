package handlers

import (
	"fmt"
	"net/http"
	"simple-jwt-golang/entities"
	"simple-jwt-golang/requests"
	"simple-jwt-golang/services"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var APPLICATION_NAME = "simple_jwt"
var LOGIN_EXPIRATION_DURATION = time.Duration(1) * time.Hour
var JWT_SIGNING_METHOD = jwt.SigningMethodHS256
var JWT_SIGNATURE_KEY = []byte("simple_signature_2893423847234")

type UserJWT struct {
	jwt.StandardClaims
	Username string `json:"Username"`
	UserRole string `json:"UserRole"`
}

type authHandler struct {
	services.AuthService
}

func NewAuthHandler(service services.AuthService) *authHandler {
	return &authHandler{service}
}

func (h authHandler) AuthLogin(c *gin.Context) {

	var LoginUserRequest requests.LoginUserRequest

	err := c.ShouldBind(&LoginUserRequest)

	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on filled %s, condition %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)

		}
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return
	}

	dataLogin := entities.AuthEntity{
		Username: LoginUserRequest.Username,
		Password: LoginUserRequest.Password,
	}

	err = h.AuthService.Login(dataLogin)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  400,
			"message": "Please check your username & password",
			"data":    "",
		})
		return
	}

	claims := UserJWT{
		StandardClaims: jwt.StandardClaims{
			Issuer:    APPLICATION_NAME,
			ExpiresAt: time.Now().Add(LOGIN_EXPIRATION_DURATION).Unix(),
		},
		Username: dataLogin.Username,
	}

	token := jwt.NewWithClaims(
		JWT_SIGNING_METHOD,
		claims,
	)

	signedToken, err_signe := token.SignedString(JWT_SIGNATURE_KEY)
	if err_signe != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  400,
			"message": err_signe.Error(),
			"data":    "",
		})
		return
	}

	tokenString := signedToken

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Authectication success",
		"data":    dataLogin,
		"token":   tokenString,
	})
}
