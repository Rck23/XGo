package bd

import (
	"context"

	"XGo/models/usuarios"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ModificoRegistro actualiza un registro de usuario en la colección "usuarios" de MongoDB.
// Toma un objeto Usuario con campos actualizados y un ID de usuario como entrada.
// Actualiza los campos del usuario en la base de datos si están presentes y diferentes de cadenas vacías.
// Devuelve un booleano indicando si la operación fue exitosa y un error si ocurrió alguno.
func ModificoRegistro(usuario usuarios.Usuario, ID string) (bool, error) {
	// Crea un contexto con timeout indefinido para la operación de actualización.
	ctx := context.TODO()

	// Obtiene la instancia de la base de datos y la colección "usuarios".
	db := MongoClient.Database(DatabaseName)
	col := db.Collection("usuarios")

	// Prepara un mapa para almacenar los campos del usuario que necesitan ser actualizados.
	registro := make(map[string]interface{})

	// Agrega los campos del usuario al mapa de actualización si están presentes y diferentes de cadenas vacías.
	if len(usuario.Nombre) > 0 {
		registro["nombre"] = usuario.Nombre
	}

	if len(usuario.Apellidos) > 0 {
		registro["apellidos"] = usuario.Apellidos
	}

	registro["fechaNacimiento"] = usuario.FechaNacimiento

	if len(usuario.Avatar) > 0 {
		registro["avatar"] = usuario.Avatar
	}

	if len(usuario.Banner) > 0 {
		registro["banner"] = usuario.Banner
	}

	if len(usuario.Biografia) > 0 {
		registro["biografia"] = usuario.Biografia
	}

	if len(usuario.Ubicacion) > 0 {
		registro["ubicacion"] = usuario.Ubicacion
	}

	if len(usuario.SitioWeb) > 0 {
		registro["sitioweb"] = usuario.SitioWeb
	}

	// Prepara la operación de actualización con el mapa de campos.
	updateString := bson.M{
		"$set": registro,
	}
	// Convierte el ID del usuario a un ObjectID de MongoDB.
	objID, _ := primitive.ObjectIDFromHex(ID)
	// Prepara el filtro para buscar el documento por ID.
	filtro := bson.M{"_id": bson.M{"$eq": objID}}
	// Intenta actualizar el documento en la colección "usuarios".
	_, err := col.UpdateOne(ctx, filtro, updateString)
	if err != nil {
		// Si ocurre un error durante la actualización, devuelve false y el error.
		return false, err

	}

	// Si la actualización es exitosa, devuelve true y nil para el error.
	return false, nil
}
