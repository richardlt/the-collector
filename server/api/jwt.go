package api

import (
	jwt "github.com/dgrijalva/jwt-go"
	errorsP "github.com/pkg/errors"
	"github.com/richardlt/the-collector/server/api/errors"
)

// ParseJWT .
func ParseJWT(secret, token string) (jwt.MapClaims, bool, error) {
	t, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if t.Method.Alg() != jwt.SigningMethodHS256.Name {
			return nil, errors.NewNotFound()
		}
		return []byte(secret), nil
	})
	if err != nil {
		return nil, false, errorsP.WithStack(err)
	}
	if claims, ok := t.Claims.(jwt.MapClaims); ok && t.Valid {
		return claims, true, nil
	}
	return nil, false, nil
}
