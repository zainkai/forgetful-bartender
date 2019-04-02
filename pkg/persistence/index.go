package persistence

import (
	"context"
	"fmt"

	// "github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
	// "github.com/mongodb/mongo-go-driver/mongo/options"

	"github.com/zainkai/forgetful-bartender/configs"
)

type DB struct {
	collectionName string
	databaseName string
	parent (*mongo.Client)
	conn (*mongo.Collection)
}

type Drink struct {
	name string
}

func Init() (*DB) {
	config.LoadConfig()
	configuration := config.Config
	client, err := mongo.Connect(context.Background(), configuration.DatabaseURL)

	if err != nil {
		fmt.Println(err)
	}
	db := client.Database("forgetful-db-v1").Collection("drinks")

	fmt.Println("Connected to database")
	return &DB{
		collectionName: "drinks",
		databaseName: "forgetful-db-v1",
		parent: client,
		conn: db}
}

func (dbPtr *DB) Disconnect() {
	fmt.Println("Disconnecting from database...")
	err := dbPtr.parent.Disconnect(context.Background())

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Disconnected from database")
}

func (dbPtr *DB) CreateDrink(data Drink) {
	res, err := dbPtr.conn.InsertOne(context.Background(), data)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res.InsertedID)
}

