package routers

import (
	"encoding/json"
	"fmt"

	"XGo/bd"
	"XGo/models"

	"github.com/aws/aws-lambda-go/events"
)

// VerPerfil es un handler para el endpoint que permite visualizar el perfil de un usuario específico.
// Recibe la solicitud de API Gateway, extrae el ID del usuario de los parámetros de consulta, busca el
// perfil correspondiente en la base de datos utilizando el ID proporcionado, y luego devuelve el perfil
// en formato JSON. Si ocurre algún error durante la búsqueda o la formación del JSON, se devuelve una
// respuesta adecuada indicando el problema.
func VerPerfil(request events.APIGatewayProxyRequest) models.ResApi {
	// Inicialización de variables
	var respuesta models.ResApi
	respuesta.Status = 400 // Código de estado predeterminado para indicar un error potencial
	// Impresión de mensaje de entrada para depuración
	fmt.Println("Entré en VerPerfil")
	// Extracción del ID del usuario de los parámetros de consulta
	ID := request.QueryStringParameters["id"]
	// Verificación de que el ID no esté vacío
	if len(ID) < 1 {
		respuesta.Message = "El ID es obligatorio"
		return respuesta
	}

	// Búsqueda del perfil del usuario en la base de datos
	perfil, err := bd.BuscoPerfil(ID)
	if err != nil {
		respuesta.Message = "Ocurrio error al buscar registro " + err.Error()
		return respuesta
	}
	// Formateo del perfil a JSON
	resJson, err := json.Marshal(perfil)
	if err != nil {
		respuesta.Status = 500
		respuesta.Message = "Error al formatear datos de usuario a JSON " + err.Error()
		return respuesta
	}
	// Actualización de la respuesta con éxito
	respuesta.Status = 200
	respuesta.Message = string(resJson)
	return respuesta
}
