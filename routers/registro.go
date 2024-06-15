package routers

import (
	"context"
	"encoding/json"
	"fmt"

	"XGo/bd"
	"XGo/models"
	"XGo/models/usuarios"
)

// Registro maneja el proceso de registro de nuevos usuarios en la aplicación.
// Toma un contexto como entrada, que debe contener el cuerpo de la solicitud en formato JSON.
// Deserializa el cuerpo de la solicitud en una estructura Usuario, valida los datos, y realiza operaciones en la base de datos.
func Registro(ctx context.Context) models.ResApi {
	// Variable para almacenar los datos del usuario obtenidos del cuerpo de la solicitud.
	var t usuarios.Usuario
	// Variable para almacenar la respuesta que será devuelta.
	var r models.ResApi
	r.Status = 400 // Inicializar el estado de la respuesta a 400 (Bad Request).

	fmt.Println("Entré a Registro")

	// Extraer el cuerpo de la solicitud del contexto.
	body := ctx.Value(models.Key("body")).(string)
	// Deserializar el cuerpo de la solicitud en una estructura Usuario.
	err := json.Unmarshal([]byte(body), &t)
	if err != nil {
		// Si hay un error al deserializar, establecer el mensaje de error en la respuesta.
		r.Message = err.Error()
		fmt.Println(r.Message)
		return r
	}

	// Validar que el correo electrónico esté presente.
	if len(t.Email) == 0 {
		r.Message = "Debe especificar el Email"
		fmt.Println(r.Message)
		return r
	}

	// Validar que la contraseña tenga al menos 8 caracteres.
	if len(t.Password) < 8 {
		r.Message = "La contraseña debe tener al menos 8 caracteres"
		fmt.Println(r.Message)
		return r
	}

	// Chequear si el usuario ya existe en la base de datos.
	_, encontrado, _ := bd.ChequeoYaExisteUsuario(t.Email)
	if encontrado {
		r.Message = "Ya existe usuario registro con ese email"
		fmt.Println(r.Message)
		return r
	}

	// Intentar insertar el nuevo registro en la base de datos.
	_, status, err := bd.InsertoRegistro(t)
	if err != nil {
		// Si hay un error al insertar, establecer el mensaje de error en la respuesta.
		r.Message = "Ocurrió un error al intentar realizar el registro del usuario " + err.Error()
		fmt.Println(r.Message)
		return r
	}

	// Verificar si la inserción fue exitosa.
	if !status {
		r.Message = "No se ha logrado insertar el registro de usuario"
		fmt.Println(r.Message)
		return r
	}

	// Si todo salió bien, establecer el estado de la respuesta a 200 (OK) y enviar un mensaje de éxito.
	r.Status = 200
	r.Message = "Registro Exitoso!"
	fmt.Println(r.Message)
	return r
}
