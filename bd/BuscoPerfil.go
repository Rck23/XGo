package bd

import (
	"context"

	"XGo/models/usuarios"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// BuscoPerfil busca un perfil de usuario en la colección "usuarios" de MongoDB por su ID.
// Devuelve el perfil de usuario encontrado, con la contraseña reemplazada por una cadena vacía para seguridad,
// y un error si ocurre alguno durante la búsqueda.
func BuscoPerfil(ID string) (usuarios.Usuario, error) {
	// Crea un contexto con timeout indefinido para la búsqueda.
	ctx := context.TODO()

	// Obtiene la instancia de la base de datos y la colección "usuarios".
	db := MongoClient.Database(DatabaseName)
	col := db.Collection("usuarios")

	// Prepara una variable para almacenar el perfil de usuario encontrado.
	var perfil usuarios.Usuario
	// Convierte el ID proporcionado a un ObjectID de MongoDB.
	objID, _ := primitive.ObjectIDFromHex(ID)

	// Define la condición de búsqueda basada en el ID del usuario.
	condicion := bson.M{
		"_id": objID,
	}
	// Realiza la búsqueda en la colección "usuarios" usando la condición definida.
	err := col.FindOne(ctx, condicion).Decode(&perfil)
	// Reemplaza la contraseña por una cadena vacía para evitar exponer información sensible.
	perfil.Password = ""
	// Si ocurre un error durante la búsqueda, se devuelve el perfil y el error.
	if err != nil {
		return perfil, err
	}
	// Si el perfil es encontrado, se devuelve el perfil sin la contraseña.
	return perfil, nil
}
