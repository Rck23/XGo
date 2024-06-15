package bd

import (
	"context"

	models "XGo/models/usuarios"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// InsertoRegistro toma un objeto Usuario, encripta su contraseña, e inserta el usuario en la colección "usuarios" de MongoDB.
// Devuelve el ID del usuario insertado como una cadena, un booleano indicando si la operación fue exitosa, y un error si ocurrió alguno.
func InsertoRegistro(user models.Usuario) (string, bool, error) {
	// Crea un contexto con timeout indefinido para la operación de inserción.
	ctx := context.TODO()

	// Obtiene la instancia de la base de datos y la colección "usuarios".
	db := MongoClient.Database(DatabaseName)
	col := db.Collection("usuarios")

	// Encripta la contraseña del usuario antes de insertarlo en la base de datos.
	user.Password, _ = EncriptarPassword(user.Password)

	// Intenta insertar el usuario en la colección "usuarios".
	resultado, err := col.InsertOne(ctx, user)
	if err != nil {
		// Si ocurre un error durante la inserción, se devuelve un mensaje vacío, false y el error.
		return "", false, err
	}

	// Obtiene el ID del usuario insertado.
	ObjID, _ := resultado.InsertedID.(primitive.ObjectID)

	// Devuelve el ID del usuario insertado como una cadena, true indicando éxito, y nil para el error.
	return ObjID.String(), true, nil
}
