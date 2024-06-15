package usuarios

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Usuario representa un usuario en la aplicación.
// Contiene información básica y opcionalmente información adicional para personalización y perfil público.
type Usuario struct {
	// ID es el identificador único asignado por MongoDB para cada documento de usuario.
	// Este campo es obligatorio y se omite durante la serialización a JSON si su valor es el valor cero para su tipo.
	ID primitive.ObjectID `bson:"_id,omitempty" json:"id"`

	// Nombre es el nombre del usuario.
	Nombre string `bson:"nombre" json:"nombre,omitempty"`

	// Apellidos son los apellidos del usuario.
	Apellidos string `bson:"apellidos" json:"apellidos,omitempty"`

	// FechaNacimiento es la fecha de nacimiento del usuario.
	FechaNacimiento time.Time `bson:"fechaNacimiento" json:"fechaNacimiento,omitempty"`

	// Email es el correo electrónico del usuario.
	Email string `bson:"email" json:"email"`

	// Password es la contraseña del usuario.
	// Dado que esta información sensible, se recomienda manejarla de manera segura y nunca exponerla.
	Password string `bson:"password" json:"password,omitempty"`

	// Avatar es la URL o el identificador del avatar del usuario.
	Avatar string `bson:"avatar" json:"avatar,omitempty"`

	// Banner es la URL o el identificador del banner del usuario.
	Banner string `bson:"banner" json:"banner,omitempty"`

	// Biografia es la biografía o descripción del usuario.
	Biografia string `bson:"biografia" json:"biografia,omitempty"`

	// Ubicacion es la ubicación del usuario.
	Ubicacion string `bson:"ubicacion" json:"ubicacion,omitempty"`

	// SitioWeb es el sitio web o perfil profesional del usuario.
	SitioWeb string `bson:"sitioweb" json:"sitioweb,omitempty"`
}
