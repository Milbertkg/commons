package mongodb

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MongoDbConnections(uri string) (*mongo.Client, error) {

	clientOptions := options.Client().ApplyURI(uri)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client, err
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client, err
	}
	log.Println("Connection to Mongodb Database is sussesfull.")
	return client, nil

}

/*ChequeoConnection es el Ping a la BD */
func CheckMDBConnection(client *mongo.Client) bool {
	err := client.Ping(context.TODO(), nil)
	if err != nil {
		return false
	}
	return true
}
