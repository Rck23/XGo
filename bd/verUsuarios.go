package bd

import (
	"XGo/models"
	"XGo/models/usuarios"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func VerUsuarios(Id string, pagina int64, buscar string, tipo string) ([]*usuarios.Usuario, bool) {
	context := context.TODO()
	db := MongoClient.Database(DatabaseName)
	col := db.Collection("usuarios")

	var resultados []*usuarios.Usuario

	opciones := options.Find()
	opciones.SetLimit(20)
	opciones.SetSkip((pagina - 1) * 20)

	query := bson.M{
		"nombre": bson.M{
			"$regex": `(?i)` + buscar,
		},
	}

	cursor, err := col.Find(context, query, opciones)
	if err != nil {
		fmt.Println(err)
		return resultados, false
	}

	var incluir bool

	for cursor.Next(context) {
		var usuario usuarios.Usuario

		err := cursor.Decode(&usuario)
		if err != nil {
			fmt.Println("Decode = " + err.Error())
			return resultados, false
		}

		var relacion models.Relacion
		relacion.UsuarioId = Id
		relacion.UsuarioRelacionId = usuario.ID.Hex()

		incluir = false

		encontrado := ConsultarRelacion(relacion)

		if tipo == "new" && !encontrado {
			incluir = true
		}
		if tipo == "follow" && encontrado {
			incluir = true
		}
		if relacion.UsuarioRelacionId == Id {
			incluir = false
		}
		if incluir {
			usuario.Password = ""
			resultados = append(resultados, &usuario)
		}
	}
	err = cursor.Err()
	if err != nil {
		fmt.Println("cursor.Err() = " + err.Error())
		return resultados, false
	}
	cursor.Close(context)
	return resultados, true
}
