package handlers

import (
	"XGo/jwt"
	"XGo/models"
	"XGo/routers"

	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
)

// Manejadores procesa las solicitudes entrantes a la API Gateway basándose en el método HTTP y el path.
// Primero, valida la autorización utilizando la función validoAuthorization. Si la autorización es válida,
// procesa la solicitud según el método HTTP y el path especificados. Retorna una respuesta de tipo models.ResApi.
func Manejadores(ctx context.Context, request events.APIGatewayProxyRequest) models.ResApi {

	fmt.Println("Voy a procesar " + ctx.Value(models.Key("path")).(string) + " > " + ctx.Value(models.Key("method")).(string))

	// Inicializa la respuesta con un estado inicial de 400.
	var respuesta models.ResApi
	respuesta.Status = 400
	// Valida la autorización de la solicitud.
	isOk, statusCode, msg, claim := validoAuthorization(ctx, request)

	// Si la autorización no es válida, establece el estado y el mensaje de la respuesta y retorna.
	if !isOk {
		respuesta.Status = statusCode
		respuesta.Message = msg
		return respuesta
	}

	// Procesa la solicitud basándose en el método HTTP y el path.
	switch ctx.Value(models.Key("method")).(string) {
	case "POST":
		switch ctx.Value(models.Key("path")).(string) {
		case "registro":
			return routers.Registro(ctx)

		case "login":
			return routers.Login(ctx)

		case "crearTweet":
			return routers.CrearTweet(ctx, claim)

		case "subirAvatar":
			return routers.SubirImagen(ctx, request, "A", claim)

		case "subirBanner":
			return routers.SubirImagen(ctx, request, "B", claim)

		case "crearRelacion":
			return routers.CrearRelacion(ctx, request, claim)

		}
		//
	case "GET":
		switch ctx.Value(models.Key("path")).(string) {
		case "verPerfil":
			return routers.VerPerfil(request)

		case "leerTweets":
			return routers.LeerTweets(request)

		case "obtenerAvatar":
			return routers.ObtenerImagen(ctx, request, "A", claim)

		case "obtenerBanner":
			return routers.ObtenerImagen(ctx, request, "B", claim)

		case "consultarRelacion":
			return routers.ConsultarRelacion(request, claim)

		case "obtenerUsuarios":
			return routers.ObtenerUsuarios(request, claim)

		case "obtenerTweetsSeguidores":
			return routers.ObtenerTweetsSeguidores(request, claim)
		}

		//
	case "PUT":
		switch ctx.Value(models.Key("path")).(string) {
		case "modificarPerfil":
			return routers.ModificarPerfil(ctx, claim)
		}
		//
	case "DELETE":
		switch ctx.Value(models.Key("path")).(string) {
		case "eliminarTweet":
			return routers.EliminarTweet(request, claim)

		case "eliminarRelacion":
			return routers.EliminarRelacion(request, claim)
		}
		//
	}

	// Si el método o path no son válidos, establece un mensaje de error en la respuesta.
	respuesta.Message = "Metodo invalido"
	return respuesta
}

// validoAuthorization valida la autorización de una solicitud entrante.
// Verifica si el path es uno de los permitidos sin token o si el token de autorización es válido.
// Retorna un booleano indicando si la autorización es válida, un código de estado, un mensaje y los claims del token.
func validoAuthorization(ctx context.Context, request events.APIGatewayProxyRequest) (bool, int, string, models.Claim) {
	// Extrae el path de la solicitud.
	pathValue := ctx.Value(models.Key("path")).(string)

	// Permite ciertos paths sin token.
	if pathValue == "registro" || pathValue == "login" || pathValue == "obtenerAvatar" || pathValue == "obtenerBanner" {
		return true, 200, "", models.Claim{}
	}

	// Verifica si el token de autorización está presente.
	token := request.Headers["Authorization"]
	if len(token) == 0 {
		return false, 401, "Token requerido", models.Claim{}
	}

	// Procesa el token de autorización.
	Claim, todoOK, msg, err := jwt.ProcesoToken(token, ctx.Value(models.Key("jwtsign")).(string))

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
