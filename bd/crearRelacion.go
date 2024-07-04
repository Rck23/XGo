package bd

import (
	"XGo/models"
	"context"
)

func CrearRelacion(relacion models.Relacion) (bool, error) {

	ctx := context.TODO()
	db := MongoClient.Database(DatabaseName)
	col := db.Collection("relaciones")

	_, err := col.InsertOne(ctx, relacion)
	if err != nil {
		return false, err
	}

	return true, nil
}
