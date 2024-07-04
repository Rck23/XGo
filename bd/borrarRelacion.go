package bd

import (
	"XGo/models"
	"context"
)

func BorrarRelacion(relacion models.Relacion) (bool, error) {
	ctx := context.TODO()
	db := MongoClient.Database(DatabaseName)
	col := db.Collection("relaciones")

	_, err := col.DeleteOne(ctx, relacion)
	if err != nil {
		return false, err
	}

	return true, nil
}
