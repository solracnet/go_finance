package util

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func ValidateToken(ctx *gin.Context, token string) error {
	claims := &Claims{}
	var jwtSignedKey = []byte("3w2j40ke")
	parsedToken, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		return jwtSignedKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			ctx.JSON(http.StatusUnauthorized, err)
			return err
		}
		ctx.JSON(http.StatusBadRequest, err)
		return err
	}

	if !parsedToken.Valid {
		ctx.JSON(http.StatusUnauthorized, "Token is not valid")
		return errors.New("Token is not valid")
	}

	return nil
}

func GetAndVerifyToken(ctx *gin.Context) error {
	authorizationHeaderKey := ctx.GetHeader("Authorization")
	if authorizationHeaderKey == "" {
		ctx.JSON(http.StatusUnauthorized, "Authorization header is required")
		return errors.New("Authorization header is required")
	}
	fields := strings.Fields(authorizationHeaderKey)
	if len(fields) != 2 || fields[0] != "Bearer" {
		ctx.JSON(http.StatusUnauthorized, "Authorization header format must be Bearer {token}")
		return errors.New("Authorization header format must be Bearer {token}")
	}
	tokenToValidate := fields[1]
	err := ValidateToken(ctx, tokenToValidate)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, err)
		return err
	}
	return nil
}
