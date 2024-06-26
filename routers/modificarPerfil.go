package routers

import (
	"XGo/bd"
	"XGo/models"
	"XGo/models/usuarios"
	"context"
	"encoding/json"
)

// ModificarPerfil es un handler para el endpoint que permite a los usuarios modificar su perfil.
// Recibe el contexto de la solicitud y las reclamaciones JWT (claim) que incluyen el ID del usuario
// autenticado. Extrae el cuerpo de la solicitud, deserializa los datos del perfil proporcionados por el
// usuario, y luego intenta modificar el registro del usuario en la base de datos. Finalmente, devuelve
// una respuesta adecuada indicando si la operación fue exitosa o no.
func ModificarPerfil(ctx context.Context, claim models.Claim) models.ResApi {
	// Inicialización de variables
	var respuesta models.ResApi
	respuesta.Status = 400 // Código de estado predeterminado para indicar un error potencial

	var user usuarios.Usuario

	// Extracción del cuerpo de la solicitud
	body := ctx.Value(models.Key("body")).(string)
	// Deserialización del cuerpo de la solicitud a la estructura Usuario
	err := json.Unmarshal([]byte(body), &user)
	if err != nil {
		respuesta.Message = "Datos incorrectos " + err.Error()
	}
	// Intento de modificación del registro del usuario
	status, err := bd.ModificoRegistro(user, claim.ID.Hex())
	if err != nil {
		respuesta.Message = "Ocurrio un error al intentar modificar el registro. " + err.Error()
		return respuesta
	}
	// Verificación de si la modificación fue exitosa
	if !status {
		respuesta.Message = "No se ha logrado modificar el registro del usuario."
		return respuesta
	}
	// Actualización de la respuesta con éxito
	respuesta.Status = 200
	respuesta.Message = "Modificación de perfil exitosa."
	return respuesta

}
