package bd

import (
	"XGo/models"
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

func VerTweetsSeguidores(Id string, pagina int) ([]*models.VerTweetsSeguidores, bool) {
	context := context.TODO()
	db := MongoClient.Database(DatabaseName)
	col := db.Collection("relaciones")

	skip := (pagina - 1) * 20

	condiciones := make([]bson.M, 0)
	condiciones = append(condiciones, bson.M{"$match": bson.M{"usuarioId": Id}})
	condiciones = append(condiciones, bson.M{
		"$lookup": bson.M{
			"from":         "tweets",
			"localfield":   "usuarioRelacionId",
			"foreignField": "usuarioId",
			"as":           "tweets",
		}})
	condiciones = append(condiciones, bson.M{"$unwind": "$tweets"})
	condiciones = append(condiciones, bson.M{
		"$sort": bson.M{"tweets.fecha": -1},
	})
	condiciones = append(condiciones, bson.M{"$skip": skip})
	condiciones = append(condiciones, bson.M{"$limit": 20})

	var resultado []*models.VerTweetsSeguidores

	cursor, err := col.Aggregate(context, condiciones)
	if err != nil {
		return resultado, false
	}

	err = cursor.All(context, &resultado)
	if err != nil {
		return resultado, false
	}

	return resultado, true

}
