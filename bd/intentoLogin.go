package bd

import (
	"XGo/models/usuarios"

	"golang.org/x/crypto/bcrypt"
)

// IntentoLogin intenta autenticar a un usuario basado en su correo electrónico y contraseña proporcionados.
// Primero, verifica si el usuario existe en la base de datos. Luego, si el usuario existe, compara la contraseña proporcionada
// con la versión encriptada almacenada en la base de datos utilizando bcrypt.
// Devuelve el usuario y un booleano indicando si el inicio de sesión fue exitoso o no.
func IntentoLogin(email string, password string) (usuarios.Usuario, bool) {
	// Busca el usuario en la base de datos por su correo electrónico.
	usu, encontrado, _ := ChequeoYaExisteUsuario(email)
	// Si el usuario no existe, devuelve el usuario y false.
	if !encontrado {
		return usu, false
	}
	// Convierte la contraseña proporcionada y la contraseña almacenada en la base de datos a slices de bytes.
	passwordBytes := []byte(password)
	passwordBD := []byte(usu.Password)

	// Compara la contraseña proporcionada con la versión encriptada almacenada.
	err := bcrypt.CompareHashAndPassword(passwordBD, passwordBytes)

	// Si hay un error durante la comparación (lo que significa que las contraseñas no coinciden),
	// devuelve el usuario y false.
	if err != nil {
		return usu, false
	}

	// Si las contraseñas coinciden, devuelve el usuario y true.
	return usu, true

}
