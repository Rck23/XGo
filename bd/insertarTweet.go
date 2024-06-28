package bd

import (
	"XGo/models"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// InsertarTweet es una función que intenta insertar un nuevo tweet en la colección "tweets" de la base de datos MongoDB.
func InsertarTweet(tweet models.CrearTweet) (string, bool, error) {
	ctx := context.TODO() // Crea un contexto sin tiempo de espera para operaciones asíncronas.

	// Obtiene la instancia de la base de datos y la colección "usuarios".
	db := MongoClient.Database(DatabaseName)
	col := db.Collection("tweets")

	// Prepara el documento a insertar en la base de datos utilizando el struct CrearTweet.
	registro := bson.M{
		"usuarioId": tweet.UsuarioId,
		"mensaje":   tweet.Mensaje,
		"fecha":     tweet.Fecha,
	}

	// Intenta insertar el tweet en la colección "tweets". La operación es asíncrona y devuelve un ID generado automáticamente si es exitosa.
	resultado, err := col.InsertOne(ctx, registro)
	if err != nil {
		// Si ocurre un error durante la inserción, se devuelve un mensaje vacío, false y el error.
		return "", false, err
	}

	// Convierte el ID generado automáticamente al tipo string para devolverlo como respuesta.
	objID, _ := resultado.InsertedID.(primitive.ObjectID)

	return objID.String(), true, nil
}
