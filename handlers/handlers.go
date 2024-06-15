package handlers

import (
	"XGo/models"
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
)

// Manejadores procesa las solicitudes entrantes basándose en el método HTTP y la ruta.
// Actualmente, solo imprime el método y la ruta, pero se pueden añadir acciones específicas
// para cada combinación de método/ruta.
func Manejadores(ctx context.Context, request events.APIGatewayProxyRequest) models.ResApi {
	// Imprimir el método y la ruta de la solicitud para depuración.
	fmt.Println("Voy a procesar " + ctx.Value(models.Key("path")).(string) + " > " + ctx.Value(models.Key("method")).(string))

	var respuesta models.ResApi

	// Inicializar el estado de la respuesta a 400 (Bad Request).
	respuesta.Status = 400

	// Determinar el método de la solicitud y ejecutar acciones correspondientes.
	switch ctx.Value(models.Key("method")).(string) {
	case "POST":
		switch ctx.Value(models.Key("path")).(string) {

		}
		//
	case "GET":
		switch ctx.Value(models.Key("path")).(string) {

		}
		//
	case "PUT":
		switch ctx.Value(models.Key("path")).(string) {

		}
		//
	case "DELETE":
		switch ctx.Value(models.Key("path")).(string) {

		}
		//
	}

	// Si el método no es reconocido, se devuelve un mensaje de error.
	respuesta.Message = "Metodo invalido"
	return respuesta
}
