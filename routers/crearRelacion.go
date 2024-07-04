package routers

import (
	"XGo/bd"
	"XGo/models"
	"context"

	"github.com/aws/aws-lambda-go/events"
)

func CrearRelacion(ctx context.Context, request events.APIGatewayProxyRequest, claim models.Claim) models.ResApi {

	var respuesta models.ResApi
	respuesta.Status = 400 // Inicializa el estado de la respuesta a 400 (Bad Request).
	// Extrae el ID del usuario de la afirmación (claim).
	ID := request.QueryStringParameters["id"]
	if len(ID) < 1 { // Verifica si el ID está presente.
		respuesta.Message = "El Id es obligatorio" // Establece un mensaje de error si el ID no está presente.
		return respuesta                           // Retorna la respuesta con el mensaje de error.
	}

	var relacion models.Relacion
	relacion.UsuarioId = claim.ID.Hex()
	relacion.UsuarioRelacionId = ID

	estatus, err := bd.CrearRelacion(relacion)

	if err != nil {
		respuesta.Message = "Ocurrio un error al intentar crear la relacion. " + err.Error()
		return respuesta
	}

	if !estatus {
		respuesta.Message = "No se ha logrado crear la relacion."
		return respuesta
	}

	respuesta.Status = 200 // Establece el estado de la respuesta a 200 (OK).
	respuesta.Message = "Relacion creada correctamente."
	return respuesta
}
