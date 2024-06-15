package models

import (
	jwt "github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Claim representa los claims personalizados junto con los claims registrados de un token JWT.
// Incluye el correo electrónico del usuario y su ID único de MongoDB.
type Claim struct {
	// Email es el correo electrónico del usuario.
	Email string `json:"email"`

	// ID es el identificador único del usuario en MongoDB.
	ID primitive.ObjectID `bson:"_id" json:"_id,omitempty"`

	// RegisteredClaims contiene los claims estándar de un token JWT, como expiración, emisor, etc.
	// Al extender jwt.RegisteredClaims, se pueden incluir estos claims estándar junto con los claims personalizados.
	jwt.RegisteredClaims
}
