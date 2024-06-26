package bd

import (
	"context"
	"fmt"

	"XGo/models"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client
var DatabaseName string

// ConectarDB intenta conectar a una base de datos MongoDB utilizando credenciales y detalles de conexión proporcionados a través de un contexto.
// Exige que el contexto contenga valores para "user", "password", "host", y "database". Utiliza estos valores para construir la URI de conexión.
// Devuelve un error si la conexión falla o si hay un problema al realizar un ping a la base de datos.
func ConectarDB(ctx context.Context) error {
	// Extrae las credenciales y detalles de conexión del contexto.
	user := ctx.Value(models.Key("user")).(string)
	password := ctx.Value(models.Key("password")).(string)
	host := ctx.Value(models.Key("host")).(string)

	// Construye la URI de conexión utilizando las credenciales y detalles extraídos.
	conexionStr := fmt.Sprintf("mongodb+srv://%s:%s@%s/?retryWrites=true&w=majority", user, password, host)

	// Crea opciones de cliente con la URI de conexión.
	var clientOptions = options.Client().ApplyURI(conexionStr)

	// Intenta conectar al servidor MongoDB.
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		// Imprime el error y lo devuelve si la conexión falla.
		fmt.Println(err.Error())
		return err
	}

	// Realiza un ping a la base de datos para verificar la conectividad.
	err = client.Ping(ctx, nil)
	if err != nil {
		// Imprime el error y lo devuelve si el ping falla.
		fmt.Println(err.Error())
		return err
	}

	// Si todo va bien, imprime un mensaje de éxito y guarda el cliente y el nombre de la base de datos.
	fmt.Println("Conexión exitosa con la BD")
	MongoClient = client
	DatabaseName = ctx.Value(models.Key("database")).(string)

	return nil
}

// BaseConectada verifica si la conexión a la base de datos MongoDB está activa realizando un ping.
// Devuelve true si la conexión está activa y false si no lo está.
func BaseConectada() bool {
	// Realiza un ping a la base de datos utilizando el cliente de MongoDB.
	err := MongoClient.Ping(context.TODO(), nil)
	// Comprueba si hubo un error durante el ping.

	return err == nil
}
