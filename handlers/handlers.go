package handlers

import (
	"XGo/jwt"
	"XGo/models"
	"XGo/routers"

	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
)

// Manejadores procesa las solicitudes entrantes basándose en el método HTTP y la ruta.
// Actualmente, solo imprime el método y la ruta, pero se pueden añadir acciones específicas
// para cada combinación de método/ruta.
func Manejadores(ctx context.Context, request events.APIGatewayProxyRequest) models.ResApi {
	// Imprimir el método y la ruta de la solicitud para depuración.
	fmt.Println("Voy a procesar " + ctx.Value(models.Key("path")).(string) + " > " + ctx.Value(models.Key("method")).(string))

	var respuesta models.ResApi

	// Inicializar el estado de la respuesta a 400 (Bad Request).
	respuesta.Status = 400

	// Validar la autorización de la solicitud.
	isOk, statusCode, msg, claim := validoAuthorization(ctx, request)

	if !isOk {
		respuesta.Status = statusCode
		respuesta.Message = msg
		return respuesta
	}

	// Determinar el método de la solicitud y ejecutar acciones correspondientes.
	switch ctx.Value(models.Key("method")).(string) {
	case "POST":
		switch ctx.Value(models.Key("path")).(string) {
		case "registro":
			return routers.Registro(ctx)
		}
		//
	case "GET":
		switch ctx.Value(models.Key("path")).(string) {

		}
		//
	case "PUT":
		switch ctx.Value(models.Key("path")).(string) {

		}
		//
	case "DELETE":
		switch ctx.Value(models.Key("path")).(string) {

		}
		//
	}

	// Si el método no es reconocido, se devuelve un mensaje de error.
	respuesta.Message = "Metodo invalido"
	return respuesta
}

// validoAuthorization valida la autorización de las solicitudes entrantes mediante tokens JWT.
// Primero, verifica si la ruta requiere autorización. Si es así, extrae el token de autorización de los headers
// de la solicitud y lo valida. Si el token es válido, permite el acceso; de lo contrario, devuelve un error
// de autorización.
func validoAuthorization(ctx context.Context, request events.APIGatewayProxyRequest) (bool, int, string, models.Claim) {
	// Verificar si la ruta requiere autorización.
	path := ctx.Value(models.Key("path")).(string)

	if path == "registro" || path == "login" || path == "obtenerAvatar" || path == "obtenerBanner" {
		return true, 200, "", models.Claim{}
	}

	// Extraer el token de autorización de los headers de la solicitud.
	token := request.Headers["Authorization"]

	if len(token) == 0 {
		return false, 401, "Token requerido", models.Claim{}
	}

	// Procesar el token JWT.
	Claim, todoOK, msg, err := jwt.ProcesoToken(token, ctx.Value(models.Key("jwtSign")).(string))

	if !todoOK {
		if err != nil {
			fmt.Println("Error en el token " + err.Error())
			return false, 401, err.Error(), models.Claim{}
		} else {
			fmt.Println("Error en el token " + msg)
			return false, 401, msg, models.Claim{}
		}
	}

	fmt.Println("Token OK")
	return true, 200, msg, *Claim
}
