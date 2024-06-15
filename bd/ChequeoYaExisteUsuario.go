package bd

import (
	"context"

	"XGo/models/usuarios"

	"go.mongodb.org/mongo-driver/bson"
)

// ChequeoYaExisteUsuario busca un usuario en la colección "usuarios" de MongoDB por su dirección de correo electrónico.
// Devuelve el usuario encontrado, un booleano indicando si el usuario existe, y el ID del usuario como una cadena.
func ChequeoYaExisteUsuario(email string) (usuarios.Usuario, bool, string) {
	// Crea un contexto con timeout indefinido para la búsqueda.
	ctx := context.TODO()

	// Obtiene la instancia de la base de datos y la colección "usuarios".
	db := MongoClient.Database(DatabaseName)
	col := db.Collection("usuarios")

	// Define la condición de búsqueda basada en el correo electrónico proporcionado.
	condition := bson.M{"email": email}

	// Prepara una variable para almacenar el resultado de la búsqueda.
	var resultado usuarios.Usuario

	// Realiza la búsqueda en la colección "usuarios" usando la condición definida.
	err := col.FindOne(ctx, condition).Decode(&resultado)
	ID := resultado.ID.Hex() // Obtiene el ID del usuario encontrado como una cadena.

	// Si ocurre un error durante la búsqueda (por ejemplo, si el usuario no existe), se devuelve el usuario vacío,
	// false y el ID (que será una cadena vacía si el usuario no existe).
	if err != nil {
		return resultado, false, ID
	}

	// Si el usuario es encontrado, se devuelve el usuario, true y el ID del usuario.
	return resultado, true, ID
}
