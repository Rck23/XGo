package routers

import (
	"XGo/awsgo"
	"XGo/bd"
	"XGo/models"
	"bytes"
	"context"
	"fmt"
	"io"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go/aws"
)

// ObtenerImagen procesa una solicitud para obtener una imagen (Avatar o Banner) de un perfil de usuario.
// @param ctx Contexto de ejecución.
// @param request Solicitud entrante de API Gateway.
// @param subirTipo Tipo de imagen a obtener ('A' para Avatar, 'B' para Banner).
// @param claim Reclamo del usuario, utilizado para autenticación y autorización.
// @return Respuesta API con el contenido de la imagen solicitada.
func ObtenerImagen(ctx context.Context, request events.APIGatewayProxyRequest, subirTipo string, claim models.Claim) models.ResApi {
	// Inicializa la respuesta con un estado predeterminado de 400 (Bad Request).
	var respuesta models.ResApi
	respuesta.Status = 400

	ID := request.QueryStringParameters["id"] // Extrae el ID del usuario de los parámetros de consulta.
	if len(ID) < 1 {                          // Verifica si el ID está presente.
		respuesta.Message = "El Id es obligatorio" // Establece un mensaje de error si el ID no está presente.
		return respuesta                           // Retorna la respuesta con el mensaje de error.
	}

	perfil, err := bd.BuscoPerfil(ID) // Intenta buscar el perfil del usuario por ID.
	if err != nil {                   // Si hay un error (por ejemplo, el usuario no existe).
		respuesta.Message = "Usuario no encontrado " + err.Error() // Establece un mensaje de error.
		return respuesta                                           // Retorna la respuesta con el mensaje de error.
	}

	var nombreArchivo string // Variable para almacenar el nombre del archivo de imagen.
	switch subirTipo {       // Determina el tipo de imagen a obtener.
	case "A":
		nombreArchivo = perfil.Avatar // Asigna el nombre del Avatar.
	case "B":
		nombreArchivo = perfil.Banner // Asigna el nombre del Banner.
	}

	fmt.Println("Archivo: " + nombreArchivo) // Imprime el nombre del archivo para depuración.

	svc := s3.NewFromConfig(awsgo.Cfg) // Crea una nueva instancia del cliente S3.

	archivo, err := descargarDeS3(ctx, svc, nombreArchivo) // Descarga el archivo de S3.
	if err != nil {                                        // Si hay un error durante la descarga.
		respuesta.Status = 500                                               // Actualiza el estado de la respuesta a 500.
		respuesta.Message = "Error al descargar imagen de S3 " + err.Error() // Establece un mensaje de error.
		return respuesta                                                     // Retorna la respuesta con el mensaje de error.
	}

	respuesta.CustomResp = &events.APIGatewayProxyResponse{ // Prepara la respuesta personalizada para API Gateway.
		StatusCode: 200,
		Body:       archivo.String(), // Establece el cuerpo de la respuesta al contenido del archivo descargado.
		Headers: map[string]string{
			"Content-Type":        "application/octet-stream",                                   // Encabezado para indicar el tipo de contenido.
			"Content-Disposition": fmt.Sprintf("Archivo adjunto; nombre=\"%s\"", nombreArchivo), // Encabezado para indicar la disposición del archivo.
		},
	}

	return respuesta // Retorna la respuesta preparada.

}

// descargarDeS3 descarga un archivo de Amazon S3 y lo retorna en un Buffer de bytes.
// @param ctx Contexto de ejecución.
// @param svc Cliente S3 para interactuar con Amazon S3.
// @param nombreArchivo Nombre del archivo a descargar.
// @return Contenido del archivo en un Buffer de bytes y un error si ocurre alguno.
func descargarDeS3(ctx context.Context, svc *s3.Client, nombreArchivo string) (*bytes.Buffer, error) {

	// Extrae el nombre del bucket del contexto.
	bucket := ctx.Value(models.Key("bucketName")).(string)

	// Intenta obtener el objeto de S3 especificado.
	obj, err := svc.GetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(nombreArchivo),
	})

	if err != nil {
		return nil, err // Si hay un error, lo retorna inmediatamente.
	}

	// Asegura que el cuerpo del objeto se cierre después de usarlo.
	defer obj.Body.Close()
	fmt.Println("Bucket name = " + bucket) // Imprime el nombre del bucket para depuración.

	// Lee todo el contenido del cuerpo del objeto.
	archivo, err := io.ReadAll(obj.Body)

	if err != nil {
		// Si hay un error durante la lectura, lo retorna inmediatamente.
		return nil, err
	}
	// Crea un nuevo Buffer de bytes con el contenido del archivo.
	buffer := bytes.NewBuffer(archivo)

	// Retorna el Buffer de bytes con el contenido del archivo y un error nulo si todo salió bien.
	return buffer, nil
}
