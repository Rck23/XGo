package models

import "github.com/aws/aws-lambda-go/events"

// ResApi representa la respuesta personalizada de una API Gateway Lambda.
// Es útil cuando necesitamos devolver datos específicos además de los estándares de respuesta HTTP.
type ResApi struct {
	// Status indica el estado de la respuesta HTTP.
	// Un valor de 200 indica éxito, mientras que otros códigos indican diferentes tipos de errores.
	Status int

	// Message es un mensaje de texto opcional que puede ser utilizado para describir el resultado de la operación.
	Message string

	// CustomResp permite especificar una respuesta personalizada utilizando la estructura APIGatewayProxyResponse de AWS Lambda.
	// Esto es útil para devolver datos en formatos específicos o para manejar casos donde necesitamos enviar respuestas no estándar.
	CustomResp *events.APIGatewayProxyResponse
}
