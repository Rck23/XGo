package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// DevuelvoTweets define el modelo de datos para un tweet que se recupera de una base de datos MongoDB.
type DevuelvoTweets struct {
	ID primitive.ObjectID `bson:"_id" json:"_id,omitempty"` // Campo que almacena el ID del tweet generado automáticamente por MongoDB.
	// La etiqueta bson:"_id" se utiliza para mapear este campo en documentos BSON, que es el formato de datos utilizado por MongoDB.
	// La etiqueta json:"_id,omitempty" se utiliza para serializar/deserializar este campo a/from JSON. Si el valor de este campo es el ID nulo predeterminado,
	// entonces este campo no se incluirá en la salida JSON.

	UsuarioId string `bson:"usuarioId" json:"usuarioId,omitempty"` // Campo que almacena el identificador único del usuario que publicó el tweet.
	// Similar a ID, esta etiqueta se utiliza para mapear y serializar/deserializar el campo a/from BSON y JSON. Si el valor de este campo es una cadena vacía,
	// entonces este campo no se incluirá en la salida JSON.

	Mensaje string `bson:"mensaje" json:"mensaje,omitempty"` // Campo que almacena el texto del tweet.
	// Este campo es fundamental para el contenido del tweet. Al igual que los anteriores, se utiliza para mapear y serializar/deserializar el campo a/from BSON y JSON.

	Fecha time.Time `bson:"fecha" json:"fecha,omitempty"` // Campo que almacena la fecha y hora en que se publicó el tweet.
	// Este campo es importante para rastrear cuándo se creó o modificó el tweet. Al igual que los otros campos, se utiliza para mapear y serializar/deserializar el campo a/from BSON y JSON.
}
