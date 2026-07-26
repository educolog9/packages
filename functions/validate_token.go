package functions

import (
	"errors"
	"os"

	"github.com/educolog9/packages/types"
	"github.com/golang-jwt/jwt/v5"
)

func ValidateToken(tokenString string) (*types.UserClaims, error) {
	// Tokens are issued with HS256 (see the user service's SignToken). Pinning
	// the accepted algorithms keeps the parser from honouring whatever `alg` the
	// token itself asks for.
	token, err := jwt.ParseWithClaims(tokenString, &types.UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	}, jwt.WithValidMethods([]string{"HS256"}))

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*types.UserClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, errors.New("invalid token")
	}
}
