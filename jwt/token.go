package jwt

import (
	"errors"
	"strings"

	jwt "github.com/golang-jwt/jwt/v5"

	"XGo/bd"
	"XGo/models"
)

var Email string
var IDUsuario string

// ProcesoToken analiza y valida un token JWT proporcionado, extrayendo y verificando sus claims.
// También busca si el email asociado al token corresponde a un usuario registrado en la base de datos.
// Retorna los claims extraídos, un booleano indicando si el usuario fue encontrado, el ID del usuario encontrado, y un error si ocurrió alguno.
func ProcesoToken(tk string, JWTSign string) (*models.Claim, bool, string, error) {
	// Obtiene la clave de firma del token del argumento proporcionado.
	miClave := []byte(JWTSign)

	// Inicializa un struct vacío para almacenar los claims extraídos del token.
	var claims models.Claim
	// Divide el token por el prefijo "Bearer" para obtener solo el token.
	splitToken := strings.Split(tk, "Bearer")
	// Verifica si el token tiene el formato correcto (contiene exactamente dos partes: el prefijo y el token).
	if len(splitToken) != 2 {
		// Retorna un error si el formato del token es incorrecto.
		return &claims, false, string(""), errors.New("Formato de token inválido")
	}

	// Elimina espacios en blanco al principio y al final del token.
	tk = strings.TrimSpace(splitToken[1])
	// Analiza el token con los claims esperados y la clave de firma.
	tkn, err := jwt.ParseWithClaims(tk, &claims, func(token *jwt.Token) (interface{}, error) {
		// Retorna la clave de firma para verificar el token.
		return miClave, nil
	})

	// Si no hay error al analizar el token, verifica si el usuario asociado al email del token existe en la base de datos.
	if err == nil {

		_, encontrado, _ := bd.ChequeoYaExisteUsuario(claims.Email)
		// Si el usuario es encontrado, guarda el email y el ID del usuario.
		if encontrado {
			Email = claims.Email
			IDUsuario = claims.ID.Hex()
		}
		// Retorna los claims, un booleano indicando si el usuario fue encontrado, y el ID del usuario.
		return &claims, encontrado, IDUsuario, nil

	}
	// Si el token no es válido, retorna un error.
	if !tkn.Valid {
		return &claims, false, string(""), errors.New("Token inválido")
	}
	// Retorna los claims y un error si el token no pudo ser analizado correctamente.
	return &claims, true, string(""), nil
}
