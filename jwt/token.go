package jwt

import (
	"errors"
	"strings"

	jwt "github.com/golang-jwt/jwt/v5"

	"XGo/models"
)

var Email string
var IDUsuario string

// ProcesoToken toma un token JWT y una clave secreta (JWTSign) como entrada.
// Analiza el token, verifica su validez, y decodifica los claims contenidos en él.
// Devuelve los claims decodificados, un booleano indicando si el token es válido, un mensaje de error si corresponde,
// y un error si ocurre algún otro problema durante el proceso.
func ProcesoToken(token string, JWTSign string) (*models.Claim, bool, string, error) {
	// Define la clave secreta para verificar la firma del token.
	miClave := []byte(JWTSign)

	// Prepara una variable para almacenar los claims decodificados del token.
	var claims models.Claim

	// Divide el token por el prefijo "Bearer " para obtener solo el token sin el tipo de bearer.
	splitToken := strings.Split(token, "Bearer")
	if len(splitToken) != 2 {
		// Si el token no sigue el formato esperado, se devuelve un error.
		return &claims, false, string(""), errors.New("Formato de token inválido")
	}

	// Elimina espacios en blanco al principio y al final del token para limpiarlo.
	token = strings.TrimSpace(splitToken[1])

	// Intenta parsear el token con los claims y la clave secreta.
	tkn, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
		// Retorna la clave secreta para verificar la firma del token.
		return miClave, nil
	})

	// Si ocurre un error durante el parsing del token, se devuelve el error.
	if err != nil {
		return &claims, false, string(""), err
	}

	// Verifica si el token es válido.
	if !tkn.Valid {
		// Si el token no es válido, se devuelve un error.
		return &claims, false, string(""), errors.New("Token inválido")
	}

	// Si el token es válido, se devuelve el objeto claims decodificado.
	return &claims, true, string(""), nil
}
