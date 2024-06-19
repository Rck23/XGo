package bd

import (
	"context"

	"XGo/models/usuarios"

	"go.mongodb.org/mongo-driver/bson"
)

func ChequeoYaExisteUsuario(email string) (usuarios.Usuario, bool, string) {
	ctx := context.TODO()
	db := MongoClient.Database(DatabaseName)
	col := db.Collection("usuarios")
	condition := bson.M{"email": email}
	var resultado usuarios.Usuario
	err := col.FindOne(ctx, condition).Decode(&resultado)
	ID := resultado.ID.Hex()
	if err != nil {
		return resultado, false, ID
	}
	return resultado, true, ID
}
