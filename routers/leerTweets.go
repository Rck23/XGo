package routers

import (
	"XGo/bd"
	"XGo/models"
	"encoding/json"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
)

// LeerTweets maneja las solicitudes HTTP para leer tweets de una base de datos MongoDB.
func LeerTweets(request events.APIGatewayProxyRequest) models.ResApi {

	// Inicializa la respuesta con un estado inicial de 400 (Bad Request).
	var respuesta models.ResApi
	respuesta.Status = 400

	// Extrae el ID del usuario y la página de la solicitud.
	Id := request.QueryStringParameters["id"]
	pagina := request.QueryStringParameters["pagina"]

	// Verifica si el ID del usuario está presente y es válido.
	if len(Id) < 1 {
		respuesta.Message = "El Id es obligatorio" // Establece un mensaje de error si el ID no está presente.
		return respuesta                           // Retorna la respuesta con el mensaje de error.
	}

	// Si la página no está presente, asume la primera página.
	if len(pagina) < 1 {
		pagina = "1"
	}

	// Convierte la página de una cadena a un entero.
	pag, err := strconv.Atoi(pagina)
	if err != nil {
		respuesta.Message = "Debe enviar el parámetro Página con valor mayor a 0"
		return respuesta
	}

	// Intenta leer los tweets del usuario y la página especificados.
	tweets, correcto := bd.LeerTweets(Id, int64(pag))

	// Verifica si la lectura de tweets fue exitosa.
	if !correcto {
		respuesta.Message = "Error al leer los tweets"
		return respuesta
	}

	// Intenta serializar los tweets a JSON.
	resJson, err := json.Marshal(tweets)
	if err != nil {
		respuesta.Status = 500
		respuesta.Message = "Error al formatear los datos de los usuarios como JSON"
		return respuesta
	}

	// Prepara la respuesta final con el estado 200 (OK) y los tweets en formato JSON.
	respuesta.Status = 200
	respuesta.Message = string(resJson)
	return respuesta // Retorna la respuesta final.
}
