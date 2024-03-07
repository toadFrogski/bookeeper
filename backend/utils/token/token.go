package token

import (
	"bookeeper/domain"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v5"
)

type Token struct {
	Token   string
	Expires int64
}

func GenerateToken(user *domain.User) (*Token, error) {
	var tokenLifespan int
	var tokenExp int64
	var token *jwt.Token
	var tokenString string

	tokenLifespan, err := strconv.Atoi(os.Getenv("TOKEN_HOUR_LIFESPAN"))
	if err != nil {
		return nil, err
	}

	tokenExp = time.Now().Add(time.Hour * time.Duration(tokenLifespan)).Unix()
	token = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"authorized": true,
		"sub":        strconv.Itoa(int(user.ID)),
		"exp":        tokenExp,
	})
	tokenString, err = token.SignedString([]byte(os.Getenv("API_SECRET")))

	return &Token{Token: tokenString, Expires: tokenExp}, err
}

func ExtractToken(c *gin.Context) string {
	if token := c.Query("token"); token != "" {
		return token
	}

	bearerToken := c.Request.Header.Get("Authorization")

	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}

func ExtractTokenClaims(c *gin.Context) (*jwt.MapClaims, error) {
	tokenString := ExtractToken(c)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("API_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		return &claims, nil
	}
	return nil, nil
}
