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
	databaseName   string
	parent         (*mongo.Client)
	conn           (*mongo.Collection)
}

type Drink struct {
	Name string `json:"name" binding:"required"`
}

func Init() *DB {
	config.LoadConfig()
	configuration := *config.Config
	dbConfig := configuration.Database
	client, err := mongo.Connect(context.Background(), dbConfig.URL)

	if err != nil {
		fmt.Println(err)
	}

	db := client.Database(dbConfig.Name).Collection(dbConfig.Collection)
	fmt.Println("Connected to database")
	return &DB{
		collectionName: dbConfig.Collection,
		databaseName:   dbConfig.Name,
		parent:         client,
		conn:           db}
}

func (dbPtr *DB) Disconnect() {
	fmt.Println("Disconnecting from database...")
	err := dbPtr.parent.Disconnect(context.Background())

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Disconnected from database")
}

func (dbPtr *DB) CreateDrink(data Drink) (*mongo.InsertOneResult, error) {
	res, err := dbPtr.conn.InsertOne(context.Background(), data)

	if err != nil {
		fmt.Println(err)
	}
	return res, err
}
