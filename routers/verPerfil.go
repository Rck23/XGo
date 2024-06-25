package routers

import (
	"encoding/json"
	"fmt"

	"XGo/bd"
	"XGo/models"

	"github.com/aws/aws-lambda-go/events"
)

func VerPerfil(request events.APIGatewayProxyRequest) models.ResApi {
	var respuesta models.ResApi
	respuesta.Status = 400

	fmt.Println("Entré en VerPerfil")
	ID := request.QueryStringParameters["id"]
	if len(ID) < 1 {
		respuesta.Message = "El ID es obligatorio"
		return respuesta
	}

	perfil, err := bd.BuscoPerfil(ID)
	if err != nil {
		respuesta.Message = "Ocurrio error al buscar registro " + err.Error()
		return respuesta
	}

	resJson, err := json.Marshal(perfil)
	if err != nil {
		respuesta.Status = 500
		respuesta.Message = "Error al formatear datos de usuario a JSON " + err.Error()
		return respuesta
	}

	respuesta.Status = 500
	respuesta.Message = string(resJson)
	return respuesta
}
