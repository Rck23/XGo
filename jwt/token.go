package jwt

import (
	"errors"
	"strings"

	jwt "github.com/golang-jwt/jwt/v5"

	"XGo/models"
)

var Email string
var IDUsuario string

func ProcesoToken(token string, JWTSign string) (*models.Claim, bool, string, error) {
	miClave := []byte(JWTSign)
	var claims models.Claim
	splitToken := strings.Split(token, "Bearer")
	if len(splitToken) != 2 {
		return &claims, false, string(""), errors.New("Formato de token inválido")
	}
	token = strings.TrimSpace(splitToken[1])
	tkn, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
		return miClave, nil
	})
	if err != nil {
		return &claims, false, string(""), err
	}
	if !tkn.Valid {
		return &claims, false, string(""), errors.New("Token inválido")
	}
	return &claims, true, string(""), nil
}
