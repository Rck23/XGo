package routers

import (
	"context"
	"encoding/json"
	"fmt"

	"XGo/bd"
	"XGo/models"
	"XGo/models/usuarios"
)

func Registro(ctx context.Context) models.ResApi {
	var t usuarios.Usuario
	var r models.ResApi
	r.Status = 400

	fmt.Println("Entré a Registro")

	body := ctx.Value(models.Key("body")).(string)
	err := json.Unmarshal([]byte(body), &t)
	if err != nil {
		r.Message = err.Error()
		fmt.Println(r.Message)
		return r
	}
	if len(t.Email) == 0 {
		r.Message = "El correo electrónico es obligatorio"
		fmt.Println(r.Message)
		return r
	}
	if len(t.Password) < 8 {
		r.Message = "La contraseña debe tener al menos 8 caracteres"
		fmt.Println(r.Message)
		return r
	}
	_, encontrado, _ := bd.ChequeoYaExisteUsuario(t.Email)
	if encontrado {
		r.Message = "Ya existe el correo electrónico en la base de datos"
		fmt.Println(r.Message)
		return r
	}
	_, status, err := bd.InsertoRegistro(t)
	if err != nil {
		r.Message = "Ocurrió un error al intentar realizar el registro del usuario " + err.Error()
		fmt.Println(r.Message)
		return r
	}
	if !status {
		r.Message = "No se ha logrado insertar el registro de usuario"
		fmt.Println(r.Message)
		return r
	}
	r.Status = 200
	r.Message = "Registro Exitoso!"
	fmt.Println(r.Message)
	return r
}
