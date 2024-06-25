package bd

import (
	"context"

	"XGo/models/usuarios"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func BuscoPerfil(ID string) (usuarios.Usuario, error) {
	ctx := context.TODO()

	db := MongoClient.Database(DatabaseName)
	col := db.Collection("usuarios")

	var perfil usuarios.Usuario
	objID, _ := primitive.ObjectIDFromHex(ID)

	condicion := bson.M{
		"_id": objID,
	}

	err := col.FindOne(ctx, condicion).Decode(&perfil)
	perfil.Password = ""
	if err != nil {
		return perfil, err
	}

	return perfil, nil
}
