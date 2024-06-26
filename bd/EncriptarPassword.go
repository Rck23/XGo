package bd

import "golang.org/x/crypto/bcrypt"

// EncriptarPassword toma una contraseña en texto plano y la encripta utilizando bcrypt con un costo específico.
// Devuelve la versión encriptada de la contraseña como una cadena y un error si ocurre alguno durante el proceso de encriptación.
func EncriptarPassword(pass string) (string, error) {

	// Define el costo de encriptación. Un valor mayor significa una encriptación más segura pero puede ser más lenta.
	costo := 8

	// Genera la versión encriptada de la contraseña utilizando bcrypt con el costo especificado.
	// Si ocurre un error durante la generación de la versión encriptada, se devuelve el mensaje de error y el error.
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), costo)
	if err != nil {
		return err.Error(), err
	}

	// Convierte la versión encriptada de bytes a una cadena y la devuelve junto con un error nulo.
	return string(bytes), nil
}
