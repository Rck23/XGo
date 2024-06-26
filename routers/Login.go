package routers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"XGo/bd"
	"XGo/jwt"
	"XGo/models"
	"XGo/models/usuarios"

	"github.com/aws/aws-lambda-go/events"
)

// Login es un handler para el endpoint de inicio de sesión. Procesa las credenciales enviadas por el cliente,
// verifica si el usuario existe y si las credenciales son correctas, genera un token JWT para el usuario
// autenticado, y devuelve una respuesta adecuada. La respuesta incluye el token JWT y se envía al cliente
// a través de una cookie para su uso en solicitudes futuras.
func Login(ctx context.Context) models.ResApi {
	// Inicialización de variables
	var usuario usuarios.Usuario
	var respuesta models.ResApi

	respuesta.Status = 400 // Código de estado predeterminado para indicar un error potencial

	// Extracción del cuerpo de la solicitud
	body := ctx.Value(models.Key("body")).(string)

	// Deserialización del cuerpo de la solicitud a la estructura Usuario
	err := json.Unmarshal([]byte(body), &usuario)
	if err != nil {
		respuesta.Message = "Credenciales incorrectas " + err.Error()
		return respuesta
	}
	// Verificación de que el correo electrónico no esté vacío
	if len(usuario.Email) == 0 {
		respuesta.Message = "El correo electrónico es requerido"
		return respuesta
	}

	// Intento de login con las credenciales proporcionadas
	usuarioData, existe := bd.IntentoLogin(usuario.Email, usuario.Password)
	if !existe {
		respuesta.Message = "Credenciales incorrectas "
		return respuesta
	}
	// Generación del token JWT para el usuario autenticado
	jwtKey, err := jwt.GenerarToken(ctx, usuarioData)
	if err != nil {
		respuesta.Message = "Ocurrio un error al intentar generar el token > " + err.Error()
		return respuesta
	}
	// Preparación de la respuesta con el token
	resp := models.RespuestaLogin{
		Token: jwtKey,
	}
	// Formateo del token como JSON
	token, err2 := json.Marshal(resp)
	if err2 != nil {
		respuesta.Message = "Ocurrio un error al intentar formatear el token > " + err2.Error()
		return respuesta
	}
	// Creación de la cookie con el token
	cookie := &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: time.Now().Add(time.Hour * 24),
	}

	cookieString := cookie.String()

	// Preparación de la respuesta HTTP
	res := &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(token),
		Headers: map[string]string{
			"Content-Type":                "application/json",
			"Access-Control-Allow-Origin": "*",
			"Set-Cookie":                  cookieString,
		},
	}

	// Actualización de la respuesta con éxito
	respuesta.Status = 200
	respuesta.Message = string(token)
	respuesta.CustomResp = res

	return respuesta

}
