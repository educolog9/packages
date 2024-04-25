package functions

import (
	"errors"
	"os"

	"github.com/educolog9/packages/types"
	"github.com/golang-jwt/jwt"
)

func ValidateToken(tokenString string) (*types.UserClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &types.UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*types.UserClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, errors.New("invalid token")
	}
}
