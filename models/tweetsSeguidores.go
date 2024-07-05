package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type VerTweetsSeguidores struct {
	ID                primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	UsuarioId         string             `bson:"usuarioId" json:"usuarioId,omitempty"`
	UsuarioRelacionId string             `bson:"usuarioRelacionId" json:"usuarioRelacionId,omitempty"`
	Tweet             struct {
		Mensaje string    `bson:"mensaje" json:"mensaje,omitempty"`
		Fecha   time.Time `bson:"fecha" json:"fecha,omitempty"`
		ID      string    `bson:"_id" json:"_id,omitempty"`
	}
}
