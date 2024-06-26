package models

import "github.com/aws/aws-lambda-go/events"

// ResApi representa la estructura de una respuesta generada por la API. Esta estructura es utilizada
// para enviar respuestas a las solicitudes HTTP realizadas a la API. Proporciona un formato consistente
// para las respuestas, facilitando la lectura y el procesamiento tanto en el lado del cliente como del servidor.
type ResApi struct {
	Status  int
	Message string

	// CustomResp es un puntero a una estructura de tipo events.APIGatewayProxyResponse. Esta estructura
	// específica de AWS Lambda (cuando se utiliza API Gateway con Lambda) permite un control más fino
	// sobre la respuesta, incluyendo encabezados personalizados, cuerpo de respuesta, y códigos de estado
	// específicos de HTTP. Este campo es opcional y solo debe ser utilizado cuando se necesita un control
	// detallado sobre la respuesta HTTP, como en casos especiales donde se requieren encabezados HTTP
	// personalizados o formatos de respuesta específicos.
	CustomResp *events.APIGatewayProxyResponse
}
