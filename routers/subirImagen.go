package routers

import (
	"XGo/bd"
	"XGo/models"
	"XGo/models/usuarios"
	"bytes"
	"context"
	"encoding/base64"
	"io"
	"mime"
	"mime/multipart"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type readSeeker struct {
	io.Reader
}

func (rs *readSeeker) Seek(offset int64, whence int) (int64, error) {
	return 0, nil
}

// SubirImagen es la función principal que maneja la subida de imágenes a un bucket de Amazon S3.
// Acepta un contexto, una solicitud API Gateway, un tipo de subida ('A' para avatares, 'B' para banners),
// y un objeto Claim que representa las afirmaciones del usuario autenticado.
func SubirImagen(ctx context.Context, request events.APIGatewayProxyRequest, subirTipo string, claim models.Claim) models.ResApi {
	// Inicializa la respuesta con un estado predeterminado de 400 (Bad Request).
	var respuesta models.ResApi
	respuesta.Status = 400
	// Obtiene el ID del usuario de las afirmaciones.
	IdUsuario := claim.ID.Hex()
	// Variables para almacenar el nombre del archivo y el objeto usuario.
	var nombreArchivo string
	var usuario usuarios.Usuario
	// Obtiene el nombre del bucket de S3 desde el contexto.
	bucket := aws.String(ctx.Value(models.Key("bucketName")).(string))
	// Construye el nombre del archivo basado en el tipo de subida y el ID del usuario.
	switch subirTipo {
	case "A":
		nombreArchivo = "avatars/" + IdUsuario + ".jpg"
		usuario.Avatar = nombreArchivo

	case "B":
		nombreArchivo = "banners/" + IdUsuario + ".jpg"
		usuario.Banner = nombreArchivo
	}
	// Intenta parsear el tipo de medio y los parámetros del encabezado 'Content-Type'.
	mediaTipo, parametros, err := mime.ParseMediaType(request.Headers["Content-Type"])
	// Si hay un error al parsear el 'Content-Type', se devuelve una respuesta con estado 500.
	if err != nil {
		respuesta.Status = 500
		respuesta.Message = err.Error()
		return respuesta
	}
	// Verifica si el tipo de medio comienza con 'multipart/', indicando un formulario multipart.
	if strings.HasPrefix(mediaTipo, "multipart/") {
		// Decodifica el cuerpo de la solicitud si es necesario.
		body, err := base64.StdEncoding.DecodeString(request.Body)
		if err != nil {
			respuesta.Status = 500
			respuesta.Message = err.Error()
			return respuesta
		}
		// Crea un nuevo lector multipart a partir del cuerpo decodificado.
		mr := multipart.NewReader(bytes.NewReader(body), parametros["boundary"])
		p, err := mr.NextPart()
		// Si hay un error al leer la siguiente parte, se devuelve una respuesta con estado 500.
		if err != nil && err != io.EOF {
			respuesta.Status = 500
			respuesta.Message = err.Error()
			return respuesta
		}
		// Si no se encuentra EOF, significa que hay al menos una parte.
		if err != io.EOF {
			if p.FileName() != "" {
				buf := bytes.NewBuffer(nil)
				// Copia el contenido de la parte de la imagen al buffer.
				if _, err := io.Copy(buf, p); err != nil {
					respuesta.Status = 500
					respuesta.Message = err.Error()
					return respuesta
				}
				// Crea una nueva sesión de AWS.
				sesion, err := session.NewSession(&aws.Config{
					Region: aws.String("us-east-1")})
				// Si hay un error al crear la sesión, se devuelve una respuesta con estado 500.
				if err != nil {
					respuesta.Status = 500
					respuesta.Message = err.Error()
					return respuesta
				}
				// Configura un uploader para S3.
				cargador := s3manager.NewUploader(sesion)
				// Intenta cargar la imagen al bucket de S3.
				_, err = cargador.Upload(&s3manager.UploadInput{
					Bucket: bucket,
					Key:    aws.String(nombreArchivo),
					Body:   &readSeeker{buf},
				})
				if err != nil {
					respuesta.Status = 500
					respuesta.Message = err.Error()
					return respuesta
				}

			}
		}
		// Actualiza el registro del usuario con la nueva URL de la imagen.
		estatus, err := bd.ModificoRegistro(usuario, IdUsuario)
		if err != nil || !estatus {
			respuesta.Status = 400
			respuesta.Message = "Error al modificar usuario " + err.Error()
			return respuesta
		}

	} else {
		// Si el 'Content-Type' no es 'multipart/', se devuelve una respuesta con estado 400.
		respuesta.Status = 400
		respuesta.Message = "Debe enviar una imagen con el 'Content-Type' 'multipart/' en el Header"
		return respuesta
	}
	// Si todo ha ido bien, se devuelve una respuesta con estado 200 y un mensaje de éxito.
	respuesta.Status = 200
	respuesta.Message = "Imagen subida exitosamente."
	return respuesta
}
