package models

// Secret almacena información sensible utilizada para autenticarse en servicios externos,
// como bases de datos o APIs protegidas mediante JWT.
type Secret struct {
	// Host es el nombre del host donde se encuentra alojado el servicio al que se accede.
	// Por ejemplo, podría ser el nombre de dominio de una base de datos o API.
	Host string `json:"host"`

	// Username es el nombre de usuario utilizado para autenticarse en el servicio.
	// Debe ser único dentro del contexto del servicio especificado.
	Username string `json:"username"`

	// Password es la contraseña asociada al nombre de usuario para autenticarse en el servicio.
	// Debe ser segura y manejada adecuadamente para evitar exposiciones.
	Password string `json:"password"`

	// JWTSign es el token JWT utilizado para autenticaciones basadas en tokens.
	// Este token debe ser generado previamente y enviado junto con las solicitudes
	// que requieran autenticación.
	JWTSign string `json:"jwtsign"`

	// Database es el nombre de la base de datos a la que se accede.
	// Este campo es opcional y solo debería ser utilizado si la estructura
	// Secret está destinada a almacenar información específica de una base de datos.
	Database string `json:"database"`
}
