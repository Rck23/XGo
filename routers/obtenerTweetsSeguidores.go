package routers

import (
	"XGo/bd"
	"XGo/models"
	"encoding/json"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
)

func ObtenerTweetsSeguidores(request events.APIGatewayProxyRequest, claim models.Claim) models.ResApi {
	var res models.ResApi
	res.Status = 400

	IdUsuario := claim.ID.Hex()

	pagina := request.QueryStringParameters["pagina"]
	if len(pagina) < 1 {
		pagina = "1"
	}

	pag, err := strconv.Atoi(pagina)
	if err != nil {
		res.Message = "La página debe ser mayor a 0 "
		return res
	}

	tweets, correcto := bd.VerTweetsSeguidores(IdUsuario, pag)
	if !correcto {
		res.Message = "Error al leer los tweets seguidores"
		return res
	}

	resJson, err := json.Marshal(tweets)
	if err != nil {
		res.Status = 500
		res.Message = "Error al formatear los datos de los tweets seguidores como JSON" + err.Error()
		return res
	}

	res.Status = 200
	res.Message = string(resJson)
	return res
}
