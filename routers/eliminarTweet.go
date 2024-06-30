package routers

import (
	"XGo/bd"
	"XGo/models"

	"github.com/aws/aws-lambda-go/events"
)

// EliminarTweet es una función que procesa una solicitud para eliminar un tweet basado en su ID.
// Recibe una solicitud HTTP a través de API Gateway y un objeto Claim que contiene información sobre el usuario autenticado.
// Retorna una respuesta formateada según los modelos de Respuesta y Reclamación.
func EliminarTweet(request events.APIGatewayProxyRequest, claim models.Claim) models.ResApi {
	// Inicializamos una variable para almacenar la respuesta que será enviada al cliente.
	var respuesta models.ResApi
	respuesta.Status = 400 // Inicializa el estado de la respuesta a 400 (Bad Request).
	// Extraemos el ID del tweet de los parámetros de consulta de la solicitud.
	ID := request.QueryStringParameters["id"]

	// Verificamos si el ID proporcionado está vacío.
	if len(ID) < 1 {
		// Si el ID está vacío, establecemos un mensaje de error y retornamos la respuesta con el estado 400.
		respuesta.Message = "El ID es obligatorio"
		return respuesta
	}

	// Intentamos eliminar el tweet utilizando el ID extraído y el ID del usuario autenticado.
	err := bd.BorrarTweet(ID, claim.ID.Hex())
	// Si ocurre un error durante la eliminación, establecemos un mensaje de error y retornamos la respuesta con el estado 400.
	if err != nil {
		respuesta.Message = "Error al borrar tweet: " + err.Error()
		return respuesta
	}
	// Si la eliminación fue exitosa, establecemos el estado de la respuesta a 200 (OK) y configuramos un mensaje de éxito.
	respuesta.Status = 200
	respuesta.Message = "Tweet eliminado "
	// Finalmente, retornamos la respuesta formateada.
	return respuesta
}
