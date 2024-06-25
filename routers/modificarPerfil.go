package routers

import (
	"XGo/bd"
	"XGo/models"
	"XGo/models/usuarios"
	"context"
	"encoding/json"
)

func ModificarPerfil(ctx context.Context, claim models.Claim) models.ResApi {

	var respuesta models.ResApi
	respuesta.Status = 400

	var user usuarios.Usuario

	body := ctx.Value(models.Key("body")).(string)

	err := json.Unmarshal([]byte(body), &user)
	if err != nil {
		respuesta.Message = "Datos incorrectos " + err.Error()
	}

	status, err := bd.ModificoRegistro(user, claim.ID.Hex())
	if err != nil {
		respuesta.Message = "Ocurrio un error al intentar modificar el registro. " + err.Error()
		return respuesta
	}

	if !status {
		respuesta.Message = "No se ha logrado modificar el registro del usuario."
		return respuesta
	}

	respuesta.Status = 200
	respuesta.Message = "Modificación de perfil exitosa."
	return respuesta

}
