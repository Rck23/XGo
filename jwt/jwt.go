package jwt

import (
	"context"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"

	"XGo/models"
	"XGo/models/usuarios"
)

func GenerarToken(ctx context.Context, usuario usuarios.Usuario) (string, error) {
	jwtSign := ctx.Value(models.Key("jwtsign")).(string)

	miClave := []byte(jwtSign)

	payload := jwt.MapClaims{
		"email":            usuario.Email,
		"nombre":           usuario.Nombre,
		"apellidos":        usuario.Apellidos,
		"fecha_nacimiento": usuario.FechaNacimiento,
		"biografia":        usuario.Biografia,
		"ubicacion":        usuario.Ubicacion,
		"sitioweb":         usuario.SitioWeb,
		"_id":              usuario.ID.Hex(),
		"exp":              time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	tokenStr, err := token.SignedString(miClave)
	if err != nil {
		return tokenStr, err
	}

	return tokenStr, nil
}
