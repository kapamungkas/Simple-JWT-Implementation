package middlewares

import (
	"errors"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var JWT_SIGNING_METHOD = jwt.SigningMethodHS256
var JWT_SIGNATURE_KEY = []byte(os.Getenv("JWT_SIGNATURE_KEY"))

func UserMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {
		const BEARER_SCHEMA = "Bearer "
		authHeader := c.GetHeader("Authorization")

		if len(authHeader) < 6 {
			respondWithError(c, 401, "API token required")
			return
		}

		tokenString := authHeader[len(BEARER_SCHEMA):]

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				// respondWithError(c, 401, "Signing method invalid")
				return JWT_SIGNATURE_KEY, errors.New("signing method invalid")
			} else if method != JWT_SIGNING_METHOD {
				return JWT_SIGNATURE_KEY, errors.New("signing method invalid")
			}

			return JWT_SIGNATURE_KEY, nil
		})
		if err != nil {
			respondWithError(c, 401, err.Error())
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			respondWithError(c, 401, err.Error())
			return
		}

		c.Set("username", claims["Username"])
		c.Next()
	}
}

func respondWithError(c *gin.Context, code int, message interface{}) {
	c.AbortWithStatusJSON(code, gin.H{
		"status":  400,
		"message": message,
		"data":    "",
	})
}
