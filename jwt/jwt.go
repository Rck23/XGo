package jwt

import (
	"context"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"

	"XGo/models"
	"XGo/models/usuarios"
)

// GenerarToken crea un token JWT para un usuario específico, incluyendo información relevante del usuario en el payload.
// Este token puede ser utilizado posteriormente para autenticar solicitudes del usuario.
// La clave utilizada para firmar el token se obtiene del contexto.
func GenerarToken(ctx context.Context, usuario usuarios.Usuario) (string, error) {
	// Obtiene la clave utilizada para firmar el token del contexto.
	jwtSign := ctx.Value(models.Key("jwtsign")).(string)
	// Convierte la clave obtenida a un slice de bytes.
	miClave := []byte(jwtSign)

	// Prepara el payload del token con información del usuario.
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

	// Crea un nuevo token con el payload y la clave de firma.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	// Genera la representación en cadena del token firmado.
	tokenStr, err := token.SignedString(miClave)
	if err != nil {
		// Si ocurre un error durante la generación de la cadena del token, se devuelve el token y el error.
		return tokenStr, err
	}
	// Si la generación del token es exitosa, se devuelve el token y nil para el error.
	return tokenStr, nil
}
