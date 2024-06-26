package models

import (
	jwt "github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Claim representa las reclamaciones asociadas a un token JWT. Incluye información sobre el usuario
// (como su email e ID) junto con las reclamaciones registradas por defecto de la biblioteca jwt-go.
type Claim struct {
	Email string             `json:"email"`
	ID    primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	jwt.RegisteredClaims
}
