package main

import (
	"context"
	"os"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	lambda "github.com/aws/aws-lambda-go/lambda"

	"XGo/awsgo"
	"XGo/bd"
	"XGo/handlers"
	"XGo/models"
	"XGo/secretmanager"
)

// main inicia la aplicación Lambda, definiendo la función EjecutoLambda como el punto de entrada.
func main() {
	// Iniciar la aplicación Lambda con la función EjecutoLambda.
	lambda.Start(EjecutoLambda)
}

// EjecutoLambda es la función principal que maneja las solicitudes entrantes a la aplicación Lambda.
// Procesa la solicitud, inicializa la configuración de AWS, valida parámetros de entorno, obtiene un secreto,
// se conecta a una base de datos, y finalmente delega el procesamiento de la solicitud a los manejadores apropiados.
func EjecutoLambda(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {

	// Configuración inicial y validación de parámetros.

	var res *events.APIGatewayProxyResponse

	awsgo.InicializoAWS()

	if !ValidoParametros() {
		res = &events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       "Error en las variables de entorno.",
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}
		return res, nil
	}

	// Obtener un secreto de AWS Secrets Manager.
	SecretModel, err := secretmanager.GetSecret(os.Getenv("SecretName"))

	if err != nil {
		res = &events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       "Error en la lectura secret. " + err.Error(),
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}
		return res, nil
	}

	// Preparar el contexto con valores relevantes para el procesamiento de la solicitud.
	path := strings.Replace(request.PathParameters["xgo"], os.Getenv("UrlPrefix"), "", -1)

	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("path"), path)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("method"), request.HTTPMethod)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("user"), SecretModel.Username)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("password"), SecretModel.Password)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("host"), SecretModel.Host)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("database"), SecretModel.Database)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("jwtsign"), SecretModel.JWTSign)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("body"), request.Body)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("bucketName"), os.Getenv("BucketName"))

	// Conectar a la base de datos.
	err = bd.ConectarDB(awsgo.Ctx)

	if err != nil {
		res = &events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "Error conectando BD. " + err.Error(),
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}
		return res, nil
	}

	// Delegar el procesamiento de la solicitud a los manejadores apropiados.
	resAPI := handlers.Manejadores(awsgo.Ctx, request)
	if resAPI.CustomResp == nil {
		res = &events.APIGatewayProxyResponse{
			StatusCode: resAPI.Status,
			Body:       resAPI.Message,
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}
		return res, nil
	} else {
		return resAPI.CustomResp, nil
	}
}

// ValidoParametros verifica si ciertas variables de entorno están definidas.
// Retorna true si todas las variables requeridas están presentes, false en caso contrario.
func ValidoParametros() bool {

	_, traeParametro := os.LookupEnv("SecretName")
	if !traeParametro {
		return traeParametro
	}

	_, traeParametro = os.LookupEnv("BucketName")
	if !traeParametro {
		return traeParametro
	}

	_, traeParametro = os.LookupEnv("UrlPrefix")
	if !traeParametro {
		return traeParametro
	}

	return traeParametro
}
