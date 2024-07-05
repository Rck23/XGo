package routers

import (
	"XGo/bd"
	"XGo/models"
	"encoding/json"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
)

func ObtenerUsuarios(request events.APIGatewayProxyRequest, claim models.Claim) models.ResApi {

	var res models.ResApi
	res.Status = 400

	pagina := request.QueryStringParameters["page"]
	tipoUsuario := request.QueryStringParameters["type"]
	buscar := request.QueryStringParameters["search"]
	IdUsuario := claim.ID.Hex()

	if len(pagina) == 0 {
		pagina = "1"
	}

	paginaTemp, err := strconv.Atoi(pagina)
	if err != nil {
		res.Message = "Debe enviar un número entero para la pagina " + err.Error()
		return res
	}

	usuarios, estatus := bd.VerUsuarios(IdUsuario, int64(paginaTemp), buscar, tipoUsuario)
	if !estatus {
		res.Message = "Error al obtener los usuarios"
		return res
	}

	resJson, err := json.Marshal(usuarios)
	if err != nil {
		res.Status = 500
		res.Message = "Error al formatear los datos de los usuarios como JSON " + err.Error()
		return res
	}

	res.Status = 200
	res.Message = string(resJson)
	return res
}
