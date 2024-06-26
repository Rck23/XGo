package models

// RespuestaLogin representa la respuesta esperada después de un intento de inicio de sesión exitoso.
// Contiene un único campo, 'Token', que es un string codificado en base64 que representa el token JWT
// (JSON Web Token) generado para el usuario. Este token se utiliza posteriormente para autenticar
// solicitudes subsiguientes del usuario al sistema.
type RespuestaLogin struct {
	// json:"token,omitempty": Los tags json especifican cómo se serializa y deserializa el campo en JSON.
	// El tag omitempty indica que el campo puede ser omitido durante la serialización si su valor es la
	// cadena vacía (""). Esto es útil para evitar enviar tokens vacíos o nulos en la respuesta.
	Token string `json:"token",omitempty`
}
