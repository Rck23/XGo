package routers

import (
	"context"
	"encoding/json"
	"fmt"

	"XGo/bd"
	"XGo/models"
	"XGo/models/usuarios"
)

// Registro es un handler para el endpoint que permite a los usuarios registrarse en la aplicación.
// Recibe el contexto de la solicitud, extrae el cuerpo de la solicitud, deserializa los datos del nuevo
// usuario, realiza varias validaciones (como verificar que el correo electrónico y la contraseña cumplen
// con los requisitos mínimos), comprueba si el correo electrónico ya existe en la base de datos, e intenta
// insertar el nuevo registro de usuario en la base de datos. Finalmente, devuelve una respuesta adecuada
// indicando si la operación fue exitosa o no.
func Registro(ctx context.Context) models.ResApi {
	// Inicialización de variables
	var user usuarios.Usuario
	var r models.ResApi
	r.Status = 400 // Código de estado predeterminado para indicar un error potencial
	// Impresión de mensaje de entrada para depuración
	fmt.Println("Entré a Registro")

	// Extracción del cuerpo de la solicitud
	body := ctx.Value(models.Key("body")).(string)
	err := json.Unmarshal([]byte(body), &user)
	if err != nil {
		r.Message = err.Error() // Mensaje de error genérico
		fmt.Println(r.Message)
		return r
	}

	// Validaciones básicas
	if len(user.Email) == 0 {
		r.Message = "El correo electrónico es obligatorio"
		fmt.Println(r.Message)
		return r
	}
	if len(user.Password) < 8 {
		r.Message = "La contraseña debe tener al menos 8 caracteres"
		fmt.Println(r.Message)
		return r
	}

	// Chequeo de existencia del correo electrónico en la base de datos
	_, encontrado, _ := bd.ChequeoYaExisteUsuario(user.Email)
	if encontrado {
		r.Message = "Ya existe el correo electrónico en la base de datos"
		fmt.Println(r.Message)
		return r
	}

	// Intento de inserción del nuevo registro de usuario
	_, status, err := bd.InsertoRegistro(user)
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

	// Actualización de la respuesta con éxito
	r.Status = 200
	r.Message = "Registro Exitoso!"
	fmt.Println(r.Message)
	return r
}
