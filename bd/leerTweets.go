package bd

import (
	"XGo/models"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func LeerTweets(ID string, pagina int64) ([]*models.DevuelvoTweets, bool) {

	ctx := context.TODO() // Crea un contexto sin tiempo de espera para operaciones asíncronas.

	db := MongoClient.Database(DatabaseName) // Accede a la base de datos.
	col := db.Collection("tweets")           // Selecciona la colección "tweets".

	// Inicializa un slice para almacenar los resultados.
	var resultados []*models.DevuelvoTweets

	condicion := bson.M{
		// Define la condición para filtrar tweets por el ID del usuario.
		"usuarioId": ID,
	}

	/* Configura las opciones de búsqueda.
	SetLimit establece el número máximo de documentos a retornar.
	SetSort ordena los resultados por fecha en orden descendente (-1).
	SetSkip calcula el número de documentos a omitir para paginar correctamente.*/
	opciones := options.Find()
	opciones.SetLimit(20)
	opciones.SetSort(bson.D{{Key: "fecha", Value: -1}})
	opciones.SetSkip((pagina - 1) * 20)

	// Realiza la consulta a la base de datos.
	conjuntoRegistros, err := col.Find(ctx, condicion, opciones)
	if err != nil {
		// Retorna los resultados y un booleano indicando fallo.
		return resultados, false
	}

	for conjuntoRegistros.Next(ctx) { // Itera sobre cada documento retornado.
		var registro models.DevuelvoTweets         // Inicializa un registro temporal para decodificar cada documento.
		err := conjuntoRegistros.Decode(&registro) // Decodifica el documento en el registro temporal.
		if err != nil {
			return resultados, false // Retorna los resultados y un booleano indicando fallo.
		}

		resultados = append(resultados, &registro) // Agrega el registro decodificado al slice de resultados.
	}

	return resultados, true // Retorna los resultados y un booleano indicando éxito.

}
