package bd

import (
	"XGo/models"
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

func ConsultarRelacion(relacion models.Relacion) bool {

	ctx := context.TODO()                    // Crea un contexto sin tiempo de espera para operaciones asíncronas.
	db := MongoClient.Database(DatabaseName) // Accede a la base de datos.
	col := db.Collection("relaciones")       // Selecciona la colección "relaciones".

	condicion := bson.M{
		"usuarioId":         relacion.UsuarioId,
		"usuarioRelacionId": relacion.UsuarioRelacionId,
	}

	var resultado models.Relacion

	err := col.FindOne(ctx, condicion).Decode(&resultado)

	if err != nil {
		return false
	}

	return true
}
