package models

// Secret almacena información sensible necesaria para conectarse a recursos protegidos, como una base de datos,
// un servicio externo, o cualquier otro recurso que requiera autenticación. Esta información incluye el host,
// el puerto (si aplica), el nombre de usuario, la contraseña, el secreto JWT para firmar tokens, y el nombre
// de la base de datos. Es importante manejar esta información con cuidado para prevenir exposiciones de datos
// sensibles.
type Secret struct {
	Host string `json:"host"`

	Username string `json:"username"`

	Password string `json:"password"`

	JWTSign string `json:"jwtsign"`

	Database string `json:"database"`
}
