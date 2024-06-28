package models

import "time"

// CrearTweet define el modelo de datos para un tweet que se va a insertar en una base de datos MongoDB.
type CrearTweet struct {
	UsuarioId string `bson:"usuarioId" json:"usuarioId, omitempty"` // Campo que almacena el identificador único del usuario que publica el tweet.
	// La etiqueta bson:"usuarioId" se utiliza para mapear este campo en documentos BSON, que es el formato de datos utilizado por MongoDB.
	// La etiqueta json:"usuarioId, omitempty" se utiliza para serializar/deserializar este campo a/from JSON. Si el valor de este campo es una cadena vacía (""),
	// entonces este campo no se incluirá en la salida JSON.

	Mensaje string `bson:"mensaje" json:"mensaje, omitempty"` // Campo que almacena el texto del tweet.
	// Similar a UsuarioId, esta etiqueta se utiliza para mapear y serializar/deserializar el campo a/from BSON y JSON.

	Fecha time.Time `bson:"fecha" json:"fecha, omitempty"` // Campo que almacena la fecha y hora en que se publicó el tweet.
	// Dado que Fecha es un tipo time.Time, que puede tener un estado inicial predeterminado (por ejemplo, el tiempo mínimo posible en Go),
	// no se utiliza omitempty aquí. Sin embargo, si se quisiera omitir este campo en la serialización a JSON cuando su valor sea el tiempo cero,
	// se podría agregar omitempty a la etiqueta json.
}
