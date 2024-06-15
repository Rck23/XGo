package bd

import (
	"context"
	"fmt"

	"XGo/models"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoClient es una variable global que guarda la conexión al cliente de MongoDB.
// Se utiliza para realizar operaciones en la base de datos.
var MongoClient *mongo.Client

// DatabaseName guarda el nombre de la base de datos a la que se conecta.
// Este nombre se obtiene del contexto durante la conexión inicial.
var DatabaseName string

// ConectarDB intenta establecer una conexión con la base de datos MongoDB utilizando credenciales y host proporcionados
// a través del contexto. Retorna un error si la conexión falla.
func ConectarDB(ctx context.Context) error {
	// Extraer credenciales y host de contexto
	user := ctx.Value(models.Key("user")).(string)
	password := ctx.Value(models.Key("password")).(string)
	host := ctx.Value(models.Key("host")).(string)

	// Construir la cadena de conexión
	conexionStr := fmt.Sprintf("mongodb+srv://%s:%s@%s/?retryWrites=true&w=majority", user, password, host)

	// Aplicar opciones de conexión
	var clientOptions = options.Client().ApplyURI(conexionStr)

	// Intentar conectar
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	// Verificar conexión
	err = client.Ping(ctx, nil)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	// Establecer variables globales y confirmar conexión exitosa
	fmt.Println("Conexión exitosa con la BD")
	MongoClient = client
	DatabaseName = ctx.Value(models.Key("database")).(string)

	return nil
}

// BaseConectada verifica si la conexión a la base de datos MongoDB aún está activa.
// Retorna true si la conexión es válida, false en caso contrario.
func BaseConectada() bool {
	err := MongoClient.Ping(context.TODO(), nil)
	return err == nil
}
