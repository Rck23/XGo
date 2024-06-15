package bd

import "golang.org/x/crypto/bcrypt"

// EncriptarPassword toma una contraseña en texto plano y la encripta usando bcrypt.
// Devuelve la versión encriptada de la contraseña y un error si ocurre alguno durante el proceso.
func EncriptarPassword(pass string) (string, error) {
	// Selecciona un costo de encriptación de 8, que es un equilibrio común entre seguridad y rendimiento.
	costo := 8

	// Genera la versión encriptada de la contraseña usando bcrypt con el costo seleccionado.
	// Si ocurre un error durante la generación, se devuelve el mensaje de error y el error.
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), costo)
	if err != nil {
		return err.Error(), err
	}

	// Convierte la versión encriptada de bytes a una cadena y la devuelve junto con un error nulo.
	return string(bytes), nil
}
