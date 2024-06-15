package awsgo

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

// Ctx es un contexto global que se utiliza para las operaciones de AWS.
var Ctx context.Context

// Cfg es la configuración global de AWS que se utiliza para las operaciones de AWS.
var Cfg aws.Config

// err es un error global que se utiliza para manejar errores durante la inicialización de AWS.
var err error

// InicializoAWS inicializa la configuración de AWS necesaria para realizar operaciones con el SDK de AWS para Go v2.
// Esta función carga la configuración predeterminada del archivo.aws/config y establece el contexto y la configuración de AWS.
func InicializoAWS() {
	// Establecer el contexto a TODO, lo que significa que las operaciones no tienen un tiempo de espera.
	Ctx = context.TODO()

	// Cargar la configuración predeterminada del archivo.aws/config.
	// Se especifica la región por defecto como "us-east-1".
	Cfg, err = config.LoadDefaultConfig(Ctx, config.WithDefaultRegion("us-east-1"))

	// Si ocurre un error al cargar la configuración, se produce un pánico.
	// Esto detendrá la ejecución del programa y mostrará el mensaje de error.
	if err != nil {
		panic("Error al cargar la configuración de.aws/config" + err.Error())
	}
}
