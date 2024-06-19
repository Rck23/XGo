package bd

import (
	"context"

	models "XGo/models/usuarios"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InsertoRegistro(user models.Usuario) (string, bool, error) {
	ctx := context.TODO()
	db := MongoClient.Database(DatabaseName)
	col := db.Collection("usuarios")
	user.Password, _ = EncriptarPassword(user.Password)
	resultado, err := col.InsertOne(ctx, user)
	if err != nil {
		return "", false, err
	}

	ObjID, _ := resultado.InsertedID.(primitive.ObjectID)

	return ObjID.String(), true, nil
}
