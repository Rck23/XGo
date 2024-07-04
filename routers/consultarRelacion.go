package routers

import (
	"XGo/bd"
	"XGo/models"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
)

func ConsultarRelacion(request events.APIGatewayProxyRequest, claim models.Claim) models.ResApi {
	var res models.ResApi
	res.Status = 400

	ID := request.QueryStringParameters["id"]
	if len(ID) < 1 {
		res.Message = "El Id es obligatorio"
		return res
	}

	var relacion models.Relacion
	relacion.UsuarioId = claim.ID.Hex()
	relacion.UsuarioRelacionId = ID

	var respuesta models.RespuestaConsultarRelacion

	hayRelacion := bd.ConsultarRelacion(relacion)
	if hayRelacion {
		respuesta.Estatus = true
	} else {
		respuesta.Estatus = false
	}

	respuestaJson, err := json.Marshal(hayRelacion)
	if err != nil {
		res.Status = 500
		res.Message = "Error al formatear los datos de los usuarios como JSON " + err.Error()
		return res
	}

	res.Status = 200
	res.Message = string(respuestaJson)
	return res
}
