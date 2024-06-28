package routers

import (
	"context"
	"encoding/json"
	"time"

	"XGo/bd"
	"XGo/models"
)

// CrearTweet maneja la creación de un nuevo tweet en la base de datos.
func CrearTweet(ctx context.Context, claim models.Claim) models.ResApi {
	// Inicializa variables locales para almacenar el mensaje del tweet y la respuesta.
	var mensaje models.Tweet

	var respuesta models.ResApi

	respuesta.Status = 400 // Inicializa el estado de la respuesta a 400 (Bad Request).

	// Extrae el ID del usuario de la afirmación (claim).
	IdUsuario := claim.ID.Hex()

	// Obtiene el cuerpo de la solicitud del contexto y lo deserializa en un objeto Tweet.
	body := ctx.Value(models.Key("body")).(string)
	err := json.Unmarshal([]byte(body), &mensaje)

	// Verifica si hubo un error durante la deserialización.
	if err != nil {
		// Si hay un error, establece un mensaje de error en la respuesta y retorna.
		respuesta.Message = "Ocurrio un error al intentar decodificar el body " + err.Error()
		return respuesta
	}

	// Crea un nuevo registro de tweet con el ID del usuario, el mensaje y la fecha/hora actual.
	registro := models.CrearTweet{
		UsuarioId: IdUsuario,
		Mensaje:   mensaje.Mensaje,
		Fecha:     time.Now(),
	}

	// Intenta insertar el nuevo tweet en la base de datos.
	_, estatus, err := bd.InsertarTweet(registro)
	// Verifica si hubo un error durante la inserción.
	if err != nil {
		// Si hay un error, establece un mensaje de error en la respuesta y retorna.
		respuesta.Message = "Ocurrio un error al insertar el registro " + err.Error()
		return respuesta
	}

	// Verifica si la inserción fue exitosa.
	if !estatus {
		// Si la inserción falló, establece un mensaje de error en la respuesta y retorna.
		respuesta.Message = "Error al crear tweet."
		return respuesta
	}

	// Si la inserción fue exitosa, configura el estado de la respuesta a 200 (OK) y establece un mensaje de éxito.
	respuesta.Status = 200
	respuesta.Message = "Tweet creado correctamente."
	return respuesta
}
