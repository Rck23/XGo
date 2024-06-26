package secretmanager

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"

	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"

	"XGo/awsgo"
	"XGo/models"
)

// GetSecret recupera un secreto almacenado en AWS Secrets Manager basado en su nombre.
// Utiliza el SDK de AWS para hacer una llamada al servicio de Secrets Manager, obtener el valor
// del secreto asociado con el nombre proporcionado, y luego deserializa ese valor de JSON a una
// estructura Secret definida en el paquete models. Retorna el secreto deserializado y un error si
// ocurre alguno durante el proceso.
func GetSecret(SecretName string) (models.Secret, error) {
	// Inicialización de variables
	var datosSecret models.Secret

	fmt.Println("> Pidiendo secreto", SecretName)
	// Configuración del cliente de Secrets Manager
	svc := secretsmanager.NewFromConfig(awsgo.Cfg)

	// Llamada al servicio para obtener el valor del secreto
	clave, err := svc.GetSecretValue(awsgo.Ctx, &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(SecretName),
	})

	if err != nil {
		fmt.Println(err.Error())
		return datosSecret, err
	}

	// Deserialización del valor del secreto a la estructura Secret
	json.Unmarshal([]byte(*clave.SecretString), &datosSecret)
	fmt.Println("> Lectura de secreto OK", SecretName)

	// Retorno del secreto deserializado
	return datosSecret, nil
}
