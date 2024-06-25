package bd

import (
	"context"

	"XGo/models/usuarios"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ModificoRegistro(usuario usuarios.Usuario, ID string) (bool, error) {

	ctx := context.TODO()
	db := MongoClient.Database(DatabaseName)
	col := db.Collection("usuarios")

	registro := make(map[string]interface{})

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

	updateString := bson.M{
		"$set": registro,
	}

	objID, _ := primitive.ObjectIDFromHex(ID)

	filtro := bson.M{"_id": bson.M{"$eq": objID}}

	_, err := col.UpdateOne(ctx, filtro, updateString)
	if err != nil {
		return false, err

	}
	return false, nil
}
