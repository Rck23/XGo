package routers

import (
	"XGo/bd"
	"XGo/models"

	"github.com/aws/aws-lambda-go/events"
)

func EliminarRelacion(request events.APIGatewayProxyRequest, claim models.Claim) models.ResApi {

	var respuesta models.ResApi
	respuesta.Status = 400

	ID := request.QueryStringParameters["id"]
	if len(ID) < 1 {
		respuesta.Message = "El Id es obligatorio"
		return respuesta
	}

	var relacion models.Relacion
	relacion.UsuarioId = claim.ID.Hex()
	relacion.UsuarioRelacionId = ID

	estatus, err := bd.BorrarRelacion(relacion)
	if err != nil {
		respuesta.Message = "Ocurrio un error al intentar borrar la relacion. " + err.Error()
		return respuesta
	}

	if !estatus {
		respuesta.Message = "No se ha logrado borrar la relacion."
		return respuesta
	}

	respuesta.Status = 200
	respuesta.Message = "Relacion borrada con exito"
	return respuesta
}
