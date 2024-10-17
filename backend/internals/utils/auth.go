package utils

import (
	"fmt"
	"time"

	"github.com/V4T54L/movie-reservation-system/internals/schemas"
)

// TODO :
var tokenStore map[string]schemas.UserToken = make(map[string]schemas.UserToken)

func GenerateToken(user schemas.UserToken) string {
	// token := time.Now().UTC().String()
	token := fmt.Sprintf("token%d", time.Now().UTC().Nanosecond())

	// TODO: Optimize this recurrsion
	// Handle token duplication
	if _, ok := tokenStore[token]; ok {
		return GenerateToken(user)
	}

	tokenStore[token] = user
	return token
}

func GetUserFromToken(token string) (schemas.UserToken, bool) {
	user, valid := tokenStore[token]
	return user, valid
}

func DecodePassword(password string) string {
	return password
}

func Hash(str string) string {
	return str
}
