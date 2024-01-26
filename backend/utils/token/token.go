package token

import (
	"fmt"
	"gg/domain"
	"gg/utils/constants"
	"os"
	"strconv"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type Claims struct {
	UserID float64
	Roles  []constants.Role
}

func GenerateToken(user domain.User) (string, error) {
	token_lifespan, err := strconv.Atoi(os.Getenv("TOKEN_HOUR_LIFESPAN"))
	if err != nil {
		return "", err
	}

	roles := make([]constants.Role, len(user.Roles))
	for i, role := range user.Roles {
		roles[i] = role.Name
	}

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["sub"] = Claims{
		UserID: float64(user.ID),
		Roles:  roles,
	}
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(token_lifespan)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(os.Getenv("API_SECRET")))
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

func ExtractTokenClaims(c *gin.Context) (*Claims, error) {
	tokenString := ExtractToken(c)
	token, err := jwt.Parse(tokenString, _validateToken)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		var roles []constants.Role

		sub := claims["sub"].(map[string]any)
		userID := sub["UserID"].(float64)
		for _, role := range sub["Roles"].([]any) {
			roles = append(roles, constants.Role(role.(string)))
		}

		return &Claims{
			UserID: userID,
			Roles:  roles,
		}, nil
	}
	return nil, nil
}

func _validateToken(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	}
	if err := token.Claims.Valid(); err != nil {
		return nil, err
	}

	return []byte(os.Getenv("API_SECRET")), nil
}
