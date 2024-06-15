package secretmanager

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"

	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"

	"XGo/awsgo"
	"XGo/models"
)

// GetSecret busca y recupera un secreto almacenado en AWS Secrets Manager.
// Devuelve el secreto deserializado como un objeto de tipo models.Secret y un error si ocurre alguno.
func GetSecret(SecretName string) (models.Secret, error) {
	// Variable para almacenar los datos del secreto obtenidos.
	var datosSecret models.Secret
	fmt.Println("> Pidiendo secreto", SecretName)

	// Crear un nuevo cliente de Secrets Manager usando la configuración global de AWS.
	svc := secretsmanager.NewFromConfig(awsgo.Cfg)

	// Realizar la solicitud para obtener el valor del secreto.
	// Se pasa el contexto global y el ID del secreto como entrada.
	clave, err := svc.GetSecretValue(awsgo.Ctx, &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(SecretName),
	})

	// Si ocurre un error al obtener el valor del secreto, se imprime el error y se retorna.
	if err != nil {
		fmt.Println(err.Error())
		return datosSecret, err
	}

	// Deserializar el contenido del secreto (en formato JSON) en la estructura models.Secret.
	json.Unmarshal([]byte(*clave.SecretString), &datosSecret)
	fmt.Println("> Lectura de secreto OK", SecretName)

	// Retornar el secreto deserializado y un error nulo si todo fue exitoso.
	return datosSecret, nil
}
