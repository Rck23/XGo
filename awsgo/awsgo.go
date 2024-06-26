package awsgo

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

var Ctx context.Context
var Cfg aws.Config
var err error

// InicializoAWS inicializa la configuración de AWS SDK para Go v2.
// Configura el contexto global y carga la configuración predeterminada del cliente de AWS,
// especificando la región por defecto ("us-east-1").
func InicializoAWS() {
	// Establece el contexto global para las operaciones de AWS.
	Ctx = context.TODO()

	// Carga la configuración predeterminada del cliente de AWS, incluyendo la región por defecto.
	Cfg, err = config.LoadDefaultConfig(Ctx, config.WithDefaultRegion("us-east-1"))
	if err != nil {
		// Si ocurre un error al cargar la configuración, se detiene la ejecución del programa.
		panic("Error al cargar la configuración de .aws/config" + err.Error())
	}
}
