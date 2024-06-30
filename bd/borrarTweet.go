package bd

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// BorrarTweet es una función que elimina un tweet específico de la base de datos basado en su ID y el ID del usuario.
// Retorna un error si algo sale mal durante la operación.
func BorrarTweet(Id string, usuarioId string) error {
	// Creamos un contexto sin valor para la operación. Esto es útil para controlar la cancelación de la operación.
	context := context.TODO()
	// Accedemos a la base de datos usando MongoClient y obtenemos una referencia a la colección "tweets".
	db := MongoClient.Database(DatabaseName)
	col := db.Collection("tweets")

	// Convertimos la cadena hexadecimal Id en un ObjectID de MongoDB para poder buscar el documento por su ID.
	objId, _ := primitive.ObjectIDFromHex(Id)

	// Preparamos las condiciones para la eliminación. Buscamos un documento cuyo campo _id sea igual al ObjectID convertido
	// y cuyo campo usuarioId coincida con el usuarioId proporcionado.
	condicion := bson.M{
		"_id":       objId,
		"usuarioId": usuarioId,
	}

	// Intentamos eliminar el documento que cumple con nuestras condiciones. Si la operación es exitosa, retorna el número
	// de documentos eliminados y un error nulo. Si hay un error, retorna el número de documentos afectados y el error.
	_, err := col.DeleteOne(context, condicion)
	// Retornamos cualquier error que haya ocurrido durante la operación de eliminación.
	return err

}
